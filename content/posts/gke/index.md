---
title: "GKE in production"
date: 2017-07-02
draft: false
slug: gke
city: Joinville
toc: true
tags: [kubernetes, google-cloud]
---

I've been working with DigitalOcean, Heroku and AWS for some years now.

Recently, I decided to give GCE (Google Compute Engine), and, more specifically, GKE (Google Container Engine) a try. In this post I intend to show a few things I learned and/or struggled with.

## SSL and Load Balancers

On AWS, we can define a Kubernetes service with a SSL certificate using annotations, something like this:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: myapp
  annotations:
    service.beta.kubernetes.io/aws-load-balancer-ssl-cert: arn:aws:iam::123:server-certificate/my_certificate
    service.beta.kubernetes.io/aws-load-balancer-backend-protocol: http
    service.beta.kubernetes.io/aws-load-balancer-ssl-ports: https
spec:
  ports:
  - name: https
    port: 443
    targetPort: 8080
  selector:
    app: myapp
  type: LoadBalancer
```

This will create an Elastic LoadBalancer with the given certificate on the port 443.

I looked into doing the same thing on GKE, but, as it turns out, GKE uses Ingresses to setup SSL, turning it into a Load Balancer, which seems odd because services also create load balancers. In the end, it turned out to work great:

```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: app
spec:
  ports:
  - name: http
    port: 80
    protocol: TCP
    targetPort: 80
  selector:
    app: app
  sessionAffinity: None
  type: NodePort
---
apiVersion: v1
kind: Service
metadata:
  name: api
spec:
  ports:
  - name: http
    port: 80
    protocol: TCP
    targetPort: 8080
  selector:
    app: api
  sessionAffinity: None
  type: NodePort
---
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: app
spec:
  tls:
  - secretName: tls-secret
  rules:
  - host: app.mysvc.com
    http:
      paths:
      - backend:
          serviceName: app
          servicePort: 80
  - host: api.mysvc.com
    http:
      paths:
      - backend:
          serviceName: api
          servicePort: 80
```

Of course, you need to create the `tls-secret` that's being used by the ingress:

```
$ kubectl create secret tls tls-secret --key=tls.key --cert=tls.crt
```

That will create a single GCE load balancer with all the host and path rules, as well as the health checks.

More on ingresses and GKE [here](https://github.com/kubernetes/ingress-gce/blob/master/docs/faq/README.md).

**PROTIP**: the healtcheck is `/` by default, so make sure all the services respond with a `200 OK` for that.

## Sending emails

One of the services I deployed uses Sendgrid to send e-mails. It's a Rails app, so it were using the default mechanisms for that, which connects on port `25` by default. It worked locally, but not on GKE. Why?

It turns out GCE blocks the port `25`. We can confirm that by using telnet, for example:

```sh
$ telnet smtp.sendgrid.net 25
Trying 108.168.190.110...
Connected to smtp.sendgrid.net.
Escape character is '^]'.
220 SG ESMTP service ready at ismtpd0001p1iad1.sendgrid.net
^]
telnet> Connection closed.

$ kubeclt exec -it app-620440047-d4tg5 /bin/bash
root@app-620440047-d4tg5:/# apt-get update && apt-get install telnet
# ...
root@app-620440047-d4tg5:/# telnet smtp.sendgrid.net 25
Trying 108.168.190.110...
Trying 108.168.190.109...
^C
```

Lucky for us, Sendgrid also replies on `2525`, which is not blocked by GCE:

```sh
root@app-620440047-d4tg5:/# telnet smtp.sendgrid.net 2525
Trying 108.168.190.110...
Connected to smtp.sendgrid.net.
Escape character is '^]'.
220 SG ESMTP service ready at ismtpd0001p1iad1.sendgrid.net
^]
telnet>
```

So, I just had to change the config to user port `2525` instead of `25`.

Another solution is to use a postfix sidecar container, which seems overkill to me.

## Connecting to a Cloud SQL database

On AWS, I usually create a RDS instance inside my VPC and connect to it using its address, a username and a password.

On GCE, however, seems like the right way of connecting to Cloud SQL databases is by using a sidecar container called `cloudsql-proxy`:

```yaml
- command:
  - /cloud_sql_proxy
  - --dir=/cloudsql
  - -instances=mysvc-123456:us-central1:mydb=tcp:5432
  - -credential_file=/secrets/cloudsql/credentials.json
  image: gcr.io/cloudsql-docker/gce-proxy:1.09
  name: cloudsql-proxy
  volumeMounts:
  - mountPath: /secrets/cloudsql
    name: cloudsql-instance-credentials
    readOnly: true
  - mountPath: /etc/ssl/certs
    name: ssl-certs
  - mountPath: /cloudsql
    name: cloudsql
```

So, in your "real container", you connect to `localhost:5432` instead of the CloudSQL address. Not very straightfoward, but works.

To be honest, I didn't like this very much, because it clutters my deployment definition. Anyway, you can learn more [here](https://cloud.google.com/sql/docs/mysql/connect-container-engine).

## Deploying

Since they are small apps with very few people working on them, I'm deploying with a simple shell script:

```sh
#!/bin/sh
set -xeo pipefail

# make sure we are on the right context
gcloud config set project mysvc-123456
kubectl config get-contexts |
  grep mysvc |
  cut -f2 -d'*' |
  awk '{print $1}' |
  xargs kubectl config use-context

# set the build tag and the base image name
build=$(date "+%Y%m%d%H%M")
image="gcr.io/mysvc-123456/mysvc"

# build and tag
docker build -t "$image:$build" .
docker tag "$image:$build" "$image:latest"

# push both the tag and latest
gcloud docker -- push "$image:$build"
gcloud docker -- push "$image:latest"

# finally, deploy and check status
kubectl set image deployments app "$image:$build" --record
kubectl rollout status deployments app
```

Pretty simple, but works really well.

If you would like to go fancy, Google Cloud Launcher has a [Spinnaker](https://console.cloud.google.com/launcher/details/click-to-deploy-images/spinnaker) click-to-deploy that you can use to deploy to Kubernetes. I didn't get to test it yet, but it should work.

## In conclusion

Creating clusters on GKE is really easy. On AWS, I've been using Kops for that, while it is a great tool and kind of easy to use, GKE can be even easier. 

Kudos to GKE.

Setting up services, however, feels easier on AWS. GKE has some particularities that required me to read docs, search though stack overflow, asks questions on their support and etc. 

Maybe it's because I'm already used to AWS, maybe it's because somethings are overcomplicated (like databases connections).

At the end, most of the time I spent figuring out GKE was because I was trying to use the same approach I use in AWS for several things, and, seeing the docs, sometimes I didn't believe the docs were updated, thinking that "there must be a newer better way of doing this". There is not.

Well, these are my experiences so far with GKE. Hope it helps you somehow!

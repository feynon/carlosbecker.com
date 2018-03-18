---
date: 2017-12-07T00:00:00Z
title: 'Kubernetes: Traefik ingresses with external-dns'
draft: true
---

In this post I'll try to show you how to use Traefik as an Ingress
controller on Kubernetes, having external-dns handling the DNS
records for you.

This is actually pretty simple, I'm only writing this because I spend some
time trying to figure it out - and maybe I can help someone spend less time
on the same thing :)

So, let's get started.

First things first, deploy [Traefik](http://traefik.io) (it is actually
written as `Træfik`). I'm using [Helm](https://github.com/kubernetes/helm)
to do that for me, but you can do it as you please.

For what is worth, my `values.yaml` looks like:

```yaml
replicas: 2
dashboard:
  enabled: true
  domain: traefik.foo.bar
  statistics:
    recentErrors: 100
```

You can take a look at the [Chart README](https://github.com/kubernetes/charts/tree/master/stable/traefik)
for more options.

Installing it is as simple as:

```sh
helm install stable/traefik --name traefik --namespace kube-system --values ./values.yaml
```

You can then create the `traefik.foo.bar` DNS record poiting to the Traefik
service LoadBalancer. You can get the address with:

```sh
kubectl describe svc traefik-traefik --namespace kube-system | grep Ingress | awk '{print $3}'
```

[external-dns](https://github.com/kubernetes-incubator/external-dns) is a
Kubernetes incubator project that can "Configure external DNS servers
(AWS Route53, Google CloudDNS and others) for Kubernetes Ingresses and
Services".

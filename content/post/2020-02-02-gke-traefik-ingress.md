---
title: "GKE using Traefik as the ingress controller"
date: 2020-02-02
draft: false
slug: gke-traefik-ingress
city: Joinville
---

I recently fall into a trap using Traefik as the ingress controller in one cluster. I decided to write about it so maybe it helps someone.

<!--more-->

## Context

We got the architecture like this:

```shell
Cloudflare -> Traefik LoadBalancer -> Traefik Pods -> App Pods
```

Traefik was running in a node pool that had only non-preemptible machines. The app pods were running in a mix of preemptible and not preemtible.

We also have other things running, so, in reality, we had a lot of nodes running neither the app nor Traefik pods, just `kube-system` stuff, monitoring, batch jobs and so on.

## What happened

Clients would sometimes report that Cloudflare returned a HTTP 520.

520 is not specified by any spec, but is used by Cloudflare to indicate "unknown connection issues between Cloudflare and the origin web server".

We though maybe something was up on our app, or maybe even on Traefik instances. We looked into our logs and monitoring, but couldn't find anything relevant.

## Despair

I was pretty much loosing my hair and was sure something was going on with Traefik and that for some reason I couldn't find what it was.

Because of that, I decided to try to create a GCE ingress - which translates to GCE Load balancer directly.

Then it hit me... after days of looking into things that had nothing to do with the problem, as most of the times, it was something obvious (once you get it).

## Traffic inside a Kubernetes cluster

Basically, any node can route to a pod running in any other node. 

That is handled by `kube-proxy` and `iptables`:

![Original: https://cloud.google.com/kubernetes-engine/docs/concepts/network-overview](/public/images/gke-traefik-ingress/1b7f6f92-a65a-4280-bf6b-66688f6ae396.png)

That, plus the following screenshot made clicked me:

![One of our load balancers showing all nodes as "green", even the ones with no Traefik pods running...](/public/images/gke-traefik-ingress/ae0eb9a1-30a7-429b-978a-fdca08ac0b6f.png)

Even though we had Traefik running only on non-preemptible nodes, all nodes, including the preemtible ones, were receiving traffic from the load balancer.

From there it is straightforward to figure out what can happen:

- Client access `myapp.foo`
- LB randomly routes to a preemptible node
- Request is still going
- Node is preempted
- Request is terminated mid way (we have some slow requests for several reasons)
- Cloudflare doesn't know what the hell happened, and they have a status code for it: 520

## Fixing things

So, what we need, ideally, is to have only the nodes that are actually running Traefik to be "green" on the load balancer, all others must fail the healthchecks.

After digging a bit, I found out about a setting called `externalTrafficPolicy`.

The default for it is `Cluster`, which makes the traffic being NATed. The main reason to change this is, usually, because you want to use  [traefik.ingress.kubernetes.io/whitelist-source-range](http://traefik.ingress.kubernetes.io/whitelist-source-range) to filter traffic, but, since by default the requests get NATed, you can't, because you won't get the correct client IP.

Setting `externalTrafficPolicy` to `Local` has two effects:

1. traffic is no longer NATed - which means  [traefik.ingress.kubernetes.io/whitelist-source-range](http://traefik.ingress.kubernetes.io/whitelist-source-range) now works;
2. oh, and also: **only the nodes that have Traefik running in then will pass the load balancer healthchecks.**

Which, yeah, fixes our issue.

Literally one line of code...

After hours and hours investigating...

🤷‍♂️

## Conclusions

It seems to me `Local` should be the default option instead of `Cluster`, as it seems to be what you will want most of the times - maybe this is an old default they don't want to change, didn't get into it.

Changing from `Cluster` to `Local` in a live cluster will cause some requests to fail. On some clusters it was only for a few seconds, in another one it took almost 1 minute to stabilize. Beware.
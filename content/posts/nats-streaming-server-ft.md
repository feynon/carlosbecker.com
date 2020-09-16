---
title: "High availability with nats-streaming-server (fault-tolerance)"
date: 2019-07-25
draft: false
slug: nats-streaming-server-ft
city: Marechal Cândido Rondon
toc: true
tags: [nats]
---

I wanted to set up a fault tolerant [nats-streaming-server](https://github.com/nats-io/nats-streaming-server), but couldn't find a "quick" guide on how to do it - so here we are.

<!--more-->

I would also recommend you to read a [previous post]({{< ref "nats-streaming-server-cluster.md" >}}) I wrote about how to do it using the clustering method.

## Why not Clustering

Clustering is simpler (in the sense of having less moving parts) compared to the Fault Tolerance strategy, but was too slow for our needs.

I don't have the exact numbers anymore because I have been delaying writing this for a few months now, but I can say for sure that single node was **several times** faster than the cluster, even if we used
faster hardware on the cluster machines.

It was somewhat curious that the cluster, while in stand-by, would still generate as much as 1 million msg/second - with no one using it. My guess is that it was happening due to the nature of [RAFT](https://raft.github.io/) -
which can be very chatty.

That alone would be OK, but once we started using it for real, the performance degraded - a lot. We would get 500k msg/sec peaks, but the majority of the messages seemed to be related to RAFT, in reality we were getting through something around 1k msg/sec.

Some of it was not fault of RAFT itself though.

## Big Messages

We had some really big messages, which don't help NATS at all. We were getting constant leader re-elections because of that, which make performance worse.

Our `max_payload` was set to 120mb/s, which is indeed too big. For reference, Amazon SQS max payload is 256kb.

If you only take one thing from this article, take this: message sizes should, ideally, max at a few kb - 256kb is a good/proven starting point.

We needed to work our backend to lower that. After some work, things got a little better, but not good enough. Meanwhile our platform was slow, and changing the message sizes and etc would take too much time due to the way our platform was working at the time.

We decided to try the fault tolerant mode after some feedback from the NATS Slack Community.

## Fault Tolerance

Another solution for deploying a highly available [NATS Streaming Server](https://github.com/nats-io/nats-streaming-server) is to use its Fault Tolerant mode - which requires a shared filesystem, for example, NFS.  On the [previous post]({{< ref "nats-streaming-server-cluster.md" >}}) I added that funny joke about NFS being the culmination of 3 lies and etc... that didn't age well. 😂

We didn't want to manage our own NFS, nor have one more database to look after, so we decided to rely on [Google Cloud Filestore](https://cloud.google.com/filestore/) - a NFS as a service offered by Google.

NATS Fault Tolerant mode works by having an active server and one or more standby servers.
The filestore is shared between all instances, so there is no replication in place on [NATS](https://github.com/nats-io/nats-server) level.

If the active server dies for any reason, one of the others takes the leadership. You can pass all [NATS](https://github.com/nats-io/nats-server) addresses to the client and it should take care of reconnecting, connecting to the active server and etc.

Configuration is also simpler, we just need to add the `routes` and `ft_group` settings:

```
; a.conf
port: 4221
cluster {
  listen: 0.0.0.0:6221
  routes: ["nats://localhost:6222"] ; node b addr
}

streaming {
  id: test-cluster
  store: file
  dir: /opt/nats_state
  ft_group: test
}
```
```
; b.conf
port: 4222
cluster {
  listen: 0.0.0.0:6222
  routes: ["nats://localhost:6221"] ; node a addr
}

streaming {
  id: test-cluster
  store: file
  dir: /opt/nats_state
  ft_group: test
}
```

Note that each config listens on different ports:

- `a`: `4221` and `6221`
- `b`: `4222` and `6222`

And then you can just start each node pointing to the specific config file:

```sh
$ ./nats-streaming-server -c a.conf
```
```sh
$ ./nats-streaming-server -c b.conf
```

We can then run our client from the [previous post]({{< ref "nats-streaming-server-cluster.md" >}}) to test it:

```sh
$ go run main.go nats://localhost:4221 nats://localhost:4222
```

Notice that on fault tolerant mode we don't need quorum. Any amount of nodes should be OK - in our example we are using only 2.

## Moneys

You may be wondering about the costs for this. Let's compare our cluster solution to our fault tolerant solution:

### cluster

3 instance with 500GB SSD each: `3 * (116.71 + (500*0.204))`

**total**: $ 656.13 / month*

### fault tolerant

2 instance with 60GB standard each + 1 STANDARD Filestore managed NFS volume with 1TB:
`(2 * (116.71 + (60*0.048))) + (1024*0.22)`

**total**: $ 464.46 / month*

> *: prices based on us-west-2
## Conclusion

The fault tolerant mode give us nearly the same performance as the single node approach, fault tolerance and costs almost $ 200 / month less than the cluster strategy.

The problem with this solution is that you have to trust on NFS - which Google advertise as 99.9% available.

It's also worth to mention that Google's Cloud Filestore caps performance at 100mb/s for reads and 100mb/s writes on our volume.
That can go up to as much as 1.2gb/s for reads and 350mb/s for writes if we upgrade to a PREMIUM volume - which of course costs way more.

As on the other post, those configs, clients and etc will be on my [nats playground](https://github.com/caarlos0/nats-streaming-server-cluster) on github. The readme there should guide you on everything.

Hope this helps somehow.

Cheers!

---

Thanks to [Bruno](https://github.com/brunocvcunha), [Alex](https://github.com/amalucelli) and the NATS Slack Community for helping figuring all this out.

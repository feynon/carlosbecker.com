---
title: "High availability with nats-streaming-server (clustering)"
date: 2019-05-16
draft: false
slug: nats-streaming-server-cluster
city: Joinville
toc: true
tags: [nats]
---

I wanted to set up a high available [nats-streaming-server](https://github.com/nats-io/nats-streaming-server) cluster,
but couldn't find a "quick" guide on how to do it.

<!--more-->

In this post I'll try to write something that would have helped me earlier.

First things first, we have 2 kinds of HA setups for [nats-streaming-server](https://github.com/nats-io/nats-streaming-server):

1. Fault Tolerance
2. Clustering

Let's dig deeper on them.

## 1. Fault Tolerance

In this mode, you setup a *active node* and one or more *stand-by nodes*.
They can share the state through NFS, for example.

> NFS is the culmination of three lies:
> 
> 1. Network
> 2. File
> 3. System
> 
> — electrified filth (@sadisticsystems) April 29, 2019

I don't like NFS, so I didn't like this option either, although the
performance may be better than the clustering option.

## 2. Clustering

Clustering uses [RAFT](https://raft.github.io/) for leader election and has no shared resources. A
write in one node will be replicated to other nodes.

This seemed like the best option for my case, so I'll go with that for now on.

[nats-streaming-server](https://github.com/nats-io/nats-streaming-server) embeds a [NATS](https://github.com/nats-io/nats-server) server too, and to cluster
[nats-streaming-server](https://github.com/nats-io/nats-streaming-server) we need to cluster [NATS](https://github.com/nats-io/nats-server) as well.

We have two alternatives here, either setup a separated [NATS](https://github.com/nats-io/nats-server) cluster
or cluster the one already embedded in [nats-streaming-server](https://github.com/nats-io/nats-streaming-server).

I choose to use the embed one.

## A simple example

Let's start with a single [nats-streaming-server](https://github.com/nats-io/nats-streaming-server) node and an example
client:

```go
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/nats-io/stan"
)

func main() {
	sc, err := stan.Connect(
		"test-cluster",
		"client-1",
		stan.Pings(1, 3),
		stan.NatsURL(strings.Join(os.Args[1:], ",")),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer sc.Close()

	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Print(".")
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer sub.Unsubscribe()

	for {
		if err := sc.Publish("foo", []byte("msg")); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(time.Millisecond * 100)
	}
}
```

It basically connects to the [nats-streaming-server](https://github.com/nats-io/nats-streaming-server) URL's passed to it,
subscribeds to a topic and keeps sending messages. A `.` is print on the
screen for each message received.

So, now we can just start both:

```sh
$ ./nats-streaming-server
```
```sh
$ go run main.go localhost:4222
```

You should see a lot of `.` being print on the screen, meaning that it is
working. If you kill the [nats-streaming-server](https://github.com/nats-io/nats-streaming-server), you'll notice that the
client will die too.

## Clustering

So, now let's stop both client and server, and start a
[nats-streaming-server](https://github.com/nats-io/nats-streaming-server) cluster.

Create 3 config files as follows:

```
; a.conf
port: 4221
cluster {
  listen: 0.0.0.0:6221
  routes: [
    "nats-route://localhost:6222",
    "nats-route://localhost:6223",
  ]
}

streaming {
  id: test
  store: file
  dir: storea
  cluster {
    node_id: "a"
    peers: ["b", "c"]
  }
}
```
```
; b.conf
port: 4222
cluster {
  listen: 0.0.0.0:6222
  routes: [
    "nats-route://localhost:6221",
    "nats-route://localhost:6223",
  ]
}

streaming {
  id: test
  store: file
  dir: storeb
  cluster {
    node_id: "b"
    peers: ["a", "c"]
  }
}
```
```
; c.conf
port: 4223
cluster {
  listen: 0.0.0.0:6223
  routes: [
    "nats-route://localhost:6221",
    "nats-route://localhost:6222",
  ]
}

streaming {
  id: test
  store: file
  dir: storec
  cluster {
    node_id: "c"
    peers: ["a", "b"]
  }
}
```

Note that each config listens on different ports:

- `a`: `4221` and `6221`
- `b`: `4222` and `6222`
- `c`: `4223` and `6223`

Also note that in each config's `cluster` we setup the routes to the other 2
instances. This cluster config is the actual [NATS](https://github.com/nats-io/nats-server) cluster.

The `streaming.cluster` config is the actual [nats-streaming-server](https://github.com/nats-io/nats-streaming-server) cluster
configuration, and only IDs each node and add the other 2 as peers.

Since we are running all nodes on the same machine, notice that the
`streaming.dir` option is different in each config.

Once that's done, we can start the 3 servers:

```sh
$ ./nats-streaming-server -c a.conf
```
```sh
$ ./nats-streaming-server -c b.conf
```
```sh
$ ./nats-streaming-server -c c.conf
```

Once all of them are up, you should see logs like the following on each of them:

```
[11361] 2019/05/16 14:03:55.994864 [INF] ::1:52022 - rid:8 - Route connection created
[11361] 2019/05/16 14:03:55.997790 [INF] ::1:52023 - rid:9 - Route connection created
```

Now, we can connect start our client again:

```sh
$ go run main.go nats://localhost:4221 nats://localhost:4222 nats://localhost:4223
```

> Notice that I'm passing the URL for all the 3 servers.

Now, play around killing some servers. You'll notice that sometimes nothing happens to the client, and other times the client also dies.

You may better handle that using a `ConnectionLostHandler`. You may
check [their repository README](https://github.com/nats-io/stan.go) for further information about this.

---

I tried to keep it as simple as possible, hope it is helpful! 🙂

If you want, you can also try this with `docker-compose`. I put all
the code (including the client) in a [GitHub Repository](https://github.com/caarlos0/nats-streaming-server-cluster).

Let me know in the comments if you have any questions!

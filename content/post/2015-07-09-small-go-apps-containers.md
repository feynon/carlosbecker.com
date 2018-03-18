---
date: 2015-07-09T00:00:00Z
slug: small-go-apps-containers
title: Small Go Apps Containers
---

Or: how to ship your app in a <20Mb container.

Well, as you may know, there is a good amount of people now building
microservices in Go and deploying them as Docker containers.

I do not yet have a lot of experience with Go and Docker, but I'll try to
share what I learned while building and shipping an internal tool, here, at
[ContaAzul](http://contaazul.com).

## The Go code example

I will assume that you know at least a little bit of Go, and, for the sake
of simplicity and brevity, I'll just use a very basic example from the
[Go wiki](https://golang.org/doc/articles/wiki/):

```go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
```

If we compile this file, the binary will have 5.5MB:

```console
$ go build
$ du -h example
5.5M	example
```

Now, let's _dockerize_ this thing!

## Let's Go with the common

What I usually see is people starting with the
[official Golang image](https://registry.hub.docker.com/_/golang/), but, we
are talking about small containers here, so, let's be hipsters and go with
some Alpine-based image.

For those who don't know, Alpine is a very minimalistic Linux distribution.
It became "famous" (at least for me) because of its adoption in the Docker
community, mainly, because its image is only 5MB in size:

[![Alpine](https://badge.imagelayers.io/alpine:latest.svg)](https://imagelayers.io/?images=alpine:latest 'Get your own badge on imagelayers.io')

We could probably use
[kiasaki/alpine-golang](https://github.com/kiasaki/docker-alpine-golang),
but, seems like it still bundling Go 1.3, and, because I wanted to use Go 1.4+,
I created a new image, based on kiasaki's image, but with Go 1.4.2.
[[source](https://github.com/caarlos0/docker-alpine-go)]

Using my image, the `Dockerfile` may look like this:

```dockerfile
FROM caarlos0/alpine-go
WORKDIR /gopath/src/app
ADD . /gopath/src/app/
RUN go get -v app
ENTRYPOINT ["/gopath/bin/app"]
```

Now let's build it:

```console
$ docker build -t caarlos0/example-small .
# ...
$ docker images
REPOSITORY                 TAG                 IMAGE ID            CREATED             VIRTUAL SIZE
caarlos0/example-small     latest              b300ec38662c        3 minutes ago       220.5 MB
caarlos0/alpine-go         latest              1cd569752ed7        3 days ago          214.7 MB
alpine                     3.2                 31f630c65071        3 weeks ago         5.254 MB
```

Yes, from ~5MB (of the alpine image) to **220.5MB**!!!!

[![Example with Go language bundled](https://badge.imagelayers.io/caarlos0/example-small.svg)](https://imagelayers.io/?images=caarlos0/example-small 'Get your own badge on imagelayers.io')

Looking at the layers, we can see that the problem is that,
well, the entire Go language and its tool, Git, Mercurial and a lot of stuff
are bundled together with the image. We only need them to compile the program,
but they will not be used anymore after that, and we can safely remove them.

I guess we could just add another `RUN` statement removing all those stuff,
right? **WRONG!**

## We have to Go deeper

Each instruction in the `Dockerfile` ends up being a
new layer, so, removing stuff from a previous layer in the current one
will not have the desired effect in the final image size.

To fix that, instead of inheriting from `caarlos0/alpine-go`, we'll have to
inherit directly from `alpine`, install what we need to compile the app,
compile the app, and, finally, remove what we don't need anymore - all this
in **one single step**.

So, we might end up with something like this:

```dockerfile
FROM alpine:3.2

ENV GOROOT=/usr/lib/go \
    GOPATH=/gopath \
    GOBIN=/gopath/bin \
    PATH=$PATH:$GOROOT/bin:$GOPATH/bin

WORKDIR /gopath/src/app
ADD . /gopath/src/app

RUN apk add -U git go && \
  go get -v app && \
  apk del git go && \
  rm -rf /gopath/pkg && \
  rm -rf /gopath/src && \
  rm -rf /var/cache/apk/*

ENTRYPOINT ["/gopath/bin/app"]
```

And, vòilá:

```console
$ docker images
REPOSITORY                 TAG                 IMAGE ID            CREATED             VIRTUAL SIZE
example-smaller            latest              6c4f2066e02e        9 seconds ago       11.43 MB
alpine                     3.2                 31f630c65071        3 weeks ago         5.254 MB
```

Just **11.43MB**:

[![Example without Go Language bundled](https://badge.imagelayers.io/caarlos0/example-smaller:latest.svg)](https://imagelayers.io/?images=caarlos0/example-smaller:latest 'Get your own badge on imagelayers.io')

I think that's literally the smaller you can get.

## Why would anyone Go with some alpine-go image then?

It works very well as a dev environment. It also don't actually hurt to have
all those stuff in your Docker image if you don't care about the image size.

But, you know, the smaller the container, the faster the push and the pull.
Besides, more free HD space is always a good thing.

## Show me the numbers

##### Empty official Golang image:
[![Empty official Golang image](https://badge.imagelayers.io/golang:latest.svg)](https://imagelayers.io/?images=golang:latest 'Get your own badge on imagelayers.io')

##### Official Alpine image:
[![Official Alpine image](https://badge.imagelayers.io/alpine.svg)](https://imagelayers.io/?images=alpine 'Get your own badge on imagelayers.io')

##### My alpine-go image:
[![My alpine-go image](https://badge.imagelayers.io/caarlos0/alpine-go:latest.svg)](https://imagelayers.io/?images=caarlos0/alpine-go:latest 'Get your own badge on imagelayers.io')

##### example-small image:
[![example-small image](https://badge.imagelayers.io/caarlos0/example-small.svg)](https://imagelayers.io/?images=caarlos0/example-small 'Get your own badge on imagelayers.io')

##### example-smaller image:
[![example-smaller image](https://badge.imagelayers.io/caarlos0/example-smaller:latest.svg)](https://imagelayers.io/?images=caarlos0/example-smaller:latest 'Get your own badge on imagelayers.io')

## Conclusion

You might have noticed that I'm removing the entire `gopath/src` folder in the
`Dockerfile`. This is not intended to free the space used by the app source
itself, but to free us the space used by the sources of the dependencies we
got with `go get -v app`.

I also remove the entire `gopath/pkg` folder for the same reason.

This two actions might have no effect in this particular example, but I decided
to leave them there as an idea of what you can do in the so-called
"real world apps".

You can also, of course, compile your app outside the container and just
`ADD` the binary to it. To do that you need to pay attention and compile it
for the same `ARCH` and `OS`, like, and, of course, have the right Go
installed:

```console
$ GOARCH=i386 GOOS=linux go build
```

The downside of this approach is that Docker Hub will not be able to
automatically build this, if you don't care about that (pushing containers by
hand), go for it.

As a final **PROTIP™**, pay attention to what you `ADD` to your image. You
might want to remove the `.git` folder, binaries, unused files and everything
you don't need. You can do this using a `.dockerignore` file.
[Reference](https://docs.docker.com/engine/reference/builder/#/dockerignore-file).

By the way, the
[app code](https://github.com/caarlos0/small-go-app-container-example)
is available on Github. Go for it!

That's all folks. :beers:

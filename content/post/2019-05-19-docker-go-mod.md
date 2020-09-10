---
title: "Faster Docker builds using go modules"
date: 2019-05-19
draft: false
slug: docker-go-mod
city: Joinville
---

Quick tip to improve the docker build speed using go modules.

<!--more-->

Normally, I would do something like this:

```docker
FROM golang as builder
ENV GO111MODULE=on
WORKDIR /code
ADD . .
RUN go build -o /app main.go

FROM gcr.io/distroless/base
EXPOSE 8080
WORKDIR /
COPY --from=builder /app /usr/bin/app
ENTRYPOINT ["/usr/bin/app"]
```

The problem with this approach is that, if I change any `.go` file and rebuild, it will download the dependencies again - which takes some time.

Taking into account that dependencies do not change very often, we can add just two lines and improve the build performance **a lot**:

```docker
FROM golang as builder
ENV GO111MODULE=on
WORKDIR /code

# add these two lines
ADD go.mod go.sum /code/
RUN go mod download

ADD . .
RUN go build -o /app main.go

FROM gcr.io/distroless/base
EXPOSE 8080
WORKDIR /
COPY --from=builder /app /usr/bin/app
ENTRYPOINT ["/usr/bin/app"]
```

This way, we'll only download the dependencies again if `go.mod` or `go.sum` changed.

This can save a lot of time on local development, even more if on a slow network connection.

That's it for today, hope it helps! 🤟
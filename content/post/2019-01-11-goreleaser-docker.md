---
title: "GoReleaser Docker support"
date: 2019-01-11T14:56:42-02:00
slug: "goreleaser-docker"
city: Joinville
tags:
- goreleaser
- golang
- docker
- ci
---

The next [GoReleaser][] version will have a more flexible [Docker][] configuration
format. In this post we'll explore it a bit.

<!--more-->

When I first added [Docker][] support on [GoReleaser][], I did a couple of
assumptions - and most of them were wrong.

At that time, the config file looked like this:

```yaml
dockers:
-
  binary: foo
	goos: linux
	goarch: amd64
	goarm: ''
	image: bar/foo
  dockerfile: Dockerfile
```

It didn't allow multiple tags, registries or anything like that... at least
not without building the image several times.

Of course, some users needed those things, and the it went through a few
interactions:

- https://github.com/goreleaser/goreleaser/pull/355
- https://github.com/goreleaser/goreleaser/pull/359
- https://github.com/goreleaser/goreleaser/pull/367
- https://github.com/goreleaser/goreleaser/pull/370
- https://github.com/goreleaser/goreleaser/pull/435
- https://github.com/goreleaser/goreleaser/pull/512
- https://github.com/goreleaser/goreleaser/pull/840
- https://github.com/goreleaser/goreleaser/pull/819
- https://github.com/goreleaser/goreleaser/pull/919

And now, I think we finally arrived at something closer to the what will be on
1.0:

```yaml
dockers:
  -
    goos: linux
    goarch: amd64
    goarm: ''
    binaries:
    - mybinary
    - otherbinary
    image_templates:
    - "myuser/myimage:latest"
    - "myuser/myimage:{{ .Tag }}"
    skip_push: false
    dockerfile: Dockerfile
    build_flag_templates:
    - "--label=org.label-schema.schema-version=1.0"
    - "--label=org.label-schema.version={{.Version}}"
    - "--label=org.label-schema.name={{.ProjectName}}"
    - "--build-arg=FOO={{.ENV.Bar}}"
    extra_files:
    - config.yml
```

This allows several things that were not possible before:

- more than one binary in the same image
- several images names, registries and tags on a single build
- extra files
- build arguments
- templating on image names
- control wether to push the images or not

Lot's of things you can do with a rather simple config format!

The only feature that is still missing is
[multiple platforms support](https://github.com/goreleaser/goreleaser/issues/530),
which should be added before 1.0.

A thing that I'm tried hard to do from the start is to keep it possible to
build docker images without running goreleaser and without changing any
configs. This was true on the first format is still true as today!

For example, imagine you have a simple project, more or less like this:

```
.
├── Dockerfile
├── config
│   └── foo.conf
├── .goreleaser.yml
└── main.go
```

Your `Dockerfile` may look like:

```Dockerfile
FROM scratch
ADD foo /usr/bin/foo
ADD config/foo.conf /etc/foo.conf
```

So, as a developer, you would `go build` your project and then
simply `docker build`:

```sh
go build -o foo
docker build -t caarlos0/goreleaser-docker-test .
```

And that would work, as expected.

You could then automate that using [GoReleaser][] by creating a
`.goreleaser.yaml` config file that looks like this:

```yaml
project_name: foo
builds:
- env:
  - CGO_ENABLED=0
dockers:
  -
    binaries:
    - foo
    image_templates:
    - "caarlos0/goreleaser-docker-test:latest"
    - "caarlos0/goreleaser-docker-test:{{ .Tag }}"
    extra_files:
    - config/foo.conf
```

And then, when you want to release, you can just run `goreleaser` and it will
do everything for you, using the same `Dockerfile`.

This allows for quick interactions while changing dockerfiles or to test
other things. Simple, yet, powerful!

> Note that it would also work with multiple binaries, given that you do the
> required changes on the config files.

So, that's it. Few words just to point out how things are evolving.
Looking forward to make releasing easier, better and more reliable for everyone!

Cheers!

[GoReleaser]: https://goreleaser.com
[Docker]: https://docker.io

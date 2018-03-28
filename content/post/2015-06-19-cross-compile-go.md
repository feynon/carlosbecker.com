---
date: 2015-06-19T00:00:00Z
slug: cross-compile-go
title: Cross-compiling Go
tags:
- golang
---

`go build` generates a binary for the platform you run it in. So, if I
build [antibody][antibody] in a Linux machine - which uses Mach-O, it will
not work in OS X - which uses ELF.

I wanted to distribute [antibody][antibody] at least for Linux and OS X, so
I went out searching for how to do this in a not so complicated way...

I found an [article from 2012](http://solovyov.net/en/2012/cross-compiling-go/),
demonstrating how to compile Go itself for multiple platforms, so you can use
the specific go to build the binary for the specific platform, which is a lot
of work to do.

Then, I discovered [gox](https://github.com/mitchellh/gox), which compiles Go
for multiple platforms and let you easily compile your application for
them. It supports these platforms:

- darwin/386
- darwin/amd64
- linux/386
- linux/amd64
- linux/arm
- freebsd/386
- freebsd/amd64
- openbsd/386
- openbsd/amd64
- windows/386
- windows/amd64
- freebsd/arm
- netbsd/386
- netbsd/amd64
- netbsd/arm
- plan9/386

And it's damn easy to use:

```console
$ go get github.com/mitchellh/gox
$ gox -build-toolchain
$ gox
```

This will install `gox`, build it's toolchain (all the Go versions) and build
your application for them. No further work required!

To automate my work a little, I also created a script within
[antibody][antibody], which will build, test, tag and release the binaries
in Github using `gox` and [github-release][github-release]. You can take a
look at the [source][releaser] if you wish.

Hope this helps somebody!

Cheers! :beers:

[antibody]: https://github.com/caarlos0/antibody
[github-release]: http://github.com/aktau/github-release
[releaser]: https://github.com/caarlos0/go-releaser/blob/81e3ceb54d321676afcba7d23e9a02a5682ed0f5/release

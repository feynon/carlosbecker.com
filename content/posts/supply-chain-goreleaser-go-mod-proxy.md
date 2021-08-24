---
title: "Supply chain integrity with GoReleaser using Go mod proxy"
date: 2021-08-23
draft: false
slug: supply-chain-goreleaser-go-mod-proxy
city: Cascavel
toc: true
tags: [goreleaser, goreleaser-pro, golang, security]
---

Since the infamous [SolarWinds attack](https://www.csoonline.com/article/3191947/supply-chain-attacks-show-why-you-should-be-wary-of-third-party-providers.html), supply chain integrity is something a lot of people are discussing and working on. 

In this post we'll see how we can verify a binary built with Go is indeed what it says it is.

## Building from Go mod proxy

### Using go install

The easiest way of doing that is using `go install`:

```sh
$ go install github.com/caarlos0/svu@v1.7.0
```

And then we can verify with `go version -m`:

```sh
$ go version -m $(which svu)
/go/bin/svu: go1.17
	path	github.com/caarlos0/svu
	mod	github.com/caarlos0/svu	v1.7.0	h1:Aqk2q+qPGRoigQWgWkMeFlsjM0cZin7QL4oCPL++xUI=
	dep	github.com/Masterminds/semver	v1.5.0	h1:H65muMkzWKEuNDnfl9d70GUjFniHKHRbFPGBuZ3QEww=
	dep	github.com/alecthomas/kingpin	v2.2.6+incompatible	h1:5svnBTFgJjZvGKyYBtMB0+m5wvrbUHiqye8wRJMlnYI=
	dep	github.com/alecthomas/template	v0.0.0-20190718012654-fb15b899a751	h1:JYp7IbQjafoB+tBA3gMyHYHrpOtNuDiK/uB5uXxq5wM=
	dep	github.com/alecthomas/units	v0.0.0-20210208195552-ff826a37aa15	h1:AUNCr9CiJuwrRYS3XieqF+Z9B9gNxo/eANAJCF2eiN4=
	dep	github.com/gobwas/glob	v0.2.3	h1:A4xDbljILXROh+kObIiy5kIaPYD8e96x1tgBhUI5J+Y=
```

In the output, we can see this line:

```sh
mod	github.com/caarlos0/svu	v1.7.0	h1:Aqk2q+qPGRoigQWgWkMeFlsjM0cZin7QL4oCPL++xUI=
```

Which is the built module, its version and its hash.

We can then verify if its the correct hash doing something like this:

```sh
$ mkdir /tmp/test
$ cd /tmp/test
$ go mod init test
$ go get -d github.com/caarlos0/svu@v1.7.0
$ grep svu go.sum
github.com/caarlos0/svu v1.7.0 h1:Aqk2q+qPGRoigQWgWkMeFlsjM0cZin7QL4oCPL++xUI=
github.com/caarlos0/svu v1.7.0/go.mod h1:x4iUUt1lKyYx8o8uKfiaRZjJSt/PHaYTSUgR13e6Zy0=
```

And we can see that the hashes match, so it is verified.

### Creating a go.mod and main.go

This is the way [GoReleaser](https://goreleaser.com) does it, its a bit more convoluted than using `go install`, with the main difference being that the binary is not installed to the `GOPATH`.

```sh
$ mkdir proxy
$ cd proxy/
$ go mod init proxy
$ go get -d github.com/caarlos0/svu@v1.7.0
$ echo '// +build main
package main

import _ "github.com/caarlos0/svu"
' > main.go
$ go build github.com/caarlos0/svu
$ go version -m svu
svu: go1.17
	path	github.com/caarlos0/svu
	mod	github.com/caarlos0/svu	v1.7.0	h1:Aqk2q+qPGRoigQWgWkMeFlsjM0cZin7QL4oCPL++xUI=
	dep	github.com/Masterminds/semver	v1.5.0	h1:H65muMkzWKEuNDnfl9d70GUjFniHKHRbFPGBuZ3QEww=
	dep	github.com/alecthomas/kingpin	v2.2.6+incompatible	h1:5svnBTFgJjZvGKyYBtMB0+m5wvrbUHiqye8wRJMlnYI=
	dep	github.com/alecthomas/template	v0.0.0-20190718012654-fb15b899a751	h1:JYp7IbQjafoB+tBA3gMyHYHrpOtNuDiK/uB5uXxq5wM=
	dep	github.com/alecthomas/units	v0.0.0-20210208195552-ff826a37aa15	h1:AUNCr9CiJuwrRYS3XieqF+Z9B9gNxo/eANAJCF2eiN4=
	dep	github.com/gobwas/glob	v0.2.3	h1:A4xDbljILXROh+kObIiy5kIaPYD8e96x1tgBhUI5J+Y=
```

We can see here that the hashes still match, success!

## Using it within GoReleaser

To proxy the module when building with GoReleaser, you only need 2 lines in your `goreleaser.yml`:

```yaml
gomod:
  proxy: true
```

And that's it! GoReleaser hides all the complexity from you and figure most of the needed parts out by itself!

## Conclusion

Assuming your tool is OpenSource, it doesn't matter where it is built, as long as its built from the Go Mod proxy, it is verifiable.

You can take this further by signing the binaries and making the result binary more reproducible by trimming the path and using a fixed timestamp (see an [example](https://github.com/caarlos0/goreleaserfiles/blob/main/build.yml)).

GoReleaser OSS itself is verifiable this way, you can try it with: 

```sh
go version -m $(which goreleaser)
```

GoReleaser Pro in the other hand is not, as it is not OpenSource. 

Still, since GoReleaser uses `go build` underneath, as long as you have `gomod.proxy` set to true, it shouldn't matter much - your builds are still verifiable. If GoReleaser Pro gets somehow corrupted and starts doing *"something funny"* with your builds, you can easily verify them using this technique.

Another side effect of using the `gomod.proxy` feature: your module gets hashed right away, so if someone deletes the tag and recreates it, the previous tag would still be used to compile.

Hope this is somewhat helpful/informative to you. 

I'll see you in the next one!

## Edit

After some more discussion in [this issue](https://github.com/goreleaser/goreleaser/issues/2391), Marcos Nils pointed out a way to hack this using `go mod vendor`:

```sh
$ go mod vendor
# change some dependency code
$ go build -mod vendor github.com/caarlos0/svu
$ go version -m svu
# same output as before
```

So, while it prevents the attacker from changing the target code, they could still change de dependencies, which **renders the idea of this post pretty much useless**.

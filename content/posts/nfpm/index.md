---
title: "Creating debs and rpms with Go"
date: 2018-02-23
draft: false
slug: nfpm
city: Joinville
toc: true
tags: [goreleaser, linux, ci-cd, golang]
---

I've been working on [GoReleaser](https://goreleaser.com/) for more than a year now, and one of the things that was bothering me the most was [fpm](https://github.com/jordansissel/fpm).

Not that fpm is bad or anything like that, is just that it can be unstable: it uses the system `tar` (GNU tar is different from BSD tar) and its written in ruby, which requires the user to deal with ruby versions, gems and all that.

I wanted something simpler, predictable and that could be used as a lib and as a binary. Couldn't find anything like that, so I went ahead and tried to implement it myself in Go, using the minimum external dependencies.

Thus, [nfpm](https://github.com/goreleaser/nfpm) was born.

NFPM stands for *Not FPM*. It's not a very good name, and I know [I suck at naming things](https://twitter.com/bepsays/status/966313876408193025).

Anyway, its usage as a lib is kind of simple and it is easy to extend:

```go
package main

import (
	"io/ioutil"
	"log"

	"github.com/goreleaser/nfpm"
	_ "github.com/goreleaser/nfpm/deb"
)

func main() {
	pkg, err := nfpm.Get("deb")
	if err != nil {
		log.Fatalln(err)
	}
	if err := pkg.Package(nfpm.Info{
		Name:        "my-pkg",
		Arch:        "amd64",
		Platform:    "linux",
		Version:     "1.2.3",
		Section:     "",
		Priority:    "",
		Replaces:    nil,
		Provides:    nil,
		Depends:     nil,
		Recommends:  nil,
		Suggests:    nil,
		Conflicts:   nil,
		Maintainer:  "",
		Description: "",
		Vendor:      "",
		Homepage:    "",
		License:     "",
		Bindir:      "",
		Files:       map[string]string{
			// "./local/path": "/path/inside/pkg",
		},
		ConfigFiles: map[string]string{
			// same thing as files, but they are treated different by deb/rpm themselves
		},
	}, ioutil.Discard); err != nil {
		log.Fatalln(err)
	}
}
```

As you can see, there are quite a few options, not all of them are required, though.

NFPM can also be installed as a binary, which you can use to generate packages for your software, or, you can use it through [GoReleaser](https://goreleaser.com/), which will automate the entire release process for you.

## Design decisions

The idea of the `Register` function and blank imports is based on the `database/sql` package.
I'm not sure this was the best idea, and I may change this in the future, but it does allow to simply add more packagers and as well to import only the ones you want.

Unfortunately, I was unable to generate rpm packages without `rpmbuild`, so there is this one dependency if you want rpm packages.

On the bright side, generating debs was almost too easy: a couple of `tar.gz` and text files inside an `ar` file and that was it, working like a charm. All done natively with Go code.

## Future ahead

Maybe eventually we could add more formats - I'm not sure. Generating `deb` and `rpm` packages cover most of the cases I believe, and Snapcraft packages are already supported by [GoReleaser](https://goreleaser.com/) in a different pipe.

We will need a full suite of acceptance tests at some point - PRs welcome!

On the [GoReleaser](https://goreleaser.com/) side of things, fpm removal is about to happen, and it will automagically default to nfpm.

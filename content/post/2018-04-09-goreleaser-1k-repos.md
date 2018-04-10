---
title: "GoReleaser: 1k repositories and beyond"
date: 2018-04-09T20:09:13-03:00
tags:
- goreleaser
- golang
---

When [first announced GoReleaser]({{< ref "post/2017-01-02-goreleaser.md" >}})
roughly 1 year ago, on January 2017, I never thought it would be somewhat
famous.

I was just solving a problem I had like I always do.

It turns out more people had the same problem, and the feedback has been
awesome!

I'm also happy to announce that my not-very-scientific script recorded that
**1000 public GitHub repositories** are using GoReleaser!

![reposistories using GoReleaser over time](https://raw.githubusercontent.com/caarlos0/goreleaser-users/master/repos.png)

That is remarkable!

There are also some big and famous repositories using GoReleaser. Here are
the top 5 by the number of stars:

![top 5 repositories using goreleaser](https://raw.githubusercontent.com/caarlos0/goreleaser-users/master/stars.png)

Speaking of stars, GoReleaser has +2,5k stars!

[![goreleaser targazers over time](https://starcharts.herokuapp.com/goreleaser/goreleaser.svg)](https://starcharts.herokuapp.com/goreleaser/goreleaser)

Well, since that January, a lot has changed in GoReleaser, let's dive in!

## What changed

Well, almost everything, to be honest.

But let's point the biggest ones:

### New Pipes!

A lot of new features were added.
These includes Docker, Signing, Snapcraft and other pipes.

### nFPM

We started supporting `deb` and `rpm` packaging through [fpm], but, after
some time (and a few issues) I decided to create [nfpm].

It is a very basic packaging tool much like [fpm], but it's written in Go
and depending on fewer external dependencies (only `rpmbuild` at the moment).

### Rewritten the core

The core was poorly written in the first versions. It worked but was hard
to change and to add new things and kinds of artifacts.

It was rewritten and is more flexible now. This allows us and
external contributors to add new kinds of artifacts easier.

### Multiple language support

GoReleaser was created with Go in mind, but we recently changed its build
pipe to theoretically support more languages.

The first one in the pipe will probably be
[Rust](https://github.com/goreleaser/goreleaser/pull/520).

### GoDownloader

[Nick Galbreath][client9] started another awesome project called [GoDownloader]
and donated it to the GoReleaser org.

It is like a reverse GoReleaser: it creates shell scripts to download and
install software from its releases.

It integrates with the `.goreleaser.yaml` file in the repository and generate
the script based on that, or guess some defaults for projects that don't
use GoReleaser.

It is an awesome way to save some time, let's say,
[downloading and installing Hugo on your CI pipeline](https://github.com/caarlos0/carlosbecker.com/blob/master/Makefile).

### New logo

I couldn't let this pass.

My [lovely wife][carine] drew this beautiful new logo for GoReleaser:

{{< figure src="https://github.com/goreleaser/artwork/raw/master/goreleaserfundo.png" alt="goreleaser new logo" height="200px" >}}

It is **awesome**. You can check out other GoReleaser art in the
[artwork repository](https://github.com/goreleaser/artwork).

Oh, I'm working on having t-shirts and stickers, if anyone is interested!

**edit**:

{{< tweet 983696191564152832 >}}

### OpenCollective

I'm not a big believer in these things but wanted to try it out anyway.

If you use GoReleaser and it saved you some time and/or you liked it,
you can now donate a few bucks in
[GoReleaser's OpenCollective page](https://opencollective.com/goreleaser).

It is not a big thing but helps to keep the maintainers motivated.

### Contributors

We had pull requests from almost 50 different contributors so far!

- [goreleaser/goreleaser contributors](https://github.com/goreleaser/goreleaser/graphs/contributors)
- [goreleaser/godownloader contributors](https://github.com/goreleaser/godownloader/graphs/contributors)
- [goreleaser/nfpm contributors](https://github.com/goreleaser/nfpm/graphs/contributors)
- [goreleaser/archive contributors](https://github.com/goreleaser/archive/graphs/contributors)

## Another 1000

It has been a great gig so far. I can't wait to see what else we will
accomplish!

Hope to see you all again in another 1000 users!

[carine]: https://twitter.com/carinemeyer
[fpm]: https://github.com/jordansissel/fpm
[nfpm]: https://github.com/goreleaser/nfpm
[client9]: https://github.com/client9
[godownloader]: https://github.com/goreleaser/godownloader

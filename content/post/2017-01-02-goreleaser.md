---
date: 2017-01-02T00:00:00Z
slug: goreleaser
title: Fast and easy Go binaries delivery
tags:
- golang
- goreleaser
- package
---

I have some apps written in Go, which I deliver as binaries for each
platform using GitHub releases. Until now, I was doing it with a very
simple [shell script](https://github.com/goreleaser/goreleaser.github.io).

<!--more-->

But, I also wanted to dist these new releases in a homebrew tap. Adding that
in the script would be kind of trivial, but I didn't like the way I was using
it all (arguments, hacks, etc).

I could, of course, change the formulas by hand, but I'm lazy and that seems
just too much work for me.

So, time to my shell script to evolve!

### goreleaser

I wanted a tool that:

- provides reproducible builds
- is fast
- is easy to wrap with the CI
- distributes itself

So, I created [goreleaser](https://github.com/goreleaser/releaser)!

![goreleaser logo](https://avatars2.githubusercontent.com/u/24697112?v=3&s=200)

GoReleaser can build and release go binaries in `tar.gz` for several platforms,
and can create/update homebrew formulaes. All that by having a simple
`goreleaser.yml` in the repository root:

```yaml
repo: goreleaser/releaser
binary_name: release
brew:
  repo: goreleaser/homebrew-formulae
build:
  oses:
    - windows
    - darwin
    - linux
```

Then, I basically wire this in the `.travis.yml`:

```yaml
after_success:
  test -n "$TRAVIS_TAG" && curl -s https://raw.githubusercontent.com/goreleaser/get/master/latest | bash
```

Which generates release like this on GitHub:

![Release screenshot](/public/images/goreleaser-release-example.png)

It is so easy that I'm already using it in all my suitable Go projects:

- [GoReleaser itself](https://github.com/goreleaser/releaser)!
- [Antibody](https://github.com/getantibody/antibody) - A faster and simpler
antigen written in Golang
- [CloneOrg](https://github.com/caarlos0/clone-org) - Clone all repos
of a github organization
- [ForkCleaner](https://github.com/caarlos0/fork-cleaner) - Cleans up old
and inactive forks on your github account
- [GithubVacations](https://github.com/caarlos0/github-vacations) -
Automagically ignore all notifications related to work when you are on vacations
- [KarmaHub](https://github.com/caarlos0/karmahub) - Compares the amount of
issues and pull requests you created with the amount of comments and code
reviews you did
- [OrgStats](https://github.com/caarlos0/org-stats) - Get the contributor
stats summary from all repos of any given organization
- [TorrentWatcher](https://github.com/caarlos0/twatcher) - Automagically
download torrent files to ~/Downloads

### Roadmap

The roadmap is pretty much what you
[find in the issues](https://github.com/goreleaser/releaser/issues),
but, IMO, the highlights are:

- bintray support
- create linux packages (with fpm or something)
- post/pre build hooks

I believe those features will enable GoReleaser to be even more useful!

### Community Feedback

I was very glad to see the results of GoReleaser!

- First page on [/r/golang](https://www.reddit.com/r/golang/) for 2 days (and
top 1 most of that time)
- 300+ stars in just a few days
- [GitHub trending on Go](https://github.com/trending/go) (still at the time
of writing)
- Several comments on the [reddit post](https://www.reddit.com/r/golang/comments/5l3i9b/deliver_go_binaries_as_fast_and_easy_as_possible/)
- [This issue](https://github.com/goreleaser/releaser/issues/26)

Thank you all!

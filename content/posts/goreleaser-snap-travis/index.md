---
title: "GoReleaser: build and push Snapcraft packages from TravisCI"
date: 2018-10-20
draft: false
slug: goreleaser-snap-travis/
city: Joinville
toc: true
tags: [goreleaser, golang, ci-cd]
---

[GoReleaser](https://goreleaser.com/) was able to build [Snapcraft](https://snapcraft.io/goreleaser) packages for a long time, but it wasn't able to push them until today. Let's see how to wrap to your [TravisCI](https://travis-ci.org/goreleaser/goreleaser) build!

---

Since [v0.28.0](https://github.com/goreleaser/goreleaser/releases/tag/v0.28.0), [GoReleaser](https://goreleaser.com/) can create snapcraft packages and
upload them to the GitHub release. On [v0.91.0](https://github.com/goreleaser/goreleaser/releases/tag/v0.91.0), we added the support
to also push those packages to the [snap](https://snapcraft.io/goreleaser) store.

On this post I'll show how to wrap that with you [TravisCI](https://travis-ci.org/goreleaser/goreleaser) config.

## GoReleaser part

Let's assume you have a Snapcraft section more or less like the below
on you `.goreleaser.yml` file:

```yaml
# .goreleaser.yaml
snapcraft:
  summary: Foo bar
  description: |
    Longer foo bar
```

The only change needed here is to add a `publish: true` to the config, like
this:

```yaml
# .goreleaser.yaml
snapcraft:
  summary: Foo bar
  publish: true # <- this line
  description: |
    Longer foo bar
```

And that should be it on the [GoReleaser](https://goreleaser.com/) side.

## TravisCI part

On your `.travis.yml`, you probably have something like this:

```yaml
language: go
go: '1.11.x'
script: make build
deploy:
- provider: script
  skip_cleanup: true
  script: curl -sL http://git.io/goreleaser | bash
  on:
    tags: true
```

So, the first step is to install snapcraft on Travis, we can do it with
a combination of `addons`, `env` and `install`:

```yaml
language: go
go: '1.11.x'

# added
addons:
  apt:
    packages:
    - snapd
env:
- PATH=/snap/bin:$PATH
install:
- sudo snap install snapcraft --classic
# end added

script: make build
deploy:
- provider: script
  skip_cleanup: true
  script: curl -sL http://git.io/goreleaser | bash
  on:
    tags: true
```

Now, we need to do an export login on Snapcraft:

```sh
snapcraft export-login snap.login
```

This will create a `snap.login` file. Make sure to add it to the `.gitignore`:

```sh
echo snap.login >> .gitignore
```

Now, we need to somehow add that file to Travis. The way we do that is to use
the [encrypt file](https://docs.travis-ci.com/user/encrypting-files/) Travis' feature:

```sh
travis encrypt-file snap.login --add
```

This will add a line more or less like this to our `.travis.yml`:

```yaml
before_install:
- openssl aes-256-cbc -K $encrypted_123abc123_key -iv $encrypted_123abc123_iv -in snap.login.enc -out snap.login -d
```

So, the only step left is to do a snapcraft login on Travis, just add this
to your `.travis.yml`:

```yaml
after_success:
- test -n "$TRAVIS_TAG" && snapcraft login --with snap.login
```

So, in the end, our `.travis.yml` will look like this:

```yaml
language: go
go: '1.11.x'
addons:
  apt:
    packages:
    - snapd
env:
- PATH=/snap/bin:$PATH
before_install:
- openssl aes-256-cbc -K $encrypted_123abc123_key -iv $encrypted_123abc123_iv -in snap.login.enc -out snap.login -d
install:
- sudo snap install snapcraft --classic
script: make build
after_success:
- test -n "$TRAVIS_TAG" && snapcraft login --with snap.login
deploy:
- provider: script
  skip_cleanup: true
  script: curl -sL http://git.io/goreleaser | bash
  on:
    tags: true
```

## Wrapping up

Now you can just do a `git tag v1.2.3; git push --tags`, and [TravisCI](https://travis-ci.org/goreleaser/goreleaser) will
[GoReleaser](https://goreleaser.com/), which will do all the work for us.

If you are creating a new snap, you will also need to register the snap
first:

```sh
snapcraft register ${APP_NAME}
```

You can also move thing around, for example, the `snap install snapcraft`,
`snapcraft login` and etc can all be on the `after_success` part, depending
on how you build your app.

I'm not sure this is the 100% best way of doing this, but it works.
You can see that the [GoReleaser v0.91.1](https://github.com/goreleaser/goreleaser/releases/tag/v0.91.1) released itself
to the [Snapcraft store](https://snapcraft.io/goreleaser) by looking into the [build logs](https://travis-ci.org/goreleaser/goreleaser/builds/444189008).

That means you can now install [GoReleaser](https://goreleaser.com/) on Linux using something like:

```sh
snap install goreleaser --classic
```

Maybe it's worth mention that this was implemented in
[Hacktoberfest Joinville](http://hacktoberfest.joinville.br/)! Great to see you all there!

Hope you enjoy it! Cheers!

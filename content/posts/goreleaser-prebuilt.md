---
title: "Using the new prebuilt builder on GoReleaser"
date: 2021-09-12
draft: false
slug: goreleaser-prebuilt
city: Cascavel
toc: true
tags: [goreleaser, goreleaser-pro, golang, ci-cd]
---

You can now [import pre-built binaries](https://goreleaser.com/customization/build/#import-pre-built-binaries) into [GoReleaser](https://goreleaser.com)! 

This feature was made with mainly two cases in mind:

1. You are migrating to GoReleaser and already have the build part covered by a `Makefile` or some other tool, and you don't want to change that
2. You want to build each platform on a different machine and join the binaries later (for performance or CGO reasons)

Let's talk about them.

## 1. Another builder

In this case, you'll likely have a `Makefile` or some other tool doing the heavy lifting of building everything for you already.

In that case, you can simply use something like the following as your build config for GoReleaser:

```yaml
# .goreleaser.yml
builds:
-
  # Set the builder to prebuilt
  builder: prebuilt

  # Specify the goos/goarch/goarm/gomips you want:
  goos: [linux, darwin]
  goarch: [amd64, arm64]

  prebuilt:
    # Set the path template of where to look for the binaries
    path: ./tmp/mybin_{{ .Os }}_{{ .Arch }}

# archive, package, sign, etc
```

This will make GoReleaser try to import into its release cycle the following binaries:

- `./tmp/mybin_linux_amd64`
- `./tmp/mybin_linux_arm64`
- `./tmp/mybin_darwin_amd64`
- `./tmp/mybin_darwin_armd64`

The only check GoReleaser will do is if these files are readable. If so, it will proceed as if they were built by GoReleaser itself.

PS: it is important that the folder in which your pre-built binaries are is either outside the repository or in your `.gitignore`.

## 2. Split builds in several machines

You might want to this to speed up/parallelize the build process, or maybe because of CGO.

Whatever your case may be, you can still use GoReleaser to build the binaries by having 2 or more GoReleaser config files, for example:

```yaml
# .builds.goreleaser.yml
dist: tmp
builds:
-
  goos: [linux, darwin]
  goarch: [amd64, arm64]
	# flags, envs, etc
```

You can then build each by running something like on each machine:

```bash
goreleaser build --single-target -f .builds.goreleaser.yml
```

Then, you can copy all the built binares into the `tmp` folder, and import those binaries using a config like this:

```yaml
# .goreleaser.yml
builds:
-
  # Set the builder to prebuilt
  builder: prebuilt

  # Specify the goos/goarch/goarm/gomips you want:
  goos: [linux, darwin]
  goarch: [amd64, arm64]

  prebuilt:
    # Set the path template of where to look for the binaries
    path: ./tmp/mybin_{{ .Os }}_{{ .Arch }}

# archive, package, sign, etc
```

And GoReleaser will proceed as usual.

PS: it is important that the folder in which your pre-built binaries are is either outside the repository or in your `.gitignore`.

## This is a new feature

This feature just landed in [GoReleaser Pro v0.179.0](https://github.com/goreleaser/goreleaser-pro/releases/tag/v0.179.0-pro), which was released today!

Here's the documentation for this feature: [link](https://goreleaser.com/customization/build/#import-pre-built-binaries).

If you want to test this feature, [here's a limited seats discount link](https://beckersoft.gumroad.com/l/goreleaser/prebuilt-article) since you read this entire article 🤘

---

Cheers! ✌️

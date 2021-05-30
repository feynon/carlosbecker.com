---
title: "GoReleaser: 4 years releasing software"
date: 2021-01-07
draft: false
slug: goreleaser-4-years
city: Cascavel
toc: true
tags: [goreleaser, golang, docker]
---

Last year, I made a blog post about G[oReleaser turning 3 years old]({{< ref "goreleaser-3-years.md" >}}).

I kind of like it, so this year I'm writing one about it turning 4. Maybe I'll make this a habit so I can quickly see how it improved over the course of a year.

## The year of the linux on the desktop

I kind of promised myself that 2020 will be the year of GoReleaser v1.0.0.

COVID and all, priorities changed for a while... and it didn't happen.

## Improvements

But that doesn't mean we didn't have improvements! 

Here are some highlights:

- removed a lot of deprecated settings
- added `GOMIPS` support
- nfpm can now sign packages
- added build hooks
- added `xz` compression
- added the `source` pipe
- refactored from `kingpin` to `cobra`
- added shell completions
- added the `goreleaser build` command
- switched from travis to GitHub Actions
- migrated [goreleaser.com](http://goreleaser.com) to Material Mkdocs
- added support for reproducible builds
- added support for closing milestones after the release
- created fileglob and replaced zglob usage with it
- added `apk` package creation
- added support to creating multi-architecture Docker images (`docker_manifests`)
- added support for using `docker buildx build` (`dockers.use_buildx`)
- `arm64` is now a default `GOARCH`
- allow to use nfpm packages when creating Docker images

Besides that, we also had some refactory (mostly on tests).

## M1 Macs

On top of all that, Apple Silicon is a thing now, and we have the support in a PR already... just waiting on the Go 1.16 release.

I do not yet have a M1 Mac to properly test those things, but I'm sure someone will come up to help with that.

## That's all folks

Stay safe, and I hope to see y'all again next year!

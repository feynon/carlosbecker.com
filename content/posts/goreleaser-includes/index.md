---
title: "Using GoReleaser includes feature"
date: 2021-07-03
draft: false
slug: goreleaser-includes
city: Cascavel
toc: true
tags: [goreleaser, goreleaser-pro, golang]
---

[GoReleaser Pro](https://goreleaser.com/pro) was released [about a month ago]({{< ref "/posts/goreleaser-pro/index.md" >}}), and with it, the [ability to include GoReleaser config files](https://goreleaser.com/customization/includes/).

In practice, within an organization, is common to have either shared pieces of configuration, if not sharing the entire files within several projects. We all been there, and keeping all those files up to date is not something fun to do.

With includes, you can include either parts or full config files from anywhere on the internet, thus enabling having only a couple of files to keep up to date.

In this post we'll explore how I'm doing that in my own projects, maybe you find something you can use on yours 🙂

## The `goreleaserfiles` repository

I created a repository called [goreleaserfiles](https://github.com/caarlos0/.goreleaserfiles) (inspired by `dotfiles` and `vimfiles`).

My approach was a more modular one: I keep one file per "feature" I want.

For instance, I have a `package.yml` and a `package_with_completions.yml`. The first one has the "simple packaging" I usually use, the other one has a more complete approach (pun intended), for tools that have completion scripts.

Both these two scripts (and some others) use [custom template variables](https://goreleaser.com/customization/templates/#custom-variables), such as `{{ .description }}` and `{{ .homepage }}`. I try to use the same names across all templates so I don't need to repeat too much stuff.

## Including

The end result looks pretty clean. Here's an example from [tasktimer](https://github.com/caarlos0/tasktimer):

```yaml
project_name: tt
variables:
  homepage: https://github.com/caarlos0/tasktimer
  description: Task Timer (tt) is a dead simple TUI task timer
includes:
  - from_url:
      url: caarlos0/goreleaserfiles/main/build.yml
  - from_url:
      url: caarlos0/goreleaserfiles/main/package_with_completions.yml
  - from_url:
      url: caarlos0/goreleaserfiles/main/release.yml
```

Here we:

- override `project_name` (would be `tasktimer` otherwise)
- set the `homepage` and `description` custom variables
- include `build.yml`, `package_with_completions.yml` and `release.yml`.

GoReleaser will recursively import those files in order when you run `goreleaser release` and the end result will be saved in your `dist` folder for further investigation.

You can also import from other directories within the same repository, which can be particularly useful with monorepos.

## That's it

That's a quick overview on [GoReleaser includes](https://goreleaser.com/customization/includes/) feature.

If you'll like to use it, check the [Pro](https://goreleaser.com/pro) page and get your license.

Also, since you read this to the end, [click here to get it with a discount](https://gumroad.com/l/goreleaser/jdny4po) (limited seats).

---

## Bonus: preventing deprecated config usage

To make sure my configs are always up to date, I wrote a script that writes a GitHub Actions workflow to run `goreleaser check -q` on every file `.yml` file.

{{< img caption="Workflow run example." src="32a74969-29ea-4949-82b8-38871369877e.png" >}}

That way I get a build error if I'm using deprecated settings or if my files are **very** wrong.

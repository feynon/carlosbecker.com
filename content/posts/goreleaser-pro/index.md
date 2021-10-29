---
title: "Announcing GoReleaser Pro"
date: 2021-05-30
draft: false
slug: goreleaser-pro
city: Cascavel
toc: true
tags: [goreleaser, golang, goreleaser-pro]
---

After [more than 4 years]({{< ref "/posts/goreleaser-4-years/index.md" >}}) working on GoReleaser, I'm launching a [Pro version](https://goreleaser.com/pro)!

## Why?

I think that's the easiest question to answer: money.

I have been working on GoReleaser basically for free for several years now. It is true I got some money from sponsors and some help from other contributors (thanks y'all 🖤), [but the majority of the work is still done by me alone](https://github.com/goreleaser/goreleaser/graphs/contributors).

GoReleaser Pro is an attempt to monetize it, so I can justify the time spent on it.

## What does it mean to GoReleaser?

GoReleaser is now split into OSS and Pro.

The Pro version is kept by me alone. Features I found to be "enterprisy" will likely be added to the Pro version, while other features and bugfixes will go into both of them.

## What are the Pro Features?

For now, a very shy start:

- Ability to [include](https://goreleaser.com/customization/includes/) config files (useful for common configurations);
- Improved global hooks and [global after hooks](https://goreleaser.com/customization/hooks/);
- [Monorepo support](https://goreleaser.com/customization/monorepo);
- [Custom template variables](https://goreleaser.com/customization/templates/#custom-variables) (goes well with [includes](https://goreleaser.com/customization/includes/)).

Will keep adding more features (assuming GoReleaser Pro actually sells).

## Enterprise support?

The thing about enterprise support is that it takes a ton of time, and time is limited.

My idea is to start selling licenses only, and eventually, if companies are interested, sell support as well, with SLAs and all that.

If your company is interested, [let me know](mailto:carlos@becker.software).

## How can I get it?

Glad you asked! Just head to [https://goreleaser.com/pro](https://goreleaser.com/pro) and hit that big button.

Or, since you read it to the end, [click here to get it with a discount](https://gumroad.com/l/goreleaser/jdny4po) (limited seats).

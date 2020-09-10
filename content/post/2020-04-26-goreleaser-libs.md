---
title: "Publishing libraries with GoReleaser"
date: 2020-04-26
draft: false
slug: goreleaser-libs
city: Joinville
---

What if I told you you can now automate the release of your libraries as well?

---

I think some of you are used to just run `goreleaser` and have your binaries compiled, archived, packaged and published... but the experience was not that great regarding libraries.

Thanks to [@leogr](https://github.com/leogr), on [v0.131.0](https://github.com/goreleaser/goreleaser/releases/tag/v0.131.0) we shipped the ability to [skip builds entirely](https://github.com/goreleaser/goreleaser/pull/1419), which means you can now release libraries!

The usage is pretty simple, just create a config setting `build.skip: true` and you're good to go:

```
before:
  hooks:
    - go mod tidy
builds:
- skip: true
changelog:
  sort: asc
  filters:
    exclude:
    - '^docs:'
    - '^test:'
```

Of course, a lot of pipes won't work, since they depend on the produced binaries. But it will [generate a release](https://github.com/caarlos0/env/releases/tag/v6.2.2), which is already better than doing it manually. 😌

Hope it helps!
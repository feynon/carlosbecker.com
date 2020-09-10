---
title: "Cleanup old GitHub Forks"
date: 2018-06-02
draft: false
slug: cleanup-old-github-forks
city: Joinville
---

I like to keep my GitHub clean. I delete forks I'm not using anymore, move old abandoned repositories to my [graveyard](https://carlosbecker.com/posts/repositories-graveyard/) and etc.

<!--more-->

I don't like to delete things manually though.

I assume I can just delete forks that match all these rules:

- have no forks - no one depends on my changes;
- have no stars - no one liked my changes 🤷‍♂️;
- have no open pull requests to upstream - my PR was either merged or closed;
- had no activity in the last 1 month - I'll probably won't use it again;

So, I was basically looking over the repository list and manually going to the fork settings, copying the name of the repo, pasting in the delete popup and etc.

Doing that from time to time is really boring.

## fork-cleaner

So, I wrote a tool for that - in **2016**. Yes, 2 years ago already.

Anyway, it does all the checks mentioned above, plus, you can customize the inactivity time (e.g. `1d` instead of `1m`), blacklist repos and include or not private forks.

### Install

You can install it On macOS:

```shell
$ brew install fork-cleaner
```

or on Linux:

```shell
$ snap install fork-cleaner
```

### Usage

You'll need a GitHub token with the `repo` scope, you can create one in the [settings page](https://github.com/settings/tokens/new).

You can then either export it as the `GITHUB_TOKEN` environment variable or using the `--token` flag.

Then just run `fork-cleaner`:

```shell
$ fork-cleaner
3 forks to delete:
 --> https://github.com/caarlos0/coinmarketcap-exporter
 --> https://github.com/caarlos0/hugo-paper
 --> https://github.com/caarlos0/prometheus_couchbase_exporter
Remove the above listed forks? (y/n) [n]:
```

Then you just need to accept or not. Waaay faster than doing it manually.

Hope it helps keeping you account clean as well 🤘
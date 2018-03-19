---
date: 2017-07-08T00:00:00Z
slug: chart-repo-stars
title: Charting Repository Stars
tags:
- github
---

I always wanted to know how stargazers of my repos increased over time.
I didn't found a good way of doing that, so I wrote an app for that™.

The app stack is simple:

- go 1.8+
- glide
- gorilla/mux
- apex/log
- go-chart
- go-redis
- heroku

It is live at [starcharts.herokuapp.com](https://starcharts.herokuapp.com) and
the code is OpenSource at
[caarlos0/starcharts](https://github.com/caarlos0/starcharts).

The charts look like this one:

[![Stargazers over time](https://starcharts.herokuapp.com/goreleaser/goreleaser.svg)](https://starcharts.herokuapp.com/goreleaser/goreleaser)

[GoReleaser](https://github.com/goreleaser) stargazers over time.

Hope you folks enjoy it!

---
title: "Charting Repository Stars"
date: 2017-08-08
draft: false
slug: chart-repo-stars
city: Joinville
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

It is live at [starchart.cc](https://starchart.cc) and the code is OpenSource at [caarlos0/starcharts](https://github.com/caarlos0/starcharts).

The charts look like this one:

![](Untitled-5bbbaab9-18ce-47b5-b1b2-03e9e1705cb8.png)

[GoReleaser](https://github.com/goreleaser) stargazers over time.

[GoReleaser](https://github.com/goreleaser) stargazers over time.

Hope you folks enjoy it!
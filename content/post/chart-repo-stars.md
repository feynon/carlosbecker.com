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

![GoReleaser stargazers over time.](/public/images/chart-repo-stars/3f40643a-8660-4f5b-804d-dc6533416978.png)

[GoReleaser](https://github.com/goreleaser) stargazers over time.

Hope you folks enjoy it!

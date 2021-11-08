---
title: "Moving to Hugo"
date: 2018-03-19
draft: false
slug: moving-to-hugo
city: Joinville
toc: true
tags: [blog, hugo]
---

After some time, I finally decided to move my blog from [Jekyll](https://jekyllrb.com/) to [Hugo](https://gohugo.io/).

---

That was probably the best thing I did in the past few weeks. My Jekyll builds were really slow, and the feedback loop was too big. I get easily distracted, so, I was not writing much.

You can read some of my fight to make it faster [here]({{< ref "/posts/jekyll-build-time/index.md" >}}). Well, with Hugo,
doesn't matter which features I use, full build time is always less than 1 second and incremental builds take ~300ms. **Fast** should be Hugo's middle name!

Also, my build was running [htmlproof](https://github.com/gjtorikian/html-proofer), which is nice, but also extremelly slow, so I only ran it on Travis. With old posts breaking randomly, I eventually gave up.

Now, I'm using [Hugo](https://gohugo.io/) with a custom version of the [hugo-type-theme](https://github.com/caarlos0/hugo-type-theme) (pull request is open though), and it's being automagically deployed to [Netlify](https://netlify.com/), with SSL by [Let's Encrypt](https://letsencrypt.org/).

I'm also linting everything with a faster port of htmlproof, written in Go. It's called [htmltest](https://github.com/wjdp/htmltest) and it is **awesome**.

On the automation end, I'm using the good-old-new-hype of `make`. So, take a look at the `Makefile` in the root of the repo if you wish.

### Next steps

Probably enable the performance stuff on Netlify and hopefully write more often again!

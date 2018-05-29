---
title: "Moving to Hugo"
date: 2018-03-19T01:21:36-03:00
tags:
- jekyll
- hugo
---

After some time, I finally decided to move my blog from [Jekyll][] to [Hugo][].

<!--more-->

That was probably the best thing I did in the past few weeks. My Jekyll
builds were really slow, and the feedback loop was too big. I get easily
distracted, so, I was not writing much.

You can read some of my fight to make it faster
[here]({{< ref "post/2017-05-14-jekyll-build-time.md" >}}). Well, with Hugo,
doesn't matter which features I use, full build time is always less than
1 second and incremental builds take ~300ms. **Fast** should be Hugo's
middle name!

Also, my build was running [htmlproof][], which is nice, but also
extremelly slow, so I only ran it on Travis. With old posts
breaking randomly, I eventually gave up.

Now, I'm using [Hugo][] with a custom version of the [hugo-type-theme][] (
pull request is open though), and it's being automagically deployed to
[Netlify][], with SSL by [Let's Encrypt][lestsencrypt].

I'm also linting everything with a faster port of htmlproof, written in Go.
It's called [htmltest][] and it is **awesome**.

On the automation end, I'm using the good-old-new-hype of `make`. So, take a
look at the `Makefile` in the root of the repo if you wish.

### Next steps

Probably enable the performance stuff on Netlify and hopefully write more
often again!

[jekyll]: https://jekyllrb.com/
[hugo]: https://gohugo.io/
[htmlproof]: https://github.com/gjtorikian/html-proofer
[hugo-type-theme]: https://github.com/caarlos0/hugo-type-theme
[netlify]: https://netlify.com
[lestsencrypt]: https://letsencrypt.org/
[htmltest]: https://github.com/wjdp/htmltest

---
title: "Turbolinks animated page transitions"
date: 2013-01-22
draft: false
slug: animating-page-transitions-in-turbolinks
city: Joinville
toc: true
tags: [ruby-on-rails]
---

Since I've seen the new [Basecamp](http://basecamp.com/), I fall in love with it.

---

It's fast, had sleek animations in page changes and so on.

Well, I put one thing in my head: "if they can, I can". So I worked. Can get a really good speed with a 1Gb RAM + SSD HD from [DigitalOcean](http://digitalocean.com/) for $ 10 per month (really cheap, I believe), and tweaked my app to use the [Puma](http://puma.io/) server (with socket, of course) + [Nginx](http://nginx.org/) (that also serves the assets), [cache-digest](https://github.com/rails/cache_digests) GEM to up the irons while caching things, and [turbolinks](https://github.com/rails/turbolinks) for a really simple AJAX (done with pushstate) (I'll talk about this architecture another day).

Well, I'm pretty happy, except that I don't had those awesome bad-ass animations on page change. So, I made it.

## [turbolinks_transitions](https://github.com/caarlos0/turbolinks_transitions)

{{< figure caption="Gif showing an example app with turbolinks-transitions." src="/public/images/animating-page-transitions-in-turbolinks/77a541f6-e264-4b6d-97a1-508dc406b22a.gif" >}}

Just a cool animated gif with the example!

Use it is pretty simple. Just add the gem and import the JS, and you're ready to go!

See the complete how to use [here](https://github.com/caarlos0/turbolinks_transitions#usage).

## Contribute

I'll be glad to hear your suggestions and merge your pull-requests! Feel free to do that.

See you soon.

---

> PS: the example above is running locally, so, maybe you will not see it at first, but, just look a little closer and you shall see a little fade while switching pages :)

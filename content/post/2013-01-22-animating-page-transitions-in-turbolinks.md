---
date: 2013-01-22T00:00:00Z
slug: animating-page-transitions-in-turbolinks
title: Turbolinks animated page transitions
---

Since I've seen the new [Basecamp][base], I fall in love with it.

It's fast, had sleek animations in page changes and so on.

Well, I put one thing in my head: "if they can, I can". So I worked. Can get a
really good speed with a 1Gb RAM + SSD HD from [DigitalOcean][digital]
for $ 10 per month (really cheap, I believe), and tweaked my app to use
the [Puma][puma] server (with socket, of course) + [Nginx][nginx] (that also
serves the assets), [cache-digest][cache] GEM to up the irons while caching
things, and [turbolinks][turbolinks] for a really simple AJAX (done with pushstate)
(I'll talk about this architecture another day).

Well, I'm pretty happy, except that I don't had those awesome bad-ass animations
on page change. So, I made it.

### [turbolinks_transitions][transitions]

![example](/public/images/turbolinks_transition.gif "Just a cool animated gif with the example!")

Use it is pretty simple. Just add the gem and import the JS, and you're
ready to go!

See the complete how to use [here][usage].

### Contribute

I'll be glad to hear your suggestions and merge your pull-requests! Feel free
to do that.

See you soon.


----

> PS: the example above is running locally, so, maybe you will not see it at first,
> but, just look a little closer and you shall see a little fade while switching
> pages :)



[base]: http://basecamp.com
[digital]: http://digitalocean.com
[puma]: http://puma.io
[nginx]: http://nginx.org/
[cache]: https://github.com/rails/cache_digests
[turbolinks]: https://github.com/rails/turbolinks
[transitions]: https://github.com/caarlos0/turbolinks_transitions
[usage]: https://github.com/caarlos0/turbolinks_transitions#usage

---
date: 2017-05-14T00:00:00Z
slug: jekyll-build-time
title: Improving Jekyll build time
tags:
- jekyll
---

I've been using Jekyll on my blog since 2012. It is great!
But, lately, its slow build times started to bother me.

<!--more-->

I consider moving to [Hugo], but that seemed like a lot
of work, and I didn't know what the real problem was.
In this post I'll show how I improved my Jekyll blog build time in about
**75%**.

[hugo]: https://gohugo.io

## Current state

We will use `jekyll build --profile` to see what's taking too much time to
build. In order to not make this too extensive, I'll keep only the top 10
files and the final time.

Let's see how much time it takes right now:

```console
Filename                                                  | Count |    Bytes |  Time
----------------------------------------------------------+-------+----------+------
_layouts/default.html                                     |   117 | 3155.23K | 9.095
_includes/head.html                                       |   118 | 2242.11K | 9.025
_layouts/compress.html                                    |   117 | 3094.18K | 0.610
_layouts/post.html                                        |    93 |  693.39K | 0.195
_includes/post_header.html                                |    93 |   19.27K | 0.069
_layouts/center.html                                      |     1 |   19.81K | 0.067
_includes/header.html                                     |   118 |   48.51K | 0.053
_includes/related_posts.html                              |    93 |   60.75K | 0.051
_includes/navigation.html                                 |   118 |   19.71K | 0.028
feed.xml                                                  |     1 |  104.09K | 0.028

                    done in 13.017 seconds.
```

13 seconds and a warning. Let's try to improve that.

## Removing HTML compress

I added it last year, while I was still not using Cloudfare.

As you can see in the profile, it was taking almost 1s, and Cloudfare
already do that, so it's kind of useless in this case.

I reverted the commit and profile again:

```console
Filename                                                  | Count |    Bytes |  Time
----------------------------------------------------------+-------+----------+------
_layouts/default.html                                     |   117 | 3157.83K | 8.535
_includes/head.html                                       |   118 | 2242.11K | 8.479
_layouts/post.html                                        |    93 |  695.99K | 0.179
_layouts/center.html                                      |     1 |   19.81K | 0.070
_includes/post_header.html                                |    93 |   19.27K | 0.066
_includes/header.html                                     |   118 |   48.51K | 0.051
_includes/related_posts.html                              |    93 |   60.75K | 0.047
_includes/navigation.html                                 |   118 |   19.71K | 0.027
feed.xml                                                  |     1 |  106.87K | 0.022
sitemap.xml                                               |     1 |   12.79K | 0.017

                    done in 11.318 seconds.
```

2 seconds! But stil it takes too much time!

## Remove CSS inlining

Some time ago I also put together a hack to inline the CSS in each page.
It is good because the entire CSS comes with the page, but at the same time,
Jekyll take years to build that. Reverting that commit seem to have done
the greater good so far:

```console
Filename                                                  | Count |    Bytes |  Time
----------------------------------------------------------+-------+----------+------
_layouts/default.html                                     |   117 | 1132.99K | 1.532
_includes/head.html                                       |   118 |  197.62K | 1.383
_layouts/post.html                                        |    93 |  698.31K | 0.236
_includes/post_header.html                                |    93 |   19.27K | 0.081
_includes/related_posts.html                              |    93 |   60.75K | 0.064
_includes/header.html                                     |   118 |   48.51K | 0.059
_includes/navigation.html                                 |   118 |   19.71K | 0.030
feed.xml                                                  |     1 |  109.32K | 0.029
sitemap.xml                                               |     1 |   12.79K | 0.020
archive.md                                                |     1 |   14.61K | 0.019

                    done in 4.806 seconds.
```

Down to 4secs! Can we improve it even more?

## Split Google Analytics into another file

I had a code block like this in my `default.html`:

```erb
{% if site.google_analytics %}
<script type="text/javascript">
  // the default google analytics script
</script>
{% endif %}
```

I've put that in a file and just included it in the `default.html` layout.

Results:

```console
Filename                                                  | Count |    Bytes |  Time
----------------------------------------------------------+-------+----------+------
_layouts/default.html                                     |   117 | 1136.08K | 1.224
_includes/head.html                                       |   118 |  197.62K | 1.092
_layouts/post.html                                        |    93 |  703.79K | 0.177
_includes/post_header.html                                |    93 |   19.27K | 0.064
_includes/header.html                                     |   118 |   48.51K | 0.050
_includes/related_posts.html                              |    93 |   60.75K | 0.047
_includes/navigation.html                                 |   118 |   19.71K | 0.027
feed.xml                                                  |     1 |  117.31K | 0.026
_includes/analytics.html                                  |   118 |   54.62K | 0.018
sitemap.xml                                               |     1 |   12.79K | 0.017

                    done in 3.554 seconds.
```

I don't really know why this is faster, but it is. Some kind of caching maybe?

## The last 3 seconds

Now, if I disable the [jekyll-seo-tag] plugin and the highlighter, I'm
down to 1.4secs.

SEO tags are kind of important, so I don't want to remove them.

The highlighter can be replaced by `highlight.js`, which also seems to
work better in some cases. In order to do that, I needed to change some
things.

First, disable the highlighter in the `_config.yml` file:

```yaml
kramdown:
  syntax_highlighter_opts:
    disable : true
```

Then, add the `script` and `link` tags into my `_post.html` layout
(including custom langs and theme):

```html
<link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.11.0/styles/darcula.min.css">
<script src="//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.11.0/highlight.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.11.0/languages/go.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.11.0/languages/erb.min.js"></script>
<script src="//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.11.0/languages/yaml.min.js"></script>
<script>hljs.initHighlightingOnLoad();</script>
```

Now, if I profile it again:

```console
Filename                                                 | Count |   Bytes |  Time
---------------------------------------------------------+-------+---------+------
_layouts/default.html                                    |   117 | 991.15K | 1.225
_includes/head.html                                      |   118 | 197.62K | 1.086
_layouts/post.html                                       |    93 | 559.20K | 0.185
_includes/post_header.html                               |    93 |  19.26K | 0.061
_includes/related_posts.html                             |    93 |  60.75K | 0.056
_includes/header.html                                    |   118 |  48.51K | 0.049
_includes/navigation.html                                |   118 |  19.71K | 0.025
feed.xml                                                 |     1 |  88.32K | 0.024
_includes/analytics.html                                 |   117 |  54.16K | 0.019
sitemap.xml                                              |     1 |  12.79K | 0.016

                    done in 2.5 seconds.
```

**2.5 seconds**. I can live with that.

## How about the users?

Well, [Cloudfare] has several neat features. It can minify files,
improve image sizes, bundle js files and so forth.

All those features actually give me faster page loads than the
bundling and etc that I had before (from ~3.5s to ~1.9s)!

I strongly recommend you take a look at it, it will even give you
free HTTPS :)

[jekyll-seo-tag]: https://github.com/jekyll/jekyll-seo-tag
[cloudfare]: https://www.cloudflare.com/

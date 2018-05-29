---
date: 2015-01-09T00:00:00Z
slug: jekyll-with-sass
title: Jekyll with Sass
---

I followed [@mdo](http://markdotto.com/) recent article
["Using Sass with Jekyll"](http://markdotto.com/2014/09/25/sass-and-jekyll/),
and wanted to point out the results.

I'm using some version of [Lanyon](http://lanyon.getpoole.com/) with some
custom stuff. So, I had 4 CSS files:

![Requests](/public/images/sass-1.png)

Summing it up, ~22K. It's not a lot, but, thinking about mobile 3G plans
that are shit (like brazilian ones), why not save some bytes and requests?

So, I moved all those files to a `_scss` subfolder, and changed their
extensions to `.scss` instead of `.css`.

Then, in my `public/css` folder, I created a `styles.scss` like this one:

```scss
---
# Needed for jekyll...
---

@import "poole";
@import "syntax";
@import "lanyon";
@import "carlos";
```

Also, I added the following section to my `_config.yml`:

```yaml
sass:
  sass_dir: _scss
  style: :compressed
```

Finally, changed my `_includes/head.html` to import only the new `styles.css`:

```html
<link rel="stylesheet" href="/public/css/styles.css">
```

And boom! It worked. With this, my previously four requests of ~22K went
to one request with 12.8K!

![Request](/public/images/sass-2.png)

Besides that, now I have all the power that
Sass provides, in my blog, without any hacks. And it works on GitHub pages!

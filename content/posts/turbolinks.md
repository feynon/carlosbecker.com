---
title: "This site is now powered by Turbolinks"
date: 2013-12-26
draft: false
slug: turbolinks
city: Marechal Cândido Rondon
toc: true
tags: [ruby-on-rails]
---

I think that [turbolinks](https://github.com/rails/turbolinks) is great: it mades it easy to add [AJAX PushState](https://www.google.com.br/search?q=AJAX+PushState) to your [Rails](http://rubyonrails.org/) Applications. The only problem with that is that we can't use it **any** WEB application, because it's a Ruby Gem. So I did some ugly-but-easy hacks and add it to this very site. I will describe the steps above.

<!--more-->

### First things first

- This site is [OpenSource](https://github.com/caarlos0/caarlos0.github.com.git).
- I'm using Jekyll alongside with less, bower, grunt and of course, node and npm.
- You can take a look at [these commits](https://github.com/caarlos0/caarlos0.github.com/commit/cb17b421e57aec67f3d7a582696e62d863c3689f) to see what I did.

With that said, let's do the thing.

### Getting turbolinks

Well, turbolinks is a rubygem, but it has a lot of dependencies, and Jekyll don't have the Rails Asset Pipeline, so I can't figure out how to use this way.

So, to get the pure turbolinks file, I used grunt-curl and compiled it to javascript using grunt-contrib-coffee:

```javascript
curl: {
  turbolinks: {
    src: 'https://raw.github.com/rails/turbolinks/master/lib/assets/javascripts/turbolinks.js.coffee',
    dest: '_assets/turbolinks.coffee'
  }
},
coffee: {
  compile: {
    files: {
      '_assets/turbolinks.coffee.js': '_assets/turbolinks.coffee'
    }
  }
},
concat: {
  dist: {
    src: [
      // other JS files
      '_assets/turbolinks.coffee.js',
      '_assets/up.js'
    ],
    dest: 'js/up.min.js'
  }
},
uglify: {
  build: {
    src: 'js/up.min.js',
    dest: 'js/up.min.js'
  }
}
```

As you can see, I get always the last file from trunk and compile to plain old JavaScript. After that, I also concatenate it with other JS files and uglify the result (using grunt-contrib-concat and grunt-contrib-uglify, respectively).

### Dealing with `style="background-image: url(image.jpg);"`

For some reason I'm now quite sure, turbolinks mess up with this kind of style declaration (which I use in the image headers). The solution I've found is kinda weird, but it works:

```javascript
var reloadImages = function() {
  var styles, style, url, _i, _len, _el;
  styles = Array.prototype.slice.call(
    document.body.querySelectorAll("[style]")
  );
  for (_i = 0, _len = styles.length; _i < _len; _i++) {
    _el = styles[_i];
    style = _el.getAttribute("style");
    if (!(style.indexOf("url(") > -1)) {
      continue;
    }
    url = style.match(/url\((.*)\)/)[1];
    _el.style.backgroundImage = "url(" + url + ")";
  }
};
$(document).on("page:change", function() {
  reloadImages();
});
```

The twitter button adds a script to the head, which turbolinks doesn't replace, so, it gets buggy. To fix that I changed a little the default button markup:

```html
<a  href="https://twitter.com/share" class="twitter-share-button"
data-lang="en" data-size="large"
data-url="{{ site.production_url }}{{ page.url }}"
data-text="{{ page.title }}">
  Tweet
</a>
<script>
!function(d,s,id){
  var js,fjs=d.getElementsByTagName(s)[0];
  if(!d.getElementById(id)){
    js=d.createElement(s);
    js.id=id;
    js.src="//platform.twitter.com/widgets.js";
    fjs.parentNode.insertBefore(js,fjs);
  }
}(document,"script","twitter-wjs");
</script>
```

And add something at the page change event:

```javascript
var twttr;
$(document).on("page:change", function() {
  if (twttr) {
    twttr.widgets.load();
  }
});
```

It was the best solution I've found.

### Dealing with Disqus load on bottom

I changed the Disqus scripts to only load the commends when the user reach the bottom of the page:

```html
<div id="disqus_thread">
  Loading Comments...
</div>
<script type="text/javascript">
  var disqus_loaded = false;
  var disqus_shortname = '{{ site.disqus_shortname }}';

  function load_disqus () {
    disqus_loaded = true;
    var dsq = document.createElement('script');
    dsq.type = 'text/javascript';
    dsq.async = true;
    dsq.src = 'http://' + disqus_shortname + '.disqus.com/embed.js';
    (document.getElementsByTagName('head')[0] || document.getElementsByTagName('body')[0]).appendChild(dsq);
  }

  window.onscroll = function(e) {
    var   currentScroll = (window.innerHeight + window.scrollY)
        , elScroll = document.getElementById('disqus_thread').offsetTop;
    if (currentScroll >= elScroll && disqus_loaded == false) {
      load_disqus()
    }
  };
</script>
<noscript>Please enable JavaScript to view the <a href="http://disqus.com/?ref_noscript">comments powered by Disqus.</a></noscript>
```

But it causes me some trouble with Turbolinks. To fix that, I simply did:

```javascript
$(document).on("page:fetch", function() {
  window.onscroll = void 0;
});
```

And it worked!

I also decided to add nprogress, a lib that provides the loading bar medium-style (and youtube-style). I had to add it to my `bower.json`, add the import in my less file (with a `(less)` prefix, so it imports it as a less file) and mix it up in my js. I also had to bind the events to nprogress, like this:

```javascript
$(document).on("page:fetch", function() {
  NProgress.start();
});

$(document).on("page:change", function() {
  NProgress.done();
});

$(document).on("page:restore", function() {
  NProgress.remove();
});

// hides the spinner
NProgress.configure({ showSpinner: false });
```

And it was working as expected.

### The Final Countdown

The result is what you're seeing right now. I found it neat and it loads really faster. Hope you like it!

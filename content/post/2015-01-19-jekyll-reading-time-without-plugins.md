---
title: "Jekyll: Reading time without plugins"
date: 2015-01-19
draft: false
slug: jekyll-reading-time-without-plugins
city: Joinville
---

Estimated reading time of a post is a feature that became popular, I believe, with Medium.

There are plenty of Jekyll plugins that address this problem, but, if you want to deploy at GitHub Pages, you can't use those plugins (GitHub will run with the `--safe` flag).

So, I created a snipped of pure Liquid code to fix that.

So, the first thing we will want to do is get the word count. That's pretty, actually:

```
{% assign words = content | number_of_words %}
```

Now, we need to divide this number with something. This something is called *Word per minute* (WPM). 

According to [Wikipedia](http://en.wikipedia.org/wiki/Words_per_minute), an average person  can read 180 words per minute in a computer monitor. Now it became really easy to do the rest:

```
{{ words | divided_by:180 }} mins
```

But, what if the post has less than 180 words? Actually, even if it has more, 350 words, for instance, when divided by 180, will result in 1.94, Liquid will round it down to 1, so, the user will see "1 mins", which is weird. 

To fix that, we have to check if it has less than 360 words, because any number great or equal 360 will result in 2+ mins, which still plural. 

That said, the solution is quite simple:

```
{% raw %}
{% if words < 360 %}
  1 min
{% else %}
  {{ words | divided_by:180 }} mins
{% endif %}{% endraw %}
```

So, to keep it organized, I put all this in a `read_time.html` in my `_includes` folder:

```
{% raw %}
<span class="reading-time" title="Estimated read time">
  {% assign words = content | number_of_words %}
  {% if words < 360 %}
    1 min
  {% else %}
    {{ words | divided_by:180 }} mins
  {% endif %}
</span>

{% endraw %}
```

And then I just `include` it in my `post` layout:

```
{% include read_time.html %}
```

And it works, as you can see here. Hope you like!
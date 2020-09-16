---
title: "UP v2"
date: 2013-10-14
draft: false
slug: up-2
city: Joinville
toc: true
tags: [jekyll, blog]
---

So, this is the all-new [UP](http://github.com/caarlos0/up) version. It is more based on my own [blog](https://carlosbecker.com/) in some way, which is a some kind of branch of the theme.

![](/public/images/up-2/7e58d31e-4f3d-4a4c-b2d3-56face03b8a7.png)

The biggest changes are:

- Assets are now managed by [bower](http://bower.io/);
- Build are now made by [grunt](http://gruntjs.com/);
- `Procfile.dev` is provided to use with `foreman` to watch changes everywhere (assets, pages, posts);
- Removed the hardcoded Jekyll from Gemfile and replaced it with the github pages Gem;
- Bumped all assets versions.

There are also some minor changes, like a new `pygments.less` and other few and small changes.

### Usage

The best way is to read the [project readme](http://github.com/caarlos0/up), but, to get started, after clone the repository and with Ruby, Bundler, Node.js and NPM already in your `PATH`, you can run the init script:

```
$ ./script/init.sh
```

It will install all needed dependencies and build the assets for you (with the `grunt` command).

You can also start it up in watch mode with:

```
$ foreman start -f Procfile.dev
```
### Update

To update your previous UP-baked blog, the easiest way is to backup your `_posts` folder and override everything else. If you don't want to do that, you will have to hack around the changes and do it by hand.

If you have any issues or suggestions, ping me at the [up project](http://github.com/caarlos0/up) or at the comments box bellow.

Have fun!

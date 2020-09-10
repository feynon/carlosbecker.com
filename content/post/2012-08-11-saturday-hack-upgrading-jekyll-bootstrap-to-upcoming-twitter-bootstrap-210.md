---
title: "Upgrading Jekyll Bootstrap"
date: 2012-08-11
draft: false
slug: saturday-hack-upgrading-jekyll-bootstrap-to-upcoming-twitter-bootstrap-210
city: Joinville
---

As you may know, this blog uses [Jekyll Bootstrap](http://jekyllbootstrap.com/) as engine. Yeah, Jekyll Bootstrap uses Twitter-Bootstrap, but, it uses the version 2.0.x (mine is 2.0.4, don't remember if I update or if it hsa like this since the begin of times), and the version 2.1.0 is comming! Like you can see in the @[twbootstrap](https://twitter.com/twbootstrap) twitter profile, 2.1.0 will be released a few days!

[https://twitter.com/getbootstrap/status/230470118562484224?ref_src=twsrc%5Etfw](https://twitter.com/getbootstrap/status/230470118562484224?ref_src=twsrc%5Etfw)

HELL YEAH, finally a new version!

So, today is saturday and I has nothing to do that could made the world a better place. Let's do this!

I'll describe here my ‘steps' to update jekyll to the new version!

### No more talks, let's do this!

First of all, you will need the Node.js Package Manager (NPM), so, install nodejs! In my archlinux box, it was simple like that:

```
$ pacman -Sy nodejs
```

Then, clone the bootstrap repo somewhere:

```
$ git clone https://github.com/twitter/bootstrap.git
$ cd bootstrap
```

Also, let's use the correct branch!

```
$ git checkout 2.1.0-wip
```

Ok! Now, we have to install the dependencies to build bootstrap. According to bootstrap github project, we will need recess (a less compiler), uglify-js, and jshint. So, run this:

```
$ sudo npm install recess uglify-js jshint -g
```

Now, the big one! Let's make it!

```
$ make
```

It should provide you a output like this:

```
[carlos@charmander bootstrap]$ make
\n##################################################
Building Bootstrap...
##################################################\n
Running JSHint on javascript...             \033[32m✔\033[39m Done
Compiling LESS with Recess...               \033[32m✔\033[39m Done
Compiling documentation...                  \033[32m✔\033[39m Done
Compiling and minifying javascript...       \033[32m✔\033[39m Done
\n##################################################
Bootstrap successfully built at 09:51PM.
##################################################\n
Thanks for using Bootstrap,
<3 @mdo and @fat\n
```

Hell yeah folks! It's much more easy than I think! So now, we have to upgrade Jekyll! Let's go?

### Upgrading Jekyll

The compiled and minified files will be at `**bootstrap home**/docs/assets`, and we will have to put them in `**jekyll home**/assets/themes/twitter/bootstrap`.

So, will made it easy for you :)

### css

- bootstrap-responsive.css
- bootstrap.css

### js

- bootstrap.min.js

### img

- glyphicons-halflings-white.png
- glyphicons-halflings.png

Place these files in the correct folder (like I said earlier). Now, check your `**jekyll home**/_includes/themes/twitter/default.html` for the imports of the css and js files, check if the filename is correct, etc.

Now, you should be able to run a `jekyll --server`, in my case it works just fine, but, my layout was modified a lot… don't sure if it works in a default jekyll. So, test it, and tell me what you got :)

I also do a lot of other hacks, using `docs.css` file as example, tweak footer, [about](https://carlosbecker.com/about/) page and other things…

Cheers!
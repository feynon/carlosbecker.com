---
date: 2012-11-28T00:00:00Z
slug: msfonts-archlinux
title: Easy-install MS Fonts in Arch Linux
draft: true
---

I just made a all-new fresh installation of archlinux a couple of days ago, and realized that my fonts are tottaly weird in the browser.

"Ok, surelly, I haven't installed msttcorefonts yet. Let's do that" I think. Sweet mistake.

Due to some license or something like this, they removed the fonts from AUR. Yeah, I know, the packages still there, but, you will need a Windows DVD, hand-copy the files to some folder, and then install the package.

> brotip: I have no time for this kind of bullshit.

So, I just made a simple and small script to download and install it. You can get it in [my github](http://github.com/caarlos0/msfonts), or, simply run:

```console
$ curl -L https://github.com/caarlos0/msfonts/raw/master/install.sh | sudo sh
```

The script will install git (if you doesn't have it in your path), download fonts to /temp, copy it to the right place and update all font-cache's that I can remember.

Yep, just that simple.

Hope it make your life easier.

Cheers.

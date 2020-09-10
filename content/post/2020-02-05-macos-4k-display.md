---
title: "4K display on MacOS: the saga"
date: 2020-02-05
draft: false
slug: macos-4k-display
city: Joinville
---

I recently got a 4K display, and it didn't work as I expected on my MacBook Pro. This is what I tried, what worked and what didn't.

<!--more-->

The display I got is the [Dell U2718Q](https://amzn.to/3aNRbDb). Its a 27", 4K, IPS display. And its pretty good.

![The display I got; Photo from Amazon.com](/public/images/macos-4k-display/b1db7860-9f52-4fcb-af25-286c3a6b61b3.png)

I wanted a LG Ultrafine 5K, of course. A single USB-C cable for both video and charging the Mac was pretty appealing. A small issue though: I live in Brazil, they don't sell it here, and if they would, it would cost like 10k BRL. It is just too much money for a display.

> FWIW: This one I got was already nearly 3k BRL, which is also a lot of money. Point being electronics are expensive af here.

Without further ado, let's get to it!

## The obvious stuff

First try was the obvious try - as it should be. 

I already had an [USB-C → HDMI, USB-C and USB-A adapter (the official from Apple)](https://www.apple.com/shop/product/MUF82AM/A/usb-c-digital-av-multiport-adapter), and just hooked up the new display hoping everything would work.

And... I got the 4K resolution! 

But this post wouldn't be any fun if this was it. 

Somethings went wrong, lets check them out...

## Lag

While the image was OK, I got the feeling it was laggy. No one deserves lag. 

Looking at the system report, we can get a pretty good idea about what the problem is:

![30 Hz instead of 60 :(](/public/images/macos-4k-display/fdab06f3-c27b-44ff-8593-030b71c893b8.png)

30 Hz. Fuck. It should be 60!

I was pretty sure my HDMI cable supported 4K @ 60Hz and it didn't had any text on it saying anything (thanks China), so I decided to look into the [adapter](https://www.apple.com/shop/product/MUF82AM/A/usb-c-digital-av-multiport-adapter) specs to rule it out, and, sure, Apple is selling poor adapters:

> 1080p at 60Hz or UHD (3840 by 2160) at 30Hz on: ...

I already hated this adapter before because, I shit you not, it causes interference on my headphones.

I have to constantly adjust it a little bit bent so it interferes less. 

You can actually find several people doing hacks like this or using aluminum foil and etc to prevent it on YouTube... with several different adapters...

Because of that, I had already ordered [an adapter from Dell](https://www.dell.com/pt-br/shop/accessories/apd/470-abmz) together with display, which, in theory, has enough bandwidth for 4K at 60Hz. But... it hasn't arrived yet (almost 3 weeks later).

And given the fact that the bluetooth interference seems to happen with several different adapters, I am pretty sure it will happen with this one too - and maybe it won't even fix the issue anyway.

That said, I tried an USB-C → miniDisplayPort cable. Notice it is not an adapter. It's a cable.

The cable arrived, I tried it in and it solved all my problems:

![@ 60 Hz!!!!!!](/public/images/macos-4k-display/3ab55fc5-3178-4aea-a46e-f27ec626cf38.png)

The only problem was this:

[IMG_6338-60d2a4e9-3fad-4b12-be3c-0098b7a82a09.mov](IMG_6338-60d2a4e9-3fad-4b12-be3c-0098b7a82a09.mov)

Flickering much?

Looking around on the internet, I found [this post](https://hackintosher.com/blog/bad-hdmi-dp-cable-can-ruin-4k-hackintosh-flickering/) of someone saying they had the same issues - but on a Hackintosh. The cable is from another brand but looks **exactly the same** as mine. Probably from China, probably lying about their bandwidth... which would be a surprise to a grand total of 0 people.

This flickering is the same thing that happens if I try to force 4K @ 60Hz on the current poor adapter + HDMI setup (you can do that with `swithresx`).

I returned it and got my money back.

## RGB

The problem here is related to a really old macOS bug. It thinks the display is a TV (for some reason), and use "TV-like colors". 

By TV-like colors I mean [YPbPR](https://en.wikipedia.org/wiki/YPbPr), which is an analog color space commonly used by DVD players and such. If you are from the 90's you'll remember this:

![YPbPR cables.](/public/images/macos-4k-display/7fccd368-6d30-45bc-82ac-81e6c8aa9700.png)

Anyway, the display uses that instead of RGB, colors looks "weird", and if I try to force RGB on the display settings I get this beautiful thing:

![Everything looks pretty pinky...](/public/images/macos-4k-display/e7cd238d-514e-4017-8a46-b7bc125d4468.png)

There is this [well known workaround](https://www.mathewinkson.com/2013/03/force-rgb-mode-in-mac-os-x-to-fix-the-picture-quality-of-an-external-monitor). I even tried it out, but on my case it only make things worse: I still got the wrong colors, and now also lost the 4K resolution.

## Finally, a good a cable

I talked to a [friend](https://github.com/marcosnils) (and also read on several places) that *Cable Matters* cables are good. So I got an [USB-C → DisplayPort cable from them](https://amzn.to/394xZiG).

When it arrived, I just plugged it in, and, like magic, **everything works**!

![4K @ 60Hz, no lag!!!!!!](/public/images/macos-4k-display/cd085cf5-0f88-4993-923b-e1e7d24a109c.png)

This also fixed the color problems:

![RGB!!!!!!!!!!!](/public/images/macos-4k-display/89866115-ca77-4c57-9f5f-ff26d70f11c0.png)

To be fair, the poor miniDP cable also fixed the colors issue.

It seems to me that the macOS algorithm that guesses which color scheme to use is more or less like the following pseudocode:

```
if (isHDMI() && isBigScreen() && isSupportYPbPR()) {
	// obviously a TV!!! use TV colors!!!
	setColors("YPbPR");
}
```

Basically, using **anything but HDMI** seems to solve the color problem.

And, obviously, you can't just change the color scheme you want. Apple just knows better.

## Conclusion

If you want to use 4K @ 60Hz on your Mac, be prepared to solve problems the Apple way: **throwing more money at them. 💸**

Do not waste time with cheap cables and USB-C → HDMI adapters. Its just not worth it. 

Get a [Cable Matters USB-C → DP cable](https://amzn.to/394xZiG), plug it in and profit.

Everything works as it should now, with the very first setup I had. Thanks Apple for making my life more difficult than it needs to be.
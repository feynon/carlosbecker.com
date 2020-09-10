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

![](asset-a3ca84d6-841e-43b9-9740-a3fe3d23b437.png)

The display I got; Photo from Amazon.com

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

![](asset-1-7bfac567-aa9d-4661-a3a7-2b3d6d2e0ad0.png)

30 Hz instead of 60 :(

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

![](asset-2-e9e7069c-61f3-47f3-a0f0-41afe5d69c50.png)

@ 60 Hz!!!!!!

The only problem was this:

[IMG_6338-60d2a4e9-3fad-4b12-be3c-0098b7a82a09.mov](IMG_6338-60d2a4e9-3fad-4b12-be3c-0098b7a82a09.mov)

Flickering much?

Looking around on the internet, I found [this post](https://hackintosher.com/blog/bad-hdmi-dp-cable-can-ruin-4k-hackintosh-flickering/) of someone saying they had the same issues - but on a Hackintosh. The cable is from another brand but looks **exactly the same** as mine. Probably from China, probably lying about their bandwidth... which would be a surprise to a grand total of 0 people.

This flickering is the same thing that happens if I try to force 4K @ 60Hz on the current poor adapter + HDMI setup (you can do that with `swithresx`).

I returned it and got my money back.

## RGB

The problem here is related to a really old macOS bug. It thinks the display is a TV (for some reason), and use "TV-like colors". 

By TV-like colors I mean [YPbPR](https://en.wikipedia.org/wiki/YPbPr), which is an analog color space commonly used by DVD players and such. If you are from the 90's you'll remember this:

![](asset-3-65a1bdb8-3be6-4139-afff-75418bf7fa3b.png)

YPbPR cables.

Anyway, the display uses that instead of RGB, colors looks "weird", and if I try to force RGB on the display settings I get this beautiful thing:

![](IMG_6334_(1)-28c3f911-341f-4569-ba8b-79ed175b26df.png)

Everything looks pretty pinky...

There is this [well known workaround](https://www.mathewinkson.com/2013/03/force-rgb-mode-in-mac-os-x-to-fix-the-picture-quality-of-an-external-monitor). I even tried it out, but on my case it only make things worse: I still got the wrong colors, and now also lost the 4K resolution.

## Finally, a good a cable

I talked to a [friend](https://github.com/marcosnils) (and also read on several places) that *Cable Matters* cables are good. So I got an [USB-C → DisplayPort cable from them](https://amzn.to/394xZiG).

When it arrived, I just plugged it in, and, like magic, **everything works**!

![](asset-5-ae328f67-4e1e-4fac-882f-6718cd66dc86.png)

4K @ 60Hz, no lag!!!!!!

This also fixed the color problems:

![](IMG_6357_(1)-a38641d0-fc59-4b7b-ae76-c6da311cb858.png)

RGB!!!!!!!!!!!

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
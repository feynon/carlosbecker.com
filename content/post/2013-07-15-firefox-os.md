---
date: 2013-07-15T00:00:00Z
slug: firefox-os
title: I just got a Firefox OS powered phone
---

So, Mozilla just gave me the [Geeksphone Keon][keon] with Firefox OS Developer
Preview, and I decided to drop here some lines about it.

First of all, as the name says, it's a Developer Preview. So, a lot of things
are buggy, others just doesn't exist yet and others of them I don't even know
if it's relative to hardware or software. I will not detail much,
it is basically what other people already said in the geeksphone's forum.

### The phone

The hardware, as you may have seen, is not that good:

> CPU Qualcomm Snapdragon S1 7225AB 1Ghz.

> UMTS 850/1900/2100 (3G HSPA).

> GSM 850/900/1800/1900 (2G EDGE).

> Screen 3.5" HVGA Multitouch.

> Camera 3 MP.

> 4 GB (ROM) and 512 MB (RAM).

> MicroSD, Wifi N, Bluetooth 2.1 EDR, Radio FM, Light & Prox. Sensor, G-Sensor,
> GPS, MicroUSB.

> Battery 1580 mAh.

I also found the touchscreen very hard to use. Sometimes the OS freezes (lock
and unlock the screen generally "unfreezes") without any apparent reason.. most
of times when I'm typing..

Another thing is that it has not a "back" button, and some apps are not ready
for that, which ends up difficulting the usage.

But, as I said before, it's a developer preview, so, its expected to have bugs
and issues of all sort.

Also, it's a pretty low-cost phone, despite the fact that in Brazil everything
is expensive due to a lot of taxes. For example, the Keon Phone costs USD 100
in EUA. To bring it to Brazil, it will cost you plus USD 100 (for transport)
plus USD ~180 (taxes), someting like USD ~400 (I'm not sure about the values,
but it's something like that). USD 400 is at least BRL 850, which is just too
much expensive. Anyways, the prices may change after some deals happens with
brazilian companies and Mozilla.

### The OS

Despite the fact that it doesn't have a calculator yet (not offically, at
least), it works as expected. Twitter and Facebook, for example, are just the
webapps that you can access from your phone at any given time. It works, but
is not that fast.

Anyway, you can call people and send messages just fine (see ya, ubuntu). There
are more configurations that I was expecting, and you can connect it through
`adb` to push apps and etc.

### Building an App

I just starting to fix Wealcash to work with Firefox OS. The work I
had for Firefox OS specifically was just to create the `manifest.webapp` and
change the nginx `mime.types` file to deal with it. The rest of the work, is
just about do the responsiveness right. That's the cool thing about Firefox OS:
you made your app responsive, it just works out of the box.

I made the needed file and the changes to the server in a couple of minutes,
and submited the app for approval...

![approved](https://pbs.twimg.com/media/BPKEcy9CMAAGdAu.png:large)

I figured it would take while to the app to be approved, since it was in the
position 62 in the queue... well, I was wrong.. ~20 minutes later it was
approved, and I just installed it in my phone:

![phone](https://pbs.twimg.com/media/BPKMNZaCUAAeOBl.png:large)

Surely, it has a lot that still need to be done. Lots of places are just
unusable due to the large amount of information, things that I don't figure out
how to fix, yet, but it is usable in its basics.

In the end, I believe that Firefox OS has much more chances to win that Ubuntu
Touch, but I'm not sure if it could beat Android or iOS... I believe that they
probably don't even want to. While most of platforms try to "lock" you in their
ecosystem, Firefox OS is just open, and doesn't not demand you to learn an
specific language/framework to build an app, just vanilla HTML, JS and CSS.

[keon]: http://www.geeksphone.com/

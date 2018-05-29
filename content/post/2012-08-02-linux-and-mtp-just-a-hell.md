---
date: 2012-08-02T00:00:00Z
slug: linux-and-mtp-just-a-hell
title: 'Linux and MTP: Just a hell'
---

Recently I bought a Samsung Galaxy X phone (brazilian Nexus). Yeah, it is a
pretty cool device, fast, etc etc... oh, yep, a little expensive too...

I was pretty happy with it, until I tried to copy some files to it
(basically, music). Then, I discover that nexus android force you to use
something called
[MTP (Media Transfer Protocol)](http://en.wikipedia.org/wiki/Media_Transfer_Protocol).
According to them, this decision is to made android better, but, I believe
that force the user to use it, is not.

User who use Linux, Mac OS and Windows XP have to install additional
packages/softwares to put it to work. And the worst part is that installing
these additional softwares not solve the problem at all.

I end up with a tecnical easy solution for arch linux, after follow a
thousand tutorials.

## The **UGLY** solution

- Get
  [this package (android-udev)](https://aur.archlinux.org/packages/android-udev-git/)
  from aur. (it will install libmtpfs and all other deps)
- Install **gvfs-gphoto2** via pacman.
- Reboot

  **PS**: If you have yaourt, you can do all this with one command:

```console
$ yaourt -S android-udev gvfs-gphoto2 --noconfirm
```

Then, if you have some luck, you can connect you device via some media player
(I used [clementine](http://code.google.com/p/clementine-player/)) and copy at
least some music to your device.

Also, you can transfer files via bluetooth (hell slow) or via wireless
(using something like Airdroid).

I'll try to mantain this post updated with new info, if you have something to
share, please leave a comment.

See ya.

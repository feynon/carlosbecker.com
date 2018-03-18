---
date: 2012-07-31T00:00:00Z
slug: updating-arch-linux-from-a-core-install
title: Updating Arch Linux from a Core Install
---

Recently I tried to made a fresh archlinux installation, with my old iso
(don't remember exactly what version is)...
Yo, first, I installed it with the local mirror only, without any update.
The installation was ok, I rebooted, then try to update, using the arch linux
home help about the */lib turn into a symlink to /usr/lib*. **Sweet** mistake.
That crap just don't work.

In last times, we have a lot of big arch linux changes:

* [changes in rc.conf and crypttab](http://www.archlinux.org/news/changes-to-rcconf-and-crypttab/)
* [grub legacy no longer supported](http://www.archlinux.org/news/grub-legacy-no-longer-supported/)
* [/lib becomes a symlink](http://www.archlinux.org/news/the-lib-directory-becomes-a-symlink/)
* [filesystem upgrade that needs manual intervention](http://www.archlinux.org/news/filesystem-upgrade-manual-intervention-required-1/)
* [pacman-key](http://www.archlinux.org/news/having-pacman-verify-packages/)

All these itens have man pages helping you to pass throug the most errors...
but, the lib symlink made a lot of people angry **never use the force (-f) in
pacman, NEVER**.

So, I tried to reinstall, but, this time, using the remote mirrors (getting all
the new packages), and it was OK, until the bootloader install part. Seem like
the installer lost the GRUB config file somewhere...

Then, I do my research, and found a
[very helpful article by Allan McRae](http://allanmcrae.com/2012/07/updating-arch-linux-from-a-core-install/).
I will just "rewrite" it here with some thought, but you can also read the
original if you want.

Remember, if you get the
[new ISO from arch linux site](http://www.archlinux.org/download/),
everything gonna be OK without any hacks.

So, after arch installed and loaded, do the following commands:

```console
$ pacman -Sy
$ rm -rf /var/run /var/lock && pacman -Sf filesystem
```

**Reboot.**

```console
$ pacman -S tzdata
$ pacman -U http://pkgbuild.com/~allan/glibc-2.16.0-1-i686.pkg.tar.xz
$ rm /etc/profile.d/locale.sh
```

In this next accept, when promped to update pacman, say **NO**, and, when
prompted to any replacement, say **YES**.

```console
$ pacman -Su --ignore glibc
$ pacman -Su
```

**Reboot** again.

So, at this point, everything should work like a boss.
See ya.

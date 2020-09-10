---
title: "Enable KMS on i915 GraphicCard"
date: 2012-07-27
draft: false
slug: how-to-enable-kms-on-i915-graphiccard-archlinux
city: Joinville
---

Acccording to [Arch Linux wiki](https://wiki.archlinux.org/index.php/Kernel_Mode_Setting), Kernel Mode Setting (KMS) is a method for setting display resolution and depth in the kernel space rather than user space. KMS enables native resolution in the framebuffer and allows for instant console (tty) switching. KMS also enables newer technologies (such as DRI2) which will help reduce artifacts and increase 3D performance, even kernel space power-saving.

So, you can see this is a good thing, uh? I'll show you the steps needed to activate this for my GraphicCard.

## Check if KMS is active

If the following command return `1`, your KMS module is already active, and you can go and drink a coffee. Otherwise, goto the next steps.

```
$ cat /sys/module/i915/parameters/modeset
```

## Install required packages

Before continue, be sure that the following packages are installed:

```
xf86-video-fbdev
xf86-video-intel
xf86-video-vesa
```

You can install it in a archlinux box with

```
$ pacman -Sy xf86-video-{fbdev,intel,vesa}
```

## Edit files

You will need to change several files. Starting with `/etc/modprobe.d/modprobe.conf`. Add (or create the file, if it doesn't exist) the following:

```
options i915 modeset=1
```

Also, edit `/etc/mkinitcpio.conf`.

- Edit the `MODULES` line, like in `MODULES="intel_agp i915"`
- Uncoment the line `FILES="/etc/modprobe.d/modprobe.conf"` and comment the `FILES=""` if it exists.

Create the `/etc/X11/xorg.conf.d/20-intel.conf` file with the following content:

```
Section "Device"
    Identifier "card0"
    Driver "intel"
    VendorName  "Intel Corporation"
    BoardName   "Intel Corporation N10 Family Integrated Graphics Controller"
    BusID       "PCI:0:2:0"
    Option      "SwapbuffersWait"    "false"
EndSection
```

## Regenerate initframs

Now, we can regen initframs with

```
$ sudo mkinitcpio -p linux
```

## Reboot

Now, reboot your system, and when it came up, run the first command again, and see if it returns `1` this time.

Also, this is my `lspci`:

```
$ lspci | grep -i vga
	00:02.0 VGA compatible controller: Intel Corporation Core Processor Integrated Graphics Controller (rev 18)
```

That's it. See ya.
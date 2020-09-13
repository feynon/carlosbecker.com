---
title: "Install JDK on OSX Yosemite"
date: 2014-08-18
draft: false
slug: jdk-yosemite
city: Joinville
toc: true
tags: []
---

For some reason, Oracle blocked the installers to run only on a fixed OSX version range with a nice and explanatory error message. This range doesn't include Yosemite, which makes sense, since nobody running Yosemite will ever want to write some Java. Anyway, here is how to fix it.

First, download and open the JDK `.dmg` file. Then, unpackage and edit the `Distribution` file:

```
$ pkgutil --expand "/Volumes/JDK 7 Update 67/JDK 7 Update 67.pkg" /tmp/jdk7.unpkg
$ cd /tmp/jdk7.unpkg
$ vim Distribution
```

> PS: I'm using vim, but you can use whatever editor you want, since it is vim.

Replace this function:

```c
function pm_install_check() {
  if (!(checkForMacOSX("10.7.3") == true)) {
    my.result.title = "OS X Lion required";
    my.result.message =
      "This Installer is supported only on OS X 10.7.3 or Later.";
    my.result.type = "Fatal";
    return false;
  }
  return true;
}
```

with this:

```c
function pm_install_check() {
  return true;
}
```

> just remove that if statement.

Then, just package and run the installer with:

```
$ pkgutil --flatten /tmp/jdk7.unpkg /tmp/jdk7.pkg
$ open /tmp/jdk7.pkg
```

The JDKs will be installed in the `/Library/Java/JavaVirtualMachines/` folder.

You can also clean up your mess:

```
$ rm -rf /tmp/jdk7.*pkg
```

You can use this tip for JDK 8 too.

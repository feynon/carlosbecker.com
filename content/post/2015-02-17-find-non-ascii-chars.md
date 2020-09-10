---
title: "Find non-ASCII chars"
date: 2015-02-17
draft: false
slug: find-non-ascii-chars
city: Joinville
---

> FYI: This is a really quick tip.

Sometimes things break because of random characters, like `'`. Those kinds of characters may break, for example, [shellcheck](http://www.shellcheck.net/) and other tools based on [hlint](https://github.com/ndmitchell/hlint):

```
hGetContents: invalid argument (invalid byte sequence)
```

Also, who likes those stupid chars? Nobody does!

### Fixing that

The first step is to find them. I had some of them in my dotfiles project, as you can see in this [pull request](https://github.com/caarlos0/dotfiles/pull/36).

So, using `grep`, I created the [nonascii function](https://github.com/caarlos0/dotfiles/commit/da1bfe4d895aad8efc9ba79cac46e2b545514576): 

```shell
#!/bin/bash
nonascii() {
 LANG=C grep --color=always '[^ -~]\+';
}
```

The usage is simple:

```
$ cat osx/set-defaults.sh | nonascii
# Don't animate opening applications from the Dock
# Don't prompt for confirmation before downloading
# Add the keyboard shortcut ⌘ + Enter to send an email in Mail.app
# Disable smart quotes as it's annoying for messages that contain code
# Don't automatically rearrange Spaces based on most recent use
# Disable smart quotes and smart dashes as they're annoying when typing code
# Disable the "Are you sure you want to open this application?" dialog
# Remove duplicates in the "Open With" menu
```

> it won't be shown here, but all nonascii chars are in red.

Fix it now is straightforward, but, without this list, it would take a big amount of time to read the entire file paying attention to this kind of detail.
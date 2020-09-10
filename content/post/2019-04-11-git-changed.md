---
title: "Git: check if a folder changed"
date: 2019-04-11
draft: false
slug: git-changed
city: Joinville
---

Often I need to "do X only if files on some folder changed" or whatever. I
always need to Google that or find it on old scripts...

This is a quick post for me to find on Google when I need it again and think
"oh its me!".

Anyway, let's get into it!

The most cleanest way I found to do it is the following:

```shell
git diff --quiet HEAD $REF -- $DIR || echo changed
```

Note that `git diff --quiet` will **exit 1** when there **are changes**.
At first I though it was confusing, but it makes sense if you think
*"if no changes = exit 0, otherwise = exit 1"*.

So, let's say you want to check if your current branch has changes in the
folder `foo` when compared to `master`:

```shell
git diff --quiet HEAD master -- foo || echo changed
```

You can also do that comparing with the previous tag, for example:

```shell
git diff --quiet HEAD "$(git describe --tags --abbrev=0 HEAD)" -- foo || echo changed
```

You can use this to "deploy only changed folders" or "lint only changed files"
or whatever else you can come up with.

Anyway, just a quick post/note.

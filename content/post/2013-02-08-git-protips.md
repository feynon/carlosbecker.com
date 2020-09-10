---
title: "Git Protips™"
date: 2013-02-08
draft: false
slug: git-protips
city: Joinville
---

Git has a lot of features, and I bet that 90% of who use it (including me) doesn't know half of them. Well, maybe, someday, one of those *"unknown features"* can *"save your life"*.

So, I decided to compile a small list of useful git commands, tips and tricks that I use or used sometime in my work.

Of course, if you have yours tricks/tips (and you probably have), comment above the post or ping me at twitter/mail/anything, I'll be glad to add your tip and put your name in it.

Let's do this.

---

### Do not rebase commits that have already been pushed to the remote repository

[Read this](http://git-scm.com/book/en/Git-Branching-Rebasing#The-Perils-of-Rebasing).

### See who and in which commit each line of a file was changed last time

```
$ git blame path/to/file
```

### See the reference log

This will show the log of your local operations in the git tree, useful to get revision code of a specific action (see next item).

```
$ git reflog
```

**HEADS UP**: `reflog` will only have the logs of actions in your local repository, more specifically, it will log each time the `HEAD` pointer changes, so, if clone a repository right now, the reflog will be empty. Also, the reflog is maintained for 30 days by default.

### Get revision code for anything

```
$ git rev-parse HEAD # last commit
$ git rev-parse HEAD~5 # 5 commits back from last commit
$ git rev-parse develop # last commit from develop branch
```

### Revert a specific file to a specific commit

```
$ git log path/to/file # to check the revision code
$ git checkout revision_code path/to/file
```

### Revert a specific file to last commit

```
$ git checkout HEAD path/to/file
```

### Revert a specific file to 3 commits back from last

```
$ git checkout HEAD~3 path/to/file
```

### Get an old version of some file

```
$ git show HEAD~3:path/to/file > path/to/file_3_commits_ago
```

### Shallow clone

According to the docs:

> Create a shallow clone with a history truncated to the specified number of revisions. A shallow repository has a number of limitations (you cannot clone or fetch from it, nor push from nor into it), but is adequate if you are only interested in the recent history of a large project with a long history, and would want to send in fixes as patches.

So, if you want only the last revision:

```
$ git clone https://github.com/twitter/bootstrap --depth 1
```

This can drastically reduce the clone. The bad thing is that you will have NO HISTORY BEFORE THE HEAD WHEN YOU DO THE CLONE.

### Move a recent commit to another branch

This is pretty useful when you're working on something, commit, and them realize that it would fit better in another branch (or that you're in the wrong branch)… this is how to move it to another branch.

```
$ git branch new
$ git reset --hard HEAD~1 #go back 1 commit, YOU WILL LOST UNCOMMITED CHANGES
$ git checkout new
```

### Revert your entire tree to the last commit state

by [@luizkowalski](http://github.com/luizkowalski)

Useful when you do some crap and want to throw it all away.

```
$ git reset --hard HEAD
```

You can also specify something like `HEAD~3` to get back 3 commits from HEAD.

### Remove files that have not been added to staging area

by [@luizkowalski](http://github.com/luizkowalski)

```
$ git clean -df
```

### Stashing

by [@thiagolenz](http://github.com/thiagolenz)

This is useful when you want to switch branches, but don't want to commit a half-done work just to get back to it later.

```
$ git stash # save the current state
# later...
$ git stash pop # apply the stash to your current HEAD and remove it from your stack.
```

More info can be found [here](http://git-scm.com/book/en/Git-Tools-Stashing).

### Merge acting like it's an unique branch

by [@thiagolenz](http://github.com/thiagolenz)

This will do the merge without creating the "merge commit", acting much like SVN.

```
$ git pull --rebase
```

### Aliases

by [@derekstavis](http://github.com/derekstavis)

You can setup aliases in git, so get smaller commands and save time. Examples:

```
# configuring...
$ git config --global alias.co checkout
$ git config --global alias.s status

# using
$ git s
$ git co https://github.com/caarlos0/up
```

### Copy a file from one branch to another

Useful if you want a file that was introduced or modified in other branch. Example:

```
$ git checkout gh-pages
$ git checkout master -- teste.js
$ git commit -m "Update test.js from master"
```

### Log deleted files

by [@ricardo_walter](https://twitter.com/ricardo_walter)

Show the commit log with the deleted files for each commit:

```
$ git log --diff-filter=D --summary
```

### Delete a branch

Local:

```
$ git branch -d name
$ git branch -D name # force delete unmerged branch
```

Remote:

by [@ricardo_walter](https://twitter.com/ricardo_walter)

```
$ git push origin :branchname
```

### Undo your last commit

by [@antonini](http://github.com/antonini)

If you do a wrong commit, revert it, but maintain the changes in your staging area.

```
$ git reset --soft HEAD^
```

Just remember: `HEAD^` is a pointer to the parent of current `HEAD`, so, you can use `HEAD~3` to go 3 commits back, for example. Also, check the `reflog` part of this post.

### Remove a file from the last commit

PROTIP: Don't do this if you already pushed the commit.

For example, if you want to remove the `wrongfile.txt`, do:

```
$ git rm wrongfile.txt
$ git commit --amend
```

### Update submodules to last commit

Instead of `cd`ing each module, `checkout master` and `pull`, in git 1.8.2 you can simply do this:

```
$ git submodule update --remote --merge
```

### Apply a patch to another file

by [@aureliojargas](http://github.com/aureliojargas)

This is useful if you have a file in two different repos for whatever reason, or if you have 2 very similar files in the same repo:

```
$ git diff HEAD^ -- hello.test > /tmp/patch
$ patch -p1 hi.test /tmp/patch
```

## Troubleshooting

### Unable to update local ref

```
! c15fe9e..eee5f38  feature1 -> origin/feature1  (unable to update local ref)`
```

This error can vary by remote and branch name. Usually, a clean up will do the job just fine:

```
$ git gc --prune=now
$ git remote -v | cut -f1 | uniq | xargs git remote prune
```

Done.

## Going Further

I **strongly** recommend the reading of the "[Pro Git](http://amzn.to/15EiEwD)" book. The book will teach tons of things: basics, importing SVN projects, advanced, branching, servers, and git internals.

---

***Let's make this list bigger! Have your own tip/trick? Something I forgot to add? Share it with us!***
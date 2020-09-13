---
title: "Dotfiles Are Meant to Be Forked"
date: 2012-11-23
draft: false
slug: dotfiles-are-meant-to-be-forked
city: Joinville
toc: true
tags: []
---

![](/public/images/dotfiles-are-meant-to-be-forked/16090780-b364-47e0-8052-6c004a9bac20.png)

Well, it has been a while since I replaced my old (but gold) bash by the great zsh.

Anyway, I have my personal computer and my job computer, and, like almost every developer, I create aliases and scripts for everything I have to do repeatedly.

![](/public/images/dotfiles-are-meant-to-be-forked/3c1791a6-8f8c-4ba0-b1b9-fbb2a33ab316.png)

Well... you can imagine.. my `.bashrc` had about 300 lines. It was really big. Almost impossible to share with others, full of personal data, full of my machine specific data, bloated with old things I left behind... well, it was a real mess.

Then, I decided to make a huge step in my life: move to [ZSH](http://www.zsh.org/)!

## ZSH

> Zsh is a powerful shell that operates as both an interactive shell and as a scripting language interpreter. While being compatible with Bash (not by default, only if you issue "emulate sh"), it offers many advantages such as: Faster, Improved tab completion, Improved globbing, Improved array handling, Fully customizable.
> 
> —- [Arch Wiki about ZSH](https://wiki.archlinux.org/index.php/Zsh)

But, well, I didn't knew where to start. So, I forked [oh-my-zsh][ohmyzsh] project. I've used it for a while, also did some contributions... but well, it has so many things I didn't use, and it doesn't had a simple way to share configuration files across computers.

Then I found [holman's dotfiles](http://github.com/holman/dotfiles). And it was perfect! Except for the fact that it was full of Mac OS X-related stuff.

So I tweaked, removed, tweaked, cleaned-up, tweaked, etc etc, and there it is, my all-new [dotfiles](http://github.com/caarlos0/dotfiles)!

They should work with both Linux and OS X (Linux lacks a little automation, tough), and I tweaked it to fit my taste.

## so get it, bro

[They are available on github](http://github.com/caarlos0/dotfiles)!

The installation is pretty straightforward. Just clone it in `~/.dotfiles` and run the `script/bootstrap` file. If you found any error, please, open an issue so I can fix them.

Also, take a look at the [readme](https://github.com/caarlos0/dotfiles#install) (pretty simple), it will make it easy to you to understand the topics and other features.

If you want, you can also read the [holman's post about his dotfiles](http://zachholman.com/2010/08/dotfiles-are-meant-to-be-forked/).

If you wanna contribute with anything, just issue a pull-request. I'll be glad to take a look at it!

That's all folks, hope to see you soon!

<!--more-->

> This post was last updated in Feb 4, 2015.

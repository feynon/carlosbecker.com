---
title: "Improving Vim and Gedit"
date: 2012-08-01
draft: false
slug: improving-vim-and-gedit
city: Joinville
---

I just started to learn Ruby and Ruby on Rails, and, as I read somewhere, most Ruby/RoR developers don't like IDEs to do their work.

Curious, I tried to use some IDEs, like IntelliJ, Netbeans and others… and really, it just slow you. A good text editor is incredible better. For example, in my VIm config, I could simply type `def[TAB]` and VIm create for me:

```
def method

end
```

with `method` already selected. So, I type the name, parameters, press `[TAB]` again, and I inside the method, then I just write it. It's really good. Also, Gedit have the same functionality, but, with the `[CTRL]+[SPACE]` combo.

## Gedit powered by GMate

GMate is a set of plugins that try to ‘transform' gedit in a ' textmate for linux'.

The installation is pretty simple:

```
$ git clone git://github.com/gmate/gmate.git
cd gmate
./install.sh
```

Then close and open your Gedit, go to settings, and take a look at themes and plugins. About themes, I **strongly** recommend *Molokai*. For me, this is the best theme ever.

## Vim powered by a gorgeus vimfiles setup

This ‘*vimfiles*' project is nothing more than a lot of plugins and some hacks in *vimrc* file, plus some themes, like Monokai (I already said thats the best theme ever!).

The installation is a little more complex than GMate, but, is not like build a Boing-747 :)

Just follow these steps:

### Clone the project from github:

```
$ git clone git://github.com/caarlos0/vimfiles.git ~/.vim
```

### Edit/Create your main `vimrc` file

```
$ vim ~/.vimrc
```

With the following content:

```
source ~/.vim/vimrc
colorscheme molokai
```

Then save and exit;

### Update the submodules:

```
$ cd ~/.vim
$ git submodule update --init
```

Done. Now, do your own tests and made your choice. I prefer VIm :)
---
date: 2013-03-16T00:00:00Z
slug: first-steps-with-mac-os-x-as-a-developer
title: First steps with Mac OS X as a Developer
---

So, I just bought my first Mac, and decided to wrote this in order to help
others =)

As a developer, the 3 apps I use most: Terminal, Text Editor and Music Player.
So, in the mac-brave-new-world, I particularly use, in order: The default
Terminal.app, SublimeText2 and iTunes.

<img src="/public/images/first-steps-osx.png" class="noshadow"
title="Mac OS X Terminal.app" alt="Mac OS X Terminal.app">

SublimeText has his own tips, and I won't write about that. The web is already
full of them. Go to the wild-web and find the best for you.

I will focus in terminal tools and other things. Also, I have a jocke about that:

> A man who calls himself as a developer says: "I'm a developer, I don't want
> to write commands, I just want to use some UI to do what I want to".

Yep, this is a joke. _(before it was sad)_

Well, let's stop this and go to what really matter.

## First things first

Install [XCode command line tools][1]. You will always need it anyway.

## Install Homebrew

[Homebrew][2] is some kind of _ports_ for Mac. As a developer, you should
install it:

```console
$ ruby -e "$(curl -fsSL https://raw.github.com/mxcl/homebrew/go)"
```

## Install some useful things via brew

```console
$ brew install grc coreutils spark z ack git
```

- grc is tool to colorize things. You can easily use it by sourcing
  `$(brew --prefix)/etc/grc.bashrc`;
- coreutils is kinda obvious;
- spark lets you echo bar charts in your terminal.. might be useful time to time;
- z shows you the most folders you access most and let's you easily access
  any of them;
- ack is kinda `grep -ril` but faster;
- git doesn't need any explanation;

## Install Homebrew Cask plugin

[Homebrew Cask][cask] let's you install normal apps with brew via command line.
You will almost-never have to manually download and install apps again.

```console
$ brew tap caskroom/homebrew-cask
$ brew install brew-cask
```

Then you can install some useful stuff with it:

```console
$ brew cask install caffeine dropbox iterm2 sequel-pro virtualbox vagrant the-unarchiver vlc google-chrome skype transmission dash cloudapp postgres divvy rdio github disk-inventory-x
```

Let's made a list:

- Caffeine: let's you prevent your mac from sleep;
- iterm2: better Terminal.app;
- sequel-pro: mysql/mariadb gui;
- vitualbox and vagrant: virtualization tools;
- the-unarchiver: extract everything;
- transmission: torrent client;
- dash: documentation visualizer (for almost every thing ever made);
- cloudapp: easy file sharing tool;
- postgres: the very simples postgresql for mac;
- divvy: tool to manage windows using keyboard;
- disk-inventory-x: tool to find files that are eating your hd.

The others are probably auto-explanatory.

[cask]: https://github.com/phinze/homebrew-cask

## Install zsh

[ZSH][3] is a pretty powerful shell for \*nix systems. As a developer, it just
changed my life. I don't even know how I lived before using it. No, really,
install it NOW:

```console
$ brew install zsh
$ chsh -s /bin/zsh
```

## Use some dotfiles

ZSH is pretty powerful and highly customizable. There are a lot of projects
around the web to achieve an easy start to it.

Some examples:

- [oh-my-zsh][4]
- [holman/dotfiles][5]
- [my dotfiles][7] (which actually do most of this stuff for me)

I've [already wrote about this before][8], in case you want to read something
about it.

In this example, let's install [my dotfiles for osx][7] (basically, the
holman's with some custom things):

```console
$ git clone  https://github.com/caarlos0/dotfiles ~/.dotfiles
$ cd ~/.dotfiles
$ script/bootstrap
$ source ~/.zshrc
$ dot_update # will install some tools and do some basic setup
```

And you should be ready to go.

## Install rbenv

[rbenv][9] is a ~lightweight~ ruby vm manager. Basically, it does the same thing
as RVM, but I found it a little bit less intrusive. If you want, you can install
it with brew:

```console
$ brew install rbenv ruby-build
$ rbenv install 2.0.0-p195
$ rbenv global 2.0.0-p195
$ rbenv rehash
```

Be sure to check your `~/.{zsh,bash}rc` file

## Install hub

[hub][10] is a github command line tool written in Ruby to improve your
git/github diary use. You can install it with `brew`:

```console
$ brew install hub
```

Example usage:

```console
$ hub clone caarlos0/up
```

Much less typing, IMHO =)

## Install RMagick gem without pain

So, in some projects I use the `rmagick` gem, and it got me some headache to
install.

Well, here the steps:

```console
$ brew install imagemagick
$ brew install pkg-config
$ C_INCLUDE_PATH=/usr/local/Cellar/imagemagick/6.8.0-10/include/ImageMagick gem install rmagick
```

And boom! It works =) Pretty tricky.

**Update:** The `C_INCLUDE` trick seems to not be needed anymore.

## Other tips:

- Everyone likes Emojis. Take this [cheat sheet][11];
- [Hidden OSX Features, tips and tricks][16];

If you want, you can also take a look at my [OSX Settings][15], which
is already available in my dotfiles (and you already have if you ran the
`dot` script).

As suggested in comments and by some friends:

- Install [iTerm2][12], a cool Terminal.app replacement;
- Install [Alfred][13], a produtivity tool.

---

Have your own tips? Share with the other fellows in comment box bellow.

Cheers!

[1]: https://developer.apple.com/devcenter/mac/index.action
[2]: http://mxcl.github.com/homebrew/
[3]: http://www.zsh.org/
[4]: https://github.com/robbyrussell/oh-my-zsh
[5]: https://github.com/holman/dotfiles/
[7]: https://github.com/caarlos0/dotfiles-osx
[8]: /posts/dotfiles-are-meant-to-be-forked/
[9]: https://github.com/sstephenson/rbenv/
[10]: https://github.com/defunkt/hub
[11]: http://www.emoji-cheat-sheet.com
[12]: http://www.iterm2.com/
[13]: http://www.alfredapp.com/
[14]: http://www.derlien.com/
[15]: https://github.com/caarlos0/dotfiles/blob/master/macos/set-defaults.sh
[16]: http://apple.stackexchange.com/questions/400/please-share-your-hidden-os-x-features-or-tips-and-tricks
[17]: http://www.irradiatedsoftware.com/sizeup/

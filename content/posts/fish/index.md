---
title: "Why I migrated to the Fish Shell"
date: 2020-12-07
draft: false
slug: fish
city: Cascavel
toc: true
tags: [shell]
---

Back in June, I started porting my dotfiles from ZSH to Fish. Here's why.

## Initial performance

Fish's performance is a lot better than ZSH's, and very similar to Bash.

We can verify that by firing up a container with limited CPU and memory, like so:

```sh
docker run --rm -it --cpus 1 --memory 100m ubuntu bash
```

And then we can install all the shells, as well as [hyperfine](https://github.com/sharkdp/hyperfine), and see how they go:

```sh
apt update
apt install zsh fish wget -y
wget https://github.com/sharkdp/hyperfine/releases/download/v1.11.0/hyperfine_1.11.0_amd64.deb
dpkg -i hyperfine_1.11.0_amd64.deb

for s in bash zsh fish; do 
	hyperfine --warmup 3 "$s -i -c 'exit 0'"
done
```

The results are **very impressive**:

{{< img caption="Fish is only a couple ms slower than Bash, but almost 4x faster than ZSH." src="febdcd3b-a9d3-42b8-acd0-9fffc8dacfba.png" >}}

I honestly never though ZSH was this slower compared to Bash... maybe it got slower over the years... I don't know, but once you use fish the first time, you'll notice it.

## Performance along the road

On both Bash and ZSH, you'll probably end up with a whole lot of `source something.sh`, which, once added up, can slow your shell quite a bit.

You also need to set your aliases, environment variables and etc, every time you start a new shell. While this is helpful to debug things, it is not very fast.

Fish allows you to set universal variables, which are shared across all shells and system restarts. You don't need to set them on every shell init, instead, you set them once (`set -U`) and they will be added to your `~/.config/fish/fish_variables`.

It also has lazy loading of completions and functions: you just put them in the right folders (`~/.config/fish/completions` and `~/.config/fish/functions`), and it loads them when you first try to use them, instead of every time you open a new shell.

This can all be confirmed by comparing my Fish setup's performance (it has some plugins, abbreaviations, completions, functions, etc) with empty ZSH and Bash:

{{< img caption="It gets a bit slower now, but still faster than an empty ZSH setup." src="ba95d0f7-5ea1-45a6-8f4e-abf36d74d06c.png" >}}

I open new shells literally hundreds of times a day, I don't want to waste those milliseconds.

Of course, performance is not the only unit of measure, otherwise I would use just plain Bash (or SH), but if I can get the same features faster... that matter a lot to me.

## Features

Some of the features that you'll need plugins on ZSH/Bash, and are native on Fish:

- Syntax Highlight
- Autosuggestions
- Man-page completions

Because of that, I only need a couple of plugins:

- `z` (`autojump`)
- `fzf` (for CTRL+R history search and other goodies)
- `grc` (colorize output of several commands)
- [lucid.fish](https://github.com/mattgreen/lucid.fish) (my prompt of choice)

Other than that, I just put a couple of files in their right folders, setup a bunch of abbreviations, and that's it.

I use [Fisher](https://github.com/jorgebucaran/fisher), because I don't really need the whole plugin framework thing from [oh-my-fish](https://github.com/oh-my-fish/oh-my-fish), but both of them are very good and you can choose whichever you like more to manage your plugins.

## Differences

Fish does not talk POSIX shell, so, some things are different.

### Scripting

Its common for me to script my way into the shell itself when dealing with issues, so, in the beginning, I struggled a bit. If you usually use just plain commands and pipes, it might not be an issue for you.

Things like `for` loops and `()` instead of `$()` took me a while to get used to, after I don't know how many years of writing things that way, but, once I got used to, it is actually simpler.

### Aliases

Fish doesn't have the concept of aliases, instead, you either wrap your command in a function (Fish has a function that does that for you, which is called `alias`), or use abbreviations.

Abbreviations expand once you type them, which is actually better in some ways:

- you can use an abbreviation similar to what you want, once it expands, edit it easily;
- you can copy-paste your terminal to someone else and they don't have to know what your 348 aliases do;
- your shell history will be saner (probably).

### argv

Instead of `$1`, `$2`, `$3` and etc, on Fish we have `$argv[1]`, `$argv[2]`, `$argv[3]`, etc.

It's not that big of a deal, but I always forget about it. 🤷‍♂️

### Setting variables

Instead of doing:

```sh
FOO="bar"
```

On Fish you need to do:

```sh
set FISH bar
```

The cool thing is that it has options to append and prepend to lists, so you don't need to do things like `PATH="/path/bin:$PATH` and can just do `set -p PATH /path/bin` instead.

It's worth mentioning that when running one liners, the plain shell syntax works just file, e.g.:

```sh
FOO=bar echo $FOO
```

## Getting started

The [Fish website](https://fishshell.com) has a lot of good docs, I suggest you start from there.

You can also take a look at how I manage my Fish [dotfiles](https://github.com/caarlos0/dotfiles.fish) for inspiration (or just fork and change to match your own taste).

You can also [play with it online](https://rootnroll.com/d/fish-shell/), without installing anything!

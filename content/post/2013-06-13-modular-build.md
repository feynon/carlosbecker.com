---
date: 2013-06-13T00:00:00Z
slug: modular-build
title: Automate your Builds
---

No one likes to do the same stuff lots of times. Copy a file from here to
somewhere, start a bunch of services, do a tricky build with maven... all this
shit are acceptable if you do one time and will never need to do it again. When
you do it two or three times in a little space of time, you will write a
script. Everyone does that. If you don't, you should.

I realized that most of this scripts are project specific. In my case, at my
daily job, I created a `scripts` folder inside project root folder, and start
coding a bunch of scripts inside it.

I end up with a lot of replicated code. Foolish things like exporting
`MAVEN_OPTS` and a lot of params to `mvn` executable. I could just put that
in my [dotfiles][dotfiles], but I it doesn't make sense, since other
projects will have other setup.

Another problem, is that lots of times I wanted to run a lot of scripts in one
line, for example:

```console
$ ./build.sh && ./publish.sh && ./start.sh
```

I hate doing this thought. So, I believed that creating a `Rakefile` would fit
good. But it doesn't make any sense to put Ruby code with Java code just for
this. Then, GNU Make looks a damn good option for me, and start hacking around
it.

In first versions, I just put all my code inside the `Makefile` itself, but it
starts to get big, and I decided to modularize.

So, I created the [modular-build project][modular-build] with some
[@merencia][lucas] contributions and ideas.

## modular-build

The base itself is not a big deal. It have just these files:

- **README.md**: Documentation, just as expected;
- **bootstrap.sh**: Instalation script;
- **Makefile**: The Makefile skelecton;
- **base.sh**: A base script that all other scripts will import. You can share
stuff here between scripts (environment variables, functions). It already came
with some funcions. You should take a look at them;
- **example.sh**: An clean example task;
- **newtask.sh**: A helper script to create tasks. `make newtask` will invoke
it, ask for the task name, then create the `.sh` file based on `example.sh` and
add it to `Makefile`.

The instalation is pretty easy, you can just call:

```console
$ wget -qO- http://git.io/63HVSg | bash
```

or, using curl:

```console
$ curl -s http://git.io/63HVSg | bash
```

After that, `make newtask` and start building your own scripts. Automate
everything, put computers to work to you, not the opposite.

You can find more info in the [project repo][modular-build] and also in
the [project site][modular-build].

Also, you can give contributions and opinions, here, or in project issues.

Happy hacking!

[modular-build]: https://github.com/caarlos0/modular-build
[lucas]: https://github.com/merencia
[dotfiles]: https://github.com/caarlos0/dotfiles

---
title: "Writing CLI applications with Golang"
date: 2017-04-13
draft: false
slug: golang-cli-apps
city: Joinville
toc: true
tags: [golang, cli]
---

Last few months I've been using Go to write quite a lot of tools. In this post I intent to show not why I chose Go over others, but how I architect those tools, what libraries I use and what kind of automation I have in place.

## Libraries

I'm sure some folks will preffer others, but, right now I'm using:

- [golang/dep](https://github.com/golang/dep) for dependency management;
- [caarlos0/spin](https://github.com/caarlos0/spin) for showing spinners when it makes sense;
- [urfave/cli](https://github.com/urfave/cli) as cli "framework";
- [stretchr/testify](https://github.com/stretchr/testify) for better testing;
- [alecthomas/gometalinter](https://github.com/alecthomas/gometalinter) for code linting;
- [goreleaser/goreleaser](https://github.com/goreleaser/goreleaser) for release automation.

At this point some will ask me why I built my own spinner library. The answer is: I didn't like the two or three libraries I tested, so I wrote my own in the way I wanted.

Folks may also argue that [spf13/cobra](https://github.com/spf13/cobra) is better than [urfave/cli](https://github.com/urfave/cli). 

Maybe it is, but, right now [urfave/cli](https://github.com/urfave/cli) is working well for me, so I didn't even tried anything else yet.

## Folder structure

A basic tree would look like:

```sh
.
├── Makefile
├── cmd
│   └── example
│       └── main.go
├── example.go
├── example_test.go
├── goreleaser.yml
├── Gopkg.lock
└── Gopkg.toml
```
- `Makefile`: contains common tasks for the project, like formating, testing, linting, etc;
- `cmd/example/main.go`: is the cli entrypoint;
- `example.go` and `example_test.go`: is the "library" of the application and its respective files. Could be more than one file, of course;
- `goreleaser.yml`: the GoReleaser configuration;
- `Gopkg.lock` and `Gopkg.toml`: dependencies locks and manifest.

Of course, creating this all the time I want to write some tool would be a lot of work, so I kind of automated it.

## Starting a new project

To help me (and hopefully others) start projects faster, I created [caarlos0/example](https://github.com/caarlos0/example), which contains the directory struture we just talked about, as well the deps, Makefile, License, readme, code of conduct and all that.

To use it, we can simply:

```sh
$ cd $GOPATH/src/github.com/myuser
$ git clone git@github.com:caarlos0/example.git myapp
$ cd myapp
$ ./script/setup myuser MyApp # notice the case on the second arg
```

It is actually a working app (that does nothing), to run it:

```sh
$ make setup
$ go run ./cmd/example/main.go -h
```

Now, we create a GitHub repository for our new app and push it:

```sh
$ git push origin master
```

If we check the README file, we'll see that there are already a few badges on the top, but some of them are not working. Let's fix them!

## Badges

Badges are important! They show cool stats about our project and may show off other tools we use.

[caarlos0/example](https://github.com/caarlos0/example) includes badges for:

- Latest release
- License
- [Travis](http://travis-ci.org/) build status
- [Codecov](https://codecov.io/) Coverage status
- [GoReportCard](https://goreportcard.com/) status
- [SayThanks.io](https://saythanks.io/), so people can thank us for our projects
- [GoReleaser](https://github.com/goreleaser/goreleaser) badge

[GoReleaser](https://github.com/goreleaser/goreleaser), [GoReportCard](https://goreportcard.com/), License and latest release don't need any extra work. 

Latest release will start to work when we launch the first release.

We need to enable our new repository on [Travis settings](https://travis-ci.org/profile/) to make the travis badge work.

Coverage reports will require us to just have an account at [Codecov](https://codecov.io/) website. 

Yep, just as simple.

Now, we need commit and push something to fire a Travis build and check if both Travis and Codecov badges work:

```sh
$ git commit --allow-empty -m 'fire travis build'
$ git push origin master
```

Wait a few minutes, and the badges should be displayed on the README!

[SayThanks.io](https://saythanks.io/) is as easy as creating an account on the site and adjusting the username on the badge.

With that done, we can focus on write the features we want!

## Writing features

First of all, we can read the `CONTRIBUTING.md` file.

It is the "newcomer guide" and usually helps having more contributions, so, it's good to make sure it's always right and up to date.

tl;dr: there are some make tasks:

```sh
$ make setup # install tools and deps
$ make fmt # format code
$ make test # runs tests
$ make lint # runs gometalinter
```

I usually write my features in the root folder. After that, I call them from my `main.go` file.

And that's it. Feature is done.

Of course, more complex tools might require more folders and stuff like that, but most simple tools I have written fit in this category.

Once we are done coding the features, we'll want to distribute our app, right?

## Release automation

For some time I had a quick-and-dirty shell script to automate this for me, but I ended up needing more power, and therefore started the
[GoReleaser project](https://github.com/goreleaser/goreleaser). It is approaching v1.0.0 very soon, and is already being used by ~50 public projects, despite being very new.

[caarlos0/example](https://github.com/caarlos0/example) is already set up to:

- build both Linux and macOS binaries for amd64;
- push a Homebrew recipe to a tap.

To enable it, we just need to create a new [GitHub token](https://github.com/settings/tokens/new) with the `repo` box checked.

Then, we need to add it as an environment variable named `GITHUB_TOKEN` on Travis project settings.

We also need a tap repository. Mine is [caarlos0/homebrew-tap](https://github.com/caarlos0/homebrew-tap), which is kind of a standard. I also created a `Formula` folder inside it.

Finally, we need to change the `gorelaser.yml` file, pointing to our new homebrew-tap repository.

Now, commit and push everything:

```sh
$ git add goreleaser.yml
$ git commit -m 'setup goreleaser'
$ git push origin master
```

To fire a release, GoReleaser expects us to create semantic versioned tags, for example, `v1.2.3`. We can do that using either the command line or GitHub web interface:

```sh
$ git tag v1.0.0
$ git push origin v1.0.0
```

Travis will build the tag, and, on the `post-success` hook, will check if its build a tag, and then, download and execute GoReleaser, which will do the heavylifting for us.

You can check how a release will look like in the [example releases](https://github.com/caarlos0/example/releases).

## Final words

This is how I work right now. I'm not saying it is perfect nor the only way of doing this kind of stuff. I only kind of automated some stuff and wanted to share it so other people may save some time as well.

I'm curious, though, what tools do you use? How do you structure your projects?

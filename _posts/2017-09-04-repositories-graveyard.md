---
layout: post
title: "A Repository Graveyard"
---

This past weekend I decided I need to clean up my GitHub profile. In this post
I'll write about why I cleaned everything up and also how I did, as well as
some initial results.

## Why

I've been using GitHub for years now, and I tend to create a lot of
repositores. Most of them, are
[PoCs or spikes](https://medium.com/production-ready/theres-nothing-like-a-good-spike-4a575686a7c5)
and are often abbandoned. Others may have been useful for a while but are
deprecated now, thus I have no reason to maintain them.

In any case: I get tons of automated scripts running, and some of them looks
into all my repositories. Since I had almost 300 of them, this was considerably
slow. It also spends a lot of quota of API calls.

## How

To address that, I decided to move all my old-not-maintained-abandoned-and-etc
repositories to a new GitHub organization I created, called
[caarlos0-graveyard](https://github.com/caarlos0-graveyard).
There isn't an API for that, so, yes, I moved **77** repositories by hand.

I also deleted some repos that were merely scaffolds with no actual code
written.

To clean it even more, I changed my
[fork-cleaner](https://github.com/caarlos0/fork-cleaner)
utility to include private forks as well, thus removing even more projects.

After doing all that, I now have only ~50 repositories!

## Other problems

In OpenSource, it is common that people abandon repositories. It is also
common that other people continue to open issues to those repos: they
don't care that you don't care.

As you don't care, it's likely that you will never respond to those requests
(maybe because you turned notifications off or something), and other folks
might get frustrated because of that.

I don't like to frustrate fellow programmers, so I created a small webhook
called [gravekeeper](https://github.com/caarlos0/gravekeeper). You can put
it in your abandoned repository webhook or even in an entire organization,
and it will reply to new pull requests and issues saying that the repository
is not being maintained anymore.

It runs as a Lambda function and costs basicaly nothing.

## Results

- I went from ~270 repositories in
[my main account](https://github.com/caarlos0) to just 54
(including private repos);
- My scripts run like 4x faster (not scientifically measured);
- My profile now shows only things I'm actually working on - or worked on
recentely - or that I intend to work on again;
- I don't need to worry about new issues and pulls on my abandoned repos;
- Less cluter == Happier Carlos :)



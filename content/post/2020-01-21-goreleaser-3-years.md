---
title: "GoReleaser: 3 years later"
date: 2020-01-21
draft: false
slug: goreleaser-3-years
city: Joinville
---

[GoReleaser](https://goreleaser.com/)‘s journey begins in December 21, 2016: the day I made its [very first commit](https://github.com/goreleaser/goreleaser/commit/8b63e6555be45234c4c2a69576ca2ddab705302c). It has been a long road since then.

<!--more-->

I didn't have big aspirations for it. Just wanted to reliably automate my workflow in another project, so, the first version was pretty hacky, but whatever, just scratching my own itch.

It turns out more people had the same itch, and GoReleaser scratched theirs as well.

I wouldn't be able to do it by myself though. I had lots of help:

- My [lovely wife](https://github.com/carinebecker) did the artwork;
- The people and companies donate, both [directly to me](https://github.com/sponsors/caarlos0/) and to the [project itself](https://opencollective.com/goreleaser/);
- The other developers that help adding features, fixing bugs, supporting new users;
- The users who report interesting new use cases I had not think of before;
- And, of course, everyone who uses it and spread the word;

**Thank you all.** 🙏

That being said, I also wanted to take a look on what we achieved so far.

## Usage

[Mike Fridman](https://twitter.com/_mfridman) did an awesome work collecting statistics about GoReleaser!

> You can check the thread on Twitter (and his work on github).

By the time Mike gather the data, there were more than **3800** unique public repositories with `/.?goreleaser.ya?ml/` file! It is… **a lot** of repositories!

What also surprised me was the amount of "high profile" users (repositories that have a lot of stars), among them:

- Apex
- Buffalo
- Cayley
- Fabio
- Hugo
- InfluxDB
- Micro
- Mock (official golang project 😱)
- NATS
- Traefik
- Vegeta
- …

You can see the full list [here](https://docs.google.com/spreadsheets/d/1zT_AIQCA7Ux7a_6eDj2eVF7LZaFaw-bbVHOHCOqkJzo/edit?usp=sharing) and the "top users" [here](https://gist.github.com/mfridman/d61502bbd837c81e50c370d2dd5a7496).

## Stars, contributors and so on

On top of that, GoReleaser now have `5000+` stars, `120+` contributors and `300+` releases.

We can see in the stars over time on the following graph that it is [growing pretty steadily](https://starchart.cc/goreleaser/goreleaser):

RenderImage when len(FileIDs) == 0 NYI
![](goreleaser-.svg)

Which is nice, I guess. 😌

## The future

All that being said, I'm happy to announce that 2020 is the year when we'll finally release `v1.0.0`.

At this point, its just a number. GoReleaser is pretty stable (both
in how it works and api-wise), so I'll just wait for the remaining [deprecation notices](https://goreleaser.com/deprecations/) to expire, remove them, and tag `v1.0.0`.

If everything goes according to the plan, this should happen by the June.

I also plan to finally tackle the problem of generating the `.sh` files, which today is [GoDownloader](https://github.com/goreleaser/godownloader)‘s job. If you have ideas about this, I invite you to participate on the [discussion](https://www.notion.so/caarlos0/8a13b651f9ff49e9b0f93f21e24563e5?v=82b7951324d84ebfa81295134ee82e48).

I have in the back of my mind an itch regarding the documentation as
well. I wanted it to be searchable, maybe add a config generator as
well… not sure yet (help/ideas appreciated here as well).

From there, the idea is to keep `v1` stable (maybe adding small features) and start thinking about `v2`, which, at this point is too fresh to have any ideas about how it should look like.

## That's all folks

Just wanted to give a quick update about how things are and where are we headed to.

Thank you all again, and happy new year!
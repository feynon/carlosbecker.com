---
title: "Monitoring GitHub releases with Prometheus"
date: 2018-12-15
draft: false
slug: monitor-github
city: Joinville
toc: true
tags: [monitoring, prometheus, github]
---

I have written some exporters to observe things on GitHub. This is how and
why.

<!--more-->

We will talk about 2 things:

1. other projects' releases;
2. my projects' releases;

So let's get started!

## Other Projects' Releases

On my work and personal projects I use **a lot** of OpenSource projects. It
is really hard to catch up with latest security updates and new features of
all of them.

Of course, most of the time you don't care, but some of them you may do. No
way to know without looking into them.

To help me with that, I wrote a [Prometheus](https://prometheus.io/) exporter called
[version_exporter](https://github.com/caarlos0/version_exporter).

With that running, I can leverage the alerting pipeline I already have
setup to let me know when it's time to update things! I just need to add
the repositories and semantic versioning constraints to a YAML file and create
the alerting rules for Prometheus, for example:

```yaml
- alert: SoftwareOutOfDate
  expr: version_up_to_date == 0
  for: 1s
  labels:
    severity: warning
  annotations:
    summary: "{{ $labels.repository }}: out of date"
    description: "latest version {{ $labels.latest }} is not within constraint {{ $labels.constraint }}"
```

I created this tool on my previous work because we had a lot of tooling in
place there, I use it on my [own prometheus instance]({{< ref prometheus-authentication-with-oauth2_proxy.md >}}) and also on my
current work.

A few weeks ago, GitHub launched a new feature so you can watch
repositories for new releases only. That's dope, but you can't warn your
entire team/organization, so this tool still have this good difference.

I hope it is useful for you and/or your team and that the README is good enough
(let me know otherwise 😂).

## My Projects Releases

I'm a curious person, I want to know how many downloads my releases get...
unfortunately, GitHub does not provide that information in its interface,
nor provide it as a time-series on the API, so... you guessed it, I wrote
another exporter: its called [github_releases_exporter](https://github.com/caarlos0/github_releases_exporter)!

On this exporter, I just set up a list of repositories I want to watch, and
then the exporter gets all assets from all releases and reports their
number of downloads as a prometheus metric.

Using [PromQL](https://prometheus.io/docs/prometheus/latest/querying/basics/) I can then get the number of downloads for each platform,
for each release, each packaging format and so on:

```
sum by(repository) (github_release_asset_download_count)
sum by(repository) (github_release_asset_download_count{name=~".*Darwin.*"})
topk(5, sum by (name, tag) (github_release_asset_download_count{repository='goreleaser/goreleaser'}))
```

It is also pretty simple, but I found it very useful to see, for example, how
many people were affected by a buggy release or something.

Hope you enjoy it!

---

## Outro

Besides that, I also have [other exporters](https://github.com/search?q=topic%3Aprometheus-exporter+user%3Acaarlos0&type=Repositories) and a few other
[github-related tools](https://github.com/search?q=topic%3Agithub-api+user%3Acaarlos0&type=Repositories). Feel free to check them out!

I've been using most of those tools for years (and I more-or-less constantly
update them)... the reason I'm writing about them just now is because I was
improving them during this weekend (because both had a very... let's say...
primitive design).

I hope at least some of these tools can be useful for you!

Feel free to ping me on the comments bellow if you have any questions.

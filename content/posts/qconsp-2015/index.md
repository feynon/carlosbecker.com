---
title: "QCon Sao Paulo - 2015: A short overview"
date: 2015-03-28
draft: false
slug: qconsp-2015
city: São Paulo
toc: true
tags: [conferences]
---

So, this week I attended to [QCon-SP](http://qconsp.com/).

The conference was great (congratulations everyone 🍻), but, I thought it would be nice to do an overview.

So, the top subjects were Microservices and Docker. A lot of Big Data too, but I like the Microservices thing more, so I didn't follow the Big Data track.

We saw a lot of company culture too, and, believe it or not, it was strongly related to Microservices.

Let me explain.

Basically, they defend small teams (5-8 people), each team owning one or more Microservices. The team is responsible for both develop and deploy those services, basically, **the teams are multidisciplinary**. The **teams**, not the **people**.

Also, a real Microservice should be **independent**. Microservices sharing the same database are not Microservices. 

The philosophy behind Microservices is based on Unix [Philosophy](http://en.wikipedia.org/wiki/Unix_philosophy):

> Make each program do one thing well. To do a new job, build afresh rather than complicate old programs by adding new "features".

Of course you probably won't do that from day-1 in your new startup. You don't even know if it will work, and, Microservices add a little of complexity you might not want to pay now.

About this, two quotes by [@randyshoup](http://twitter.com/randyshoup):

> "To improve is to change; to be perfect is to change often."

And

> "If you don't end up regretting your early technology decisions, you probably over-engineered."

There is a lot of cool things that you can do with it right now, and there will probably be more soon, like running desktop softwares inside a Docker container and freezing user space to turn on a Docker container in a "warm" state - which seems nice if you think about JVM JIT, for example.

I wish I had attended to the [Docker tutorial by Jerome](http://qconsp.com/sp2015/tutorial/docker-and-containers-fast-paced-introduction.html), but, unfortunately, that was not possible. There were very little practical stuff about Docker in talks and keynotes, but they were nice anyway.

See you next year!

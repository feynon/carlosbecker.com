---
title: "Find the slowest tests of a Java project"
date: 2014-04-28
draft: false
slug: slow-tests
city: Joinville
---

I found that it's pretty hard to have a project with high test coverage and fast build... if the tests are slow, people will feel the need to skip them to speed up the build, and will probably write less tests than they should, afraid that the build will became even slower. You go out for a walk and when  you come back no one is running or writing tests anymore...

One way to avoid that is to track down the slowest tests and fix them. The less dependencies your test have, the fast it will run. Just saying.

Paraphrasing Uncle Bob:

> "The first rule about tests is that they should be fast. The second rule about tests is that they should be faster than that."

** Actually he said that about method sizing, but, I think this fits just fine.

To help with that, I created a [project](https://github.com/caarlos0/junit-slowest-tests) with scripts that can generate a list of the slowest JUnit tests in a project. It should give you some insight in where to attack.

The [project's README](https://github.com/caarlos0/junit-slowest-tests) is pretty complete and the scripts are really simple and easy to use.

Besides that, you can use [maven-profiler](https://github.com/takari/maven-profiler) to find other slow parts of your build and fix them.

One last tip: if you have problems with low test coverage in your project, try [coverage-maven-plugin](https://carlosbecker.com/posts/mvn-pr-coverage-blammer/),
a maven plugin that will blame pull request with coverage bellow a specified amount.

That's it for today, happy hacking!

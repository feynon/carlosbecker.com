---
date: 2015-02-22T00:00:00Z
slug: introducing-jvm
title: Introducing JVM - The Java Version Manager
city: Joinville
tags:
- shell
---

Last years it becomes more and more common to work in different projects
running on different versions of Java. There still some running on Java 6,
and there are tons already running on Java 8.

That's nice!

What isn't nice is to change `JAVA_PATH` and `PATH` every time you work on
other project.

[JVM][jvm] aims to help you with that.

It's basically a shell script that you load on your `bashrc` or `zshrc`. It
will then listen on folder change events. If you `cd` to a project running
on Java 7, and you have Java 8 on your `PATH`, JVM will do the ~magic~ for you.
:sparkles:

But how it works?

Basically, JVM will search for `pom.xml` files. If any is found, search for a
`java.version` property in it. If it's found, search for matching installed
Java, otherwise, JVM will look for `.java-version` file, and, if found, try
to set a matching Java version to your `PATH`, otherwise do nothing.

Currently, it works Mac OS X and Ubuntu only, but, of course, pull requests
are welcome.

For further info, take a look at the [project's README][jvm].

[jvm]: https://github.com/caarlos0/jvm

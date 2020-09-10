---
title: "Dealing with Maven dependency hell"
date: 2016-09-03
draft: false
slug: maven-dependency-hell
city: Joinville
---

Every now and then an active java-based project enters in a "dependency hell" state. That usually happens because people keep adding dependencies without checking what comes in transitively nor if that dependency is declared somewhere else already.

It also seems to happen more with really big projects, which only make everything event worse when it comes to the time that a fix is needed.

This is a story on how I (try) to sanitize a big java project's dependencies.

## Get rid of the obvious stuff first

Yeah, start with the basics!

Let's use `maven-depedency-plugin` to find obvious (and easy) stuff to fix:

### 1. duplicated dependencies

```
mvn dependency:analyze-duplicate
```

Just remove what's duplicated and that's it.

### 2. used and declared; used and undeclared; unused and declared

```
mvn dependency:analyze
```

This will take a while and may warn you about a lot of depedency errors. 

You will be able to remove dependencies that are not used anymore and add dependencies that are directly used.

> **Beware**: This might show some false-positives.

## Bring the big guns

### 1. multiple versions of the same depedency

This is the harder, as maven doesn't have a plugin to find them easily. 

So, I used a little bit of bash trickery to do the job for me:

```
mvn dependency:list -Dsort=true |
  grep "^\[INFO\]    " |
  awk '{print $2}' |
  cut -f1-4 -d: |
  sort |
  uniq |
  cut -f1-3 -d: |
  uniq -c |
  grep -v '^ *1 '
```

Line-by-line, explained version:

```
mvn dependency:list -Dsort=true | # list all deps
  grep "^\[INFO\]    " |          # grep for the deps list only
  awk '{print $2}' |              # remove the INFO prefix
  cut -f1-4 -d: |                 # removes the dep scope
  sort |                          # sort (duh)
  uniq |                          # remove duplicates
  cut -f1-3 -d: |                 # removes the version
  uniq -c |                       # count line groups
  grep -v '^ *1 '                 # grep groups that repeat
```

You'll end up with a list like this:

```
2 com.foo:foo-modules-commons:jar
    3 com.fasterxml.jackson.core:jackson-annotations:jar
    3 com.fasterxml.jackson.core:jackson-core:jar
    2 com.fasterxml.jackson.core:jackson-databind:jar
    2 com.squareup.okio:okio:jar
    2 commons-collections:commons-collections:jar
    2 commons-fileupload:commons-fileupload:jar
    2 joda-time:joda-time:jar
    2 org.mapstruct:mapstruct:jar
```

Now, for each of these deps, you'll have to run:

```
mvn dependency:tree -Dincludes=DEP
```

as in

```
mvn dependency:tree -Dincludes=org.mapstruct:mapstruct
```

This will show you all the places this dependency is being used, so all you need to do is fix it.

### 2. duplicated classes

This is something that I have seen a lot in Apache Commons libraries.

As far as I know, the easiest way to find these problems is by using Jboss' [Tattletale Maven Plugin](http://docs.jboss.org/tattletale/userguide/1.2/en-US/html/maven.html).

While the plugin seems to be abandoned, it still works. Just add it to your parent pom and run it to get the reports.

These kind of problems may require manually removing `.class` files from jars, excluding dependencies or, sometimes, just ignoring some of them.

## Preventive actions

Act when the everything went to shit already usually is harder than avoiding small mistakes.

So, my tips are:

- Give more attention to it in code reviews;
- Run `maven-dependency-plugin` and `tattletale-maven` in the build for every pull request to block problematic changes;
- You can tweak the script to find dependencies with multiple versions to `exit 1` in thoses cases, and just add it to the build too.

That's it. Hope it helps!
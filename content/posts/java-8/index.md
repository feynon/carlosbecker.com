---
title: "Java 8"
date: 2014-08-04
draft: false
slug: java-8
city: Joinville
toc: true
tags: [java]
---

Earlier this year, the new version of the Java Programming Language was released. Finally, it enters in the field of the "cool peeps" with some features it should have since years ago, like Lambdas.

---

> "Java developers are now Hipsters again!"
> 
> — Seen somewhere in the wild internet.

Anyways, I have been reading about and building simple algorithms with it since day one, but, now that I'm using it in a "real world" project, I would like to say some words about how the experience is going.

But firsts things first. The new features!

## The new features!

Well, this version surely have a lot of cool new features, so, I would like to point out my personal favorites!

### 1. Lambda Expression

This is really awesome. A lot of languages have this feature for years now, and for years I've been talking about how Java should have this. Now, finally, we have:

```java
Runnable runnable = () -> System.out.println("Run forrest, run!");
```

Even if the implementation "behind the scenes" is not quite the best, in my humble opinion it is better than the interfaces with all that weird inner classes that we were using before:

```java
Runnable runnable = new Runnable() {
  @Override
  public void run() {
    System.out.println("Run forrest, run!");
  }
};
```

### 2. The Stream collection types

I have to admit, this is awesome! We can do all sort of things with a DSL that is actually enjoyable to use:

```java
Map<String, String> hash = new HashMap<>();
hash.entrySet().stream()
    .map(entry -> entry.getValue())
    .filter(entry -> entry.startsWith("C"))
    .distinct()
    .count();
```

There is also methods for parallel sorting and other stuff (`parallelStream`). 

Check also the `IntStream` class.

### 3. The new Date and Time API

This is another one that was required for years. Basically, they moved the Date/Time API to the `java.time` package and followed the `Joda Time` format, with the good thing that most classes are Threadsafe and immutable. 

I just can't wait for Java 8 and the Money API! :moneybag:

### 4. String.join

I still can't believe that it take so long...

```java
System.out.println(String.join(". ", "FINALLY", "String", "joining"));
// FINALLY. String. joining
```

### 5. Optional

I posted about this in my Facebook timeline and it ends up becoming a little noisy. Anyway, I would really prefer a [ruby-ish approach]({{< ref "/posts/ruby-nil/index.md" >}}) to this problem, but, well, this is better than nothing. I have used it in some `filter` calls like this, for example:

```java
List<String> names = new ArrayList<>();
names.stream()
    .filter(entry -> entry.startsWith("C"))
    .findFirst() // returns an optional
    .get(); // tries to get its value (might throw NoSuchElementException)
```

## Working with it in the wild

Update all the things! That's it.

I have to admit that I was expecting for this. Almost every library I was using before (with Java 6 or 7) had to be updated.

Some of them aren't even working, and needed some special attention, like opening a pull-request with the fixes and the deploy of custom version with the fix in local nexus repository until the new official version is released, but, nothing that can't be fixed.

This isn't really bad, it's just labor intensive... and a little tedious. Looking for the bright side, you always learn a little. ✨

## Resources

Well, you have probably read something about Java 8 already, but, if you didn't, or just want to read more, I recommend you to read the [Java 8 Friday](http://blog.jooq.org/tag/java-8/) at the jOOQ's blog. They have plenty of good study material about this manner. 

I could also recommend you a book that I have read: [Java 8 Lambdas: Pragmatic Functional Programming](https://www.amazon.com/gp/product/1449370772/ref=as_li_tl?ie=UTF8&camp=1789&creative=390957&creativeASIN=1449370772&linkCode=as2&tag=carlbeck-20&linkId=FLJGZ6WNDZWK7EQK).

If you want to read some Java 8 code, I have solved some katas using it just for fun. They are probably pretty bad, so feel free to suggest improvements.

You can take a look at the source code [here](https://github.com/caarlos0/java-katas).

I have also ran a small Ruby Sinatra service with the last version of jRuby under JDK8 and it worked gracefully. You can take a look at it [here](https://github.com/caarlos0/danfe-server).

Well, this is all for now! Cheers!

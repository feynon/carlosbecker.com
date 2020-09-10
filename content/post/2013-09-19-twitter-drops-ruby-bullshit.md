---
title: "Twitter drops Ruby - The Bullshit"
date: 2013-09-19
draft: false
slug: twitter-drops-ruby-bullshit
city: Joinville
---

This post is about the Twitter change from Ruby to Java, some years ago, fact that, sadly, is still being used by some people to say that Ruby sucks and Java is the best language in the world.

Before you read that, a little background on me would be good:

- Java is the first language I worked with. I started 4 years ago and I'm still using it **every day**;
- No, I don't hate Java, it is just not my favorite language;
- I started with Ruby about 2 years ago;
- A lot of friends and coworkers try to troll me with that bullshit about Twitter and others.

With that in mind, I decided to drop a few lines about this here.

## Twitter dropped Ruby

Yeah, in the [original post](https://blog.twitter.com/2011/Twitter-search-now-3x-faster) Twitter says that they get 3x more performance using Lucene and Java. The post is OK. The arguments are right. Your interpretation may be wrong.

People like to generalize things. Like "Twitter says they get more performance with Java, so, Ruby sucks", which is not correct. If it was, maybe you could also use PHP, since Facebook uses it and it works for them.

## Scalability isn't everything

Java is, in fact, more scalable than Ruby. And Erlang is more scalable than Java. So if scalability is your only argument and concern, let's Erlang everything.

![](/public/images/twitter-drops-ruby-bullshit/567e1df8-9b29-4dcc-80e7-50753ce2e891.png)

Oh, are you talking about performance now? Why don't you write your entire application in C using [raphters](https://github.com/DanielWaterworth/Raphters)?

Exactly: because of **productivity**.

The fact is that Ruby and Java have different focuses in some points. Ruby is way more focused in readability and productivity than Java, for example. But this argument itself doesn't make one language better or worse than other. They are just different.

## Twitter never said that Ruby sucks

Never. You could read that post how many times you want.

What happens is that a lot of republications about that post do.

Usually, they don't even know what they are talking about. They just link to that post, talk a lot of bullshit around the subject, and say something pretentious, like, "thanks to java Twitter handles the U.S. elections", which is also usually the title. I'm very inclined to believe that someone is paying them to say that.

That's bullshit. Maybe, if they just stopped using MySQL and started using Lucene and refactored their Ruby code - which they said in the post that had a huge technical debt - it would be fast enough. Maybe not. We can't be sure about it based on assumptions.

## You don't have Twitter's problems

You probably don't and probably will never have numbers even close to Twitter's. You probably don't have the same problems. The solutions that work for Twitter may not work for your application.

## Java != Java EE

Oracle, obviously, took advantage of this situation to promote their technology, Java and Java Enterprise Edition.

The thing is that Twitter doesn't use any JEE technology. They implemented their own "platform", called Blender, using things like Thrift, Netty and NIO, as they say in the post:

> Blender is a Thrift and HTTP service built on Netty, a highly-scalable NIO client server library written in Java that enables the development of a variety of protocol servers and clients quickly and easily.

JEE is heavy and has a lot of features that Twitter will never use. Maybe for an ERP it could make sense and work pretty well. But not for Twitter, not for everyone.

## Let's talk languages

So, what's the real concurrency problem with Ruby?

The main problem is that Ruby uses a Global Interpreter Lock (GIL), which means that even if you use multiple threads, there will be only one lock per interpreter process. In other hand, Java has a very well tested and stable concurrency model with one lock per thread, as well other concurrecy models (like Actor model using things like Akka).

![](/public/images/twitter-drops-ruby-bullshit/eaebed20-d049-4173-aa31-e4f5daaca3b3.png)

The implications of Ruby's behavior is that if you want a 10 level concurrency, you will need to start 10 application processes. Assuming that each application instance uses around 200mb of memory, 10 x 200 = 2Gb~ (which is indeed a lot of memory), while in Java it would probably use less, depending on the application server.

The main problem is the memory usage. If you need something like 1k level concurrency, you are screwed. But this is not the only problem. Since Ruby has the GIL, if one thread locks it, the other threads running in that interpreter instance would have to wait. This behavior **could** also lead to a lot of CPU idle time.

The good part is that some Ruby application servers, like unicorn, rely on unix process forking to achieve concurrency. Basically, it means that instead of 10 processes with separate memory spaces, the server "clones" the process into 10. They share this part of memory, but they can also modify their own memory without affecting other forks.

Note that what I said above is about Ruby MRI. You can also try jRuby, which has all the advantages of JVM, or even Rubinius.

## What's the point?

The point is the same old conclusion that everyone keeps repeating but nobody is listening: **There is no silver bullet**.

I like to see comparisons between languages. I don't mind if I read somewhere that "*[some lang I like][random feature]* sucks", s long as they have valid arguments.

I also don't mind if someone says that *[lang X]* is better than *[lang Y]* in some matter, as long as they prove me their point.

Now, saying that "*[lang X] is better than [lang Y]* ***because of [reason Z]***" doesn't make any sense.
---
title: "What is good code?"
date: 2015-12-27
draft: false
slug: good-code
city: Marechal Cândido Rondon
toc: true
tags: [software-engineering]
---

I've been wanting to write this for a long time, just to clarify my thoughts on the subject. Now, on vacations, I took a couple of days and finally did it. This is a personal opinion based on my personal experience and tons of books I have read, and I am not, by any means, the supreme holder of the true, so you will probably disagree with me at some point.

That said, let's try to answer the question: "**What good is code?"**

I always tried to write small methods, and was constantly discussing with some coworkers about it. "Why write a giant 3000 lines method that no one can fully understand if you can break it down in more  classes and methods so it all just makes sense?", I asked. "Why in hell don't you guys want to write unit tests to assert your code is doing the right thing?", I complained. "Let's refactor this class together, its
enormous...", I tried.

I was young and I was stupid. That kind of discussion improved the code in nothing, because I spent a lot of energy focusing in the wrong part of the problem. One day, I gave up and moved to another job where [Clean Code](http://amzn.to/1O6J19h) was the standard.

I remember the first time I read [The Clean Code](http://amzn.to/1O6J19h) - it all seemed so obvious: small methods and classes, good method and variable names, refactoring... and some stupid irrelevant bullshit like removing the `{}` from one line `if` statements just to make Java look less verbose.

So, when I started working at this new job, for some time I was realized: People writing unit tests, small functions, small classes... and all of them with somehow meaningful names... I mean, that's the dream, right? **Good code!**

I struggled to write the best code I could write, after all, now I have plenty of energy to invest in this kind of stuff because I didn't need to discuss "obvious stuff" with nobody.

After some time, I read [Object Thinking](http://amzn.to/1Yuyf3h). More than half of the book was extremely philosophical and, at the end of this part, I started to fully understand what object oriented programming means.

Then, I re-read some parts of [The Clean Code](http://amzn.to/1O6J19h), and some of those parts just didn't make sense anymore.

> "The first rule of functions is that they should be small. The second
>  rule of functions is that they should be smaller than that. Functions 
> should not be 100 lines long. Functions should hardly ever be 20 lines 
> long."
> 
> — Clean Code: A Handbook of Agile Software Craftsmanship

For me, the problem in this statement is that it focus more in the amount of lines than in the amount of responsibilities.

Let me explain:

If we did good object oriented design, in theory, our object's methods (not functions) would not need to be big - they would not have any reason to. You just won't have that much things to write in a method to make it that big. The methods would probably have only one `return` statement which will probably be another object or a composition of objects.

That's because in good object oriented programming, we won't have a `WhateverService` or `ThingController` or anything like that, just because those things don't exist in the domain's problem. Instead, we would have a lot of objects, each one doing their part of the work - totally decentralized, instead of objects acting as coordinators.

The Clean Code tries to make us get closer to good design by forcing us to break methods (and classes) that do too much things in methods (and classes) that do fewer things with the excuse of **lowering the LOC** of those methods (and classes) - and assuming we will think in what you're doing.

Of course, to make a method smaller we will have to, eventually, make it do fewer things - which is a good thing. But is it good design? Not necessarily.

**If we break a shitty method into smaller shitty methods, we are not making the code any better, we're just spreading shit all over the place.**

Let me be clear: **the goal should not be "to have small methods", but "to rightfully decompose the domain problem" instead.**

To properly decompose our methods, we need to think in terms of objects and their contracts, not in the number of lines. We should probably refactor the entire modeling, not only one method or class. If a method is too big, it's a consequence of poor design choices: we didn't decompose the problem enough/right.

The same goes for classes.

So, if you ask me "**what good is code?**", I would not say that it has small functions - I would rather say its problem's domain was properly decomposed and that the contracts are well defined - and, of course, that it follows the language standards, has unit tests, etc...

That is what good code looks like to me.

---

Books I recommend:

- [The Clean Code](http://amzn.to/1O6J19h);
- [Object Thinking](http://amzn.to/1Yuyf3h).

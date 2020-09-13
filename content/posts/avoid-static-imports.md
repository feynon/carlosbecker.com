---
title: "Avoid static imports"
date: 2015-07-05
draft: false
slug: avoid-static-imports
city: Marechal Cândido Rondon
toc: true
tags: []
---

There are a lot of Java API's and Frameworks which rely in `static` methods and the sort. Arguably, this is a bad OOP practice, but lets not enter this particular subject just yet.

For example, it is pretty common to write unit tests like this:

```java
package blah;

import static org.junit.Assert.assertTrue;

import org.junit.Test;

public class SomeTest {

  @Test
  public void testSomething() {
    assertTrue(1 + 1 == 2);
  }
}
```

While it's kind of pretty to `import static org.junit.Assert.assertTrue`  and use `assertTrue` directly instead of `Assert.assertTrue`, one might argue that it's difficult to know where this `assertTrue` method came from. I don't fully agree with that in this particular case because this is an isolated and well known context (the test) and every Java developer knows (or should know) Junit, therefore the `assertTrue` should be obvious.

What I believe not to be a good thing is this:

```java
package foo;

public enum FooType {
  BAR, ANOTHER;
}
```

Later:

```java
package bar;

import static foo.FooType.*;
import foo.BarService;

public class BarService {
  private FooService foo;

  public void update(Long id) {
    foo.update(id, BAR);
  }
}
```

Now, when other folks read this code, they will see this `BAR` there and will have to discover where it come from. It's not the `Foo` context anymore and it's not a well known API, so, how can people know where this came from?

It would be much simpler if it was written as `FooType.BAR`, right? Besides, we just polluted the namespace with all `FooTypes`, which is also a bad thing.

I know... static importing some stuff can make the code "prettier" or look simpler. Yeah, nope. If your code is not simple or pretty enough, it's not the job of static imports to fix it. Probably your design is just bad.

I'm not the only one who believe in this:

- There is a [checkstyle rule for it](http://checkstyle.sourceforge.net/config_imports.html#AvoidStaticImport);
- [Java Documentation about static imports](http://docs.oracle.com/javase/1.5.0/docs/guide/language/static-import.html).

Check this out:

> So when should you use static import? **Very sparingly**! Only use it when you'd otherwise be tempted to declare local copies of constants, or to abuse inheritance (the Constant Interface Antipattern). In other words, use it when you require frequent access to static members from one or two classes. If you overuse the static import feature, it can make your program unreadable and unmaintainable, polluting its namespace with all the static members you import. Readers of your code (including you, a few months after you wrote it) will not know which class a static member comes from. Importing all of the static members from a class can be particularly harmful to readability; if you need only one or two members, import them individually. Used appropriately, static import can make your program more readable, by removing the boilerplate of repetition of class names.

Of course, if you follow the checkstyle rule, you will have no static imports in your code. This can be a little "too much", but, if I had to choose between static imports everywhere or no static imports at all, I'll surely choose the later. 

But, what I really prefer, is the common sense.

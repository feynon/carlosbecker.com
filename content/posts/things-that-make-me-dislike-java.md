---
title: "Things that make me dislike Java"
date: 2012-12-15
draft: false
slug: things-that-make-me-dislike-java
city: Joinville
toc: true
tags: [java]
---

So, I just compiled a little list of things that, I believe, make me dislike Java everyday a little more.

This is not a hateful flamewar-guided post. Just my thoughts about a subject.

Java has a lot of good points. Starting with community, specs, lot of libraries and so on. I use it everyday in my work and have been used it for almost 4 years. The things I'll talk about today, is more related to concepts adopted by Java language itself.

So, if you only programmed in Java until today and/or have Java as a mother for you and can't hear any kind of shit about it, **go away now**. If you're not, keep reading. (or go away too, I don't really care)

## OOP aspects

Even though people say that Java is an object oriented language, well, it is not that much. It has some concepts of OO, like classes, object, inheretance, and so on... but, it still not be a pure OOP language.

For example, in OO, **everything** should be an object. Java primitives, for example, are not.

## autoboxing and null behaviors

Like I said before, primitives are not objects, which means that you can't use they with Generics, for example. You can't make a `List` of `int`s, you have to use `Integer` wrapper class instead.

The big problem is that the behavior of the primitives and their respective wrappers are inconsistent in my point of view.

IMHO, these wrapper classes should not be wrappers and primitives should not exist. One of the seven qualities to satisfy the OOP concept is that all predefined types are objects, which not applies to Java. Anyway, I wasn't there when Java engineers take these decisions, so, I'll not judge. But, I believe that they decide this because in the beggining, the Java Virtual Machine was slow, and have primitive types looks good to "make it fast". Maybe it was a bad decision. Anyway, I really wish that at least these wrappers had the same behavior that their respective primitives. That would be awesome.

Well, If you still thinking that's not inconsistent, I'll give you what I think is the worst example:

Everyone knows that a `boolean` should have only two states, in Java, for example, they're `true` and `false`. That's right. Now, let's take a look at the `boolean` wrapper class: `Boolean`.

`Boolean`s let you have three states: `true`, `false` and `null`, which means that's not actually a trully bolean value. Two states != Three states.

Yeah, I know, all that thing about reference. `null` means that the `Boolean` Object doesn't point to anything, for me, it would be nice if it automagically binds to `false`, avoiding people to do wrong things.Check if a `Boolean` isn't `null` doens't make any sense for me, sorry about that.

Let's take an example from Ruby, which is considered a pure OOP language. If you try to evaluate an expression like `puts "not nil" if not nil`, it will works, and will print `not nil`, as expected. This behavior is expected because in ruby, everything is an object, including `nil` itself (of `NilClass` type). In example above, `nil` are automagically evaluated to `false`, which also is an object (of `FalseClass` type), that's just beautiful.

That's the behavior I expect from Java. It would be really nice to use Java `null` like a object, isn't? Take another Ruby example:

```ruby
t = nil
puts t.class # NilClass
puts t.nil? # true
puts "t isn't nil" if t # nothing, t is nil
```

In Ruby, [every object has a](http://ruby-doc.org/core-1.9.3/Object.html#method-i-nil-3F) `[nil?](http://ruby-doc.org/core-1.9.3/Object.html#method-i-nil-3F)` [method](http://ruby-doc.org/core-1.9.3/Object.html#method-i-nil-3F), and `nil` also is an object, which is awesome. I realy miss that behavior in Java.

## Closures

For God sake, I really miss this. Closures, IMHO, are the most usefull thing in a lot of languages, like JavaScript, for example. Ruby has the blocks concepts, that is basically the same thing. In Java, you can get this concept with an interface, doing something like a Command Pattern to encapsulate the action. Anyway, the syntax sucks. You probably will use anonymous classes to get something more readable, but it still pretty ugly.

I also know that is something that's comming with Java 8, but, most languages have this functionally for years, and Java 8 is not ready yet. It kinda sucks.

Anyway, I'll really like to believe that someday I'll could use Java Methods as objects, like it should be in a pure OOP lang... (fingers crossed)

## metaprogramming

Well, would be really nice been able to do that. For example, we don't have a coalesce method in Java's Object, and I really want it! So, why not been able to do something like:

```java
class Object
  public static <T> T coalesce(T ...objects) {
    for(T t : objects) if(t != null) return t;
    return null;
  }
}

Person realPerson = Object.coalesce(person1, person2, person3);
```

We use it a lot in Ruby. We can also "teach" some object a behavior! That's awesome. A Ruby example:

```ruby
class Person
end
p = Person.new
p.respond_to? 'walk' # false
def p.walk
  puts "walking..."
end
p.respond_to? 'walk' # true
p.walk # walking...
```

I really like this behavior, think about the possibilities! For example, we can teach a `map` and/or `filter` function for any kind of iterable thing. That's very, very useful.

## Method Missing

Ruby has a feature heavily used by Rails ActiveRecord, the `method_missing`. It let you made some regex or another logic if user try to use a method that doesn't exist. ActiveRecord's `find_by_anything` feature is based on that. And it's awesome!

If Java also has a `method_missing`, like ruby, you can made an "intelligent" `DAO`, like ActiveRecord. For example:

```java
public abstract class AbstractDAO<T> implements DAO<T> {
  private static final String FIND_BY = "findBy";

  public <T> T methodMissing(String name, Object ...params) {
    int i = name.indexOf(FIND_BY);
    if(i < 0) {
      return super.methodMissing(name, params);
    }
    String[] fields = magicMethodThatExtractTheFields();
    return magicMethodThatBuildAndExecuteAQuery(fields);
  }
}

// somewhere
personDao.findByName("carlos"); // and methodMissing do the magic
```

You can do a lot of things with that. It's very useful.

## DRY

Well, I feel like Java is always making me repeat myself. An example that surely you have been noticed: The ";".

For me, is pretty obvious that if I start coding in a new line, the previous line is completed. A lot of languages doesn't force you to put a ";" at the end of the line: Ruby, CoffeeScript, Haskell, Groovy, (almost infinite list...).

And it's just one example. There is a lot more. You will feel it a lot working with collections.

## The end

I'll like to didn't have to say that, but, well, most of Java programmers have Java more like a religion than a work tool, so, I'm not saying that Java is wrong. I'm discussing some behaviors, comparing with other languages, etc. Please folks, don't crucify me.

Also, some of the concepts I talk about here could be simulated, sometimes with a weird syntax, using [Google Guava Libraries](http://code.google.com/p/guava-libraries/). I've [talked about]({{< ref "rocking-out-with-google-guava.md" >}}) it sometime ago.

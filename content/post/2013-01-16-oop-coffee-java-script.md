---
date: 2013-01-16T00:00:00Z
slug: oop-coffee-java-script
title: 'JavaScript OOP: A beginners guide'
---

Well, there's a lot of tutorials, writings, blog posts, etc in the web
about this subject, but, I'm writing this in a way that I found better to
understand what is **really** happening while you code.

We will code the most part of the examples directly in
[CoffeeScript][coffeescript], so, we will have much less code in this article.
Don't know or don't like CoffeeScript? Don't worry, I'll link the pure
JavaScript for you.

## Getting started

![OOP](/public/images/oop.png "OOP")

In JS, a class is nothing more than a function, since functions (and everything
in JS) are functions. So, basically, you open a function scope, and start code:

```js
function Test() {
  Test.prototype.t = function() {
    alert("test");
  };
}

var o = new Test();
o.t();
```

**What we just done?** We created a scope (`function Test`) and then we add to
the `prototype` of the scope (which is called `Test`) the function `t`.

## Getting started with CoffeeScript

This same piece of code, in CoffeeScript:

```coffeescript
class Test
  t: ->
    alert 'test'

o = new Test()
o.t()
```

> [side by side js and coffee][first]

So, we just learn that prototyped methods can be accessed with an object
instance. Let's go to the next level.

## Constructors

In JS, just like in Java and other languages, the constructor is a method
with the name of the class and no return:

```js
var Test = (function() {
  function Test(p) {
    this.p = p;
  }

  Test.prototype.t = function() {
    return alert("test " + this.p);
  };

  return Test;
})();

var o = new Test("this!");
o.t();
```

**Notice**: this time we had to wrap the function context inside an anonymous
function that "calls itself".

Now, the same example in CoffeeScript:

```coffeescript
class Test
  constructor: (@p) ->

  t: ->
    alert "test #{@p}"

o = new Test("that!")
o.t()
```

> [side by side js and coffee][second]

That's pretty cool, but, what if we want a default value to the parameters at
the constructor? If you code Ruby, just do it in the same way:

```coffeescript
class Test
  constructor: (@p = "what?") ->

  t: ->
    alert "test #{@p}"

o = new Test()
o.t()

oo = new Test("this!")
oo.t()
```

> [side by side js and coffee][third]

## Scopes

Like every OOP language, you have scopes to play with. Most of languages has
`private`, `protected` and `public`, but, JS has only `public` and `private`
scopes (but you can always [hack some code](http://nemisj.com/protected-javascript/)
to archieve that.

#### What you can do and what you can't

- You **can** call a private funcion from a public one;
- You **can't** call a public function from a private one.

Which basically means that this:

```coffeescript
class Test
  tt = ->
    alert 'tt'

  t: ->
    tt()

new Test().t()
```

> [side by side js and coffee][fourth]

while this:

```coffeescript
class Test

  ttt: ->
    alert 'ttt'

  tt = ->
    @ttt()

  t: ->
    tt()

new Test().t()
```

> [side by side js and coffee][fifth]

don't, like I just said above.

**That's it.**

## Inheritance

Yeah, JS also has inheritance support, even if it much like a hack.

Archieve that with CoffeeScript is **really** simple, pretty much like Java
or other language:

```coffeescript
class Test
  t: ->
    alert 'as'


class TT extends Test
  tt: ->
    @t()
  ttt: ->
    alert 'tttt'

t = new TT()
t.tt()
t.ttt()
```

> [side by side js and coffee][sixth]

The example is self-explanatory. You can call super methods from your inherited
class, and, obviously, add more methods.

## Bye

Well, I wrote this to be simple. It's pretty simple right now, but, if you have
any question, suggestion, correction, completions (etc etc etc), please, use
the comment box below and left your feedback.

I really hope it help you guys into the wild world of JS-OOP.

[coffeescript]: http://coffeescript.org/
[first]: http://coffeescript.org/#try:class%20Test%0A%20%20t%3A%20-%3E%0A%20%20%20%20alert%20'test'%0A%0Ao%20%3D%20new%20Test()%0Ao.t()
[second]: http://coffeescript.org/#try:class%20Test%0A%20%20constructor%3A%20(%40p)%20-%3E%0A%0A%20%20t%3A%20-%3E%0A%20%20%20%20alert%20%22test%20%23%7B%40p%7D%22%0A%0Ao%20%3D%20new%20Test(%22that!%22)%0Ao.t()
[third]: http://coffeescript.org/#try:class%20Test%0A%20%20constructor%3A%20(%40p%20%3D%20%22what%3F%22)%20-%3E%0A%0A%20%20t%3A%20-%3E%0A%20%20%20%20alert%20%22test%20%23%7B%40p%7D%22%0A%0Ao%20%3D%20new%20Test()%0Ao.t()%0A%0Aoo%20%3D%20new%20Test(%22this!%22)%0Aoo.t()
[fourth]: http://coffeescript.org/#try:class%20Test%0A%20%20tt%20%3D%20-%3E%0A%20%20%20%20alert%20'tt'%0A%0A%20%20t%3A%20-%3E%0A%20%20%20%20tt()%0A%0Anew%20Test().t()
[fifth]: http://coffeescript.org/#try:class%20Test%0A%0A%20%20ttt%3A%20-%3E%0A%20%20%20%20alert%20'ttt'%0A%0A%20%20tt%20%3D%20-%3E%0A%20%20%20%20%40ttt()%0A%0A%20%20t%3A%20-%3E%0A%20%20%20%20tt()%0A%0Anew%20Test().t()
[sixth]: http://coffeescript.org/#try:class%20Test%0A%20%20t%3A%20-%3E%0A%20%20%20%20alert%20'as'%0A%0A%0Aclass%20TT%20extends%20Test%0A%20%20tt%3A%20-%3E%0A%20%20%20%20%40t()%0A%20%20ttt%3A%20-%3E%0A%20%20%20%20alert%20'tttt'%0A%0At%20%3D%20new%20TT()%0At.tt()%0At.ttt()

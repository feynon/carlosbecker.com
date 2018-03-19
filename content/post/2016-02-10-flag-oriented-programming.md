---
date: 2016-02-10T00:00:00Z
slug: flag-oriented-programming
title: Flag-oriented Programming
---

`booleans`, right? What a wonderful piece of technology! They help us solve so
many problems...

> I just need this method to behave slightly different when
> some condition is `true`.

Nice, what's the problem in that?

We had a method like this:

```ruby
class Coffee
  def initialize(size)
    @size = size
  end

  def make
    # do stuff to make coffee with the given @size
  end
end

Coffee.new(:large).make
```

Now, we have something like this:

```ruby
class Coffee
  def initialize(size)
    @size = size
  end

  def make(milk)
    # do stuff to make coffee with the given @size
    if milk
      # actually make that previous coffee be a latte
    end
  end
end

Coffee.new(:large).make(true)
```

So, what's wrong?

At the first sight, you might think it is OK: "Just a single flag. This
wouldn't hurt anybody..."

The problem is that, later on, someone else will like to add another feature,
for example, whether to add or not sugar...

```ruby
class Coffee
  def initialize(size)
    @size = size
  end

  def make(milk, sugar)
    # do stuff to make coffee with the given @size
    if milk
      # actually make that previous coffee be a latte
    end
    if sugar
      # actually make that previous coffee or latte, a sweet
      # coffee or sweet latte
    end
  end
end

Coffee.new(:large).make(true, false)
```

Now the method is slightly more complicated than before (imagine that real
code will be in there instead of comments).

What happens next?

Of course, another feature:

```ruby
class Coffee
  def initialize(size)
    @size = size
  end

  def make(milk, sugar, cream)
    # do stuff to make coffee with the given @size
    if milk
      # actually make that previous coffee be a latte
    end
    if sugar
      # actually make that previous coffee or latte, a
      # sweet coffee or sweet latte
    end
    if cream
      # actually make that coffee, sweet coffee, latte or
      # sweet latte a coffee with cream, sweet coffee with
      # cream, latte with or sweet latte with cream
    end
  end
end

Coffee.new(:large).make(true, true, true)
```

I dare you to guess what happens next...

At this point, the code is already a mess. Maybe the developers will attempt
to apply some clean code and put the contents of each `if` statement in
separated functions... but, is this [good design]({{< ref "2015-12-27-good-code.md" >}}), at all?

Of course not. This is imperative, flag-oriented programming.

Also: look at the method call. Who will know in 3 months from now what each
one of those booleans are?

Exactly: **nobody**.

Maybe this problem could be better solved with decorators, for example:

```ruby
class Coffee
  def initialize(size)
    @size = size
  end

  def make
    # make a coffee with the given @size
  end
end

class Latte
  def initialize(coffee)
    @coffee = coffee
  end

  def make
    @coffee.make
    # adds milk to the given @coffee
  end
end

class SweetCoffee
  def initialize(coffee)
    @coffee = coffee
  end

  def make
    @coffee.make
    # adds sugar to the given @coffee
  end
end

class WhippedCreamCoffee
  def initialize(coffee)
    @coffee = coffee
  end

  def make
    @coffee.make
    # adds whipped cream to the given @coffee
  end
end

# just plain old black coffee
Coffee.new(:small).make

# a latte
Latte.new(
  Coffee.new(:large)
).make

# a latte with whipped cream
WhippedCreamCoffee.new(
  Latte.new(
    Coffee.new(:medium)
  )
).make

# a sweet latte with whipped cream
WhippedCreamCoffee.new(
  Latte.new(
    SweetCoffee.new(
      Coffee.new(:large)
    )
  )
).make
```

Yes, it is a little more verbose, but in the other hand it is more readable,
maintainable, extendable, testable, etc... and, of course, more object oriented.

I know it feels like a lot of work (and of classes), but that is how Objects
are intended to be: small, doing one thing very well, composable, and that's it.

## Just blame the next guy...

Looking back at the beginning of the post, one can argue that the second
programmer shouldn't have added more flags.

Yeah, maybe... but... **humans**... right?

We often do what we know is not the right thing because... _reasons_. I have
no intend to approach the psychology aspect of this, but I do recommend
[this book][irrational], if you want to learn something about it.

The big question is: will you trust that, given a method with one flag already,
the next programmer right away fix the mess and do the right thing?

I'm sorry, but I believe it is more likely to happen this way:

- Programmer X (the next guy - given `next = n`), arrives late at the office
because of a flat tire - not very happy already;
- Programmer X needs to add another feature to the coffee making class,
let's say, whether or not to add cinnamon, of course, as fast as possible;
- Programmer X opens the `Coffee.rb` file, takes a look around and start
cursing someone - because he couldn't stand to look at that hundred `if`
statements;
- Programmer X then refactor the entire code - commit message has some
curse words in it;
- Programmer X then finally implements the Cinnamon feature and its tests -
commit message again has some curse words in it;
- At this time, Programmer X's day is already ruined. He is mad. And he knows
it was you (`git blame coffee.rb`). And he knows where you live.

Why is he mad?

Because he spent 3 hours refactoring the code to finally being able to
implement his feature in a decent way, while he could have done that in 30
minutes **if the previous guys weren't lazy**.

Do you want to be the "Programmer X" in this case? Do you want to be one of the
lazy guys?

I know I don't.

---

This is my point of view on this subject, if you disagree,
please, comment above, let's discuss how to write better code.

And, just to be clear, there could be N other implementations, decorators are
**not** a silver bullet. Use them wisely.

[irrational]: http://amzn.to/1SH7rJV

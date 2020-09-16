---
title: "Seven Languages in Seven Weeks: Erlang"
date: 2012-12-18
draft: false
slug: seven-languages-in-seven-weeks-erlang
city: Joinville
toc: true
tags: [erlang]
---

So, I just bought *["Seven Languages in Seven Weeks"](http://amzn.to/17vnhdb)*. I've read the Ruby chapter, not big deal at all, so I slipt it. I also skip Io, Prolog and Scala (for now), and then, fall in Erlang!

I'll explain it better another day, but, as you may know, Erlang is a concurrency language, that used to turn difficult thing easy and easy things difficult. There is a price to be paid for solve the world first problems.

![](/public/images/seven-languages-in-seven-weeks-erlang/e5a7c888-be2e-4d2d-ba3d-00222a37437a.png)

So, I'll share with you guys my solutions to the problems proposed by the book.

## Day 1

### Problem 1

> Write a function that uses recursion to return the number of words in a string.

Solution:

```erlang
list_length([]) -> 0;
list_length(String) ->
  [_ | Tail] = String,
  1 + list_length(Tail).
count_words(Text) ->
  R = re:split(Text, " "),
  list_length(R).
```

Just a simple pattern matching and a tail recursion in list length, and in `count_words` we has regex to split the text by space, to actually get the words. Sure, there is simple ways to solve that, but, the fun is to do in a "complicated" way. Erlang already has a function to count words, FYI.

### Problem 2

> Write a function that uses recursion to count to ten.
```erlang
count_until(Val, Max) when Val < Max ->
  io:fwrite("~w~n", [Val]),
  count_until(Val + 1, Max);
count_until(_, Max) ->
  io:fwrite("~w~n", [Max]).
count_until(Max) ->
  count_until(0, Max).
```

Nothing new here. At all.

### Problem 3

> Write a function that uses matching to selectively print "success" or "error: message" given input in the form `{error, Message}` or success.

What a easy problem. Take the solution:

```
print_msg(success) -> io:fwrite("Success~n");
print_msg({error, Message}) ->
  io:fwrite("sir... we got an error: ~s~n", [Message]).
```
## Day 2

### Problem 1

> Consider a list of keyword-value tuples, such as `[{erlang, "a functional language"}, {ruby, "an OO language"}]`. Write a function that accepts the list and a keyword and returns the associated value for keyword.

Solution with a simple list comprehension:

```erlang
lang(Tuples, Key) ->
  [TupleValue || {TupleKey, TupleValue} <- Tuples, (Key == TupleKey)].
```
### Problem 2

> Consider a shopping list that looks like `[{item quantity price}, ...]`. Write a list comprehension that build a list of items of the form `[{item total_price}, ...]`, where `total_price` is `quantity` times `price`.

Pretty easy, huh? Solution:

```erlang
full_price(List) ->
  [{Name, Price*Quantity} || {Name, Quantity, Price} <- List].
```
### Problem 3

I was pretty lazy, it looks big, I don't want to do it.

## Day 3

Well, day 3 was pretty big. We will now work with multi concurrency.

![](/public/images/seven-languages-in-seven-weeks-erlang/3976f7fc-ec94-4569-b8a8-ca1311a6c5eb.gif)
### Problem 1

> Monitor the translate_service and restart if should it die.

That is a more complex example, so I'll put the [entire file](https://github.com/caarlos0/erlang-playground/blob/master/day3_examples_exs/translate_service.erl):

```erlang
-module(translate_service).
-export([loop/0, translate/2, watch/0]).

loop() ->

  receive
    {From, "casa"} ->
      From ! "house",
      loop();
    {From, "blanca"} ->
      From ! "white",
      loop();
    {From, _} ->
      From ! "Don't get your point..",
      exit("Adiós muchacho!")
end.

translate(To, Word) ->
  To ! {self(), Word},
  receive
    Translation -> Translation
end.


watch() ->
  process_flag(trap_exit, true),
  receive
    new ->
      io:format("~nCreating new etc..~n"),
      register(translator, spawn_link(fun loop/0)),
      watch();
    {'EXIT', From, Reason} ->
      io:format("~nTranslator ~p dies with the reason ~p", [From, Reason]),
      io:format("~nRestarting this crap..."),
      self() ! new,
      watch()
end.

%
% Usage
%
% Translator = spawn(fun translate_service:watch/0).
% Translator ! new.
% translate_service:translate(translator, "casa").
% translate_service:translate(translator, "blanca").
% translate_service:translate(translator, "asdasd").
```

FYI: The original [translate_service](https://github.com/caarlos0/erlang-playground/blob/master/day3_examples/translate_service.erl) implementation.

### Problem 2

> Make the `Doctor` process restart itself if it should die.

That's just like the example above:

```erlang
-module(doctor).
-export([loop/0, watch/0]).

loop() ->
  process_flag(trap_exit, true),
  receive
    new ->
      io:format("Creating and monitoring process.~n"),
      register(revolver, spawn_link(fun roulette:loop/0)),
      loop();
    {'EXIT', From, Reason} ->
      io:format("The shooter ~p died with reason ~p.~n", [From, Reason]),
      io:format("Restarting... ~n"),
      self() ! new,
      loop()
end.

watch() ->
  process_flag(trap_exit, true),
  receive
    new ->
      io:format("Creating and monitoring new Doctor process.~n"),
      register(doctor, spawn_link(fun loop/0)),
      doctor ! new,
      watch();
    {'EXIT', From, Reason} ->
      io:format("The shooter doctor ~p died with reason ~p.~n", [From, Reason]),
      io:format("Restarting... ~n"),
      self() ! new,
      watch()
end.

% usage:
%
% Doc = spawn(fun doctor:loop/0).
% % we must create a new revolve before use it!
% Doc ! new.
% % there we go:
% revolver ! 2.
% revolver ! 3.
```
## Done

That's all for now folks. Yeah, I know that there is some problems missing, obviously, even these I made should had errors. But, well, it works. As I starting to learn erlang now, It's pretty normal made mistakes. The thing is: **LEARN WITH THEM**.

Hope you enjoy Erlang.

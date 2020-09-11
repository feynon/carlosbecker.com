---
title: "I wrote Antigen in Go: Antibody"
date: 2015-06-06
draft: false
slug: go-antibody
city: Marechal Cândido Rondon
---

Learning Go was in my TODO list for a while, and finally I did something about it. This post contains some thoughts about it...

Before we start, I should familiarize you with antigen.

[Antigen](http://github.com/zsh-users/antigen) is particularly slow. With the plugins I have, it takes an eternity to load. I saw an attempt to re-write Antigen in Haskell, but I always believed that Antigen does too much, so I decided that it would be
a nice thing to play with. 

And I also wanted to learn Go... so... there it is: [Antibody](http://github.com/caarlos0/antibody)

## The "strange at first sight"

Most of the language might look strange for people who just came from Java or Ruby (like myself).

I think the worst of them is the first capitalized character modifier:

```go
func privateFunction() {
  fmt.Println("private function...")
}

func PublicFunction() {
  fmt.Println("public function...")
  privateFunction()
}
```

Yes, anything starting with a capitalized letter will be public. Doesn't matter if it's a `struct`, a `function` or a `const`.

It's also a little strange the way you declare variables and method parameters:

```go
// declared an empty variable
var name string

// if you declare and assign in the same line, go can infer the type
age := 21

// constants assignment work with `=` only.
const Pi = 3.14

// name comes first in arguments... and the return type is the last
func sum(a int, b int) int {
  return a + b
}
```

The way you declare a method in a `struct` is also different from what I was used to in other languages:

```go
// this is how you define a type
type Person struct {
  Name string
}

// defining the Surname function for Person
func (p *Person) Surname() string {
  name := strings.Split(p.Name, " ")
  return name[len(name)-1]
}

// and this is how you use them:
person := Person{"Carlos Alexandro Becker"}
fmt.Println(person.Surname())
```

This all really seamed strange in the beginning, but now I'm actually liking it. 😌

## The Niceties

The niceties includes `go fmt`. Ever heard about those endless discussions about 2 or 4 spaces or tabs, or whether to put a `;` or not, or where the `{`  should be? Well, `go fmt` solves that. It's a standard tool to format code in the Go defaults. 

You cannot change it, so, it is what it is.

Other tool above my expectations was `go build`. `go build` will break if there are unused imports or variables (as well if there is any error in your code). I wish I had something like this for all languages I ever worked with. Java just turned 20 years old, and we don't yet have any of this...

## The Nicest

But, the nicest thing I played with until now surely is `goroutines`. 

A `goroutine` is basically a lightweight thread of execution. I used it to support a list of bundles coming from `STDIN` (like Antigen does):

```sh
$ antibody apply < bundles.txt
```

So, given that `bundles.txt` contains a valid list of plugins, it will load all of them.

Antibody will iterate the lines and check if the respective folder exists, in case it doesn't, it will clone it.

This (cloning stuff) can be very slow, so, I did it with `goroutines`, so the lines could be checked concurrently. I also used a `sync.WaitGroup` to wait for all lines to complete processing before exit (avoiding premature termination):

```go
func process(bundle string, home string, wg *sync.WaitGroup) {
  defer wg.Done()
  Bundle(bundle, home)
}

func ProcessStdin(stdin io.Reader, home string) {
  var wg sync.WaitGroup
  bundles, _ := ioutil.ReadAll(stdin)
  for _, bundle := range strings.Split(string(bundles), "\n") {
    if bundle != "" {
      wg.Add(1)
      go process(bundle, home, &wg)
    }
  }
  wg.Wait()
}
```

I found this very simple to use compared with other languages. Oh, did I mention the ability to send messages through `channels`? It kind of make me remember [Erlang's message passing](http://erlang.org/doc/getting_started/conc_prog.html).

## So, What's next?

Well, now it's time to start learning how to do web apps and how to deal with databases and stuff. This might be actually fun. 🍺

## Summing up

After some time learning the language basics and getting used to the syntax, Go now is for me very pleasant to write, besides all the "strange" stuff.

Also, by strange, I don't mean that Go is strange. I mean that it might be strange for people coming from Java, for example.

About Antibody: I know, the code probably is not the nicest piece of Go code you ever read. It is my first Go project, so I yet have some stuff to learn (any feedback will be very welcome, by the way). Anyway, it was a nice experience for a useless holiday, and as a bonus it improved my shell performance!

> I also shared it in both [zsh](http://www.reddit.com/r/zsh/comments/38lt3h/caarlos0antibody_a_faster_and_simpler_version_of/) and [golang](http://www.reddit.com/r/golang/comments/38t3r7/caarlos0antibody_a_faster_and_simpler_antigen/) Reddits.

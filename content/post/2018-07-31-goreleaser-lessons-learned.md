---
title: "GoReleaser: lessons learned so far"
date: 2018-07-31
slug: goreleaser-lessons-learned
city: Joinville
tags:
- goreleaser
- go
- opensource
---

I've started [GoReleaser] almost 2 years ago. This is a summary of (some)
things I've learned down the road.

<!--more-->

{{< figure src="https://github.com/goreleaser/artwork/raw/master/goreleaserfundo.png" alt="goreleaser new logo" height="200px" >}}

I already talked about [GoReleaser] here a few times, if you feel like reading
about it first:

- [Announce]({{< ref "post/2017-01-02-goreleaser.md" >}});
- [1k users]({{< ref "post/2018-04-09-goreleaser-1k-repos.md" >}}).

I tried to organize things in subtopics, some of them are bigger may go
further into the subject than others.

Without further due, let's get started!

# Naming things is hard

> There are only two hard things in Computer Science: cache invalidation and
> naming things.
>
> --- Phil Karlton

People misreading [GoReleaser] as "gore leaser" is quite common.
I thought it was a really good name because it has "Go" and "Releaser", so,
yeah, it releases Go projects, pretty easy to figure out what is going on
there.

But I totally forgot about trying it all lowercase and without context.

On the bright side, "gore leaser" would be a very good heavy metal band name,
as someone already pointed out on Twitter (I just need a band and more
guitar skills 😂):

{{< tweet 1017028524840976385 >}}

I still think the name is OK though, after all, it is easy to write and it
tells what the software does. What can be bad is that the scope is very
reduced: what if I decide to also release Rust projects? [^fn:rust]

Maybe I'm just real bad at naming things. 🤷‍♂️

# The internal context package

At first glance, it seemed like a good idea. I needed to transport the
"context" of the release through the pipeline, and also needed the features of
the "regular" `context` package (like cancellation). So I created an internal
`context` package, which holds all the data I need plus a `context.Context`
instance, so I can use them interchangeably.

The good thing is that things are checked at compile time: I know I
that can access `ctx.Config.ProjectName` and that it will work.

The bad thing about it is that it can be confusing for new contributors,
as most of them will expect `context` to be the language's
[context package](https://golang.org/pkg/context/).

I'm not sure that was a good idea, but it is confusing. Maybe I could at
least rename the package - but that would require renaming things in a lot
of files, so I keep postponing it.

# Tests using the same fake data

[GoReleaser] can do a lot of things and because of that it also has a lot of
tests, and some of them are kind of complex.

A good example is the Docker pipe tests, in which I have things like this:

```go
var table = map[string]struct {
  dockers     []config.Docker
  publish     bool
  expect      []string
  assertError errChecker
}{
  // a lot of test cases
}
// later:
for name, docker := range table {
  t.Run(name, func(tt *testing.T) {
    // actually run the tests
  })
}
```

This suite has several tests that create a lot of images. All of which
were sharing the same image name and binary name.

When a test failed, it was hard to figure things out based on logs - especially
when you have table-driven tests [^fn:tdt] - which is the case for this example.

I've been slowly fixing this kind of thing, so the fake data is unique - most
times by using the test name or something like that. It helps **a lot**.

A cool trick I've learned reading other's people code: the `errChecker`
interface I have on this suite:

```go
type errChecker func(*testing.T, error)
```

I also have some "helper" functions:

```go
var shouldErr = func(msg string) errChecker {
  return func(t *testing.T, err error) {
    assert.Error(t, err)
    assert.Contains(t, err.Error(), msg)
  }
}
var shouldNotErr = func(t *testing.T, err error) {
  assert.NoError(t, err)
}
```

Then, on my cases I can have things like:

```go
"successfull test case": {
  // omited details for the sake of brevity
  assertError: shouldNotErr,
},
"bad template test case": {
  // omited details for the sake of brevity
  assertError: shouldErr(`template: tmpl:1: unexpected "}" in operand`),
},
```

With that I can have a lot of complex test cases and wrap them all without
repeating code.

Of course, you can use that for other things as well instead of only errors.

# Dependencies on 3rd parties and weird environments

This was the main reason I decide to [nfpm]. I was using [fpm], which is good,
but **people will have weird environments**. Random versions of things - or
just very old versions of things, weird `PATH` setups and a whole lot of
things you were not expecting.

You can either "guard" your software against that - say, works only on version
v1.4.2 of something, make the code work with multiple versions, or remove the
dependency.

What I've learned is that the less 3rd parties you depend on, the better.
More importantly, I've learned that sometimes it is just not worth it to
remove some dependency from a 3rd party.

For example, [GoReleaser] depended on [fpm]. At some point, I decided to remove
the dependency entirely because it was generating a lot of bugs and random
build failures.

To be able to do that, I wrote [nfpm], which is guarded against multiple
versions of `rpmbuild`, its only external dependency.

By removing [fpm], I also removed the dependencies on `tar`, Ruby, `gem` and
possibly others.

I could probably write `rpmbuild` in Go as well, so I could get rid of that
one last dependency, but, is it worth it?

I think it would have a bad return over the investiment of my time, so, **no**.
There is no lib to deal with that, and RPM packages seem to be really complex
to generate. Plus, `rpmbuild` is already distributed for all major platforms.

Was removing [fpm] a good investment of my time? I think so, yes: I have
way fewer reports of "deb packaging not working" and way less
unstable builds.

When writing [nfpm], I've also learned:

- how debs are packaged [^fn:debs];
- that RPM is nearly impossible to package without `rpmbuild` [^fn:rpmbuild];
- people have some really complex packaging needs - and that's why packaging
software is complex.

> **Important**: is not that [fpm] is not good, it is awesome software! I just
> didn't want to guard [GoReleaser] against all the combinations of things that
> could go wrong and I didn't need all its features either. If you need to
> package your software in a lot of formats using a single tool,
> [fpm] for the win!

# Documentation is hard

I've learned that it is very hard to get just the exact amount of
documentation so it doesn't suck.

If you write too much, people probably won't read. It is also likely to get
too complex, thus also hard to grasp.
If you write too little, people maybe will read, but will not learn all things
they need.

Writing more docs also eventually leads to more complicated and confusing
docs - just like writing more code, who knew!?

I tried to provide some kind of commented config examples, thinking it may
be straightforward enough (as people could copy and change), but, just as
most commented config files out there, most people won't read them.
And I don't blame them, I should probably do a better job on that.

I still don't know the right/best way of doing this, just learned one more
non-optimal way. If you do, I would love to chat about it though!

# Monorepo or not?

In the beginning, I split the [archive] package into another repo, because I
though it would be useful to other people as well.

[I was wrong](https://github.com/search?l=Go&q=%22goreleaser%2Farchive%22&type=Code).

> I'll eventually move it into goreleaser's tree very soon.

I don't usually work with monorepos, so I haven't planned [GoReleaser] to work
in that way. A few things are already fixed, but it still won't support
releasing each [artifact with a separated tag][i723], for example.

I think that in the particular case of [GoReleaser], a monorepo will be
easier to manage:

- most changes on [nfpm] require a `dep ensure -update` on [GoReleaser];
- if the above issue was fixed, it would be easy enough to still release
[nfpm] in its own tag;
- [GoDownloader][] already stopped working several times due to changes on
[GoReleaser];
- the [archive] package doesn't make sense alone;
- the docs source are already in the same repository, which is very handy;
- all issues would be in a single place.

So, yes, I probably should have gone with a monorepo.

Is not that a monorepo is the fix for all problems though.

I think the real issue is that I forced myself to split things too early.
Maybe [nfpm] and [godownloader] are less wrong, but [archive] for sure was a
mistake. Premature optimization... of sorts.

# Data structures and relationships

Someone way smarter than me once said:

> Bad programmers worry about the code.
> Good programmers worry about data structures and their relationships.
>
> --- Linus Torvalds

Remember `context` from before, right? It was even worse. The way the
artifacts were stored in context was in way that work great for a few
artifacts, not so great with several artifacts and several kinds of
artifacts. The way it was lead to a lot of bad code all spread across almost
all pipes.

Later on, I've added the `artifact` package, which abstract this into a
more proper data structure, adding also a nice DSL to filter artifacts
by kind and etc - and [I had to refactor a lot of things for that][pr463].

So, now if I want to upload all Linux packages for `amd64` in a pipe,
for example, I can:

```go
ctx.Artifacts.Filter(
  artifact.And(
    artifact.ByGoarch("amd64"),
    artifact.ByType(artifact.LinuxPackage),
  ),
).List()
```

I'm still not sure that current form is the best solution, but for sure is
better than the first ones. Another good catch here is that I can replace the
internals of how it works while keeping the way pipes access it.

To summarize, I've learned, in the hard way, that bad decisions on the data
structures lead to bad decisions on the interactions between them,
which lead to bad design in general.

# Yes is forever

It is not easy to be the product owner and the developer at the same time.

As a developer, it is easy to say "yes" to things. It is just more code, right?
It reminds me of this Tweet:

{{< tweet 532572556159488000 >}}

Assuming that `errors = (more code)^2` and `more features = more code` are both
true, how do I decide if something goes in or not?

I have been trying to follow what Solomon Hykes [^fn:hykes] once said:

{{< tweet 715277134978113536 >}}

Some things are, like, just... cute... you will regret adding those.
Other people also learned adding things because it's kind of cool may not be
a very good idea:

{{< youtube id="M3BM9TB-8yA" >}}

The entire talk is great, but the point I refer to is at 14:38.

The reasoning for the "yes is forever" thing is that if you say _no_ to
something, you can always go back later and say _yes_ if you change your mind,
but once something is in, you can never take it out without breaking changes.

So, having that in mind, when I look at a feature request, I always try to ask
myself:

> --- Am I 100% sure that this should go in?
>
> --- Am I 100% sure that this is the way we should do it?

I usually don't rush into deciding that. I take a look, think what I have to
think about it, give it a day or two (or more), take another look and think
again. If the answer to both questions was yes both times, I "say yes" to the
PR/feature request/whatever. Otherwise, I suggest what I think that needs to
change (and explain why), or just say why I don't want that as nicely
as I can and close as `wontfix`.

Even with all that in mind, it is hard to say _no_. People spend time working
on that pull request and they are full of good intentions.
Maybe it is their first pull request ever - or the person is just really
excited about it, who knows?!

Most maintainers try to be nice - I surely do, but I know that it can still
be a bad experience.

One thing that I've learned and that I think helps: open an issue first,
ask if the maintainer would be interested in a PR implementing the feature
you want to implement. Someone did that on [GoReleaser] (I think) and I
believe it is a great way of saving everyone's time!

If the maintainer says no but you still really do want that feature,
keep a fork. Everybody wins. 🙂

# Announce breaking changes to your users is hard

Technically, [GoReleaser] is still not v1, so it should mean that I could just
break stuff... of course, I don't want to do that. I want the transitions
to be as easy and painless as they can be.

That's why I've added deprecation notices to the
[docs](https://goreleaser.com/deprecations/), and when you run
`goreleaser` with a deprecated config, it will put a `WARN` log pointing to
that URL.

I still think there should be a way of making that more visible to users...
maybe adding a 1min sleep or something can be a valid approach, but I'm still
not sure about it.

I've also learned it is hard to find if someone is using that thing you
want to deprecate. Maybe I should add some kind of tracking? Don't know.

I've learned that I just don't know how to handle those things on GoReleaser
because people may not even read the log unless it fails (e.g. running on
the CI).

# No one owes anyone anything

We (people) usually don't read terms, licenses and etc because it is boring.
I release most of my software under the [MIT][] license, including [GoReleaser].
Summarizing: the MIT license says you can basically do whatever with my
software **as is**.

That means that you can open an issue with some bug or feature request and
I could literally play [Battlefield][bf] forever instead of fixing it
(I really could, just look at those shiny stats that would be way better
if I played sober most of the time haha).

I think I've learned things on both sides of the coin:

1. **as an user**: don't push. If you are in a hurry, fix it yourself, provide a
PR and use your custom build meanwhile. Don't even expect it to get fixed -
ever, it may not. The maintainer doesn't owe you anything.
2. **as a maintainer**: don't push yourself too much. It is fine to just rest,
this is not part of your daily job, you're not getting paid for it (probably).
You don't owe your users anything.

Or, as I like to say, **no one owes anyone anything**.

# Slack is not the best place to ask questions

In the issue template, I ask people to ask questions on Slack. I think that
is not optimal, as most people will try to search for they problem on
Google, and Google does not index Slack conversations.

I've learned that probably the best place for questions on rather small
communities is GitHub issues. I've seen bigger ones, like Hugo's, using
[Discourse](https://discourse.gohugo.io/), and it seems to serve them well.
On GoReleaser's case that seems overkill.

# Famous last words

I could probably write about more things, for sure. Some topics were still
really hard for me to externalize in words in a form that makes sense, so I
end up removing them... at least for now.

Anyway, I hope that the reading of my many screw ups was interesting and that
you enjoyed it. If not, please feel free to comment/complain bellow or contact
me in any way (except maybe phone calls haha)! I'll be glad to discuss about
it and maybe learn that I was wrong about one more thing.

> **Spoiler**: I'll talk more or less about those topics at
> [GopherCon Brazil][con], so your feedback is greatly appreciated!
> Hope to see you there!

<!-- footnotes -->
[^fn:debs]: Basically two `tar.gz` files inside an `ar` file, one of the `tar.gz` files has the software itself, the other has control files.
[^fn:rpmbuild]: `rpmbuild` is a CLI provided by RedHat to build `.rpm` packages from a spec file.
[^fn:tdt]: Table-driven tests are useful if you need to copy/paste some structure on a lot of tests. You learn more about it [here](https://github.com/golang/go/wiki/TableDrivenTests).
[^fn:rust]: Adding Rust support to [GoReleaser] was discussed [some times](https://github.com/goreleaser/goreleaser/issues?utf8=%E2%9C%93&q=rust), but was still not implemented.
[^fn:hykes]: One of the founders of Docker and dotCloud.

<!-- links -->
[bep]: https://github.com/bep
[tj]: https://github.com/tj
[fpm]: https://github.com/jordansissel/fpm
[nfpm]: https://github.com/goreleaser/nfpm
[archive]: https://github.com/goreleaser/archive
[GoDownloader]: https://github.com/goreleaser/godownloader
[i723]: https://github.com/goreleaser/goreleaser/pull/723
[pr463]: https://github.com/goreleaser/goreleaser/pull/463
[MIT]: https://choosealicense.com/licenses/mit/
[bf]: https://battlefieldtracker.com/bf1/profile/xbox/caarlos0
[con]: https://2018.gopherconbr.org/
[goreleaser]: http://goreleaser.com

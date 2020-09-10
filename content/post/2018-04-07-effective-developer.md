---
title: "On being an effective developer"
date: 2018-04-07
draft: false
slug: effective-developer
city: Joinville
---

Over the years I read several articles on how to be effective, and how the 10x engineer thing is or is not a lie and all that.

---

I think this "being effective" varies a lot from one to another, so it is probably really hard to be 100% accurate with anything.

That being said, friends and coworkers eventually ask me what I do to be effective/faster, and I never have anything solid to say.

Overall, I’m just another engineer. I’m not smart and I for sure don’t know everything that is to know about any topic. I’m for sure not the most effective engineer out there either. But, I think we can all teach something to others, so this is my take on the subject.

---

I tried to break it down into a few topics on things I **try** to do:

- study with focus;
- try to really understand the problem;
- fix things early;
- improve my dev environment;
- reduce noise;
- k.i.s.s.;
- practice.

Hope this helps some of you and that you like it.

---

## Study with a focus

Choose one thing to learn and learn it well.

Study until you can discuss it with someone who you know and knows a lot about that topic. Discuss the topic with that person.

Study until you can explain how the thing works - not only how to use it. Explain it to someone interested.

The idea is simple: how can you be effective at using something you don’t understand? Sure it is easy to use a framework that gives you everything, but what will you do when the framework fails and you don’t know why?

Probably copy some crappy hack from StackOverflow.

Don’t get me wrong, I learned a lot on StackOverflow, but you need to want to learn, not only to fix your current problem.

It’s also OK to use frameworks if you understand how they work.

The point is, choose something to learn, and go deep.

## Try to really understand the problem

It’s kind of common to have a situation in which a problem occurs, and someone will give you a playbook, a script, or something else to follow to fix that problem.

**Ask why.**

Why does this script fix that problem? What is the cause of the problem? How does this thing fix that?

Sometimes its just some bad config or misuse and you can fix the problem on its roots. If it is not the case, you’ll learn something new, which is always good.

## Fix things early

> "I’ll fix this later"

No, you will probably not. Stop lying to yourself.

My take on this is:

- if the problem is small, fix it;
- if the problem is big, document your finding on a GitHub issue or a Jira ticket or whatever you use, bring it up when you can.

Other than that, it’s all excuses - and we all do that because it is **so easy** to say "I’ll fix it later" and forgetting about it.

To hack myself, I usually think something like:

> If I leave this for later, I’ll forget - but the issue will still be there. It will haunt me in the worst way and timing possible, probably while I’m out with my wife having a good time or watching a movie or doing anything else. I don’t want that to happen.

Of course, having it documented and not fixing it won’t prevent this, but here is a curious thing: if I write things down, it’s easier for me to remember it later.

It’s like when you start asking a question on StackOverflow and actually solves the issue while writing it down.

Just like a rubber duck.

The problem sometimes looks bigger than it really is.

If it is really big (which can happen), I now have a GitHub issue for it - and I don’t like to leave issues open.

## Improve my dev environment

I don’t mean buying new hardware or changing OSes or anything like that.

I spend most of my time on my terminal, so it makes sense to know how to use it.

It is amazing the amount of work you can avoid by knowing some shell script and piping things. If you read my blog and attend to details, you’ll see that sometimes I post my shell witchcraft here, for example, [when I migrated a lot of repositories from one CI to another]({{< ref "2017-03-11-travis-to-buildkite.md" >}}).

I also have a lot of automation in place. Common tasks should be fast to execute or execute themselves.

My [dotfiles repository](https://github.com/caarlos0/dotfiles) has a lot of things in the way I like: aliases, helpers, plugins and etc.

I also occasionally write new tools for problems I need. For example:

- [clone-org](https://github.com/caarlos0/clone-org): to clone all repositories in of given GitHub org/user
- [antibody](http://getantibody.github.io/): a faster zsh plugin manager
- [goreleaser](https://goreleaser.com/): automates release processes

And there are more.

My thought process on this is simple: since I’m in theory a software engineer, I should solve problems with software, not by repeatedly clicking things.

Summing up:

- learn your tools well
- make them faster
- write new tools if there is none

I would also advocate learning to shell script, but, whatever works best for you, learn it and then improve it. **Make the machine work for you**, not the other way around.

## Reduce noise

I try hard to keep the noise and distraction at a low level - which is extremely hard.

The problem is that, nowadays, most of the available information is bullshit, and I don’t want to spend my time on that. [Death to bullshit](http://deathtobullshit.com/)!

As James Gleick once said:

> When information is cheap, attention becomes expensive.

I have been doing some things to guard myself against this bullshit river:

- filter crap emails (also Gmail’s mute feature is awesome);
- [#QuitFacebook](https://twitter.com/hashtag/QuitFacebook);
- try to keep my [#InboxZero](https://twitter.com/hashtag/inboxZero) - most of the time by ignoring most of the emails;
- mute all notifications (both on OS and smartphone) while working;
- disable notifications from social media apps;
- mute slack channels (or just close it sometimes);
- if the office is too noisy, work from home;
- if I didn’t read that tab in a 1 week, I won’t read it anymore: close it;
- good headphones.

Not perfect, but it has been helping me so far.

## K.I.S.S.

Keep things simple.

You probably don’t need spreadsheets, todo lists and a lot of things to fill with tasks that you probably won’t do anyway.

Have one source of truth for each "kind" of thing.

I use:

- GitHub issues for software (and `TODO`s on the code);
- Calendar for the rest (I recommend [Fantastical](https://flexibits.com/fantastical)).

Again, not perfect, but it’s way better than 10 apps and 321 lists if you ask me.

## Practice

Like any other thing in the word, it takes practice - and practice takes time.

So, don’t be too hasty. Practice every day, read other people’s code, maybe join some OpenSource project… eventually, you will get there.

---

That’s what I have to say. Hope it helps someone and that it makes sense to someone.

If you have your own tips (I know you do) and feel like sharing them, please, leave a comment - it will be greatly appreciated.

I also want to leave a very special thanks to [Tiago Matias](https://github.com/tmatias) for the review and for the unintentional inspiration to write this.
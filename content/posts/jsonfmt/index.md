---
title: "Keeping json files formatted"
date: 2018-11-12
draft: false
slug: jsonfmt
city: Joinville
toc: true
tags: [golang, json, cli]
---

I'm working in a project that uses Chef, so it has a lot of JSON files.

I like everything to be preperly formatted. The choice of format itself does not bother me much, giving that there is one. 

This project did not (although it was surprisingly not that bad).

---

So, I decided to properly format all files. The first thing that came to mind
was [Prettier](https://prettier.io/), but, since I wanted to add it to the CI as well,
I would rather have something statically compiled and that don't require me
to install 3k dependencies on Jenkins.

I could also use `jq`, but to do all I wanted to do (auto write changes,
print the diff, custom "tab" character and etc) it would require some
probably ugly shell scripting.

So, I searched for `jsonfmt` (as I thought that would be an obvious name),
finding only half-backed and abandoned stuff.

At this point I was like "duh it must be easy let me write this thing",
already having in mind that I would definitely learn something new.

I decided to use Go because I'm used to it and it has been great to write
CLI tools using it.

## First attempt

My first thought was to unmarshal to a `map[string]interface{}` and then
marshal it back, something like:

```go
var tmp map[string]interface{}
_ := json.Unmarshal(input, &tmp)
output, _ := json.MarshalIndent(tmp, "", "  ")
```

To my surprise, it reordered all keys.

The JSON spec does not define key order, so we shouldn't depend on it. While
it's ok and we don't, it would reformat all files differently every time I run
then, so, not what I wanted...

## Second attempt

Searching for things like `json keep order golang` I came across the
the [ordered-json](https://github.com/virtuald/go-ordered-json) lib,
and tried to use it:

```go
var tmp json.OrderedObject
_ := json.Unmarshal(input, &tmp)
output, _ := json.MarshalIndent(tmp, "", "  ")
```

Now the root keys' order were kept, but inner structures had the same issue
as before.

## Third attempt

I decided to step back and take another look at the [JSON package docs](https://golang.org/pkg/encoding/json/)
instead of trusting my memory, and remembered that `json.RawMessage` exists.
It seamed reasonable to me that it would not mess around with the order of the
keys, so I tried it out:

```go
var tmp json.RawMessage
_ := json.Unmarshal(input, &tmp)
output, _ := json.MarshalIndent(tmp, "", "  ")
```

And **BOOM**, it works!

But is it good enough? Do I really need to `unmarshal` and then `marshal` back?

## Fourth attempt

Read the docs for a little while, and found about the `json.Indent` method.
That seems to be it! Let's give it a try:

```go
var out bytes.Buffer
_ = json.Indent(&out, input, "", "  ")
```

And it works almost perfectly: it keeps extra empty lines in the end of the
file.

We can fix that by trimming the white spaces from the input:

```go
var out bytes.Buffer
_ = json.Indent(&out, bytes.TrimSpace(input), "", "  ")
```

And finally, we may add an empty line to the end of the file
([because that's the right thing to do](https://stackoverflow.com/questions/2287967/why-is-it-recommended-to-have-empty-line-in-the-end-of-file)):

```go
out.Write([]byte{'\n'})
```

And that's it!

Once again, the best solution was simpler than all others!

## How I use it

So, basically I have something like this on my makefile:

```makefile
fmt:
	jsonfmt -w

check:
	jsonfmt
```

On the CI, I run `make check` and locally I run `make fmt` (on `pre-commit`).

So, I can have everything on the same formatting, and if someone does not
follow the rule, the CI build will fail and print the differences.

## Can I have it?

Sure you can! The code is OpenSource and can be found on [GitHub](https://github.com/caarlos0/jsonfmt).

You can also install it using several methods:

**homebrew**:

```sh
brew install caarlos0/tap/jsonfmt
```

**snapcraft**:

```sh
snap install jsonfmt
```

**docker**:

```sh
docker run -v $PWD:/data --workdir /data caarlos0/jsonfmt -h
```

**deb/rpm**:

Download the `.deb` or `.rpm` from the [releases page](https://github.com/caarlos0/jsonfmt/releases) and
install with `dpkg -i` and `rpm -i` respectively.

**manually**:

Download the pre-compiled binaries from the [releases page](https://github.com/caarlos0/jsonfmt/releases) or
clone the repo build from source.

## That's all folks

Hope you like it and find it useful somewhere. Feel free to give any feedback
on the comments section below!

I would also like to thank the folks from the [Golang subreddit](https://www.reddit.com/r/golang) for their
great feedback!

Cheers!

---
title: "Integrating Minitest with Shippable"
date: 2015-02-04
draft: false
slug: integrating-minitest-with-shippable
city: Joinville
---

I know, everyone uses Travis. I have nothing against it. But in case you want to test and/or use [Shippable](http://shippable.com/), this might be just the guide for you. I will also show how to setup that nice tabs with the test and coverage reports.

## Rationale

I'm writing this post because I found no blog post or anything compiling all the needed steps, just some parts here and there. I want to make it easy for new users to start testing their stuff using an Continuous Integration system.

So, this is my attempt to compile all that in a *how-to* like blog post.

## The good stuff

First thing we need is the `.shippable.yml` file. Mine looks like this:

```yaml
language: ruby
rvm:
  - 2.1.2
script:
  - bundle exec rake test
env:
  global:
    - CI_REPORTS=shippable/testresults COVERAGE_REPORTS=shippable/codecoverage
```

The `env` is required for that nice tabs I said before to work.

Don't forget to add the required stuff to `Gemfile`:

```ruby
source "https://rubygems.org"
gem "rake" # make sure you have rake here!
# ...
group :development, :test do
  gem "minitest"
  gem "simplecov"
  gem "simplecov-csv"
  gem "minitest-reporters"
end
```

I also created a `Rakefile` (so I can call it in the `.shippable.yml` file):

```ruby
require 'rake/testtask'
Rake::TestTask.new do |t|
  t.pattern = "spec/*_spec.rb"
end
```

The last thing we need is the `spec/spec_helper.rb` file:

```ruby
require "simplecov"
require "simplecov-csv"
SimpleCov.formatters = [
  SimpleCov::Formatter::HTMLFormatter,
  SimpleCov::Formatter::CSVFormatter,
]
SimpleCov.coverage_dir(ENV["COVERAGE_REPORTS"])
SimpleCov.start # you might want to pass 'rails' and/or a block here
require "minitest/autorun"
require "minitest/reporters"
MiniTest::Reporters.use!([
  MiniTest::Reporters::DefaultReporter.new,
  MiniTest::Reporters::JUnitReporter.new(
    ENV["CI_REPORTS"] || "coverage/ci"
  )
])
require_relative "../lib/mylibrary" # change to your library here
```

Most of this stuff is for Shippable to correctly parse test and coverage results. The order of the stuff here **does matter**. Don't change it if you don't know what you're doing.

With that being said, now, our specs will just have to `require` our `spec_helper`, and everything should work like magic when you run `bundle exec rake test`.

## Troubleshooting

I had some trouble with compatibility issues related to Ruby's `minitest`.

To clarify: there is the `minitest` gem and the `minitest` inside Ruby itself. If you use RVM (Shippable does), you will find it under `~/.rvm/rubies/ruby-VERSION/lib/ruby/2.1.0/minitest/`.

Unfortunately for me, I use `rbenv`, and, for some reason I don't yet understand, `rake test` worked beautifully. With Shippable and RVM, I got this:

```
/home/shippable/.rvm/gems/ruby-2.1.2/gems/minitest-5.5.1/lib/minitest/assertions.rb:17: warning: already initialized constant MiniTest::Assertions::UNDEFINED
/home/shippable/.rvm/rubies/ruby-2.1.2/lib/ruby/2.1.0/minitest/unit.rb:80: warning: previous definition of UNDEFINED was here
Emptying /home/shippable/workspace/src/bitbucket.com/repo/repo/shippable/testresults
```

To fix that, I just had to change the `rake test` to `bundle exec rake test` (as already demonstrated in the `.shippable.yml` file in the beginning of this post). 

Simple fix for a hard-to-understand problem.
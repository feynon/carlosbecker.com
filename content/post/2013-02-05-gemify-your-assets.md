---
title: "Gemify your assets"
date: 2013-02-05
draft: false
slug: gemify-your-assets
city: Joinville
---

Rails 3.1 introduced the asset pipeline, which make it easy to include versioned external assets as application dependencies.

Probably you will find almost any JS library you want, already *Gemified*, but, if it is not the case, you can *Gemify* those libraries by your own, and I can help you with it. So, let's do it!

## 1. Create the project skeleton

Bundler makes it simple to create the files and folders necessary for building a Gem:

```
$ bundle gem somelibrary
```

This command will create basically the following tree:

```
$ tree
somelibrary
├── Gemfile
├── lib
│   ├── somelibrary
│   │   └── version.rb
│   └── somelibrary.rb
├── LICENSE.txt
├── Rakefile
├── README.md
└── somelibrary.gemspec
```

But, I usually made some changes to it:

### 1. Delete the `version.rb` file and it's parent folder:

```
$ rm -rf lib/somelibrary
```

> **Heads up**: Don't forget to change the `lib/someligrary.rb` and `somelibrary.gemspec` files!

### 2. As I said in item 1, change the `somelibrary.gemspec` file…

Remove the references to `version.rb`, put a static version, replace the TODO's by your own texts.

### 3. As I said in item 1, change the `lib/somelibrary.rb` file…

You had to force it to load the rails engine:

```
module Somelibrary
    class Engine < ::Rails::Engine
    end
end
```

## 2. Create the `vendor/assets` folder

That's easy!

```
$ mkdir -p vendor/assets/{javascripts,stylesheets}
```

## 3. Copy the assets to the folders

Even easier, just copy the respective `js` and `css` files to `javascripts` and `stylesheets` folders, respectively.

## 4. Test

Now, create some rails app, and add the following to the Gemfile:

```
gem ‘somelibrary', path: ‘absolute/path/to/your/gem'
```

And, off course, add the respective imports in your `application.js` and `application.css` files.

## 5. Document

Yeah, users would like some documentation and a few appointments about how to use your Gem. So, please, **do it**.

## 6. Release

Well, now that the Gem is crafted, tested and documented, we could finally release it.

Create a github repository for your gem, and do the basics:

```
$ git init
$ git add -A
$ git commit -m 'first'
$ git remote add https://github.com/USER/REPO
$ git push origin master
```

And then, finally, release it using Rake:

```
$ rake release
```

This will create a tag for the version specified in your gemspec, push it to github and to rubygems.org.

Simple, uh?

## 7. Enjoy

Now, just use it as a normal gem :)

Hope you enjoyed the post.

Cheers
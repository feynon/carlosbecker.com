---
date: 2016-08-07T00:00:00Z
slug: glide-circleci
title: Setting up a Go build with Glide on CircleCI
tags:
- ci
- golang
---

I've lost a considerable amount of time trying to bind those things together,
so I decided to write this quick post about it, so others could also benefit
from it.

<!--more-->

The tips are:

- Don't install Glide using `go get`. Use `apt`;
- Override `dependencies` and `test` sections;
- You'll have to do some link sorcery so import paths won't be messed up.

The final `circle.yml` would look like this:

```yaml
machine:
  environment:
    IMPORT_PATH: "/home/ubuntu/.go_workspace/src/github.com/myuser"
    APP_PATH: "$IMPORT_PATH/myproject"
dependencies:
  override:
    - sudo add-apt-repository ppa:masterminds/glide -y
    - sudo apt-get update
    - sudo apt-get install glide -y
test:
  pre:
    - mkdir -p "$IMPORT_PATH"
    - ln -sf "$(pwd)" "$APP_PATH"
    - cd "$APP_PATH" && glide install
  override:
    - cd "$APP_PATH" && go test -cover $(glide nv)
```

Replace `myuser` and `myproject` to the username and repository name of your
project on github.

There is probably some easy way to make this work, but I coulnd't find it.

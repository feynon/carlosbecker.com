---
title: "From Travis Enterprise to BuildKite in 50 repositories"
date: 2017-03-11
draft: false
slug: travis-to-buildkite
city: Joinville
---

At [ContaAzul](http://contaazul.com/), we use the CI infrastructure **a lot**. We open several pull requests in several projects every day, and we block the merge until the build pass. We consider our `master` branches are sacred, and we can't afford too much waiting to change them.

## Travis Enterprise

For the past year, we have been using [Travis Enterprise](https://enterprise.travis-ci.com/). While it is pretty good, it also has some problems:

- No worker autoscale. We solved this by writing some Ruby and Terraform scripts, but it was
giving us random build failures;
- It was expensive for us. Dollars are very expensive here in Brazil, and Travis Enterprise costs a lot of them.

So, I went out looking for cheaper and better alternatives.

## BuildKite

My friend [Caio](https://github.com/caiofbpa) told me they were using [Buildkite](https://buildkite.com/) at [MeusPedidos](https://meuspedidos.com.br/) and that it was working well.

I decided to give it a try. My plan was:

1. Setup Buildkite for some repositories;
2. Let both Buildkite and Travis run for a few days;
3. Interview the team to ask them what they think;
4. If the results were good, migrate all repositories and `terminate` some machines.

Setting up the master and the workers was easy. Buildkite provides an Open Source tool called
[Elastic CI Stack for AWS](https://github.com/buildkite/elastic-ci-stack-for-aws).

It works great, letting us even use Spot Instances (which most of the time are cheaper).

I spent some time reading all its code before using it (trust no one). Then, launched a cluster on our sandbox AWS account. After that Buildkite was already working.

Then, I set Buildkite up on the top 3 most changed repositories and let it run for a few days.

## Decision making

After a week, I was sold. I'm not the only user though, so I asked the team to reply a small Google form. Here are the results:

Classify, from 0 to 10, your satisfaction with Travis:

![](Untitled-ac9f82f9-b165-4d44-9aa6-690ab05c380d.png)

Classify, from 0 to 10, your satisfaction with Buildkite:

![](Untitled-bed83eb3-53b5-452f-a992-f1b1ad48ceca.png)

How would you feel if we fully migrate to Buildkite?

![](Untitled-7b085d05-1c44-4484-9683-6b3f0d9e0741.png)

All the scores were good:

All the scores were good:

- Buildkite had higher satisfaction;
- No one voted against migrating to Buildkite.

The only downside I was able to see was the migration process.

## The migration process

I had 50 projects to migrate.

The migration process was:

1. Clone the repository;
2. Remove the `.travis.yml` file;
3. Create a new pipe on Buildkite;
4. Setup GitHub integration;
5. Create `.buildkite/pipeline.yml` and other required files;
6. Open a Pull Request;
7. Check if it worked.

7 steps x 50 repositories = A lot of work.

It would be easier if:

- Buildkite had better GitHub integration;
- Buildkite accepted `.travis.yml` as a fallback or had a conversion tool of some sort.

But none of those were true, so I automated some steps of the process.

## Automating the work

First of all, I cloned all repositories with the help of [clone-org](http://github.com/caarlos0/clone-org):

```shell
$ clone-org --org ContaAzul --destination /tmp/ca
```

Second, I created both Buildkite and GitHub API tokens to script the link between them. Of course I made the script [open-source](https://github.com/caarlos0/github-buildkite-wire).

After that, I created some template files for Java Buildkite builds. Since most of our projects are in Java, I figured that this will do most of the work for me. This step included creating a Docker image with Java, Maven and other tools.

With all that in place, I wrote an "one-liner" and a helper function that would do most of the work for me:

```shell
# helper function
setup_buildkite() {
  cd ~/Code/github-buildkite-wire &&
    ./setup-pipeline.sh ContaAzul "$(basename "$1")"
}
```
```shell
# the one liner
$ find . -depth 1 -type d | while read -r folder; do
  cd "$folder"
  test -f '.travis.yml' &&
    setup_buildkite "$folder" &&
    rm .travis.yml &&
    git checkout -b rm-travis &&
    git commit -am 'removed travis.yml' &&
    cp -r ~/Code/buildkite/template/* . &&
    git add -A &&
    git commit -m 'buildkite :wrench:' &&
    git pr
done
```

PS: The `git pr` alias calls [this plugin](https://github.com/caarlos0/zsh-open-pr)

`git pr` opened one pull request for each project. The majority worked, a few were completely wrong and some needed some tweaks. But, I automated most of my work, allowing me to complete days of work in about 2 hours.

## Final thoughts

Our custom Travis elastic stack wasn't working very well. 

We experienced a lot of random job failures and jobs stuck in queues for no
obvious reason.

Buildkite native elastic stack works great. Their web interface feels faster and cleaner. We also didn't experience any random job failures. 

Having a base Docker image for our builds also allowed us to have faster builds. Travis' base image had a lot of stuff we didn't need and configurations we didn't knew about. We tweaked ours for our use case, dropping the average build time in ~20%.
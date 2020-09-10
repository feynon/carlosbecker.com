---
title: "Configuring a git server with gitolite"
date: 2012-10-03
draft: false
slug: howto-setup-a-git-server
city: Joinville
---

## Setup

You will need:

- git
- git-daemon-run
- ssh
- gitolite

Install everything:

```
$ apt-get install git ssh gitolite git-daemon-run
```

## Generate a SSH key pair

The server needs a public/private key pair. So, you will need to generate it:

```
$ ssh-keygen -t rsa
```

Now, will copy it to `/tmp` folder, we will need it there later.

```
$ cp ~/.ssh/id_rsa.pub /tmp/local.pub
```

## Setup your git profile

```
$ git config --global user.name "Your Name"
$ git config --global user.email your@email.com
```

## Creating the `git` user:

```
$ sudo adduser --system --shell /bin/bash --gecos 'git version control' --group --disabled-password --home /home/git git
$ sudo su git
$ echo "PATH=$HOME/bin:$PATH" > ~/.bashrc
$ gl-setup /tmp/local.pub
```

When you run `gl-setup` command, it will open a file in edit mode.. probably with `vim`. We don't need to change anything here right now. Just save and exit (press `ESC` and type `:wq!`).

## Configuring user access and repositories

Now, go back to the previous user:

```
$ exit
```

Go to some folder, e.g., `~/code`, and clone the configuration gitolite repository:

```
$ git clone git@ubuntu:gitolite-admin.git && cd gitolite-admin
```

Take a look the `conf/gitolite.conf` file. In this file you will be able to configure groups, users and their access to repositories. By example:

```
repo foo
	RW+	= local
	R 	= carlos william
```

In that example, I created a repository called `foo`, with `RW+` access (read and write) for the `local` user (wich represents the `keydir/local.pub` public key) and with `R` access for `carlos` and `william` (`keydir/carlos.pub` and `keydir/william.pub`).

`local`, `carlos` and `william` are machines with that I copied the public key to the `keydir` folder (in the `gitolite-admin` folder), meaning that for every machine you want to give access, you will have to copy their public key.

After doing your changes, commit and push:

```
$ git add -A
$ git commit -m 'added repo foo'
$ git push origin master
```

Now, you will be able to clone the `foo` repo from the configured machines, wich something like:

```
$ git clone git@SERVER:foo.git
```

## Final thoughts

In this tutorial, I used a `ubuntu-12.10-beta2` box, but I believe that it works with older versions too.

If you have any doubt, please share in the comment box above.

Cheers.
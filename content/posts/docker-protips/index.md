---
title: "Docker Protips™"
date: 2015-05-10
draft: false
slug: docker-protips
city: Joinville
toc: true
tags: [docker]
---

Like my old post [on git]({{< ref "/posts/git-protips/index.md" >}}), this is somewhat a collection of useful Docker commands/tricks/whatever.

Feel free to leave yours in the comments!

### Stop all containers

```
$ docker stop $(docker ps -qa)
```
- `ps -qa` will output the `CONTAINER_ID` of all containers;
- `stop` will get `ps -qa` as input and stop all of them.

You can also `kill` all running containers instead of stop them.

### Delete all stopped containers

```
$ docker rm $(docker ps -qa -f="exited=0")
```
- `ps -qa` will output the `CONTAINER_ID` of all containers;
- `-f="exited=0"` flag will tell `docker ps` to filter by exited containers only;
- `rm` will remove the container.

This command is particularly useful if you tested a lot of stuff in your machine which is now running out of disk space.

### Delete all images

```
$ docker rmi -f $(docker images -q)
```
- `images -q` will output the `IMAGE_ID` of all known images;
- `rmi -f` will force delete all the given images.

### Delete unused images

```
$ docker run --rm -v /var/run/docker.sock:/var/run/docker.sock bobrik/image-cleaner
```
- [Read more](https://github.com/bobrik/docker-image-cleaner)

### Kitematic

If you use a Mac, managing `boot2docker` by hand can be a little "boring", you can use [Kitematic](https://kitematic.com/) for that.

---

***Let's make this list bigger! Have your own tip/trick? Something I forgot to add? Share it with us!***

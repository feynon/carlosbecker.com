---
title: "Docker: The very basics"
date: 2015-05-12
draft: false
slug: docker-basics
city: Joinville
---

Or "what the hell is this Docker thing?"

### Intro

According to their [website](https://www.docker.com/),

> Docker is an open platform for developers and sysadmins to build, ship, and run distributed applications. Consisting of Docker Engine, a portable, lightweight runtime and packaging tool, and Docker Hub, a cloud service for sharing applications and automating workflows, Docker enables apps to be quickly assembled from components and eliminates the friction between development, QA, and production environments. As a result, IT can ship faster and run the same app, unchanged, on laptops, data center VMs, and any cloud.

The main idea is that, instead of shipping wars, ears, tars, debs or whatever package system you might think of, you ship a standardized container, which can run "anywhere" without much effort.

The terminology is based on real containers. A container, if you look at [Wikipedia](http://en.wikipedia.org/wiki/Intermodal_container), is:

> An intermodal container (also known as a container, freight container, ISO container, shipping container, hi-cube container, box, sea container, container van) is a standardized reusable steel box. Intermodal containers are used to store and move materials and products efficiently and securely in the global containerized intermodal freight transport system. "Intermodal" indicates that the container can be used across various modes of transport, (from ship to rail to truck) without unloading and reloading its contents.

So, the idea is that, you can "ship" a container to production, which can be the same container you used before to test the app on your own machine, maybe changing only the database URL or something.

It is common to "link" Docker to Microservices and vice-versa, but, even if it's nice to, you don't need to use Docker to ship Microservices. Docker is just the tool commonly used for that.

If you think outside the box of putting web apps in production using Docker, you will see that you can ship basically any software using it, from [scripts](https://github.com/bobrik/docker-image-cleaner) to
[selenium nodes](https://github.com/SeleniumHQ/docker-selenium), for example. 

It is very important to note that the containers are stateless, meaning, they don't store any data within them. With that said, it is no surprise that CI companies are using Docker to "isolate" and distribute the builds.

### How it works

I will assume that you are familiar with the "default" virtualization softwares, like, let's say, VirtualBox. This is how they work:

![](/public/images/docker-basics/f468d0d9-51d2-40ce-bb5f-0defd407c955.png)

Basically, each new VM you need to run will run isolated and load up the entire "guest" OS, using the "host" hardware through the "Hypervisor".

This all may sound nice in some cases, but this strategy uses a lot of resources and may be a little slow too.

Docker, in the other hand, doesn't have a Hypervisor. Instead, it does an operating-system-level virtualization, which is a server virtualization method where the kernel of an operating system allows for multiple isolated user space instances, instead of just one.

![](/public/images/docker-basics/13400398-9f3b-4d77-a473-f0f7f83829ea.png)

LXC, which was used as default by Docker until version [0.9](https://blog.docker.com/2014/03/docker-0-9-introducing-execution-drivers-and-libcontainer/)), is nothing new. The first commit was made in [Aug 6, 2008](https://github.com/lxc/lxc/commit/5e97c3fcce787a5bc0f8ceef43aa3e05195b480a).

The thing is that LXC is a little hard to work with, and Docker "abstracted" a lot of stuff with their `libcontainer`, which should allow us to, for example, [run Docker on Windows Server](http://www.pcworld.com/article/2834132/microsoft-to-bring-docker-to-windows-server.html) in a near future.

Another key feature, in my humble opinion, is AuFS. 

AuFS (AnotherUnionFS) is a multi-layered filesystem that implements union mount, which basically allows several filesystems or directories to be simultaneously mounted and visible through a single mount point, appearing to be one filesystem to the end user (in this case, us).

Supposing you have 10 Docker containers based on, let's say, a 1Gb Ubuntu Server image, they will all use only 1Gb plus their specific data, instead of 10+Gb like it would use if running on VirtualBox, for  example.

So, if you pull, let's say, the `bobrik/image-cleaner` image, you will see something like:

```shell
$ docker pull bobrik/image-cleaner
Pulling repository bobrik/image-cleaner
28b7cd17052f: Download complete
511136ea3c5a: Download complete
a5b60fe97da5: Download complete
9bff7ebd6f58: Download complete
5381e678f99a: Download complete
Status: Downloaded newer image for bobrik/image-cleaner:latest
```

Each of those checksums are nothing but AuFS "layers". So, if I decide to do another image adding some stuff to this one, I will not have to push the entire image, just the "diff"… more or less like a git  commit.

Also, `docker pull`?

### The Docker Registry

Yeah, you can push and pull container images to/from the Docker registry, which is something like a "GitHub" for Docker container images.

If you look at some public image in the registry, like [this one](https://registry.hub.docker.com/u/bobrik/image-cleaner/), for example, you will see there is a `Source project page` link, which leads to the [GitHub repository](https://github.com/bobrik/docker-image-cleaner) of that specific image, where you can finally see the `Dockerfile` for that image.

A Dockerfile is a text file that contains all the commands you would normally execute manually in order to build a Docker image. It extends another image, which can be a "clean" ubuntu server or really any other image.

The sintax is pretty simple (of course this is a very basic example), for example:

```docker
FROM busybox
ENTRYPOINT echo "Hello world"
```

You can then build this image, tagging it as `hello-world`:

```shell
$ docker build -t hello-world .
Sending build context to Docker daemon 2.048 kB
Sending build context to Docker daemon
Step 0 : FROM busybox
latest: Pulling from busybox
cf2616975b4a: Pull complete
6ce2e90b0bc7: Pull complete
8c2e06607696: Already exists
Digest: sha256:38a203e1986cf79639cfb9b2e1d6e773de84002feea2d4eb006b52004ee8502d
Status: Downloaded newer image for busybox:latest
 <!--more-->> 8c2e06607696
Step 1 : ENTRYPOINT echo Hello World
 ---> Running in b29590127d4f
 ---> 7fa687f18c73
Removing intermediate container b29590127d4f
Successfully built 7fa687f18c73
```

Now you can finally run it:

```shell
$ docker run -t hello-world
Hello World
```

If you create a Docker Registry account and if this wasn't a totally useless image, you could a so push it with something like:

```shell
$ docker push caarlos0/hello-world
```

You can learn more about Dockerfile and the Docker CLI [here](https://docs.docker.com/engine/reference/builder/) and [here](https://docs.docker.com/engine/reference/commandline/).

### That's all folks!

I think this enough information for now. You can basically Google and learn more about all those fancy names I cited as well those in the images if you feel like needing more info.

Also, feel free to ask questions and/or share your thoughts in the comments box bellow. 🍻
---
date: 2015-05-15T00:00:00Z
slug: selenium-grid-docker-compose
title: Running a Selenium Grid with docker-compose
---

At [ContaAzul](http://contaazul.com), we had 31 Windows machines powering our
Selenium tests - one running the grid and 30 more running clients. Needless
to say, this is **very** expensive.

As we are already using [Docker](http://docker.io/) to run our builds (on
[Shippable](http://shippable.com)), we decided to try it out to run Selenium
tests too.

It was no surprise that Selenium guys already made a ready-to-go set of Docker
images. There is an image for the Selenium Grid itself, and the browser images -
Chrome and Firefox, also with debug versions - which allow you to access them
using VNC to _"see what's happening there"_. You can check them out in their
[GitHub repository](https://github.com/SeleniumHQ/docker-selenium).

## Making it work

Basically, I created a `c3.4xlarge` EC2 machine and installed both
Docker and docker-compose, following their respective README's:

```console
# install docker:
$ wget -qO- https://get.docker.com/ | sh

# install docker-compose:
$ curl -L https://github.com/docker/compose/releases/download/1.2.0/docker-compose-`uname -s`-`uname -m` > $ /usr/local/bin/docker-compose
$ chmod +x /usr/local/bin/docker-compose
```

With docker up and running, I created a simple `docker-compose.yml` file, which
describes my containers and how they interact with each other. It ended up
being something like this:

```yaml
hub:
  image: selenium/hub
  ports:
    - "4444:4444"
firefox:
  image: selenium/node-firefox
  links:
    - hub
chrome:
  image: selenium/node-chrome
  links:
    - hub
```

This uses only the very basics of docker-compose syntax. You can always take a
look at the [docs](https://docs.docker.com/compose/) to learn a bit
more about it.

Now, we can start the service running:

```console
$ docker-compose up
```

After that, it is just a matter of telling the selenium test runner to connect
on the Docker host machine on port `4444` and everything will just work.

But, well, we had 30 machines before... now I only have 2 selenium instances
running (one with Firefox and the other with Chrome). How can I fix that?
Well, I'm glad you asked:

```console
$ docker-compose scale chrome=15 firefox=15
```

Around 10 seconds later, 30 selenium instances up and running. **Sweet**.

## Let's talk money

![show me the money gif](/public/images/show-me-the-money.gif)

The objective was to decrease our costs with EC2 instances in our development
VPC.

With this changes, we dropped our monthly EC2 billing related to Selenium
by ~**77%**! Ok, ok, we have also changed the main OS where Selenium runs.
Well, even if the instances were already Linux boxes before, it
would still be a cut of ~**57%**:

![ec2 values](/public/images/docker-selenium-ec2-values.png)

It is also important to notice that we pay the Amazon bill in USD, and we pay
around BRL 4.5 per USD. That said, USD 1161 costs us around BRL 5224.5, which can
buy here ~411L of beer (using BRL ~12.69/L).

Well, **MISSION ACCOMPLISHED**.

> 50% usage because we only use them about 12 hours per day (business hours,
give or take).

## Try it out

In order to make it easier for you guys to put all this to work (save you
a bunch of copy-and-paste combos), I created a simple set of shell scripts that
can easily put a selenium grid up and running.

To test that, you can start a fresh linux machine (or not, your call) and hit:

```console
$ git clone https://github.com/caarlos0/selenium-grid-on-docker-example.git grid && \
  cd grid && sudo ./install.sh
```

This will download the scripts and install docker and docker compose. When you
install Docker, it will suggest you to add your user to the `docker` group.
You should really do that. I help you, it's something like this:

```console
$ sudo usermod -aG docker your-username
```

Now, let's put it to run:

```console
$ ./run.sh
```

This command will pull and run 3 containers: hub, firefox and chrome. You can
scale things up with:

```console
$ ./scale.sh 10
```

This will scale the grid to 10 Chrome containers and 10 Firefox containers
(be advised that it will eat a lot memory - it's 20 browsers, after all).

On my Mac, I scaled it to 4 Chrome and 4 Firefox instances, and it works:

![Running 4 Firefox and 4 Chrome instances on my laptop](/public/images/docker-on-my-machine-scale-4.png)

Just 5 seconds to start 8 containers. _Neat_.

## Conclusions

**Docker is great!** (I can't say this enough)

Some people don't yet trust Docker enough to put it in production, or are
scared of it because of the lack of knowledge. I can only suggest you to
start testing it in development machines, CI environments and so forth. It is
safe and you will surely learn a lot (and love almost every second of it).
You can also read my
[previous post about how Docker works]({{< ref "post/2015-05-12-docker-basics.md" >}}).

The best part: need more speed? Just change the instance type and let
docker-compose scale it up!
The other best part (yes there is more than one): you can put 30 Selenium
machines to run in a few seconds. Try that with 30 Windows machines.

## Future

Maybe autoscale the cluster using spot instances?

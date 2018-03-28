---
date: 2015-09-29T00:00:00Z
slug: cd-github-circleci-docker-beanstalk
title: Continuous Delivery with GitHub, CircleCI, Docker and AWS Elastic Beanstalk
tags:
- ci
- aws
---

This is just a quick overview of how I did it in [antibody's homepage][1].

![antibody home page screenshot](/public/images/antibody-site.png)

The site has a very simple `index.html` plus a
service that can discover and download the latest
[antibody](https://github.com/getantibody/antibody) version from GitHub
releases. Not sure how relevant this is, but the service is written in Go.

Currently, every commit pushed to the `master` branch of the site repo will be
automatically pushed to production. There is no option to skip that.

Let's see how this can be done using
[GitHub](http://github.com),
[CircleCI](https://circleci.com/),
[Docker](http://docker.com) and
[Amazon ElasticBeanstalk](https://aws.amazon.com/elasticbeanstalk/).

### Docker

The very first part of all this process is the `Dockerfile`. This file describes
the software needed to run our service, also how the service is deployed inside
it:

```Dockerfile
FROM alpine:3.2

ENV GOPATH=/gopath \
  SRC=/gopath/src/github.com/getantibody/getantibody

WORKDIR $SRC
ADD . $SRC
EXPOSE 3000

RUN apk add -U git go && \
  go get -v -d ./... && \
  go get -v github.com/GeertJohan/go.rice/rice && \
  /gopath/bin/rice embed-go -i ./server && \
  go install -v ./... && \
  apk del git go && \
  rm -rf /gopath/src /gopath/pkg /var/cache/apk/*

CMD /gopath/bin/server
```

As you can see, I get all dependencies, do what I need with them and
clean it up, in one single step, just to ensure that the image will be as
small as it can, currently, around 20Mb.

### Elastic Beanstalk

For the `awsebcli` to work, after the app is already created, we need a
`Dockerrun.aws.json` file. Mine looks like this:

```json
{
  "AWSEBDockerrunVersion": "1",
  "Image": {
    "Name": "caarlos0/getantibody:%BUILD_NUM%",
    "Update": "true"
  },
  "Ports": [
    {
      "ContainerPort": "3000"
    }
  ]
}
```

Nothing extraordinary here, I just expose the port `3000` and set the image
name, which has a `%BUILD_NUM%` expression in it... but I'll explain that later.

I put this file inside a `.deploy` folder. The only thing in that folder
is this single file, because `awsebcli` zips and uploads everything
inside the folder you are when you run `eb deploy`, and we need just that
one file.

### CircleCI - and the integrations

The `circle.yml` file must do the following:

- start the Docker service;
- install Python - which is needed to install `awsebcli`;
- install `awsebcli`;
- build the Docker image and push it to Docker Hub;
- deploy the image to Beanstalk.

That said, mine looks like this:

```yaml
machine:
  services:
    - docker
  python:
    version: 2.7
dependencies:
  pre:
    - pip install awsebcli
deployment:
  production:
    branch: master
    commands:
      - docker login -e "$DOCKER_EMAIL" -u "$DOCKER_USER" -p "$DOCKER_PASS"
      - docker build -t "caarlos0/getantibody:$CIRCLE_BUILD_NUM" .
      - docker push "caarlos0/getantibody:$CIRCLE_BUILD_NUM"
      - sed -i'' -e "s;%BUILD_NUM%;$CIRCLE_BUILD_NUM;g" ./.deploy/Dockerrun.aws.json
      - cd .deploy && eb init -r us-east-1 getantibody
      - cd .deploy && eb deploy -l $CIRCLE_BUILD_NUM
```

The most interesting part here surely is the `deployment` section.
The `DOCKER_*` environment variables are configured in the
CircleCI panel, and are used to login to Docker Hub. The `CIRCLE_BUILD_NUM`
environment variable is provided by CircleCI itself and is used
to tag the image and build (more on that later).

You can also see there the `sed` expression that replaces the `%BUILD_NUM%`
expression by the `CIRCLE_BUILD_NUM` variable inside `Dockerrun.aws.json`,
so that way we link the Beanstalk deploy version with the Docker image
tag.

### So, what's new on build `x`?

I tag the image and the deployment with the CircleCI build number,
so I can track down which commits are in production, just by
looking at the tag names.

For example, at the time I wrote this, tag `63` was deployed to Elastic
Beanstalk:

![antibody on elastic beanstalk screenshot](/public/images/antibody-eb.png)

Which I know is deploying the Docker image tag `63`:

![antibody on dockerhub screenshot](/public/images/antibody-dockerhub.png)

Which I know was built in build number `63` in CircleCI:

![antibody on circle screenshot](/public/images/antibody-circle.png)

In which were added the following changes (click the "compare" link on
CircleCI):

![comparing changes on github screenshot](/public/images/antibody-github.png)

Isn't that great? You can track down a version deployed in production
directly to the source code that was introduced with it. Since the Docker
image config is in the repo too, we can also kind of track infrastructure
changes!

### Conclusion

I know, it is a very simple workflow and might not work for everyone, but it is
simple enough so you can basically copy-paste on your own pet project and
try it!

BTW: I'm using AWS free tier for the app, so, yup, it's free... at least for
now.

Last but not least, you can see the entire source code
[here](https://github.com/getantibody/getantibody).

Cheers! :beers:

[1]: http://getantibody.github.io/

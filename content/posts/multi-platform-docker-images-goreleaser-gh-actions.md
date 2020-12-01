---
title: "Multi-platform Docker images with GoReleaser and GitHub Actions"
date: 2020-11-30
draft: false
slug: multi-platform-docker-images-goreleaser-gh-actions
city: Cascavel
toc: true
tags: [docker, golang, goreleaser, ci-cd]
---

[GoReleaser v0.148.0 is out](https://github.com/goreleaser/goreleaser/releases/tag/v0.148.0), and with it, the ability to release multi-platform Docker images, a.k.a. *[Docker Manifests](https://docs.docker.com/engine/reference/commandline/manifest/)*.

In this guide we'll explore how to use it with [GitHub Actions](https://github.com/features/actions), and how GoReleaser releases itself in this way.

## An example project

I created an example project showing with all the code needed for everything to work. You can check it out [here](https://github.com/caarlos0/goreleaser-docker-manifest-actions-example).

## A simple `main.go`

For our example, we'll have a very simple `main.go` file:

```go
// main.go
package main

import (
	"fmt"
	"runtime"
)

var version = "dev"

func main() {
	fmt.Println("example", version, runtime.GOOS)
}
```

### Our mighty `Dockerfile`

GoReleaser builds Docker images by copying the previously built binaries to the images (instead of building the binary inside Docker itself). This guarantees that the binary inside the image and the one you download from the releases page is the same.

To account for multiple platforms, we either create several `dockerfiles`, or, when possible, make the platform a parameter in a single one. We'll use the second approach.

Our very basic `Dockerfile` looks like this:

```dockerfile
# Dockerfile
ARG ARCH
FROM ${ARCH}/alpine
COPY example /usr/bin/example
ENTRYPOINT ["/usr/bin/example"]
```

We can test it without GoReleaser by running:

```bash
GOOS=linux GOARCH=amd64 go build -o example .
docker build -t testimage:amd64 . --build-arg ARCH=amd64

GOOS=linux GOARCH=arm64 go build -o example .
docker build -t testimage:arm64v8 . --build-arg ARCH=arm64v8
```

With that in place, let's check our GoReleaser config file.

### Our very own `.goreleaser.yml`

Our GoReleaser config file is very simple:

- One item in the `builds` section which will build for multiple platforms;
- Two items in the `dockers` section, building one image for `amd64` and another for `arm64`;
- One item in the new `docker_manifests` section, tying together the two images in a single manifest.

It looks like this:

```yaml
# .goreleaser.yml
project_name: example
builds:
  - env: [CGO_ENABLED=0]
    goos:
      - linux
      - windows
      - darwin
    goarch:
      - amd64
      - arm64
dockers:
- image_templates: ["ghcr.io/caarlos0/goreleaser-docker-manifest-actions-example:{{ .Version }}-amd64"]
  binaries: [example]
  dockerfile: Dockerfile
  build_flag_templates:
  - --build-arg=ARCH=amd64
  - --label=org.opencontainers.image.title={{ .ProjectName }}
  - --label=org.opencontainers.image.description={{ .ProjectName }}
  - --label=org.opencontainers.image.url=https://github.com/caarlos0/goreleaser-docker-manifest-actions-example
  - --label=org.opencontainers.image.source=https://github.com/caarlos0/goreleaser-docker-manifest-actions-example
  - --label=org.opencontainers.image.version={{ .Version }}
  - --label=org.opencontainers.image.created={{ time "2006-01-02T15:04:05Z07:00" }}
  - --label=org.opencontainers.image.revision={{ .FullCommit }}
  - --label=org.opencontainers.image.licenses=MIT
- image_templates: ["ghcr.io/caarlos0/goreleaser-docker-manifest-actions-example:{{ .Version }}-arm64v8"]
  binaries: [example]
  goarch: arm64
  dockerfile: Dockerfile
  build_flag_templates:
  - --build-arg=ARCH=arm64v8
  - --label=org.opencontainers.image.title={{ .ProjectName }}
  - --label=org.opencontainers.image.description={{ .ProjectName }}
  - --label=org.opencontainers.image.url=https://github.com/caarlos0/goreleaser-docker-manifest-actions-example
  - --label=org.opencontainers.image.source=https://github.com/caarlos0/goreleaser-docker-manifest-actions-example
  - --label=org.opencontainers.image.version={{ .Version }}
  - --label=org.opencontainers.image.created={{ time "2006-01-02T15:04:05Z07:00" }}
  - --label=org.opencontainers.image.revision={{ .FullCommit }}
  - --label=org.opencontainers.image.licenses=MIT
docker_manifests:
- name_template: ghcr.io/caarlos0/goreleaser-docker-manifest-actions-example:{{ .Version }}
  image_templates:
  - ghcr.io/caarlos0/goreleaser-docker-manifest-actions-example:{{ .Version }}-amd64
  - ghcr.io/caarlos0/goreleaser-docker-manifest-actions-example:{{ .Version }}-arm64v8
```

> You can check more options for [builds](https://goreleaser.com/customization/build), [docker](https://goreleaser.com/customization/docker) and [docker manifests](https://goreleaser.com/customization/docker_manifest) on [GoReleaser's website](https://goreleaser.com).
> 
> The labels added to the images are optional, but in the specific case of `ghcr.io`, they allows GitHub to know which image is built from which repository and other metadata.

We can now verify this locally with:

```bash
goreleaser release --snapshot --rm-dist
```

GoReleaser will use defaults for a lot of things, you can check the full config (with the defaults) in at `dist/config.yaml`.

### GitHub Actions

Here we pretty much copy what's already in [GitHub Actions](https://goreleaser.com/ci/actions/) section in the [GoReleaser's website](https://goreleaser.com/):

```yaml
# .github/workflows/goreleaser.yml
name: goreleaser

on:
  pull_request:
  push:

jobs:
  goreleaser:
    runs-on: ubuntu-latest
    env:
      DOCKER_CLI_EXPERIMENTAL: "enabled"
    steps:
      -
        name: Checkout
        uses: actions/checkout@v2
        with:
          fetch-depth: 0
      -
        name: Allow arm Docker builds # https://github.com/linuxkit/linuxkit/tree/master/pkg/binfmt
        run: sudo docker run --privileged linuxkit/binfmt:v0.8
      -
        name: Docker Login
        uses: docker/login-action@v1
        with:
          registry: ghcr.io
          username: ${{ github.repository_owner }}
          password: ${{ secrets.GH_PAT }}
      -
        name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.15
      -
        name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v2
        if: startsWith(github.ref, 'refs/tags/')
        with:
          version: latest
          args: release --rm-dist
        env:
          GITHUB_TOKEN: ${{ secrets.GH_PAT }}
      -
        name: Clear
        if: always()
        run: |
          rm -f ${HOME}/.docker/config.json
```

### Important things to notice

- We need to set `DOCKER_CLI_EXPERIMENTAL=enabled` for the `docker manifest` command to work;
- We need to use `linuxkit/binfmt` to allow the GitHub Actions worker to create Docker images other than `linux/amd64`. More info [here](https://github.com/linuxkit/linuxkit/tree/master/pkg/binfmt);
- We need to login into the GitHub Container Registry with a Personal Access Token (PAT), since the default `GITHUB_TOKEN` does not have enough permissions.

And that's pretty much it!

## Releasing

Now, we just need to push a tag, sit back, relax and watch the [GoReleaser Action](https://github.com/marketplace/actions/goreleaser-action) do everything.

In the end, you should have a release more or less like this:

{{< figure caption="GitHub Release - the changelog, assets and our multi-platform Docker image" src="/public/images/multi-platform-docker-images-goreleaser-gh-actions/54b3b826-1a23-4b6f-a9a9-aaba029fabd1.png" >}}

You should also be able to see that the image is in fact multi-platform in the container registry:

{{< figure caption="GitHub Container Registry showing both OS/Arch combinations we provided." src="/public/images/multi-platform-docker-images-goreleaser-gh-actions/359d79b4-7e66-451d-9272-e1212497c342.png" >}}

We can now run our image:

```bash
$ docker run --rm ghcr.io/caarlos0/goreleaser-docker-manifest-actions-example:1.0.0
example 1.0.0 linux
```

It works! 🎉

## That's it!

That's it! I hope this is useful somehow.

Don't forget to check out [GoReleaser's documentation](https://goreleaser.com) for more details. Also make sure to take a look at [Docker's manifest documentation](https://docs.docker.com/engine/reference/commandline/manifest/).

## Special Thanks

- [@CrazyMax](https://crazymax.dev) for the review;

---
title: "Multi-platform Docker images with GoReleaser and GitHub Actions"
date: 2020-11-30
draft: false
slug: multi-platform-docker-images-goreleaser-gh-actions
city: Cascavel
toc: true
tags: [docker, golang, goreleaser, ci-cd, github]
---

[GoReleaser v0.148.0 is out](https://github.com/goreleaser/goreleaser/releases/tag/v0.148.0), and with it, the ability to release multi-platform Docker images, a.k.a. *[Docker Manifests](https://docs.docker.com/engine/reference/commandline/manifest/)*.

In this guide we'll explore how to use it with [GitHub Actions](https://github.com/features/actions), and how GoReleaser releases itself in this way.

PS: You will need GoReleaser version 0.152.0 or later for this guide to work properly.

## An example project

I created an example project showing with all the code needed for everything to work. You can check it out [here](https://github.com/caarlos0/goreleaser-docker-manifest-actions-example).

### A simple `main.go`

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
	fmt.Println("example", version, runtime.GOOS, runtime.GOARCH)
}
```

### Our mighty `Dockerfile`

GoReleaser builds Docker images by copying the previously built binaries to the images (instead of building the binary inside Docker itself). This guarantees that the binary inside the image and the one you download from the releases page is the same.

Our very basic `Dockerfile` looks like this:

```dockerfile
# Dockerfile
FROM alpine
COPY goreleaser-docker-manifest-actions-example \
	/usr/bin/goreleaser-docker-manifest-actions-example
ENTRYPOINT ["/usr/bin/goreleaser-docker-manifest-actions-example"]
```

To account for multiple platforms, we either create several `dockerfiles`, or use the `--platform` build flag. We'll use the second approach in our example.

We can test it without GoReleaser by running:

```bash
GOOS=linux GOARCH=amd64 go build -o example .
docker buildx build -t testimage:amd64 . --platform=linux/amd64

GOOS=linux GOARCH=arm64 go build -o example .
docker buildx build -t testimage:arm64v8 . --platform=linux/arm64/v8
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
- image_templates: ["ghcr.io/caarlos0/{{ .ProjectName }}:{{ .Version }}-amd64"]
  dockerfile: Dockerfile
  use: buildx
  build_flag_templates:
  - --platform=linux/amd64
  - --label=org.opencontainers.image.title={{ .ProjectName }}
  - --label=org.opencontainers.image.description={{ .ProjectName }}
  - --label=org.opencontainers.image.url=https://github.com/caarlos0/{{ .ProjectName }}
  - --label=org.opencontainers.image.source=https://github.com/caarlos0/{{ .ProjectName }}
  - --label=org.opencontainers.image.version={{ .Version }}
  - --label=org.opencontainers.image.created={{ time "2006-01-02T15:04:05Z07:00" }}
  - --label=org.opencontainers.image.revision={{ .FullCommit }}
  - --label=org.opencontainers.image.licenses=MIT
- image_templates: ["ghcr.io/caarlos0/{{ .ProjectName }}:{{ .Version }}-arm64v8"]
  goarch: arm64
  dockerfile: Dockerfile
  use: buildx
  build_flag_templates:
  - --platform=linux/arm64/v8
  - --label=org.opencontainers.image.title={{ .ProjectName }}
  - --label=org.opencontainers.image.description={{ .ProjectName }}
  - --label=org.opencontainers.image.url=https://github.com/caarlos0/{{ .ProjectName }}
  - --label=org.opencontainers.image.source=https://github.com/caarlos0/{{ .ProjectName }}
  - --label=org.opencontainers.image.version={{ .Version }}
  - --label=org.opencontainers.image.created={{ time "2006-01-02T15:04:05Z07:00" }}
  - --label=org.opencontainers.image.revision={{ .FullCommit }}
  - --label=org.opencontainers.image.licenses=MIT
docker_manifests:
- name_template: ghcr.io/caarlos0/{{ .ProjectName }}:{{ .Version }}
  image_templates:
  - ghcr.io/caarlos0/{{ .ProjectName }}:{{ .Version }}-amd64
  - ghcr.io/caarlos0/{{ .ProjectName }}:{{ .Version }}-arm64v8
- name_template: ghcr.io/caarlos0/{{ .ProjectName }}:latest
  image_templates:
  - ghcr.io/caarlos0/{{ .ProjectName }}:{{ .Version }}-amd64
  - ghcr.io/caarlos0/{{ .ProjectName }}:{{ .Version }}-arm64v8
```

You can check more options for [builds](https://goreleaser.com/customization/build), [docker](https://goreleaser.com/customization/docker) and [docker manifests](https://goreleaser.com/customization/docker_manifest) on [GoReleaser's website](https://goreleaser.com).

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
  push:
    tags:
      - '*'

permissions:
  contents: write

jobs:
  goreleaser:
    runs-on: ubuntu-latest
    env:
      DOCKER_CLI_EXPERIMENTAL: "enabled"
    steps:
      - name: Checkout
        uses: actions/checkout@v2
        with:
          fetch-depth: 0
      - name: Set up QEMU
        uses: docker/setup-qemu-action@v1
      - name: Docker Login
        uses: docker/login-action@v1
        with:
          registry: ghcr.io
          username: ${{ github.repository_owner }}
          password: ${{ secrets.GH_PAT }}
      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.16
      - name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v2
        with:
          version: latest
          args: release --rm-dist
        env:
          GITHUB_TOKEN: ${{ secrets.GH_PAT }}
```

### Important things to notice

- We need to set `DOCKER_CLI_EXPERIMENTAL=enabled` for the `docker manifest` command to work;
- We need to setup `qemu` and `buildx` in order to build Docker images in platforms other than `linux/amd64` using `docker buildx build`;
- We need to login into the GitHub Container Registry with a Personal Access Token (PAT), since the default `GITHUB_TOKEN` does not have enough permissions.

And that's pretty much it!

## Releasing

Now, we just need to push a tag, sit back, relax and watch the [GoReleaser Action](https://github.com/marketplace/actions/goreleaser-action) do everything.

In the end, you should have a release more or less like this:

{{< img caption="GitHub Release - the changelog, assets and our multi-platform Docker image" src="54b3b826-1a23-4b6f-a9a9-aaba029fabd1.png" >}}

You should also be able to see that the image is in fact multi-platform in the container registry:

{{< img caption="GitHub Container Registry showing both OS/Arch combinations we provided." src="359d79b4-7e66-451d-9272-e1212497c342.png" >}}

We can now run our image:

```bash
$ docker run --rm --platform linux/amd64 \
	ghcr.io/caarlos0/goreleaser-docker-manifest-actions-example
example 1.0.2 linux amd64
```

We can also test the arm64 image:

```bash
$ docker run --rm --platform linux/arm64/v8 \
	ghcr.io/caarlos0/goreleaser-docker-manifest-actions-example
example 1.0.2 linux arm64
```

It works! 🎉

## That's it!

That's it! I hope this is useful somehow.

Don't forget to check out [GoReleaser's documentation](https://goreleaser.com) for more details. Also make sure to take a look at [Docker's manifest documentation](https://docs.docker.com/engine/reference/commandline/manifest/).

## Special Thanks

- [@CrazyMax](https://crazymax.dev) for the review;

## Updates

- Aug 1, 2021: overall update on code to use newer syntax/techniques;
- Jan 04, 2021: using `buildx` and setting `use_buildx` on the Docker config and using `docker/setup-qemu-action`;
- Dec 29, 2020: using `--platform` instead of `ARCH` build arg, added more `docker run` examples;

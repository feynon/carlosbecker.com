---
title: "Multi-platform Docker images with GoReleaser, Podman and GitHub Actions"
date: 2021-08-01
draft: false
slug: goreleaser-actions-podman
city: Cascavel
toc: true
tags: [goreleaser, goreleaser-pro, docker, podman, github, ci-cd]
---

A few months ago, I published a post on [Multi-platform Docker images with GoReleaser and GitHub Actions]({{< ref "multi-platform-docker-images-goreleaser-gh-actions.md" >}}). Today's post has the same idea, but using [Podman](https://podman.io) instead of [Docker](https://www.docker.com).

The main advantage of Podman is that you can run in *rootless* mode (e.g. inside a container) and that it doesn't require a daemon.

## An example project

I created an example project showing with all the code needed for everything to work. You can check it out [here](https://github.com/caarlos0/goreleaser-podman-actions-example).

### A simple `main.go`

For our example, we'll have a very simple `main.go` file:

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

### A `Dockerfile`

GoReleaser builds Docker images by copying the previously built binaries to the images (instead of building the binary inside Docker itself). This guarantees that the binary inside the image and the one you download from the releases page is the same.

Our very basic `Dockerfile` looks like this:

```dockerfile
# Dockerfile
FROM alpine
COPY goreleaser-podman-actions-example /usr/bin/goreleaser-podman-actions-example
ENTRYPOINT ["/usr/bin/goreleaser-podman-actions-example"]
```

To account for multiple platforms, we either create several `dockerfiles`, or use the `--platform` build flag. We'll use the second approach in our example.

We can test it without GoReleaser by running:

```sh
GOOS=linux GOARCH=amd64 go build .
podman build -t testimage:amd64 . --platform=linux/amd64

GOOS=linux GOARCH=arm64 go build .
podman build -t testimage:arm64v8 . --platform=linux/arm64/v8
```

With that in place, let's check our GoReleaser config file.

### The `.goreleaser.yml` file

Our GoReleaser config file is very simple:

- One item in the `builds` section which will build for multiple platforms;
- Two items in the `dockers` section, building one image for `amd64` and another for `arm64`, setting the `use` property to `podman`;
- One item in the new `docker_manifests` section, tying together the two images in a single manifest and also setting the `use` property to `podman`.

It looks like this:

```yaml
# .goreleaser.yml
builds:
- env: [CGO_ENABLED=0]
  goos:
  - linux
  - darwin
  - windows
	goarch:
  - amd64
  - arm64
dockers:
- image_templates: ["ghcr.io/caarlos0/{{ .ProjectName }}:{{ .Version }}-amd64"]
  dockerfile: Dockerfile
  use: podman
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
  use: podman
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
  use: podman
  image_templates:
  - ghcr.io/caarlos0/{{ .ProjectName }}:{{ .Version }}-amd64
  - ghcr.io/caarlos0/{{ .ProjectName }}:{{ .Version }}-arm64v8
- name_template: ghcr.io/caarlos0/{{ .ProjectName }}:latest
  use: podman
  image_templates:
  - ghcr.io/caarlos0/{{ .ProjectName }}:{{ .Version }}-amd64
  - ghcr.io/caarlos0/{{ .ProjectName }}:{{ .Version }}-arm64v8
```

*You can check more options for **[builds](https://goreleaser.com/customization/build)**, **[docker](https://goreleaser.com/customization/docker)** and **[docker manifests](https://goreleaser.com/customization/docker_manifest)** on **[GoReleaser's website](https://goreleaser.com/)**.*

> *The labels added to the images are optional, but in the specific case of **`ghcr.io`**, they allows GitHub to know which image is built from which repository and other metadata.*

We can now verify this locally with:

```sh
goreleaser release --snapshot --rm-dist
```

GoReleaser will use defaults for a lot of things, you can check the full config (with the defaults) in at `dist/config.yaml`.

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
    steps:
      - name: Checkout
        uses: actions/checkout@v2
        with:
          fetch-depth: 0
      - name: Set up QEMU
        uses: docker/setup-qemu-action@v1
      - name: Log in to ghcr.io
        uses: redhat-actions/podman-login@v1
        with:
          username: ${{ github.repository_owner }}
          password: ${{ secrets.GH_PAT }}
          registry: ghcr.io
      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.16
      - name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v2
        with:
          distribution: goreleaser-pro
          version: latest
          args: release --rm-dist
        env:
          GITHUB_TOKEN: ${{ secrets.GH_PAT }}
          GORELEASER_KEY: ${{ secrets.GORELEASER_KEY }}
```

### Important things to notice

- We need to setup `qemu` in order to build Docker images in platforms other than `linux/amd64` using `podman build`;
- We need to login into the GitHub Container Registry with a Personal Access Token (PAT), since the default `GITHUB_TOKEN` does not have enough permissions;
- Podman is [GoReleaser Pro](https://goreleaser.com/pro) feature, so we need to use its distribution and pass a valid `GORELEASER_KEY`.

And that's pretty much it!

## **Releasing**

Now, we just need to push a tag, sit back, relax and watch the [GoReleaser Action](https://github.com/marketplace/actions/goreleaser-action) do everything.

In the end, you should have a release more or less like this:

{{< figure caption="GitHub Release - the changelog, assets and our multi-platform Docker image" src="/public/images/goreleaser-actions-podman/96ea6580-1d9f-4610-9046-45c726feb930.png" >}}

You should also be able to see that the image is in fact multi-platform in the container registry:

{{< figure caption="GitHub Container Registry showing both OS/Arch combinations we provided." src="/public/images/goreleaser-actions-podman/f6369315-c810-4db5-9072-0be9aa7d207d.png" >}}

We can now run our image with either Docker or Podman:

```sh
$ docker run --rm --platform linux/amd64 \
		ghcr.io/caarlos0/goreleaser-podman-actions-example
example 0.0.1 linux amd64

$ podman run --rm --platform linux/amd64 \
		ghcr.io/caarlos0/goreleaser-podman-actions-example
example 0.0.1 linux amd64
```

We can also test the arm64 image:

```sh
$ docker run --rm --platform linux/arm64 \
		ghcr.io/caarlos0/goreleaser-podman-actions-example
example 0.0.1 linux arm64

$ podman run --rm --platform linux/arm64 \
		ghcr.io/caarlos0/goreleaser-podman-actions-example
example 0.0.1 linux arm64
```

It works! 🎉

## **That's it!**

That's it! I hope this is useful somehow.

Don't forget to check out [GoReleaser's documentation](https://goreleaser.com/) for more details. Also make sure to take a look at [Docker's manifest documentation](https://docs.docker.com/engine/reference/commandline/manifest/).

Also check out the [GoReleaser Pro](https://goreleaser.com/pro) offering!

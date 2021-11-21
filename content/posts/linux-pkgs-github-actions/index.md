---
title: "Testing Linux packages on GitHub Actions"
date: 2021-11-21
draft: false
slug: linux-pkgs-github-actions
city: Cascavel
toc: true
tags: [goreleaser, nfpm, linux, github, ci-cd]
---

One issue we had from time to time on GoReleaser was related to its Linux packages.

We had a single map from `GOARCH` to Linux arch, when in fact, each package manager might have their own.

That led to less popular packages (e.g. arm) to report the wrong architecture and thus be not installable. 

This was recently moved to [nFPM](https://nfpm.goreleaser.com), and we now just pass a concatenation of `GOARCH+GOARM+GOMIPS`, and each packager handles it on nFPM.

With that change, we added tests for more arches on nFPM, and also added tests to GoReleaser packages themselves.

## The idea

The idea consists in 2 steps:

1. Run the release and cache its outputs
2. Spawn another job for each format

### 1. Cache the ouput

This was easily done using GitHub Actions Cache:

```yaml
- uses: actions/cache@v2
  with:
    path: |
      ./dist/*.deb
      ./dist/*.rpm
      ./dist/*.apk
    key: ${{ runner.os }}-go-${{ hashFiles('**/*.go') }}-${{ hashFiles('**/go.sum') }}
```

This will cache all `deb`, `rpm` and `apk` files on `dist` after the release.

### 2. Spawn a job for each format

Here we used the matrix feature of GitHub actions, as well as Docker, to install the package on different platforms:

```yaml
jobs:
  goreleaser-check-pkgs:
    runs-on: ubuntu-latest
    env:
      DOCKER_CLI_EXPERIMENTAL: "enabled"
    needs:
      - goreleaser
    if: github.ref == 'refs/heads/main'
    strategy:
      matrix:
        format: [ deb, rpm, apk ]
    steps:
      - uses: actions/checkout@v2
        with:
          fetch-depth: 0
      - uses: arduino/setup-task@v1
        with:
          version: 3.x
          repo-token: ${{ secrets.GITHUB_TOKEN }}
      - uses: docker/setup-qemu-action@v1
      - uses: actions/cache@v2
        with:
          path: |
            ./dist/*.deb
            ./dist/*.rpm
            ./dist/*.apk
          key: ${{ runner.os }}-go-${{ hashFiles('**/*.go') }}-${{ hashFiles('**/go.sum') }}
      - run: task goreleaser:test:${{ matrix.format }}
```

Notice that it needs the `goreleaser` job, and only runs on `main`.

After setting up things we need (docker, task, qemu), we run a task named `goreleaser:test:${{ matrix.format }}`.

Let's take a look on how that looks like.

## The test task

Putting simply, each task spawns a Docker container for several platforms, install the package and checks if `goreleaser --version` succeeds.

If something fails, it will exit 1, so we'll know something is wrong.

```yaml
goreleaser:test:rpm:
  desc: Tests rpm packages
  cmds:
    - docker run --platform linux/386 --rm -v $PWD/dist:/tmp/ centos:centos7 sh -c 'rpm --nodeps -ivh /tmp/goreleaser-*.i386.rpm && goreleaser --version'
    - docker run --platform linux/amd64 --rm -v "$PWD/dist":/tmp/ fedora sh -c 'rpm --nodeps -ivh /tmp/goreleaser-*.x86_64.rpm && goreleaser --version'
    - docker run --platform linux/arm64 --rm -v "$PWD/dist":/tmp/ fedora sh -c 'rpm --nodeps -ivh /tmp/goreleaser-*.aarch64.rpm && goreleaser --version'
    - docker run --platform linux/arm/6 --rm -v "$PWD/dist":/tmp/ fedora sh -c 'rpm --nodeps -ivh /tmp/goreleaser-*.armv6hl.rpm && goreleaser --version'
    - docker run --platform linux/arm/7 --rm -v "$PWD/dist":/tmp/ fedora sh -c 'rpm --nodeps -ivh /tmp/goreleaser-*.armv7hl.rpm && goreleaser --version'

goreleaser:test:deb:
  desc: Tests rpm packages
  cmds:
    - docker run --platform linux/amd64 --rm -v "$PWD/dist":/tmp/ ubuntu bash -c 'dpkg --ignore-depends=git -i /tmp/goreleaser*_amd64.deb && goreleaser --version'
    - docker run --platform linux/arm64 --rm -v "$PWD/dist":/tmp/ ubuntu bash -c 'dpkg --ignore-depends=git -i /tmp/goreleaser*_arm64.deb && goreleaser --version'
    - docker run --platform linux/arm/6 --rm -v "$PWD/dist":/tmp/ debian bash -c 'dpkg --ignore-depends=git -i /tmp/goreleaser*_armel.deb && goreleaser --version'
    - docker run --platform linux/arm/7 --rm -v "$PWD/dist":/tmp/ ubuntu bash -c 'dpkg --ignore-depends=git -i /tmp/goreleaser*_armhf.deb && goreleaser --version'

goreleaser:test:apk:
  desc: Tests rpm packages
  cmds:
    - docker run --platform linux/386 --rm -v "$PWD/dist":/tmp/ alpine ash -c 'apk add --allow-untrusted -U /tmp/goreleaser*_x86.apk && goreleaser --version'
    - docker run --platform linux/amd64 --rm -v "$PWD/dist":/tmp/ alpine ash -c 'apk add --allow-untrusted -U /tmp/goreleaser*_x86_64.apk && goreleaser --version'
    - docker run --platform linux/arm64 --rm -v "$PWD/dist":/tmp/ alpine ash -c 'apk add --allow-untrusted -U /tmp/goreleaser*_aarch64.apk && goreleaser --version'
    - docker run --platform linux/arm/6 --rm -v "$PWD/dist":/tmp/ alpine ash -c 'apk add --allow-untrusted -U /tmp/goreleaser*_armhf.apk && goreleaser --version'
    - docker run --platform linux/arm/7 --rm -v "$PWD/dist":/tmp/ alpine ash -c 'apk add --allow-untrusted -U /tmp/goreleaser*_armv7.apk && goreleaser --version'
```

Note that different formats tests against different platforms.

{{< img caption="How it looks like on GitHub Actions" src="f45b161c-a91e-441e-bdf6-a45434c251f0.png" >}}

That's mostly because not all Linux distributions have all platforms available on Docker Hub - e.g. no Fedora arm/6 nor Ubuntu 386.

Still, testing against those should cover most of the use cases, and I think other projects that worry about this might benefit for something like this themselves.

## References

- [https://github.com/goreleaser/nfpm/pull/398](https://github.com/goreleaser/nfpm/pull/398)
- [https://github.com/goreleaser/goreleaser/pull/2640](https://github.com/goreleaser/goreleaser/pull/2640)
- [https://github.com/goreleaser/goreleaser/blob/35eb844f937d3c18f308b3eff85e1c5abf092f4d/Taskfile.yml#L128-L153](https://github.com/goreleaser/goreleaser/blob/35eb844f937d3c18f308b3eff85e1c5abf092f4d/Taskfile.yml#L128-L153)
- [https://github.com/goreleaser/goreleaser/blob/35eb844f937d3c18f308b3eff85e1c5abf092f4d/.github/workflows/build.yml#L12-L38](https://github.com/goreleaser/goreleaser/blob/35eb844f937d3c18f308b3eff85e1c5abf092f4d/.github/workflows/build.yml#L12-L38)
- [https://github.com/goreleaser/goreleaser/actions/runs/1487207845](https://github.com/goreleaser/goreleaser/actions/runs/1487207845)

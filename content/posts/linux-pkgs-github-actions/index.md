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
goreleaser:test:pkg:
  desc: Test a package
  cmds:
    - docker run --platform linux/{{.Platform}} --rm --workdir /tmp -v $PWD/dist:/tmp {{.Image}} sh -c '{{.Cmd}} && goreleaser --version'

goreleaser:test:rpm:
  desc: Tests rpm packages
  vars:
    rpm: 'rpm --nodeps -ivh'
  cmds:
    - task: goreleaser:test:pkg
      vars:
        Platform: '386'
        Image: centos:centos7
        Cmd: '{{.rpm}} goreleaser-*.i386.rpm'
    - task: goreleaser:test:pkg
      vars:
        Platform: 'amd64'
        Image: fedora
        Cmd: '{{.rpm}} goreleaser-*.x86_64.rpm'
    - task: goreleaser:test:pkg
      vars:
        Platform: 'arm64'
        Image: fedora
        Cmd: '{{.rpm}} goreleaser-*.aarch64.rpm'
    - task: goreleaser:test:pkg
      vars:
        Platform: 'arm/6'
        Image: fedora
        Cmd: '{{.rpm}} goreleaser-*.armv6hl.rpm'
    - task: goreleaser:test:pkg
      vars:
        Platform: 'arm/7'
        Image: fedora
        Cmd: '{{.rpm}} goreleaser-*.armv7hl.rpm'

goreleaser:test:deb:
  desc: Tests rpm packages
  vars:
    dpkg: 'dpkg --ignore-depends=git -i'
  cmds:
    - task: goreleaser:test:pkg
      vars:
        Platform: 'amd64'
        Image: ubuntu
        Cmd: '{{.dpkg}} goreleaser*_amd64.deb'
    - task: goreleaser:test:pkg
      vars:
        Platform: 'arm64'
        Image: ubuntu
        Cmd: '{{.dpkg}} goreleaser*_arm64.deb'
    - task: goreleaser:test:pkg
      vars:
        Platform: 'arm/6'
        Image: debian
        Cmd: '{{.dpkg}} goreleaser*_armel.deb'
    - task: goreleaser:test:pkg
      vars:
        Platform: 'arm/7'
        Image: ubuntu
        Cmd: '{{.dpkg}} goreleaser*_armhf.deb'

goreleaser:test:apk:
  desc: Tests rpm packages
  vars:
    apk: 'apk add --allow-untrusted -U'
  cmds:
    - task: goreleaser:test:pkg
      vars:
        Platform: '386'
        Image: alpine
        Cmd: '{{.apk}} goreleaser*_x86.apk'
    - task: goreleaser:test:pkg
      vars:
        Platform: 'amd64'
        Image: alpine
        Cmd: '{{.apk}} goreleaser*_x86_64.apk'
    - task: goreleaser:test:pkg
      vars:
        Platform: 'arm64'
        Image: alpine
        Cmd: '{{.apk}} goreleaser*_aarch64.apk'
    - task: goreleaser:test:pkg
      vars:
        Platform: 'arm/6'
        Image: alpine
        Cmd: '{{.apk}} goreleaser*_armhf.apk'
    - task: goreleaser:test:pkg
      vars:
        Platform: 'arm/7'
        Image: alpine
        Cmd: '{{.apk}} goreleaser*_armv7.apk'
```

Note that different formats tests against different platforms.

{{< img caption="How it looks like on GitHub Actions" src="f45b161c-a91e-441e-bdf6-a45434c251f0.png" >}}

That's mostly because not all Linux distributions have all platforms available on Docker Hub - e.g. no Ubuntu 386.

Still, testing against those should cover most of the use cases, and I think other projects that worry about this might benefit for something like this themselves.

## References

- [https://github.com/goreleaser/nfpm/pull/398](https://github.com/goreleaser/nfpm/pull/398)
- [https://github.com/goreleaser/goreleaser/pull/2640](https://github.com/goreleaser/goreleaser/pull/2640)
- [GoReleaser's Taskfile](https://github.com/goreleaser/goreleaser/blob/12066dd8d9179ef19dde061d67354f5ebc081685/Taskfile.yml#L128-L226)
- [GoReleaser Build Workflow](https://github.com/goreleaser/goreleaser/blob/35eb844f937d3c18f308b3eff85e1c5abf092f4d/.github/workflows/build.yml#L12-L38)
- [Example Actions Run](https://github.com/goreleaser/goreleaser/actions/runs/1487207845)

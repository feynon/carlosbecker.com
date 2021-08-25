---
title: "Signing artifacts with cosign and GoReleaser"
date: 2021-08-24
draft: false
slug: goreleaser-cosign
city: Cascavel
toc: true
tags: [goreleaser, goreleaser-pro, security, golang]
---

In GoReleaser v0.176.0 (both OSS and Pro), we released the ability to sign Docker images - with [cosign](https://github.com/sigstore/cosign) in mind, and also did small quality-of-life improvements in the artifact signing feature.

In this post we'll explore how to quickly add this to your GoReleaser config so your users can verify the artifacts they download.

## cosign

You'll need to install [cosign](https://github.com/sigstore/cosign), and then generate a key pair with it:

```sh
cosign generate-key-pair
```

It'll ask for a password and its confirmation - and that's it.

PS: the private key itself is encrypted, so it's safe to push it to your git repository if you have a descent password. I recommend using some password generator and create a really big password.

## GoReleaser config

Here things are fairly simple as well, you just need to pass the correct settings to the `signs` section, and declare a `docker_signs` section as well.

I also recommend setting the `checksums.name_template` setting to simply `checksums.txt`.

Here's how it looks like:

```yaml
checksum:
  name_template: 'checksums.txt'

signs:
- cmd: cosign
  stdin: '{{ .Env.COSIGN_PWD }}'
  args: ["sign-blob", "-key=cosign.key", "-output=${signature}", "${artifact}"]
  artifacts: checksum

docker_signs:
- artifacts: manifests
  stdin: '{{ .Env.COSIGN_PWD }}'
```

In this example I'm signing only the `checksum` file and the produced Docker manifests, but you can change that as you see fit. Check the documentation on [signs](https://goreleaser.com/customization/sign/) and [docker signs](https://goreleaser.com/customization/docker_sign/) for more details.

Notice that in this example you would still need to install `cosign`, have the `cosign.key` in the repository root folder and export the key password as `COSIGN_PWD`.

## Verifying

Let's say you sign only the checksum, as the example above. In that case, your user will need:

- the public key used to sign (`cosign.pub`)
- the `checksums.txt` file that was uploaded to the release
- the `checksums.txt.sig` file that was uploaded to the release
- whatever artifacts they need (e.g. the linux amd64 debian package)

Then, they can verify using `cosign`:

```sh
cosign verify-blob \
  -key cosign.pub \
  -signature checksums.txt.sig \
  checksums.txt
```

If the validation succeeds, it means the `checksums.txt` was not tampered with, so we can use it to verify the hashes of other files using `sha256sum`:

```sh
sha256sum --ignore-missing -c checksums.txt
```

Doing that, we can verify our binaries are correct.

## Verifying Docker images

We can also verify the Docker image signature. Here `cosign` actually embeds the signature in the image manifest, so we only need the public key used to sign it in order to verify its authenticity:

```sh
cosign verify -key cosign.pub user/image
```

And that's it!

## Conclusion

This is just a quick "how-to" on these new features, I hope you find them useful!

FWIW, GoReleaser itself is also signed with `cosign` now! You can check the documentation on how to verify the artifacts [here](https://goreleaser.com/install/#verifying-the-binaries) and the Docker images [here](https://goreleaser.com/install/#verifying-docker-images).

## Credits

I already had the idea of integrating GoReleaser with `cosign` in my backlog, but seeing [Engin's article on how to do it in previous GoReleaser versions](https://blog.ediri.io/build-trust-with-signing-your-cli-binary-and-container) gave me the extra motivation I needed to actually sit down and do it.

So, thanks Engin for the article and overall support of OpenSource and GoReleaser, its really appreciated! 🖤

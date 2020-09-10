---
title: "Easy private Helm repositories"
date: 2020-02-21
draft: false
slug: private-helm-repo
city: Joinville
---

Can we have a private Helm repository using GCS as backend? Yes we can!

<!--more-->

The easier way to create a public helm repository is to just upload your artifacts to a bucket somewhere. 

I usually use a script like the following to do that:

```shell
#!/bin/sh
set -e

helm init --client-only

mkdir -p upstream
gsutil -m rsync gs://my-charts upstream

find charts -maxdepth 1 -mindepth 1 -type d | while read -r CHART; do
	helm dep update "${CHART}"
	helm package "${CHART}" --destination upstream
done

helm repo index upstream/ --url "https://my-charts.storage.googleapis.com/"

gsutil -m rsync upstream gs://my-charts
```

This will download the previous releases to a `upstream` folder (needed to recreate the `index.yaml` file), package all charts inside the `charts` folder and then `rsync` the `upstream` folder back to the bucket.

If the bucket is public, you can just:

```shell
helm repo add test https://my-charts.storage.googleapis.com/
helm repo update test
helm search test
```

And use it normally.

## Privacy

But what if you want a private repository, for whatever reason? Or just an authenticated one?

We created [storage-auth-proxy](https://github.com/totvslabs/storage-auth-proxy) to do just that: auth and proxy requests to a private bucket.

You can define several `user:password` combos, point to a bucket, and that's it:

```shell
./storage-auth-proxy \
	-listen 0.0.0.0:8080 \
	-bucket gs://my-private-charts \
	-authorize foo:bar \
	-authorize carlos:secret
```

You can then expose the service (e.g. `helm.mycompany.com`) and change our script a bit:

```shell
# change the URL
helm repo index upstream/ --url "https://helm.mycompany.com/"
```

And then, finally, just use the repository and use it:

```shell
helm repo add mycompany https://helm.mycompany.com \
	--username carlos \
	--password secret
helm repo update mycompany
helm search mycompany
```

**And that's it!**

It supports out of the box:

- GCS
- S3
- Azure

You can authenticate as normally would for each provider (default environment variables, `gcloud auth`, etc).

## Conclusion

This is a pretty simple and cheap solution. If you only want a public repo, its likely you can run on the free-tier of AWS or GCP.

If you want it private and don't have too much traffic, you can run a single very small VM/container and that's it.

Hope that's somehow useful for you! 🙂

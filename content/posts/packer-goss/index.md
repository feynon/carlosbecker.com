---
title: "Using Goss to validate Packer builds"
date: 2018-12-06
draft: false
slug: packer-goss
city: Joinville
toc: true
tags: [packer, ci-cd]
---

Ever wanted to validate your [Packer](https://packer.io/) image with [Goss](http://goss.rocks/)? Well, you can!

---

I was looking into ways to make sure that the image I just provisioned with
[Packer](https://packer.io/) and [Chef](https://www.chef.io/) is working as expected. After some research, I found
[Goss](http://goss.rocks/), which is a tool to validate servers.

So, I just needed to glue all that together.

Lucky me, there is a [Packer plugin](https://github.com/YaleUniversity/packer-provisioner-goss) for that, made by the folks
at Yale.

I found some issues within the plugin, so I fixed them in my [own fork](https://github.com/caarlos0/packer-provisioner-goss),
though I have high hopes that they will accept my suggestions soon.

You can download the binary from the releases page and save it as
`~/.packer.d/plugins/packer-provisioner-goss` with `0755` permissions.

The usage is very simple, jsut add a `goss` provisioner to your [Packer](https://packer.io/)
build:

```json
{
  "builders": [
    {
      "type": "googlecompute",
      "project_id": "foo",
      "zone": "us-west2-b",
      "machine_type": "n1-standard-2",
      "preemptible": true,
      "ssh_username": "ubuntu",
      "source_image_family": "ubuntu-1804-lts"
    }
  ],
  "provisioners": [
    {
      "type": "chef-solo",
      "version": "14.7.17",
      "cookbook_paths": [
        "cookbooks"
      ],
      "chef_environment": "myenv",
      "data_bags_path": "data_bags/myenv",
      "environments_path": "environments",
      "roles_path": "roles/myenv",
      "run_list": [
        "role[foo]"
      ]
    },
    {
      "type": "goss",
      "retry_timeout": "5m",
      "sleep": "5s",
      "tests": [
        "goss/goss.yml"
      ]
    }
  ]
}
```

The `goss/goss.yml` is a regular [Goss](http://goss.rocks/) test suite file.

And thats it. Run `packer build` and it will run everything and check it
with [Goss](http://goss.rocks/) before pushing the image to Google Cloud (in this example).

Very simple and quick tip, but I found it really useful and I hope you enjoy
it!

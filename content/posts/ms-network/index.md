---
title: "Riding a MS Network with Linux"
date: 2013-06-03
draft: false
slug: ms-network
city: Joinville
toc: true
tags: [linux]
---

I work in a company that have an enterily MS-based network, which means, all that domain login crap and everything that comes with it: outlook, MS Communicator, proxies, monitoring and etc.

I particurarly hate to use MS Windows anytime, even more for work. So, I installed an Ubuntu Linux with the windows that came in my machine. Doing this came with some weird issues and workarounds, and this post will describe some of them.

## Network

This was the trickiest: I needed a "GOD-mode certificate", you should ask the domain admins for something like it. After that, the setup was pretty easy:

{{< img caption="" src="cdaf2ee3-24a4-4c9b-ad5b-c28aa977c11d.png" >}}

## Proxy

You can either set the proxy direct on Network Proxy Setup and give the password ten thousand times a day or use CNTLM. I strongly suggest you to use CNTLM.

You can download it from [here](https://sourceforge.net/projects/cntlm/files/cntlm/cntlm%200.92.3/). Setup is a little tricky, but I will guide you. You shall have a `/etc/cntlm.conf` file, setup the URL, domain, user and exceptions.

After that, you may have to setup the password. I recommend you to encrypt it with the following command:

```
$ sudo cntlm -H -c /etc/cntlm.conf
```

Then put you root password followed by the domain password. Copy the given hashes to the `/etc/cntlm.conf` and restart the service with `/etc/cntlm restart`, setup your proxy as `127.0.0.1:3128` (8010 for SOCKS) and you are ready to go (remember to apply system wide):

{{< img caption="" src="af7cdd85-525a-4326-8683-aa6752f2fec4.png" >}}

This will fix browsing, apt and most of the things (everything that reads `http_proxy` environment variable). You may also want to `chmod 0600` the config file to avoid others access.

## Email

For email setup, you will probably need to call the network admin/support to ask the addresses and ports, then, setup it in Thunderbird (my preferred email client). After that, the setup will be easy.

The only trick I could give you here is about new message composing: Outlook create replies putting the *new content* above the previous emails, while Thunderbird put it below.

To change it, go to your account settings, and change the "Automatically quote the original message when replaying Then," to *"start my reply above the quote"* in the "Compositing & Addressing" tab:

{{< img caption="" src="d7dcf823-9741-4f28-bc70-1675d9f68c6d.png" >}}

## Lync / Communicator

For this -crapy- thing we have a web version available (as well as email), but you can also setup this (if you have luck) in pidgin, using pidgin-sipe (you can get all the needed packages from apt). You can follow this (link is now offline) tutorial, skipping the installation (use apt).

I don't have luck with that, so I still with web version.

The rest work seamlessly without any tweaks. If you have any thought, please share it in comment box above.

Thanks.

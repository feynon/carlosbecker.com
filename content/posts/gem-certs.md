---
title: "Fixing Rubygems Certificates"
date: 2013-11-28
draft: false
slug: gem-certs
city: Joinville
toc: true
tags: []
---

Today, once again, my environment start throwing that SSL cert error.

I followed the basics of [this common link](http://railsapps.github.io/openssl-certificate-verify-failed.html) (and also very good, by the way), but sadly id didn't solve the issue. So, I came across this blog post which solve my problem. I just did:

```
$ brew update
$ brew install openssl
$ brew link openssl --force
$ brew install curl-ca-bundle
# this next line makes the difference for me!
$ export SSL_CERT_FILE=/usr/local/opt/curl-ca-bundle/share/ca-bundle.crt
```

And now it seems to be working again.

I'll just leave this here, in case anyone ran through the same issue.

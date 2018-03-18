---
date: 2014-04-01T00:00:00Z
slug: spotify-everyone
title: Spotify for Everyone
---

I personally love Spotify. I can hear music everywhere and even
sync some music to my devices, and it works pretty well even in
slow connections.

But in lot of countries, including Brazil (where I currently live),
Spotify would not let you create an account, due to some copyright
bullshit.

Lucky for us, this is one case where proxies are actually useful.
To begin, you will need a proxy server in USA, or, you could just
open a SSH connection to some server in US and use it as a SOCKS5
proxy.

I used the later and it is working as expected for a month or two
now. So, I will guide you through this process.

## 1. Get some server

First thing, you need a server in US. I recommend DigitalOceal for
that. It is really cheap, easy to sign up, easy and fast to create
and destroy the server.

[Create an account](https://www.digitalocean.com/?refcode=7e8e9efb2f77)
and create a server. You can put anything you want in the
hostname and leave the rest as it is.

You will receive an email with the password to access your server,
and it will be up and running in about 60 seconds.


## 2. Create a SOCK5 proxy

You will receive from DigitalOcean an email like this:

    You can access it using the following credentials:
    IP Address: X.X.X.X
    Username: root
    Password: pwd

To create the proxy, you will need OpenSSH installed. It is default
in Linux and Mac OS X boxes, I don't know about Windows, but,
given the history, probably not. In that case, you are by yourself
to get this working.

Now, back to what really matters: to create the proxy server, run this:

```console
$ ssh root@X.X.X.X -NCD 8080
```

Here, X.X.X.X is the IP from the DigitalOcean email. When asked, enter
also the password provided in the very same email.

Now you should have a proxy server running!

## 3. Setup your machine

You will have to setup SOCKS proxy to address to `127.0.0.1` and the
port to `8080`.

I recommend you to use Mozilla Firefox for that, so you can change the
settings for the browser only.

[How to setup the proxy in Firefox](http://www.wikihow.com/Enter-Proxy-Settings-in-Firefox)

Don't forget to check the SOCKSv5 option if asked and to uncheck
the "Use this proxy server for all protocols" if it is checked.

Now, point to Google and search for the term "my ip". You should
see the IP you received by email.

## 4. Create the account, finally

Now, just point the browser to the Spotify website and register. You
will notice that you will not be able to change your country. This
is expected. Just proceed and create your account.

Now:

- Go to the console where you ran that command I said before,
and press CTRL+C. This will stop the proxy server;
- Go to Firefox settings and check the "No Proxy" option.

Now, access your Spotify Payment Method Setup and configure everything.
Yes, you will have to pay or they will block you out after some time.

## 5. Enjoy!

Download the app for your OS or access the
[web player](https://play.spotify.com) and enjoy your music!

## 6. Clean up

If you will not need the server anymore, you can shut it down.
To do that, access your droplet in DigitalOcean, click in
the destroy tab and follow the instructions.

## Thats all folks

Hope you enjoy this not-very-technical post. Cheers!

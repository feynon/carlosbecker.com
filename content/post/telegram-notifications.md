---
title: "Notify your team using Telegram"
date: 2014-11-26
draft: false
slug: telegram-notifications
city: Joinville
---

In an ideal world, applications will never go down, for example. In the real world, shit happens. Every second counts.

Let's take the example of a server going down for some reason (which shouldn't happen, ever). I want the team to be notified as soon as possible to mitigate the issue. To do that in a easy and free way, I decided to use [Telegram](https://telegram.org/).

### Wait, wait... what is this Telegram thing again?

According to [their website](https://telegram.org/):

> Telegram is a cloud-based mobile and desktop messaging app with a focus on security and speed.

It's basicaly a *Whatsapp-like* messenger, with an open API and more security.

### Wiring it up

First, install the [telegram-client](https://telegram.org/dl/cli) following their README.

Then, write your script. Mine looks like this:

```bash
#!/bin/bash
TG=/opt/tg

# $1: Server Name (used in the message body)
# $2: URL to test
# $3: Group/User to notify
function check() {
	wget -q $2 -O /dev/null
	if [ ! $? -eq 0 ]; then
		echo "msg $3 WARN: $1 IS DOWN!!!" | $TG/bin/telegram-cli \
			-k $TG/tg-server.pub -W
	fi
}

check "my server" "http://myserver.blah.fake.address/check" "Server1"
check "another server" "http://myserver2.blah.fake.address/check" "Server2"
check "Google" "http://google.com" "General"
```

Sure, you can script it in order to notify you about anything, including some business specific things, dependency on third-party systems... well, use your imagination.

I also added it to the `crontab`, so it will run every minute.

And, sure enough, it works:

![](/public/images/telegram-notifications/390610a2-e32f-49c7-9c79-595af9224e9c.png)

Hope it helps!

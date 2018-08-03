---
date: 2013-08-09T00:00:00Z
slug: hubot-pokemon
title: Charmander, our Hubot
---

We just started to use the awesome [Github Hubot][hubot], "a customizable
kegerator-powered life embetterment robot".

Basically, it's a Campfire bot crafted with Node.js which you can setup to
hear your chat and do things.

These things are basically CoffeeScript files using the Robot API. You can
see a lot of examples [here][hubot-scripts].

These scripts vary from super useful ones, like starting a Jenkins build, to
others even more use useful, like get a Carlton Dancing Gif or parse the
emojis in the message (yes, emojis are very important).

<!--

![carlton](http://gifsoup.com/webroot/animatedgifs/131815_o.gif)

-->

We named our personal bot "Charmander". He is kinda untrained yet, but it still
do a lot of things. We can easily command him to deploy a branch in our staging
server:

> charmander deploy wealcash/awesome-feature to staging

And he will lovely do all needed work to made it happen.

We also use it to get NewRelic and Gaug.es statuses, weather, movie suggestions
and tons of other things. Better than that, we can tell him to burn things:

> charmander burn that crap

![charmander burning things](https://24.media.tumblr.com/69c947ad72c230f72d55a93bed958417/tumblr_mhthj8aPR21s595b5o1_500.gif)

Dude, this is awesome!

## Installation

Hubot installation is pretty simple. Basically, you will need node.js and npm
installed in your server. After that, you can follow [this guide][hubot-install]
to put it to work.

After that, write scripts to automatize everything. Srsly, stop doing repetitive
and boring jobs.

## Campfire

Maybe you don't want to pay for Campfire. There is no problem. You can use tons of
alternatives.

We are using one called [talker][talker]. It's pretty good and free. If you worry
about your data, you could install kandan in your server, it even works
in [heroku][heroku] if you want to.

There is tons of other hubot compatible chats out there. You might want to take a
look at them to see which best fit your needs. For the sake of information, we are
using talker and, except that it doesn't show us emojis by default,
it is awesome and works well for us.

[hubot]: http://hubot.github.com/
[hubot-scripts]: https://github.com/github/hubot-scripts
[hubot-install]: https://github.com/github/hubot/tree/master/docs
[heroku]: https://heroku.com/
[talker]: http://talkerapp.com/

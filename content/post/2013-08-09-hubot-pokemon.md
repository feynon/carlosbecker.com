---
title: "Charmander, our Hubot"
date: 2013-08-09
draft: false
slug: hubot-pokemon
city: Joinville
---

We just started to use the awesome [Github Hubot](http://hubot.github.com/), "a customizable kegerator-powered life embetterment robot".

Basically, it's a Campfire bot crafted with Node.js which you can setup to hear your chat and do things.

These things are basically CoffeeScript files using the Robot API. You can see a lot of examples [here](https://github.com/github/hubot-scripts).

These scripts vary from super useful ones, like starting a Jenkins build, to others even more use useful, like get a Carlton Dancing Gif or parse the emojis in the message (yes, emojis are very important).

We named our personal bot "Charmander". He is kinda untrained yet, but it still do a lot of things. We can easily command him to deploy a branch in our staging server:

> charmander deploy wealcash/awesome-feature to staging

And he will lovely do all needed work to made it happen.

We also use it to get NewRelic and Gaug.es statuses, weather, movie suggestions and tons of other things. Better than that, we can tell him to burn things:

> charmander burn that crap
![](/public/images/hubot-pokemon/ebfd3e1f-755c-4d2f-85d9-df1a2047261c.gif)

Dude, this is awesome!

## Installation

Hubot installation is pretty simple. Basically, you will need node.js and npm installed in your server. After that, you can follow [this guide](https://github.com/github/hubot/tree/master/docs) to put it to work.

After that, write scripts to automatize everything. Srsly, stop doing repetitive and boring jobs.

## Campfire

Maybe you don't want to pay for Campfire. There is no problem. You can use tons of alternatives.

We are using one called [talker](http://talkerapp.com/). It's pretty good and free. If you worry about your data, you could install kandan in your server, it even works in [heroku](https://heroku.com/) if you want to.

There is tons of other hubot compatible chats out there. You might want to take a look at them to see which best fit your needs. For the sake of information, we are using talker and, except that it doesn't show us emojis by default, it is awesome and works well for us.
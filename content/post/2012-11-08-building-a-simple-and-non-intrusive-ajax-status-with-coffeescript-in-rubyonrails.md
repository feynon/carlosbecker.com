---
date: 2012-11-08T00:00:00Z
slug: building-a-simple-and-non-intrusive-ajax-status-with-coffeescript-in-rubyonrails
title: Simple RoR AJAX Status
---

## tl;dr

> A simple tutorial explaining how to made a simple coffeescript that can
> automagically show and hide a loading spin while doing a ajax call.

## Hi

I just built a simple and non intrusive AJAX status indicator with [
spin.js](http://fgnass.github.com/spin.js/) and
[CoffeeScript](http://coffeescript.org), and I think that you people may like
it :)

By simple, I meant that the code itself is simple. You don't need to know JS/Coffee
like a ninja.

By non intrusive, I meant that the indicator does not block user interaction with the app,
and you will see that you don't need to change your code to fit in my indicator ~rules~, so,
it's also code-non-intrusive. (LOL)

Let's do it!

## What we need

Basically, the only dependency is the spin.js library. You probably use something
like `jquery-ujs`, `turbolinks` or even the `jQuery` API itself to made ajax calls.

In my case, I use all of them.

## Let's code!

My `Gemfile` looks like this:

```ruby
(...)
group :assets do
  gem 'sass-rails',   '~> 3.2.3'
  gem 'coffee-rails', '~> 3.2.1'
  gem 'therubyracer', :platforms => :ruby
  gem 'uglifier', '>= 1.0.3'
  gem 'turbolinks'
  gem 'jquery-turbolinks'
  gem 'spinjs-rails'
end

# asset related
gem 'jquery-rails'
(...)
```

You will also need to import them in you `application.js`.

## events

We all know that something about events triggered by those APIS:

- `jQuery.{ajax,post,get}`: `ajax{Start,Stop,Complete,Error,Success}`
- `turbolinks`: `page:{fetch,load,restore,change}`
- `jQuery-ujs`: `ajax:{beforeSend,success,error,complete,aborted:required,aborted:file}`

There's a lot of events to deal, uh? Not really.
According to my tests, basically, all `jQuery`events will fire `ajaxStart` and
`ajaxComplete` at the start and end of request repectivelly. In the same way, `tubolinks`
will trigger `page:fetch` and `page:change`, an at last `jQuery-ujs` will fire
`ajax:beforeSend` and `ajax:complete`. So, we will have something like:

```coffeescript
$(document).on 'ajax:before ajaxStart page:fetch', ->
  # show spin

$(document).on 'ajax:complete ajaxComplete page:change', ->
  # hide spin
```

But we also have another problem: if some user interaction fire multiple ajax calls?
In my app, I don't need to worry too much about that for now, because I don't need to
do more than one `$.get` (for example) in the same user interaction. I only have
some `$.get` (for example) callbacks that call `Turbolinks.visit`, so, imagine something like:

        ajax:beforeSend--------------ajax:complete------------------>
        -------------------page:fetch------------------page:change-->

The spin will shows up in `beforeSend`, and hide in `complete`, but there's another
ajax going on in `page:fetch`. To solve that, I just cache in a variable the
last "event type" (ajax or page), and only hide when this call ends.

To fix the issue about multiple ajax calls for the same type, we can made a dirty hack
to store in a hash or something the count of events that should end to spin disapear.
Is pretty simple, but I'll let this for you guys :)

## Complete code

The most "difficult" part of the solution is already explained, so, get the code,
read it, and use the comment box if you have any doubts =)

---

`ajax.spin.js.coffee` code:

```coffeescript
opts = {
  lines: 7, # The number of lines to draw
  length: 6, # The length of each line
  width: 3, # The line thickness
  radius: 5, # The radius of the inner circle
  corners: 1, # Corner roundness (0..1)
  rotate: 0, # The rotation offset
  color: '#000', # #rgb or #rrggbb
  speed: 1.1, # Rounds per second
  trail: 100, # Afterglow percentage
  shadow: false, # Whether to render a shadow
  hwaccel: false, # Whether to use hardware acceleration
  className: 'spinner', # The CSS class to assign to the spinner
  zIndex: 2e9, # The z-index (defaults to 2000000000)
  top: 'auto', # Top position relative to parent in px
  left: 'auto' # Left position relative to parent in px
}

# save the lastEvent type that was called
lastEvent = undefined
# the element where the spinner should appear
$n = undefined

# I can only pop the $n var when document is ready
$(document).ready ->
  $n = $('.navbar.navbar-static-top')

# get the event type, ex: a "page:change" will return only 'page'
eventType = (event) ->
  return false if not event
  type = event.type
  if type.indexOf(':') > -1
    type.split(':')[0]
  else
    type.match(/[A-Z]?[a-z]+|[0-9]+/g)[0]

# show the spinner
loadState = (event) ->
  lastEvent = eventType event
  $n.spin opts

# hide the spinner
doneState = (event) ->
  if eventType(event) == lastEvent
    lastEvent = undefined
    $n.spin false

# bind some states (will see if it is more needed)
$(document).on 'ajax:before ajaxStart page:fetch', (event) ->
  loadState event
$(document).on 'ajax:complete ajaxComplete page:change', (event) ->
  doneState event
```

---

The compiled `ajax.spin.js` file, for those wo don't use CoffeeScript:

```js
(function() {
  var $n, doneState, eventType, lastEvent, loadState, opts;

  opts = {
    lines: 7,
    length: 6,
    width: 3,
    radius: 5,
    corners: 1,
    rotate: 0,
    color: "#000",
    speed: 1.1,
    trail: 100,
    shadow: false,
    hwaccel: false,
    className: "spinner",
    zIndex: 2e9,
    top: "auto",
    left: "auto"
  };

  lastEvent = void 0;

  $n = void 0;

  $(document).ready(function() {
    return ($n = $(".navbar.navbar-static-top"));
  });

  eventType = function(event) {
    var type;
    if (!event) {
      return false;
    }
    type = event.type;
    if (type.indexOf(":") > -1) {
      return type.split(":")[0];
    } else {
      return type.match(/[A-Z]?[a-z]+|[0-9]+/g)[0];
    }
  };

  loadState = function(event) {
    lastEvent = eventType(event);
    return $n.spin(opts);
  };

  doneState = function(event) {
    if (eventType(event) === lastEvent) {
      lastEvent = void 0;
      return $n.spin(false);
    }
  };

  $(document).on("ajax:before ajaxStart page:fetch", function(event) {
    return loadState(event);
  });

  $(document).on("ajax:complete ajaxComplete page:change", function(event) {
    return doneState(event);
  });
}.call(this));
```

Hope it help somebody, enjoy.

Cheers

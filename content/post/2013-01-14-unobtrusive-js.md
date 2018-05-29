---
date: 2013-01-14T00:00:00Z
slug: unobtrusive-js
title: Unobtrusive JS
---

![Unobtrusive car repair](/public/images/repair.jpg "Unobtrusive car repair")

One of the principles of Unobtrusive JS is the _"separation of functionality
(the "behavior layer") from a Web page's structure/content and presentation"_.

#### A small example:

##### Wrong:

```html
<a href="#" onclick="SomeObj.someAction(2)">Do some action on Obj 2</a>
```

##### better:

```html
<a href="#" class="someaction" data-id="2">Do some action on Obj 2</a>
```

And in our JS file:

```js
$(
  (function() {
    $(document).on("click", ".someaction", function() {
      SomeObj.someAction($(this).data("id"));
    });
  })(jQuery)
);
```

Notice that with this, you don't even need to expose the `someAction` method from
`SomeObj`, you can do these bindings in the object declaration.

#### An example with charts (and morris.js)

A pretty common mistake (I think) occurs while adding some charts to a page (
genarally a dashboard). I commited this mistake too.

##### A bad example with Rails, will be something like:

> dashboard.html.erb

```erb
<script type="text/javascript">
Morris.Donut({
  element: "my-chart-placeholder",
  data: <%= raw ModelController.chart_data.to_json %>
})
</script>
<div id="my-chart-placeholder"></div>
```

As you can see, I'm doing the JS right into the page, so, I can concatenate it
with a Ruby method call, and get the data without use a Restful call.

Pretty tricky, but, you can do it better.

##### The better way

Use data attributes!

> dashboard.html.erb

```erb
<div id="my-chart-placeholder" data-chart="<%= ModelController.chart_data.to_json %>">
</div>
```

> dashboard.js

```js
$(
  (function() {
    var id = "my-chart-placeholder";
    var chart = $("#" + id);
    // verify if the placeholder exists
    if (chart.lenght > 0) {
      Morris.Donut({
        element: id,
        data: chart.data("chart")
      });
    }
  })(jQuery)
);
```

Much better, huh? We now have our presentation and behavior more separated.

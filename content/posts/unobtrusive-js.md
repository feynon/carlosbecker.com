---
title: "Unobtrusive JS"
date: 2013-01-13
draft: false
slug: unobtrusive-js
city: Joinville
toc: true
tags: [javascript]
---

![](/public/images/unobtrusive-js/4aebb97c-0ec2-437e-b6f6-426395e89121.png)

One of the principles of Unobtrusive JS is the *"separation of functionality (the "behavior layer") from a Web page's structure/content and presentation"*.

### A small example:

```html
<a href="#" onclick="SomeObj.someAction(2)">Do some action on Obj 2</a>
```

### better:

```html
<a href="#" class="someaction" data-id="2">Do some action on Obj 2</a>
```

And in our JS file:

```javascript
$(
  (function() {
    $(document).on("click", ".someaction", function() {
      SomeObj.someAction($(this).data("id"));
    });
  })(jQuery)
);
```

Notice that with this, you don't even need to expose the `someAction` method from `SomeObj`, you can do these bindings in the object declaration.

### An example with charts (and morris.js)

A pretty common mistake (I think) occurs while adding some charts to a page ( genarally a dashboard). I commited this mistake too.

### A bad example with Rails, will be something like:

```html
<script type="text/javascript">
Morris.Donut({
  element: "my-chart-placeholder",
  data: <%= raw ModelController.chart_data.to_json %>
})
</script>
<div id="my-chart-placeholder"></div>
```

As you can see, I'm doing the JS right into the page, so, I can concatenate it with a Ruby method call, and get the data without use a Restful call.

Pretty tricky, but, you can do it better.

### The better way

Use data attributes!

```html
<div id="my-chart-placeholder" data-chart="<%= ModelController.chart_data.to_json %>">
</div>
```

> dashboard.js
```javascript
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

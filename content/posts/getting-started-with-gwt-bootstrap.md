---
title: "Getting started with GWT-Bootstrap"
date: 2012-06-26
draft: false
slug: getting-started-with-gwt-bootstrap
city: Joinville
toc: true
tags: [java, gwt]
---

[GWT-Bootstrap](http://gwtbootstrap.github.com/) is a project that aims to provide all the [Twitter Bootstrap](http://getbootstrap.com/) styles and widgets to GWT applications. I had this idea in Jan/2012, and it was not a library, initially, I just made a simple `InputText` and `Button` widgets as a proof of concept, and then I think: *"Why not?"*.

So, here we go. The [initial release is almost done](https://github.com/gwtbootstrap/gwt-bootstrap/issues?milestone=3&page=1&state=open), and has a lot of widgets and functionally ported and working, including the awesome *responsiveness* and the great *NavBar*.

So, let's start using it.

## Get the jar

There are 2 ways to get the jar:

- [Maven](https://github.com/gwtbootstrap/gwt-bootstrap/wiki/Using-GWT-Bootstrap-Maven-Repository)
- [Downloading the jar by hand and put it in you app classpath](https://github.com/gwtbootstrap/gwt-bootstrap/downloads)
## Configuring you module

You will have to configure your *App.gwt.xml* file to *inherit* the GWT-Bootstrap widget library. You can do it like this:

```xml
<inherits name="com.github.gwtbootstrap.Bootstrap"/>
```
## UiBinder usage

Assuming that you're using UIBinder, add the following namespace to the `<ui:UIBinder>` element:

```xml
xmlns:b="urn:import:com.github.gwtbootstrap.client.ui"
```

Then you can easily use the `b` namespace in your widget design, like `Heading`:

```xml
<b:heading size="2">Hello World GWT-Bootstrap</b:heading>
```
## Datepicker

We also have the [bootstrap datepicker](https://github.com/eternicode/bootstrap-datepicker) ported into GWT-Bootstrap, but, in a different module and namespace.

Add to your *App.gwt.xml*:

```xml
<inherits name='com.github.gwtbootstrap.datepicker.Datepicker' />
```

And in your *Widget.ui.xml* file:

```xml
xmlns:b2="urn:import:com.github.gwtbootstrap.datepicker.client.ui"
```

Then, just use the widget like this:

```xml
<b2:DateBox format="dd/mm/yyyy" autoClose="true" />
```
## Pros and Cons

### Pros

- Easy to setup and use
- Themeable
- Almost all widgets ported
- Active forum and development
### Cons

- Doesn't support `less` files ([#29](https://github.com/gwtbootstrap/gwt-bootstrap/issues/29))
- Load all the *JavaScript* files in the bootstrap of the app ([#70](https://github.com/gwtbootstrap/gwt-bootstrap/issues/70))

**Happy coding!**

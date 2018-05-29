---
date: 2012-06-26T00:00:00Z
aliases:
- /code/2012/06/26/getting-started-with-gwt-bootstrap/
slug: getting-started-with-gwt-bootstrap
title: Getting started with GWT-Bootstrap
---

[GWT-Bootstrap](http://gwtbootstrap.github.com) is a project that aims to
provide all the [Twitter Bootstrap](http://getbootstrap.com) styles
and widgets to GWT applications.
I had this idea in Jan/2012, and it was not a library, initially, I just
made a simple `InputText` and `Button` widgets as a proof of concept, and then
I think: _"Why not?"_.

So, here we go. The
[initial release is almost done](https://github.com/gwtbootstrap/gwt-bootstrap/issues?milestone=3&page=1&state=open),
and has a lot of widgets and functionally ported and working, including the
awesome _responsiveness_ and the great _NavBar_.

So, let's start using it.

## Get the jar

There are 2 ways to get the jar:

- [Maven](https://github.com/gwtbootstrap/gwt-bootstrap/wiki/Using-GWT-Bootstrap-Maven-Repository)
- [Downloading the jar by hand and put it in you app classpath](https://github.com/gwtbootstrap/gwt-bootstrap/downloads)

## Configuring you module

You will have to configure your _App.gwt.xml_ file to _inherit_ the
GWT-Bootstrap widget library. You can do it like this:

```xml
<inherits name="com.github.gwtbootstrap.Bootstrap"/>
```

## UiBinder usage

Assuming that you're using UIBinder, add the following namespace to
the `<ui:UIBinder>` element:

```xml
xmlns:b="urn:import:com.github.gwtbootstrap.client.ui"
```

Then you can easily use the `b` namespace in your widget design, like `Heading`:

```xml
<b:heading size="2">Hello World GWT-Bootstrap</b:heading>
```

## Datepicker

We also have the
[bootstrap datepicker](https://github.com/eternicode/bootstrap-datepicker)
ported into GWT-Bootstrap, but, in a different module and namespace.

Add to your _App.gwt.xml_:

```xml
<inherits name='com.github.gwtbootstrap.datepicker.Datepicker' />
```

And in your _Widget.ui.xml_ file:

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
- Load all the _JavaScript_ files in the bootstrap of the app ([#70](https://github.com/gwtbootstrap/gwt-bootstrap/issues/70))

**Happy coding!**

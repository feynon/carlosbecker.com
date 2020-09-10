---
title: "Theming GWT-Bootstrap"
date: 2012-06-27
draft: false
slug: using-a-custom-bootstrap-theme-in-gwt-bootstrap
city: Joinville
---

First of all, if you haven't done it yet, read [Getting started with GWT-Bootstrap]({{< ref getting-started-with-gwt-bootstrap.md >}}).

## Create the GWT-Project

```shell
$ mvn archetype:generate \
   -DarchetypeRepository=repo1.maven.org \
   -DarchetypeGroupId=org.codehaus.mojo \
   -DarchetypeArtifactId=gwt-maven-plugin \
   -DarchetypeVersion=2.4.0
```

I created my project with the following properties:

```
Define value for property 'groupId': : com.github.caarlos0
Define value for property 'artifactId': : Example
Define value for property 'version':  1.0-SNAPSHOT: :
Define value for property 'package':  com.github.caarlos0: :
Define value for property 'module': : Example
```

Now, lets add the `GWT-Bootstrap` dependency to the `pom.xml` file:

### Add the repository:

```xml
<repositories>
  <repository>
    <id>gwt-bootstrap</id>
    <name>GWT-Bootstrap SNAPSHOTS</name>
    <url>http://gwtbootstrap.github.com/maven/snapshots</url>
  </repository>
</repositories>
```

### And the dependency itself:

```xml
<dependency>
  <groupId>com.github.gwtbootstrap</groupId>
  <artifactId>gwt-bootstrap</artifactId>
  <version>2.0.4.0-SNAPSHOT</version>
</dependency>
```

And then, update your project with a `$ mvn clean install`.

## Configure GWT-Bootstrap

In this point, we have to setup the `*.gwt.xml` file and our `UIBinder` XML file. You can follow [getting started tutorial]({{< ref getting-started-with-gwt-bootstrap.md >}}). to do that in the right way.

## Get a custom Bootstrap Theme

You can get a custom `bootstrap.min.css` file in several ways:

- Making your own CSS changing the `.less` files and re-generating the files
- [Customizing your download](http://getbootstrap.com/customize) (Basically the option above in a easy way)
- [Downloading it somewhere](https://www.google.com.br/search?q=twitter+bootstrap+themes)

For this example, I'll use [this theme](http://bootswatch.com/slate/). Download the `bootstrap.min.css` from the site.

## Clean the example

By default, the Maven GWT Archetype will generate a lot of junk, "by example", for you. You can clean it up.

You can remove:

- All classes inside `shared` folder;
- All classes inside `server` folder;
- All the content in the `EntryPoint` class (`Example`, in our case);
- The servlet declarations from `web.xml` file;
- Test related classes and files;
- Messages files in `resources/client` folder.

At this point we will have a structure like this:

```
$ tree
|-- src
|   |-- main
|   |   |-- java
|   |   |   `-- com
|   |   |       `-- github
|   |   |           `-- caarlos0
|   |   |               |-- client
|   |   |               |   |-- Example.java
|   |   |               |-- server
|   |   |               `-- shared
|   |   |-- resources
|   |   |   `-- com
|   |   |       `-- github
|   |   |           `-- caarlos0
|   |   |               `-- Example.gwt.xml
|   |   `-- webapp
|   |       |-- Example.css
|   |       |-- Example.html
|   |       `-- WEB-INF
|   |           `-- web.xml
|   `-- test
`-- target
```

## Create our example

Now, lets create a UIBinder class to made our amazing test widget! Create a new UiBinder class/xml combo called `ExampleUiBinder`, with the following content:

### ExampleUiBinder.ui.xml:

```xml
<ui:UiBinder xmlns:ui='urn:ui:com.google.gwt.uibinder'
             xmlns:g='urn:import:com.google.gwt.user.client.ui'
             xmlns:b="urn:import:com.github.gwtbootstrap.client.ui">
  <g:HTMLPanel>
    <b:Container>
        <b:Navbar>
          <b:Brand>Bootstrap</b:Brand>
          <b:Nav>
            <b:NavLink>Test</b:NavLink>
            <b:NavLink href="http://www.google.com">Another Test</b:NavLink>
          </b:Nav>
          <b:NavForm size="1"/>
          <b:Nav alignment="RIGHT">
            <b:NavLink>Test</b:NavLink>
          </b:Nav>
          <b:NavText alignment="RIGHT">Right</b:NavText>
          <b:NavSearch size="2" placeholder="Search" alignment="RIGHT"/>
      </b:Navbar>

      <b:Hero>
        <b:Heading size="1">Hello Custom Theme!</b:Heading>
        <b:Button type="WARNING">Warning</b:Button>
      </b:Hero>
    </b:Container>
  </g:HTMLPanel>
</ui:UiBinder>
```
```java
public class ExampleUiBinder extends Composite {
  interface ExampleUiBinderUiBinder
      extends UiBinder<HTMLPanel, ExampleUiBinder> {
  }

  private static ExampleUiBinderUiBinder ourUiBinder =
      GWT.create(ExampleUiBinderUiBinder.class);

  public ExampleUiBinder() {
    initWidget(ourUiBinder.createAndBindUi(this));
  }
}
```

At this point, if everything is OK, we will get a window like this:

![](/public/images/using-a-custom-bootstrap-theme-in-gwt-bootstrap/e0639978-f394-4e1f-be9c-4371150bd951.png)

## Hacking

Right now, we will have to write our own `Resources` and `Configuration` classes. I'll advise you that it's a boring thing to do, but the result could be really awesome. So, let's go.

We will have to create a `resources` folder under the same folder of our `*.gwt.xml` file. Just to you understand better, the folders `client`, `shared`, `server` and `resources` **must be** in the same hierarchical level, just like this:

```
$ tree
src/main/java/com/github/caarlos0/
|-- client
|-- resources
|-- server
`-- shared
```

### Creating the needed files

Assuming that we will only change the CSS file, inside your `resources` file, create a `css` folder, and paste the `bootstrap.min.css` file downloaded before inside it. Yes, the file name **must** be `bootstrap.min.css`.

As said before, we also need a `Resources` and `Configuration` files. This files must be inside our `resources` folder too. The content is the following:

### ExampleResources.java

```java
package com.github.caarlos0.resources;

import com.github.gwtbootstrap.client.ui.resources.Resources;
import com.google.gwt.resources.client.TextResource;

public interface ExampleResources extends Resources {
  @Source("css/bootstrap.min.css")
  TextResource bootstrapCss();
}
```

### ExampleConfigurator.java

```java
package com.github.caarlos0.resources;

import com.github.gwtbootstrap.client.ui.config.Configurator;
import com.github.gwtbootstrap.client.ui.resources.Resources;
import com.google.gwt.core.client.GWT;

public class ExampleConfigurator implements Configurator {
  public Resources getResources() {
    return GWT.create(ExampleResources.class);
  }

  public boolean hasResponsiveDesign() {
    return false;
  }
}
```

At this point, the structure should be something like this:

```
$ tree
src/main/java/com/github/caarlos0/
|-- client
|   |-- Example.java
|   |-- ExampleUiBinder.java
|   |-- ExampleUiBinder.ui.xml
|-- resources
|   |-- css
|   |   `-- bootstrap.min.css
|   |-- ExampleConfigurator.java
|   `-- ExampleResources.java
|-- server
`-- shared
```

Now, we have to do a little hack in our `*.gwt.xml`. We will need to replace `com.github.gwtbootstrap.client.ui.config.Configurator` with our Configurator, and setup the resources dir. So, in the end, we will have something like this:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<module rename-to='Example'>
    <inherits name='com.google.gwt.user.User'/>
    <inherits name='com.google.gwt.junit.JUnit'/>
    <inherits name='com.google.gwt.user.theme.standard.Standard'/>
    <inherits name="com.github.gwtbootstrap.Bootstrap"/>
    <entry-point class='com.github.caarlos0.client.Example'/>
    <source path='client'/>
    <source path='shared'/>

    <!--pay attention in this part-->
    <source path='resources'/>
    <replace-with class="com.github.caarlos0.resources.ExampleConfigurator">
      <when-type-is class="com.github.gwtbootstrap.client.ui.config.Configurator"/>
    </replace-with>
    <public path="resources">
      <exclude name="** /*.java"/>
      <exclude name="** /*.class"/>
    </public>

</module>
```

That's it :)

![](/public/images/using-a-custom-bootstrap-theme-in-gwt-bootstrap/b67ca540-e6e9-4f32-828b-7e7a4f45fc67.png)

## Considerations

Sometimes GWT caches everything, and seems like it doesn't work. In this cases, do the following:

- `mvn clean`;
- Delete the `webapp/Example` folder (in our case, `Example` is the Module name);

If it still dont working, do this:

- Comment the `<source path='resources'/>` tag in your `*.gwt.xml`;
- Run the app, you will got a exception;
- Uncomment the line again, and run you app again.

For me, this process always works.

[Download the code](http://github.com/gwtbootstrap/custom-theme-example).

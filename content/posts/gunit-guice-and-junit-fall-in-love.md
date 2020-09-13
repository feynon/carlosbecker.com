---
title: "Guice Junit test-runner"
date: 2012-08-07
draft: false
slug: gunit-guice-and-junit-fall-in-love
city: Joinville
toc: true
tags: []
---

Yesterday, I wrote a [small article]({{< ref guice-and-junit.md >}}) talking about Guice and JUnit, so, this time, I'll just say how to use the small lib that I build (not big deal, one class, one annotation =] )

So, I dont push it to maven central yet, so, you will need to do some work to made it work. Yep, you will need to build and install it to your local repository.

```sh
$ git clone git://github.com/caarlos0/gunit.git
$ cd gunit
$ mvn install
```

Now, just add it in your `pom.xml` dependencies:

```xml
<dependency>
	<groupId>com.github.caarlos0</groupId>
	<artifactId>gunit</artifactId>
	<version>1.0.0</version>
	<scope>test</scope>
</dependency>
```

And follow the instructions code example and the motivation of doing this in [this article]({{< ref guice-and-junit.md >}}), but, basically, your tests will look like this:

```java
@RunWith(GuiceTestRunner.class)
@GuiceModules(FooModule.class)
public class FooTests {

	@Inject Bar bar;

	@Test
	public void testBar(){
	  assertTrue(bar.thisShouldReturnTrue());
	}
}
```

[Get the code](https://github.com/caarlos0/guice-junit-test-runner).

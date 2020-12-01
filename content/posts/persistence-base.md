---
title: "JPA2 with Guice"
date: 2013-06-26
draft: false
slug: persistence-base
city: Joinville
toc: true
tags: [java]
---

Some time ago, I [posted here]({{< ref "modular-persistence.md" >}}) about a simple project that I've done in my post-graduation classes. Since I need it in other projects, and also some friends ended up using it in small projects, I decided to evolve it a little bit.

So, this is just another post (better, I hope) about [that project](https://github.com/caarlos0/persistence-base). It's basically, a "base" for your Java projects that will somehow need to persist data to some database (almost any app). It's baked by Guice, Guava, EclipseLink and Apache B-Val, so, you must have almost everything you need to get started almost instantly (you still need to add the JDBC driver).

In this post, I'll use the version **0.0.5** as base, but it should work exactly the same way in currently **0.0.6-SNAPSHOT** version and hopefully in further versions.

## Project creation

Create a new Maven project and add the following to your `pom.xml`:

```xml
<dependencies>
    <dependency>
      <groupId>com.github.caarlos0</groupId>
      <version>0.0.5</version>
      <artifactId>persistence-base</artifactId>
    </dependency>
  </dependencies>

  <repositories>
    <repository>
      <id>caarlos0-releases</id>
      <url>https://github.com/caarlos0/maven/raw/master/releases</url>
    </repository>
  </repositories>
```

Nice, you now have almost-all needed dependencies (Guava, EclipseLink, Guice, etc). In fact, if you run `mvn dependency:tree`, you will get something like:

```
[INFO] com.carlosbecker:persistence-base-example:jar:1.0-SNAPSHOT
[INFO] \- com.github.caarlos0:persistence-base:jar:0.0.5:compile
[INFO]    +- javax.validation:validation-api:jar:1.0.0.GA:compile
[INFO]    +- org.eclipse.persistence:javax.persistence:jar:2.0.3:compile
[INFO]    +- org.eclipse.persistence:eclipselink:jar:2.4.0:compile
[INFO]    |  \- commonj.sdo:commonj.sdo:jar:2.1.1.v201112051852:compile
[INFO]    +- org.glassfish.external:commons-codec-repackaged:jar:3.0:compile
[INFO]    +- org.apache.bval:bval-core:jar:0.5:compile
[INFO]    |  +- org.apache.commons:commons-lang3:jar:3.1:compile
[INFO]    |  \- commons-beanutils:commons-beanutils-core:jar:1.8.3:compile
[INFO]    +- org.apache.bval:bval-jsr303:jar:0.5:compile
[INFO]    +- com.google.inject:guice:jar:3.0:compile
[INFO]    |  +- javax.inject:javax.inject:jar:1:compile
[INFO]    |  \- aopalliance:aopalliance:jar:1.0:compile
[INFO]    +- com.google.inject.extensions:guice-persist:jar:3.0:compile
[INFO]    \- com.google.guava:guava:jar:14.0.1:compile
```

**HEADS UP**: Don't forget to add your database dependency and the `persistence.xml` file.

## Create a model and a DAO

Just for example, I will create a very simple Person model, something like this:

```java
@Entity
public class Person extends Bean {
  private static final long serialVersionUID = 1L;

  private String firstname;

  // getters, setters and etc are hidden to use less space
}
```

Notice that, since it extends `Bean`, the model already has an `id` and a `version` attribute. In current SNAPSHOT, you will also have the `TimestampedBean`, which provides `created_at` and `updated_at` attributes plus the `id` and `version`.

<!--more-->

**HEADS UP:** Don't forget to add the class to `persistence.xml` file!

---

Now, you should be able to bind a `GenericDao` for this Entity. To do that, create some class exteding `AbstractPersistentModule`, like the following:

```java
public class MainModule extends AbstractPersistentModule {

  @Override
  protected void configure() {
    // 'example' is your persistence unit name
    install(new PersistenceModule("example"));

    // this is a shorthand/syntax sugar to bind a Dao<?> to GenericDao<?>
    bindGenericDaoFor(Person.class);
  }
}
```

Now, in your app main class, you should be able to inject the `GenericDao<Person>`, like the following:

```java
public class App {

  @Inject
  public Dao<Person> personDao;

  public static void main(String[] args) {
    App app = new App();
    Injector injector = Guice.createInjector(new MainModule());
    injector.injectMembers(app);

    Person p = new Person("Carlos");
    daoPerson.save(p);
  }
}
```

You will probably want to create your own `Dao` methods eventually, instead of just use the Generic ones. You can easily achieve this by creating your own `Dao` interface extending `Dao` and the specific `Dao` impl:

```java
public interface AnimalDao extends Dao<Animal> {
  void somecustomMethod(Animal a);
}

public class AnimalDaoImpl extends GenericDao<Animal> implements AnimalDao {
  public void somecustomMethod(Animal a) {
    // impl
  }
}
```

---

**HEADS UP:** Notice that you should be able to get an `EntityManager` instance inside any `GenericDao` specification by calling the `em()` method from the superclass. This method relies on an `EntityManagerProvider`, avoiding some weird session issues.

---

And, of course, you will need to bind it in your module:

```java
bind(AnimalDao.class).to(AnimalDaoImpl.class);
```

After that, you can simply inject it whenever you want (repecting Guice, obviously) with:

```java
@Inject AnimalDao animalDao;
```

And you will be able to call `animalDao.someCustomMethod(animal)` =)

## WIP

This is a almost-infinite-work-in-progress, so, feel free to make pull-requests, suggestions and report eventual bugs.

Cheers!

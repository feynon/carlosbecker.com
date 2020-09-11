---
title: "Don't use junit.framework"
date: 2014-12-14
draft: false
slug: dont-user-junit-framework
city: Joinville
---

Every once in a while I see somewhere a wrong import to the old, deprecated `junit.framework` instead of `org.junit`.

I particularly believe they should remove this package and put it in a `junit-compat` jar or something... but, while they don't do that, we can avoid those old imports with 2 basic steps:

### 1. Replace all old imports with the new imports

This is pretty straightforward with a simple bash script:

```sh
#!/bin/bash
for file in $(git grep --break --heading "import junit." | grep java); do
  sed -i.bak s/junit.framework/org.junit/g $file
  rm -rf $file.bak
done
```

### 2. Don't accept it in new code

The second step is basically to break the build whenever someone try to use those old imports. 

A simple way to do that is using the [restrict-maven-plugin](https://github.com/yamanyar/restrict-maven-plugin).

An example of configuration for maven will look like this:

```xml
<plugin>
  <groupId>com.yamanyar</groupId>
  <artifactId>restrict-maven-plugin</artifactId>
  <version>0.6</version>
  <executions>
    <execution>
      <phase>process-classes</phase>
      <goals>
        <goal>restrict</goal>
      </goals>
    </execution>
  </executions>
  <configuration>
    <continueOnError>false</continueOnError>
    <restrictions>
      <restriction>com.carlosbecker.* to junit.*</restriction>
    </restrictions>
  </configuration>
</plugin>
```

Sure it is a simple issue, which might never cause you problems, but, if one day the Junit team remove the old packages (probably soon), you will have a headache. Besides that, writing new code using deprecated classes?

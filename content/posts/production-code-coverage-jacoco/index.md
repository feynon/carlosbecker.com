---
title: "Measuring production code coverage with JaCoCo"
date: 2017-03-20
draft: false
slug: production-code-coverage-jacoco
city: Joinville
toc: true
tags: [java]
---

Microservices is the new fancy way of doing applications. Yet, most companies still have big and old monoliths in production. In fast evolving software of this size, it's usual to have lines of code which are never executed in production. Production code coverage reports can help
us find those lines.

The strategy described here can be used in virtually any Java software, but what led me to do this was the difficulty to find code that was not being used in a monolith.

## ContaAzul, the monolith

The monolith in question here is [ContaAzul](http://contaazul.com/). It started as a Struts2 application and evolved to a Java EE application, but still has Struts2 code running in it. You can check [these slides](https://pt.slideshare.net/marcoswp3x/tdc-2015-java-from-old-school-to-moder-art) by [@Marcos](https://github.com/marcos) to get an idea on how it evolved.

According to `cloc`, it has ~5k Java files, summing ~400k lines of code. That's a lot of code and I'm not even counting other kinds of source files (like SQL, templates, etc).

In an attempt to remove useless Java code, I had the idea of running the [JaCoCo](https://github.com/jacoco/jacoco) agent in production server for a while to generate the coverage reports.

## Wait, JaCoCo what?

[JaCoCo](https://github.com/jacoco/jacoco) stands for "Java Code Coverage". It is usually used to measure code covered by unit tests.

It is an agent that you pass to the `java` command in the `javaagent` flag. It instruments the *bytecode* and generates a binary report file. Having the source and binary report files, you can generate
human readable HTML reports .

At the build phase, Maven plugins usually does all this work for us. But I don't want the test coverage, I want production coverage. So, I need to add the JaCoCo agent to a production server to instrument production code.

## Instrumenting production code

We don't use Maven in production, so I had to put the agent to run myself.

We use Puppet to manage our servers configuration. So, I wrote a manifest that would download and instrument our Wildfly application server if a `coverage` tag was set on the target node:

```ruby
$jacoco_version='0.7.9'

if tagged('coverage') {
  exec { 'download-jacoco':
    command => "/usr/bin/curl -L -o /storage/environment/jacoco-${jacoco_version}.zip https://repo1.maven.org/maven2/org/jacoco/jacoco/${jacoco_version}/jacoco-${jacoco_version}.zip",
    creates => "/storage/environment/jacoco-${jacoco_version}.zip",
  }~>
  exec { 'unzip-jacoco':
    command     => "/usr/bin/unzip /storage/environment/jacoco-${jacoco_version}.zip -d /storage/environment/jacoco",
    refreshonly => true,
    creates     => '/storage/environment/jacoco',
  }
  file_line { 'jacoco_java_opts':
    path => '/storage/environment/contaazul/bin/standalone.conf',
    line => 'JAVA_OPTS="$JAVA_OPTS -javaagent:/storage/environment/jacoco/lib/jacocoagent.jar=destfile=/storage/environment/contaazul/jacoco.exec,output=file,append=true,dumponexit=true"',
  }
}
```

With that in place, I could tag the nodes that I want coverage reports for:

```ruby
node api {
  tag('coverage')
}
```

But running JaCoCo is only part of the problem. We also need to gather the binary files and generate the HTML reports from them.

## Compiling the reports

The first step is to get the report from the server. The most straightforward way to do that is to use `scp` and that was what I did.

```sh
scp theserver:/storage/environment/contaazul/jacoco.exec .
```

Then, I downloaded and extracted JaCoCo to my `/tmp` folder:

```sh
$ wget https://repo1.maven.org/maven2/org/jacoco/jacoco/0.7.9/jacoco-0.7.9.zip -O /tmp/jacoco.zip
$ unzip /tmp/jacoco.zip -d /tmp/jacoco
```

With all that in place, I wrote a very simple Ant task to generate the HTML reports:

```xml
<project name="ContaAzul" xmlns:jacoco="antlib:org.jacoco.ant">
    <taskdef uri="antlib:org.jacoco.ant" resource="org/jacoco/ant/antlib.xml">
        <classpath path="/tmp/jacoco/lib/jacocoant.jar"/>
    </taskdef>
    <jacoco:report>
        <executiondata>
            <file file="jacoco.exec"/>
        </executiondata>
        <structure name="ContaAzul">
            <classfiles>
                <dirset dir="." includes="**/target/classes"/>
            </classfiles>
            <sourcefiles encoding="UTF-8">
                <dirset dir="." includes="**/src/main/java"/>
            </sourcefiles>
        </structure>
        <html destdir="report"/>
    </jacoco:report>
</project>
```

Then, fire `ant` up and open the report in your web browser:

```sh
$ ant
$ open report/index.html
```

You should see a report like this (this is a fake one):

{{< img caption="Coverage of several packages" src="ef7cbae0-80b1-4187-8829-38332d6cb183.png" >}}

ProTip™: if the Ant task fails, try to run it with `-v` for a more verbose output.

## To automate, or not to automate

At this point I wanted to automate it all to get daily reports (or something like that).

After giving it some thought, I decided that it wouldn't worth it. The reasoning behind this decision was:

- JaCoCo may add some overhead;
- People won't look at these reports every day;
- Running it 24/7 would add this kind of complexity:
    - Merging reports;
    - Dealing with code that is always changing;
    - Ant failures due to class name conflicts (which may happen on big multi-module projects).

I'm also aware that running it this way have its own setbacks:

- I (or someone else) have to manually generate and share the reports;
- Some code may be used only in some days or periods of the day (seasonality), which may induce to humans removing code that is used.

Anyway, with all this code in place it's kind of easy to run it for a while in a server and generate the reports later.

## Results

Having this coverage report is not a free pass to set on fire code that shows up in it. It's more like an inspirational map of things that we can remove. Using it as map, the team was able to open ~20 pull requests which removed ~5k lines of useless code in a very short time.

I'm happy with this result and looking forward to generate more reports and remove even more code.

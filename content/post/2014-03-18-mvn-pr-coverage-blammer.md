---
date: 2014-03-18T00:00:00Z
aliases:
- /posts/mvn-pr-coverage/
slug: mvn-pr-coverage-blammer
title: PullRequest Coverage Blammer Maven Plugin
---

At the [company I work][ca] Pull Requests are part of our culture.
When someone opens a Pull Request, we do Code Review.
If we think it's OK, we comment "+1", or "-1" otherwise.
We usually only merge a PR when it has 3 or more "+1" comments.

Part of this review is to check for tests. We used to manually
look at our code coverage statuses, and see if our recently
added lines are with enough coverage. But this is boring,
and we are developers, and developers automate things,
and so we did.

[@velo][velo] came with the idea of doing a maven plugin to automatically
report bad code coverage in new lines added in pull requests.
In the weekend, I sit next to a pack of beers and a considerable
amount of coffee and made it work. Well, not the 100% one perfect
solution, but we are improving it. And it is [OpenSource][repo]!
You can see a Pull Request example [here][pr]. All my comments
in this PR are actually the plugin working via the Github API
:heart:.

![coverage](/public/images/coverage.png)

We already have [Sonar][sonar] and
[Code Formatter][formatter] Maven Plugins (both written
by [@velo][velo]) that will fail the build in case of code style
changes and report code with bugs and other issues,
respectively, so, it makes sense to us to add coverage
to the party.

We configured those plugins to run within our CI Server,
so, everything is automated: we just open a pull request
and the CI do the rest.

Now, with test coverage reports automated for new code, we
only need to discuss about what really matters: the business.

[ca]: http://github.com/ContaAzul
[velo]: http://github.com/velo
[formatter]: http://lepaysmaudit.blogspot.com.br/2014/03/one-formatter-to-rule-them-all.html
[sonar]: http://lepaysmaudit.blogspot.com.br/2014/03/getting-pull-request-and-sonar-playing.html
[repo]: https://github.com/caarlos0/coverage-maven-plugin
[pr]: https://github.com/caarlos0/coverage-maven-plugin/pull/16

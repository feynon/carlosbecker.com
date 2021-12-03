---
title: "Changelog grouping with GoReleaser v1.1"
date: 2021-12-03
draft: false
slug: goreleaser-changelog-groups
city: Cascavel
toc: true
tags: [goreleaser, golang, semantic-versioning]
---

In the [v1.1 release](https://github.com/goreleaser/goreleaser/releases/tag/v1.1.0), GoReleaser introduced a new feature called ["changelog groups"](https://goreleaser.com/customization/changelog/). This is a quick post to spread the word.

This feature allows you to "organize" your changelog in categories by using regular expressions.

Using it with `use: github` and some exclusion filters yields pretty good looking release notes:

{{< img caption="" src="cfcb5a8b-6fd9-408c-90e7-251b9c93d8aa.png" >}}

Here's a quick usage example:

```yaml
# .goreleaser.yaml
changelog:
  sort: asc
  use: github
  filters:
    exclude:
    - Merge pull request
    - Merge remote-tracking branch
    - Merge branch
  groups:
    - title: 'New Features'
      regexp: "^.*feat[(\\w)]*:+.*$"
      order: 0
    - title: 'Bug fixes'
      regexp: "^.*fix[(\\w)]*:+.*$"
      order: 10
    - title: Other work
      order: 999
```

GoReleaser uses it on its [own releases](https://github.com/goreleaser/goreleaser/releases), so you can always poke around its [.goreleaser.yaml config file](https://github.com/goreleaser/goreleaser/blob/main/.goreleaser.yaml). Also make sure to give the [documentation](https://goreleaser.com/customization/changelog/?h=groups) a read!

---

## Acknowledgements

This feature was added by [@dirien](https://github.com/dirien) in [this PR](https://github.com/goreleaser/goreleaser/pull/2670)! Thanks for all the hard work! 🙏

# [carlosbecker.com](https://carlosbecker.com) [![Build Status](https://travis-ci.org/caarlos0/carlosbecker.com.svg?branch=master)](https://travis-ci.org/caarlos0/carlosbecker.com)

Sources to my blog.


## create new post

```sh
hugo new "post/$(date +%Y-%m-%d)-$(filter-out $@,$(MAKECMDGOALS)).md"
```

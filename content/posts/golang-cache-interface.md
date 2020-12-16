---
title: "Golang: cache things using interfaces"
date: 2018-12-17
draft: false
slug: golang-cache-interface
city: Joinville
toc: true
tags: [golang]
---

Caching things can be hard to do and hard to test. In this post I'll demonstrate a convenient way of doing that using interfaces.

---

## The problem

Let's suppose we want to cache calls to the GitHub API. Let's say we want
to get the my repository list for whatever reason:

```go
package client

type Repository struct {
	Name string `json:"name"`
}

func GetRepositories() ([]Repository, error) {
	// TODO: do a call to https://api.github.com/users/caarlos0/repos
	return []Repository{}, nil
}
```

Let's say this will be called from another website, so, it could easily
rate-limit. This information also does not change much, and users won't care
if they see data for 5 minutes ago.

We could easily do that using an in-memory cache, like [go-cache](https://github.com/patrickmn/go-cache). The
first thing that comes to mind is to something like:

```go
var cache = cache.New(5*time.Minute, 5*time.Minute)

func GetRepositories() ([]Repository, error) {
	cached, found := cache.Get("my-repos")
	if found {
		return cached.([]Repository), nil
	}
	// TODO: do a call to https://api.github.com/users/caarlos0/repos
	// result := blah
	c.cache.Set(repo, result, cache.DefaultExpiration)
	return []Repository{}, nil
}
```

But this comes with a couple of problems:

- cache is global;
- the client now has a hard dependency on the cache;
- can't test the caching logic separately.

## Interfaces to the rescue

One way to solve all those problems is to create an interface and decorate
the "real" client implementation with another implementation that only handles
caching.

For our example, we could create an interface like this:

```go
// client.go
package client

type Repository struct {
	Name string `json:"name"`
}

type Client interface {
	GetRepositories() ([]Repository, error)
}
```

And then the "real" implementation:

```go
// github.go
package client

func NewGithubClient() Client {
	return ghClient{}
}

type ghClient struct {}

func (ghClient) GetRepositories() ([]Repository, error) {
	// TODO: do a call to https://api.github.com/users/caarlos0/repos
	return []Repository{}, nil
}
```

And finally, a cached implementation that wraps any other `Client`
implementation:

```go
// cache.go
package client

func NewCachedClient(client Client, cache *cache.Cache) Client {
	return cachedClient{
		client: client,
		cache:  cache,
	}
}

type cachedClient struct {
	client Client
	cache *cache.Cache
}

func (c cachedClient) GetRepositories() ([]Repository, error) {
	cached, found := c.cache.Get("my-repos")
	if found {
		return cached.([]Repository), nil
	}
	// call the underlying client
	live, err := c.client.GetRepositories()
	c.cache.Set(repo, result, cache.DefaultExpiration)
	return live, err
}
```

And that would be it. This also enabled us to test the caching only.

## Enter testing

In our example, we can test the cache implementation pretty easily: we just
need to create a fake client implementation and wrap it in `cachedClient`,
and then write some tests for it.

Code example of a very simple implementation:

```go
// cache_test.go
package client

type cacheTestClient struct {
	result *[]Repository
}

func (f cacheTestClient) GetRepositories() ([]Release, error) {
	return *f.result, nil
}

func TestCachedClient(t *testing.T) {
	var cache = cache.New(1*time.Minute, 1*time.Minute)
	var expected = []Repository{
		{ Name: "caarlos0/version_exporter" },
		{ Name: "caarlos0/dotfiles" },
		{ Name: "caarlos0/carlosbecker.com" },
	}

	var cli = NewCachedClient(cacheTestClient{result: &expected}, cache)

	// test getting from out fake client
	t.Run("get fresh", func(t *testing.T) {
		res, err := cli.GetRepositories()
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})


	// here we change the inner fake client result, but the result
	// should be the cached one
	t.Run("get from cache", func(t *testing.T) {
		var oldExpected = expected
		expected = append(rel, Repository{Name: "caarlos0/env"})
		res, err := cli.GetRepositories()
		require.NoError(t, err)
		require.Equal(t, oldExpected, res)
	})


	// here we flush the cache and verify that the result is the one
	// from the fake client
	t.Run("flush cache", func(t *testing.T) {
		c.Flush()
		res, err := cli.GetRepositories()
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})
}
```

Althought this is a simple example, it is also functional.

You can for sure write a smarter fake client (the example doesn't handle errors
for example), and can also use this strategy for a Redis-backed cache, for
any API calls and also for SQL databases for example.

This interface + decoration strategy can be used for other features, for
example, for circuit breakers and things like that. It makes it easier to
decouple implementations and to test them.

Hope this was useful for you.

> Side note: the examples provided here are based on real code from
> my version_exporter repository.

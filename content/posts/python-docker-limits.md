---
title: "Making Python respect Docker memory limits"
slug: python-docker-limits
city: Cascavel
date: 2020-11-20
toc: true
draft: false
slug: python-docker-limits
tags: [docker, kubernetes, python]
---


If you run Python inside containers, chances are you have seen Linux's OOMKiller working at least a couple of times.

---

This happens because Python sees the entire host's resources as if they were available for its use.

Then, it may try to allocate more memory than it is allowed to, which causes Linux to kill the process.

To fix that, we may read the actual limits from `/sys/fs/cgroup/memory/memory.limit_in_bytes` and set it as the process max address space area.

Example:

```python
import resource
import os
import time

if os.path.isfile('/sys/fs/cgroup/memory/memory.limit_in_bytes'):
    with open('/sys/fs/cgroup/memory/memory.limit_in_bytes') as limit:
        mem = int(limit.read())
        resource.setrlimit(resource.RLIMIT_AS, (mem, mem))

x = bytearray(900*1024*1024) # allocate 900mb

print('ok')
time.sleep(20)
```

Save it as `mem.py`, and lets run some tests:

```bash
$ docker run \
  --rm \
  -m 1G \
  -v $PWD:/tmp \
  python:rc-alpine \
  python /tmp/mem.py
```

If you run `docker stats` in another container, you'll see it takes ~900Mb of RAM, and everything is fine.

With or without setting the `RLIMIT_AS`, this would already work.

Now, let's try to change:

```diff
-x = bytearray(900*1024*1024) # allocate 900mb
+x = bytearray(4000*1024*1024) # allocate 4000mb
```

And we'll see:

```bash
$ docker run --rm -m 1G -v $PWD:/tmp python:rc-alpine python /tmp/mem.py
Traceback (most recent call last):
  File "/tmp/mem.py", line 10, in <module>
    x = bytearray(4000*1024*1024) # allocate 4000mb
MemoryError
```

Now, let's try something different:

```bash
docker run \
  --rm \
  -m 1G \
  -v $PWD:/tmp \
  python:rc-alpine \
  python -c "x = bytearray(4000*1024*1024); print('ok')"
```

You'll see the container gets killed without printing 'ok'.

I think most of the time you'll want a `MemoryError` instead of a `SIGKILL`, if that's the case, running the first snipped in the `__init__` of your code might be a good workaround.

## -XX:+UseContainerSupport

I think Python should probably copy Java and either make this the default behavior or hide it in a flag.

In my perception, most people would expect this to be the default... as well as a lot of other tools (`top` for example).

## Workarounds

As stated before, this is just a workaround.

Its possible it doesn't work on all distributions, or that it stops working in the future.

Hopefully this becomes a native option on Python soon.

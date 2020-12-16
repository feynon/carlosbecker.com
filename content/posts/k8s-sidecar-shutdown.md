---
title: "Kubernetes Jobs: shut down sidecar when main container finishes"
date: 2020-02-28
draft: false
slug: k8s-sidecar-shutdown
city: Joinville
toc: true
tags: [kubernetes]
---

Kubernetes Pod lifecycle does not cover everything just yet.

---

I'm working on an app that listens to Buildkite `job.scheduled` webhook and creates a Kubernetes Job to execute it.

The idea is to have a "scale-to-zero" approach, with the max number of nodes being defined by the the node pool settings.

Buildkite has a pretty nice CLI which allows us to do start the agent with something like:

```sh
buildkite-agent-entrypoint start \
	--disconnect-after-job \
	--disconnect-after-idle-timeout=50
```

This will start the agent, it will either get a job from the queue, execute it and exit, OR wait for 50 seconds, and if no job arrives for it, exit.

The problem is that some of our builds (if not most of them) require a sidecar with `docker:dind`. So, OK, the agent container `exit 0`, but the `dind` sidecar is still there, waiting for nothing.

There is an [open proposal on sig-apps](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/sidecarcontainers.md#proposal) to add better lifecycle to this, but its not implemented yet.

So, here we go with the ugly hack!

## Making it work

The idea is actually pretty simple:

- we add an `emptyDir` volume to both containers
- `agent` writes a file to it before exiting
- `dind` waits for that file, and when it is created, `exit 0`

So, what we do, in YAML, is:

```yaml
# ...
containers:
- name: agent
  command: ["/bin/sh", "-c"]
  args:
  - |
    trap 'touch /usr/share/pod/done' EXIT
    buildkite-agent-entrypoint start \
      --disconnect-after-job \
      --disconnect-after-idle-timeout=50
# ...
  volumeMounts:
  - mountPath: /usr/share/pod
    name: tmp-pod
- name: dind
  command: ["/bin/sh", "-c"]
  args:
  - | 
    dockerd-entrypoint.sh &
    while ! test -f /usr/share/pod/done; do
      echo 'Waiting for the agent pod to finish...'
      sleep 5
    done
    echo "Agent pod finished, exiting"
    exit 0
# ...
  volumeMounts:
  - mountPath: /usr/share/pod
    name: tmp-pod
    readOnly: true
# ...
volumes:
- emptyDir: {}
  name: tmp-pod
# ...
```

What we do:

- `trap` the `EXIT` to create a `/usr/share/pod/done` file and then start the agent;
- start docker in background (`&`) and wait for `/usr/share/pod/done` to exist, when it happens, `exit 0`.

Important details:

- on `dind`, we mount the volume as `readOnly`;

---

Its a pretty simple, although hacky, solution. Hope its useful for you somehow.

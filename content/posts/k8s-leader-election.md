---
title: "Leader Election inside Kubernetes"
date: 2020-03-14
draft: false
slug: k8s-leader-election
city: Joinville
toc: true
tags: [kubernetes]
---

Ever needed a simple leader election mechanism on something that will run on a Kubernetes cluster? There's an easy way to do that!

<!--more-->

A couple of days ago I was working on an app that needed to do some work from time to time, but only in a single replica.

Easy thing to do if you have only one replica running. That was not my case though.

I considered a couple of different things, like creating a `CronJob` and another docker image which will do only this job... but I wanted something more simple and self-contained.

After a some minutes thinking about running an `etcd` backend and all that, I though "well maybe Kubernetes has something for that already".

And then I found the `coordination.k8s.io` API, and the [leader election package in the Go SDK](https://pkg.go.dev/k8s.io/client-go/tools/leaderelection).

## Leader Election

As you can read in [the docs](https://pkg.go.dev/k8s.io/client-go/tools/leaderelection?tab=doc), it does not do fencing and there may be more than one nodes leading. On the app I was working on, this is OK. I just didn't want to have all nodes doing that thing all the time.

Anyway, here's how I did it.

### Identity

We'll need to be able to tell which replica is holding the lease. You could just generate a new UUID every time, but then it's hard to know, just by `kubectl describe lease`, which node is holding it right now - it will also change if the pod restarts for whatever reason.

I prefer to use the pod name.

To do that, we can accept a flag, like this:

```go
package main

import "flag"

func main() {
	var nodeID = flag.String("node-id", "", "node id")
	flag.Parse()
}
```

And then pass it on the the deployment:

```yaml
spec:
  template:
    spec:
      containers:
        - args:
            - "-node-id=$(POD_NAME)"
          env:
            - name: POD_NAME
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: metadata.name
```

This way, the identifier will always be the pod name.

### Electing the leader

First thing we need to do is to create a `context`:

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

Then, we need to create the `lock` object:

```go
var lock = &resourcelock.LeaseLock{
	LeaseMeta: metav1.ObjectMeta{
		Name:      "my-lock",
		Namespace: "default",
	},
	Client: clientset.CoordinationV1(),
	LockConfig: resourcelock.ResourceLockConfig{
		Identity: nodeID,
	},
}
```

Finally, we use the `leaderelection` API to do the rest:

```go
leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
	Lock:            lock,
	ReleaseOnCancel: true,
	LeaseDuration:   15 * time.Second,
	RenewDeadline:   10 * time.Second,
	RetryPeriod:     2 * time.Second,
	Callbacks: leaderelection.LeaderCallbacks{
		OnStartedLeading: func(ctx context.Context) {
			log.WithField("id", nodeID).Info("started leading")
			// work
		},
		OnStoppedLeading: func() {
			log.WithField("id", modeID).Info("stopped leading")
			// stop working
		},
		OnNewLeader: func(identity string) {
			if identity == options.NodeID {
				return
			}
			log.WithField("id", nodeID).
				WithField("leader", identity).
				Info("new leader")
		},
	},
})
```

You can tune the parameters as you wish. You'll also need to implement the *work* and *stop working* logic. You can do that using, for example, `atomic.StoreInt32` and `atomic.LoadInt32` to create a local "lock".

But that's pretty much it.

### RBAC

The default Service Account does not have access to the coordination API, so we'll need to create another and set up RBAC accordingly.

Something like this:

```yaml
apiVersion: v1
automountServiceAccountToken: true
kind: ServiceAccount
metadata:
  name: mysa

---

apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: myrole
rules:
  - apiGroups:
      - coordination.k8s.io
    resources:
      - leases
    verbs:
      - '*'

---

apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: mysa-myrole
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: myrole
subjects:
  - kind: ServiceAccount
    name: mysa
```

And then we just need to use that account on our deployment:

```yaml
# ...
spec:
  template:
    spec:
      automountServiceAccountToken: true
      serviceAccount: leaderz
      serviceAccountName: leaderz
# ...
```

And that should be it.

## Running

Deploying all that, you can see the lease being created and changing over time:

```bash
$ kubectl describe lease
Name:         my-lock
Namespace:    default
Labels:       <none>
Annotations:  <none>
API Version:  coordination.k8s.io/v1
Kind:         Lease
Metadata:
  Creation Timestamp:  2020-03-05T20:17:00Z
  Resource Version:    118745761
  Self Link:           /apis/coordination.k8s.io/v1/namespaces/default/leases/my-lock
  UID:                 d68c26d7-2c19-47cb-bc8a-c69cf2c67e7b
Spec:
  Acquire Time:            2020-03-06T19:38:10.890381Z
  Holder Identity:         my-app-59d4f98568-6twdw
  Lease Duration Seconds:  15
  Lease Transitions:       12
  Renew Time:              2020-03-09T18:59:33.590502Z
Events:                    <none>
```

We can see the pod `my-app-59d4f98568-6twdw` is holding the lease, so this is pod leading right now.

## A complete example

You can find a complete working example on [this GitHub Repository](https://github.com/caarlos0/leaderz).

You can apply it with:

```sh
$ kubectl apply -f kube/
```

And then watch the pods logs:

```sh
$ stern leaderz
```

And see the lease:

```sh
$ kubectl describe lease my-lock
```

You can then kill pods and watch the lease eventually be assigned to another pod.

## Final words

Its not a perfect solution for leader election, as stated before, but depending on your needs, its a simple enough way of having it, without needing to run more things on the cluster or coding a lot.

I hope you find it useful. 🍻

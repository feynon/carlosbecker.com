---
title: "Kubernetes pod shutdown lifecycle"
date: 2020-07-30
draft: false
slug: k8s-pod-shutdown-lifecycle
city: Joinville
---

I always forget the details about Kubernetes pod shutdown lifecycle is something when I need them, so this is my now made public notes on the subject.

<!--more-->

## TL;DR

When a pod is signaled to terminate (deleted, for example), what happens is:

1. it enters in `Terminating` phase
2. run all `preStop` lifecycle hooks
3. sends a `SIGTERM`

This entire process has a timeout defined in the `terminationGracePeriodSeconds` (30 seconds by default). After that, Kubernetes sends a `SIGKILL` to the container.

During this entire process, `readiness` and `liveness` probes will still be probed, but their failure will not cause the killing of the container, as it is already being killed.

## Playground

Let's setup a playground so we can explore this.

The sources I used are bellow:

```yaml
# deploy.yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: test
spec:
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: test
    spec:
      terminationGracePeriodSeconds: 18
      containers:
      - name: test
        image: caarlos0/sinkhole
        ports:
        - containerPort: 8080
        lifecycle:
          preStop:
            exec:
              command: ["sleep", "8"]
        livenessProbe:
          httpGet:
            path: /live
            port: 8080
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
```

The `test` image can be generated with this source code:

```go
// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var addr = ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if sleep := r.URL.Query().Get("sleep"); sleep != "" {
			d, err := time.ParseDuration(sleep)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			time.Sleep(d)
		}

		defer r.Body.Close()
		log.Println(r.URL)
		fmt.Fprintln(w, "ok")
	})

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("server started on " + addr)

	<-done
	log.Println("server stop requested")

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed:%+v", err)
	}
	log.Print("server exit properly")
}
```
```docker
# Dockerfile
FROM golang:alpine as build
WORKDIR /src
COPY . .
RUN go build -o server main.go

FROM alpine
COPY --from=build /src/server server
ENTRYPOINT ["./server"]
```

And then just run:

```shell
docker build -t test .
```

## Testing

We can deploy our testing image with

```shell
kubectl apply -f deploy.yaml
```

And forward it to our local machine:

```shell
kubectl port-forward `kubectl get po -l app=test -oname` 8080:8080
```

From now on, we'll need a couple of terminals.

### Terminal #1

Simulate normal requests:

```shell
while true; do 
	curl localhost:8080/fake
	sleep 1
done
```

### Terminal #2

Simulate a long living request:

```shell
curl localhost:8080/slow\?sleep=30m
```

### Terminal #3

Watch the pod's events:

```shell
while true; do 
	kubectl describe -l app=test
	sleep 1
done
```

### Terminal #4

Watch the pods logs:

```shell
kubectl logs -f -l app=test
```

### Terminal #5

Delete the pod:

```shell
kubectl delete po -l app=test
```

## What happens

As soon as we delete the pod in **Terminal #5**, our describe on **Terminal #3** will show something like this:

```shell
Normal  Killing    5s    kubelet, node-12108603-9g16  Stopping container test
```

This will trigger our `preStop` hook, which is just a 8 seconds sleep.

After some time, we see more events:

```shell
Warning  Unhealthy  8s    kubelet, node-12108603-9g16  Readiness probe failed: Get http://172.16.6.7:8080/ready: dial tcp 172.16.6.7:8080: connect: connection refused
Warning  Unhealthy  5s    kubelet, node-12108603-9g16  Liveness probe failed: Get http://172.16.6.7:8080/live: dial tcp 172.16.6.7:8080: connect: connection refused
```

Also, on our pod's logs, we see a `server stop requested`, meaning we got a `SIGTERM`, and the server is not accepting new incoming request anymore. Our slow request on **Terminal #2** is still there though.

After another ~10s (`terminationGracePeriodSeconds`'s 18s - `preStop` hook's 8s), the pod receives a `SIGKILL` because it hasn't finished in time due to our slow request.

We can now run this scenario again, but this time, have the slow request with a smaller time (e.g. `15s`), and the pod exit, and we'll see in its logs that it `exit 0`:

```shell
2020/07/30 18:47:03 server stop requested
2020/07/30 18:47:03 server exit properly
```

That's because all requests finished, so our process finished as well.

## Final words

Kubernetes is powerful, and because you can do a lot of things, its normal to be confused by it at times. 

Particularly, the shutdown behavior confused me more times than I can count ("does the pre-stop hook account on the grace period? I don't remember" et al), that's why I noted this for quick reference, and thought I might as well make it public.

The Go app used in the example is an app that I use often to test/debug things that need dummy a web server, you can just use `caarlos0/sinkhole` instead of building locally if prefer (also check the [full source code](https://github.com/caarlos0/sinkhole)).

That's it, hope it helps. 🙂

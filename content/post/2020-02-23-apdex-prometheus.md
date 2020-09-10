---
title: "Calculating the Apdex score with Prometheus"
date: 2020-02-23
draft: true
slug: apdex-prometheus
city: Joinville
---

Apdex is an industry standard to measure application performance.

<!--more-->

The only thing we need to define is a `target` response time. With that, we can use the following formula to calculate the apdex of a given app:

```bash
apdex = (satisfactory + (tolerable / 2)) / total
```

In this formula:

- All non 5xx requests served under `target` are considered satisfactory
- All non 5xx request server above `target` and under `target*4` are considered tolerable
- All requests served above `target` are considered intolerable
- All 5xx requests are considered intolerable

OK, but how can we do this with Prometheus?

## A sample web app

We will need 2 things to calculate the apdex of our service:

1. a gauge that says what is the app's `target`
2. a histogram that has the buckets for `target` and `target * 4`

Here's an example in Go:

```go
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/campoy/unique"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var addr = flag.String("listen-address", "127.0.0.1:8080", "The address to listen on for HTTP requests.")
var target = flag.Duration("target", time.Millisecond*100, "apdex target response time")

func main() {
	flag.Parse()

	// set up the buckets
	var buckets = prometheus.DefBuckets
	buckets = append(buckets, target.Seconds())
	buckets = append(buckets, target.Seconds()*4)
	// needs to be ordered and must not have duplicated items
	unique.Slice(&buckets, func(i int, j int) bool {
		return buckets[i] < buckets[j]
	})

	// set up the apdex target metric
	promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "http",
		Subsystem: "apdex",
		Name:      "target_seconds",
	}).Set(target.Seconds())

	var mux = http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	// test endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch rand.New(rand.NewSource(time.Now().UnixNano())).Intn(4) {
		case 0:
			fmt.Fprintln(w, "ok")
		case 1:
			time.Sleep(*target * 2)
			fmt.Fprintln(w, "ok but slow")
		case 2:
			time.Sleep(*target * 5)
			fmt.Fprintln(w, "ok but very slow")
		case 3:
			http.Error(w, "error", http.StatusInternalServerError)
		}
	})

	var observer = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Subsystem: "request",
		Name:      "duration_seconds",
		Buckets:   buckets,
	}, []string{"code", "method"})

	fmt.Println("listening on", *addr)
	log.Fatal(http.ListenAndServe(*addr, promhttp.InstrumentHandlerDuration(observer, mux)))
}
```

OK, let's also setup a Prometheus server monitoring our app, create a `prometheus.yaml` file like this:

```yaml
global:
  scrape_interval: 10s
  scrape_timeout: 10s
  evaluation_interval: 10s
rule_files:
  - rules.yaml
scrape_configs:
- job_name: app
  static_configs:
  - targets:
    - localhost:8080
```

And a `rules.yaml` like this (for now):

```yaml
groups:
- name: apdex
  rules: []
```

You should be able to start both the app and prometheus, and metrics should show up:

```shell
$ go run main.go
$ prometheus --config.file ./prometheus.yaml
```

IMAGEM AQUI

OK, now let's start working on our recording rules!

### Calculating the Apdex

We may have several instances of the same app running, and they may have different targets. This should be a temporary thing, so, lets just assume the max:

```yaml
- record: job:http_apdex_target_seconds:max
  expr: max(http_apdex_target_seconds) BY (job)
```

OK, next step is to have a recording rule for both `target` and `target * 4`:

```yaml
- record: job_le:http_apdex_target_seconds:max
  expr: |
    clamp_max(
      count_values(
        "le",
        job:http_apdex_target_seconds:max
      ) BY (job),
      1
    )

- record: job_le:http_apdex_target_seconds:4_times_max
  expr: |
    clamp_max(
      count_values(
        "le",
        job:http_apdex_target_seconds:max * 4
      ) BY (job),
      1
    )
```

This will create 2 metrics with the prometheus bucket label (`le`) and both will have the value `1`. We'll use them to match on our apdex query. We could also have those metrics being exported in the app itself.

We now need to calculate a couple of things:

**Satisfactory:**

```yaml
sum(
  rate(http_request_duration_seconds_bucket{status_code!~"5.."}[1m])
  *
  ON(job, le) GROUP_LEFT() job_le:http_apdex_target_seconds:max
) BY (job)
```

**Tolerable:**

```yaml
sum(
  rate(http_request_duration_seconds_bucket{status_code!~"5.."}[1m])
  *
  ON(job, le) GROUP_LEFT() job_le:http_apdex_target_seconds:4_times_max
) BY (job)
```

**Total:**

```yaml
sum(rate(http_request_duration_seconds_count[1m])) BY (job)
```

That's all we need!

The end formula will look like this:

```yaml
- record: job:http_apdex
  expr: |
    (
      (
        sum(
          rate(http_request_duration_seconds_bucket{status_code!~"5.."}[1m])
          *
          ON(job, le) GROUP_LEFT() job_le:http_apdex_target_seconds:max
        ) BY (job)
        +
        sum(
          rate(http_request_duration_seconds_bucket{status_code!~"5.."}[1m])
          *
          ON(job, le) GROUP_LEFT() job_le:http_apdex_target_seconds:4_times_max
        ) BY (job)
      ) / 2
    )
    /
    sum(rate(http_request_duration_seconds_count[1m])) BY (job)
```

We could probably also alert when Apdex is bad. Bellow `0.8` is pretty standard, so let's do it:

```yaml
- alert: HTTPApdexViolation
  expr: job:http_apdex < 0.8
  for: 5m
  annotations:
    summary: '`{{ $labels.job }}`: apdex is too low'
    description: 'ApDex has dropped below 0.8: {{printf "%.2f" $value}}'
```

And that's it!

The full `rules.yaml` will look like this:

```yaml

```
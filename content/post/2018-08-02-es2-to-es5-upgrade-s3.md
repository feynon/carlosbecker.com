---
title: "Upgrading ElasticSearch 2 to 5: S3 snapshot/restore strategy"
date: 2018-08-02T10:08:53-03:00
city: Joinville
slug: es2-to-es5-upgrade-s3
tags:
- elasticsearch
---

Migrating an ElasticSearch cluster from version 2 to 5 can be challenging,
even more if it is a big cluster.

In this post we will explore one of the strategies that can be used to do
such migration, setting a up a playground environment in which we can learn
the migration procedures and test things out safely.

<!--more-->

# Definition List

List of terms we will use:

ES2
: ElasticSearch v2.x.y

ES5
: ElasticSearch v5.x.y

# Introduction

Let's say you need to do a migration like that, you may come up with three
strategies:

1. use rack awaraness to split the cluster and selectively upgrade one rack
to ES5 later on;
1. snaphost on ES2 and restore on ES5 using NFS;
1. snapshot on ES2 and restore on ES5 using S3 (this post subject).

The rack awareness strategy requires more manual intervertions and therefore
is probably less safe (there will be a post about that). About NFS... well, S3
is simpler and probably safer.

Since production is no place to play around and test things out, I
created two docker-compose environments to learn, test and polish the
procedure so I can eventually do it in a production cluster with more
security.

So, without further due, let's get started!

# ES2 Cluster Setup

First, we need an ES2 cluster. For testing purposes, we'll create a 4 node
cluster, one master and three data nodes:

```yaml
# es2/docker-compose.yml
version: '2.2'

services:
  master:
    build: .
    image: es2-plugs
    env_file:
    - .env
    volumes:
    - ./master.yml:/usr/share/elasticsearch/config/elasticsearch.yml
    ports:
    - "9200:9200"
  data:
    image: es2-plugs
    env_file:
    - .env
    volumes:
    - ./data.yml:/usr/share/elasticsearch/config/elasticsearch.yml
    scale: 3
    depends_on:
    - master
```

Note that this docker-compose file builds a new image called `es2-plugs` and
also depends on some config files.

The `Dockerfile` looks like this:

```dockerfile
# es2/Dockerfile
FROM elasticsearch:2.4.6-alpine
# cloud-aws is required for s3 snapshots
RUN /usr/share/elasticsearch/bin/plugin install --batch cloud-aws
# elasticsearch-migration shows us what we need to change in order to migrate
# to es5
RUN /usr/share/elasticsearch/bin/plugin install --batch https://github.com/elastic/elasticsearch-migration/releases/download/v2.0.4/elasticsearch-migration-2.0.4.zip
```

We also need an `.env` file:

```env
AWS_ACCESS_KEY_ID=your key
AWS_SECRET_KEY=your secret
```

> **PROTIP**: If you don't wan't to use a real S3 bucket, you can start a local
> [minio][] server and change set the `cloud.aws.s3.endpoint` on all
> `elasticsearch.yml` files.

[minio]: https://www.minio.io/

The `data.yml` file:

```yaml
# es2/data.yml
network.host: 0.0.0.0
cluster.name: beckerz
discovery.zen.ping.unicast.hosts:
  - master:9300
node.master: false
node.data: true
```

And finally the `master.yml` file:

```yaml
network.host: 0.0.0.0
cluster.name: beckerz
discovery.zen.minimum_master_nodes: 1
node.master: true
node.data: false
```

So, we should have a `es2` folder with the following structure:

```sh
.
└── es2
    ├── .env
    ├── Dockerfile
    ├── data.yml
    ├── docker-compose.yml
    └── master.yml
```

Now, we can just up our env with:

```sh
docker-compose up
```

You can check that the cluster is green:

```sh
curl -s "localhost:9200/_cluster/health"
```

It will build the image, launch one master and three data nodes talking to that
master.

# Adding some fake data

Let's create a `customer` index and add some fake data to it:

```sh
curl -sXPUT 'http://localhost:9200/customer/?pretty' -d '{
  "settings" : {
      "index" : {
          "number_of_shards" : 6,
          "number_of_replicas" : 2
      }
  }
}'

while ! curl -s "localhost:9200/_cat/indices?v" | grep green; do
  sleep 0.1
done

for i in `seq 1 1000`; do
  curl -sXPUT "localhost:9200/customer/external/$i?pretty" -d "
  {
    \"number\": $i,
    \"name\": \"John Doe - $i\"
  }"
done
```

Once that is done, time to snapshot it!

# Snapshotting ES2

The first thing you need to do is to create a repository, let's call it
`backups`:

```sh
curl -sXPUT "localhost:9200/_snapshot/backups?pretty" -d'
{
  "type": "s3",
  "settings": {
    "bucket": "my-bucket",
    "base_path": "my-subfolder"
  }
}
'
```

You can check it with:

```sh
curl -s "localhost:9200/_snapshot?pretty"
```

Now, let's snapshot everything:

```sh
curl -sXPUT "localhost:9200/_snapshot/backups/snapshot_1?wait_for_completion=true&pretty"
```

Once it is done, you can already restore it on another cluster.

# Check the elasticsearch-migration plugin

You can explore the migration plugin by going to
http://localhost:9200/_plugin/elasticsearch-migration

# Setup an ES5 cluster

We'll create another 4-node cluster using docker-compose, this time running
ES5:

```yaml
# es5/docker-compose.yml
version: '2.2'

services:
  master:
    build: .
    image: es5-plugs
    env_file:
    - .env
    environment:
    - "ES_JAVA_OPTS=-Xms512m -Xmx512m"
    volumes:
    - ./master.yml:/usr/share/elasticsearch/config/elasticsearch.yml
    ports:
    - "9400:9200" # this master wil bind to your local 9400 port
  data:
    image: es5-plugs
    depends_on:
    - master
    scale: 3
    env_file:
    - .env
    environment:
    - "ES_JAVA_OPTS=-Xms512m -Xmx512m"
    volumes:
    - ./data.yml:/usr/share/elasticsearch/config/elasticsearch.yml
```

We can use the same `.env`, `master.yml` and `data.yml` file from our ES2 env.
The `Dockerfile` is different though:

```dockerfile
# es5/Dockerfile
FROM elasticsearch:5.6.10-alpine
# repository-s3 is required for s3 snapshots
RUN /usr/share/elasticsearch/bin/elasticsearch-plugin install --batch repository-s3
```

So we will have the following structure now:

```sh
.
├── es2
│   ├── .env
│   ├── Dockerfile
│   ├── data.yml
│   ├── docker-compose.yml
│   └── master.yml
└── es5
    ├── .env
    ├── Dockerfile
    ├── data.yml
    ├── docker-compose.yml
    └── master.yml
```

Great, let's fire this cluster up!

```sh
docker-compose up
```

You can check that the cluster is green:

```sh
curl -s "localhost:9400/_cluster/health"
```

> **Attention**: Note that the master port is now 9400.

We should now have one master and three data nodes running ES5 as well!

# Restore ES2 snapshot into ES5

For that, we need to create a repository with the same arguments we used on
ES2.

```sh
curl -sXPUT "localhost:9400/_snapshot/backups?pretty" -d'
{
  "type": "s3",
  "settings": {
    "bucket": "my-bucket",
    "base_path": "my-subfolder"
  }
}
'
```

And then, finally, restore that snapshot:

```sh
curl -sXPOST "localhost:9400/_snapshot/backups/snapshot_1/_restore?wait_for_completion=true&pretty"
```

You can always check that documents count and settings match and etc:

```sh
curl -s "localhost:9200/customer/_settings?pretty"
curl -s "localhost:9400/customer/_settings?pretty"
curl -s "localhost:9200/_count?pretty"
curl -s "localhost:9400/_count?pretty"
```

# Incremental snapshots

So, on a _real_ scenario, snapshots would probably take a lot of time (let's
say the cluster have hundreds of terabytes), and your users will still want to
use the app meanwhile.

To minimize downtimes, you can probably do the following:

1. full snapshot
1. restore full snapshot
1. incremental snapshot 1
1. restore incremental snapshot 1
1. bring app down or put it in read-only mode
1. incremental snapshot 2
1. restore incremental snapshot 2
1. bring app up again, pointing to the new cluster

Doing it like that, your downtime will be reduced to the time between steps
5 and 8. If your app have such feature, you can leave it in read-only mode for
some time to validate more things and avoid a split brain scenario, and once
you are confident, finnaly enable everything - from that time on we can't go
back anymore.

# Timings

I tested the same documents in both Amazon S3 and DigitalOcean Spaces.
Amazon S3 seemed orders of magnitude faster (70% or so), so I would
probably recommend you use S3 for production environments - or just measure
it yourself 🤷‍♂️.

# Conclusion

Using this strategy is fairly easy and you can basically abort it at any time
if you find bugs in your app, for example.

I hope that the receipts for the test environment serve you well and that you
can also use them to practice in a safe environment before working on "real"
environments.

# Next steps

On the next posts we will explore the rack awareness strategy.

# Links

- [ElasticSearch documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-snapshots.html)

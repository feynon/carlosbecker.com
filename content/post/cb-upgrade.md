---
title: "Couchbase: rolling upgrade from 4.5.x to 5.1.x"
date: 2018-10-10
draft: false
slug: cb-upgrade
city: Joinville
---

I have an old Couchbase 4.5.x cluster, and I though it would be nice to upgrade it. This are my notes and the tests I did before doing it "in production"™️.

<!--more-->

## Breaking changes

First thing I did was going through the changelogs and look for breaking
changes.

- [5.5.x](https://docs.couchbase.com/server/5.5/release-notes/relnotes.html)
- [5.1.x](https://docs.couchbase.com/server/5.1/release-notes/relnotes.html)
- [5.0.x](https://docs.couchbase.com/server/5.0/release-notes/relnotes.html)
- [4.5.x](https://docs.couchbase.com/server/4.5/release-notes/relnotes.html)

Turns out there is only one: [MB-18042 FTS UI: Remove Byte Array Converter from UI - Couchbase](https://issues.couchbase.com/browse/MB-18042).

Since I don't use the feature, not a problem, yay! 🚀

The last community version available at the time of writing is 5.1, so we
will use that one!

## Procedure

As per [documentation](https://docs.couchbase.com/server/5.5/install/upgrade-strategies.html),
we have several options. XDCR was my first idea, but on a big cluster it may
not work properly (at least on Couchbase `4.5.x`). Graceful failover and
delta recovery seemed too "manual" for me, as I didn't want to manually upgrade
existing instances.

That leaves us with with the Swap Rebalance option.

### Swap rebalance

Swap rebalance is an automatic kind of rebalance that Couchbase does.
It happens when you are adding and removing the same amount of nodes and do
a rebalance.

If that happens, Couchbase will copy the vBuckets from one node being
removed to one node being added, which is faster that what it usually does.

If a swap rebalance fails or is stopped for whatever reason, you can rebalance
again, but it will not be a swap rebalance anymore and thus will take more time.

## Testing it with Docker

We can try this out in a test cluster using Docker! Create a
`docker-compose.yaml` file like the following:

```yaml
version: '2'

services:
  cb1:
    image: couchbase/server:community-4.5.0
    networks:
      cbnet:
        ipv4_address: 172.21.0.10
    ports:
    - 8091:8091
  cb2:
    image: couchbase/server:community-4.5.0
    networks:
      cbnet:
        ipv4_address: 172.21.0.11
  cb3:
    image: couchbase/server:community-4.5.0
    networks:
      cbnet:
        ipv4_address: 172.21.0.12
  cb4:
    image: couchbase/server:community-5.1.1
    networks:
      cbnet:
        ipv4_address: 172.21.0.20
    ports:
    - 9091:8091
  cb5:
    image: couchbase/server:community-5.1.1
    networks:
      cbnet:
        ipv4_address: 172.21.0.21
  cb6:
    image: couchbase/server:community-5.1.1
    networks:
      cbnet:
        ipv4_address: 172.21.0.22

networks:
  cbnet:
    driver: bridge
    ipam:
     config:
      - subnet: 172.21.0.0/16
        gateway: 172.21.0.1
```

If you are using Docker 4 Mac, I recommend increasing the CPU limits:

![](/public/images/cb-upgrade/33b1abdf-dc37-4857-bf07-b10787fc2270.png)

Once that's done, fire the machines up:

```sh
docker-compose up
```

`docker-compose up`

Initialize the Couchbase `4.5.x` cluster:

```sh
docker exec cb_cb1_1 couchbase-cli cluster-init \
  --cluster-username=adm \
  --cluster-password=secret \
  --cluster-port=SAME \
  --cluster-ramsize=256 \
  --cluster-index-ramsize=256 \
  --services=data,index,query,fts
```

Add the other nodes and rebalance:

```sh
docker exec cb_cb1_1 couchbase-cli rebalance \
  -c localhost -u adm -p secret \
  --server-add=172.21.0.11 \
  --server-add=172.21.0.12 \
  --server-add-username=adm \
  --server-add-password=secret \
  --services=data,index,query,fts

docker exec cb_cb1_1 couchbase-cli server-list \
  -c localhost -u adm -p secret
```

Create a bucket and insert some data:

```sh
docker exec cb_cb1_1 couchbase-cli bucket-create \
  -c localhost -u adm -p secret \
  --bucket=customer \
  --bucket-ramsize=100 \
  --enable-index-replica=1 \
  --bucket-replica=2


for i in `seq 1 100`; do
  echo "inserting $i"
  docker exec cb_cb1_1 cbq --quiet=true -output=/dev/null \
    -u adm -p secret -engine=http://127.0.0.1:8091/ \
    --script="INSERT INTO customer (KEY, VALUE) VALUES (\"$i\", { \"id\": $i, \"name\": \"John Doe -$i\"})"
done
```

You can confirm everything by going to the
Couchbase Console.

### Swap rebalance 1 node

Let's do a single node swap rebalance first to see what happens:

```sh
docker exec cb_cb1_1 couchbase-cli rebalance \
  -c localhost -u adm -p secret \
  --server-remove=172.21.0.12 \
  --server-add=172.21.0.20 \
  --server-add-username=adm \
  --server-add-password=secret \
  --services=data,index,query,fts
```

You can see your cluster using the new Couchbase 5 interface by going to its
Couchbase Console.

### Compatibility mode

Since the cluster has nodes with version `4.5.x` and `5.1.x`,
it will run in "4.5 compatibility mode", which is OK:

![](/public/images/cb-upgrade/59cdd3e6-9575-4c90-83db-223670c57d41.png)

While in that mode, we can add and remove nodes from both versions. Once all
nodes are on Couchbase 5.1, you can't add old version nodes anymore.

### Swap rebalance all other nodes

Let's finish the upgrade by swap rebalancing the last nodes:

```sh
docker exec cb_cb4_1 couchbase-cli server-add \
  -c localhost -u adm -p secret \
  --server-add=172.21.0.21 \
  --server-add-username=adm \
  --server-add-password=secret \
  --services=data,index,query,fts

docker exec cb_cb4_1 couchbase-cli server-add \
  -c localhost -u adm -p secret \
  --server-add=172.21.0.22 \
  --server-add-username=adm \
  --server-add-password=secret \
  --services=data,index,query,fts

docker exec cb_cb4_1 couchbase-cli rebalance \
  -c localhost -u adm -p secret \
  --server-remove=172.21.0.10 \
  --server-remove=172.21.0.11
```

Once the rebalance is over, we should have a full 5.1 Couchbase cluster -
without any downtimes!

The compatibility warning should go away as well:

![](/public/images/cb-upgrade/2aba0b3d-8b3d-4ab3-9858-14c783f626cf.png)

### Cleanup

Kill all containers and remove them.

```sh
docker-compose kill
yes | docker-compose rm
```

## How about production

Well, the general idea is the same:

- add a few Couchbase 5 nodes;
- remove a few Couchbase 4 nodes;
- rebalance;
- rince and repeat;

It will take some time, sure, but there are a lot of improvements. The ones
I'm more interested on are related to XDCR, which seems to have a received a
lot of improvements!

---
date: 2015-02-19T00:00:00Z
slug: dump-postgres-table-inserts
title: Dump a PostgreSQL table as insert statements
city: Joinville
tags:
- postgresql
---

> FYI: Like the [previous post]({{< ref "post/2015-02-17-find-non-ascii-chars.md" >}}),
> this is a really quick tip.

This week I'm working closely to the "front-end guy". Not that I don't know
how to front end, but he was helping me.

We are developing an internal tool, that, for this first version, will use
only a few tables of one of our databases.

Doing the "back end" part of it, I created tons of rows in my local database,
and, in order to properly test the front end part, the "front end guy"
needed some data. He could just create a lot of rows in that table. That would
be easy. And boring. And inaccurate.

Dump an entire database is also boring, and usually takes a lot of time.
Besides, he only needs one table.

So, I dumped that one table in the form of inserts. `pg_dump`, with some
parameters, can easily do that:

```console
$ pg_dump \
  -h localhost \
  -p 5432 \
  -U user -W \
  --table="table-name" \
  --data-only \
  --column-inserts \
  database-name > table.sql
```

So, I just had to send him this `table.sql` file somehow, and he had to
execute that SQL file in his database, which can be easily done with the `pg`
command:

```console
$ psql \
  -h localhost \
  -p 5432 \
  -U user \
  database-name \
  -f table.sql
```

That's it. A quick and useful tip that I have used many times and will probably
use many more.

---

**ad**: Work from anywhere with a
[Cloud hosted desktop](https://www.clouddesktoponline.com) and stay connected
to your team with [hosted SharePoint](https://www.cloudappsportal.com) services.
Take a risk free trial today.

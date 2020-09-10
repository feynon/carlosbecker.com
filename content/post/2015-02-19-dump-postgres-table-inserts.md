---
title: "Dump a PostgreSQL table as insert statements"
date: 2015-02-19
draft: false
slug: dump-postgres-table-inserts
city: Joinville
---

> FYI: Like the previous post, this is a really quick tip.

This week I'm working closely to the "front-end guy". Not that I don't know how to front end, but he was helping me.

We are developing an internal tool, that, for this first version, will use only a few tables of one of our databases.

Doing the "back end" part of it, I created tons of rows in my local database, and, in order to properly test the front end part, the "front end guy" needed some data. He could just create a lot of rows in that table. That would be easy. And boring. And inaccurate.

Dump an entire database is also boring, and usually takes a lot of time. 

Besides, he only needs one table.

So, I dumped that one table in the form of inserts. `pg_dump`, with some parameters, can easily do that:

```
$ pg_dump \
  -h localhost \
  -p 5432 \
  -U user -W \
  --table="table-name" \
  --data-only \
  --column-inserts \
  database-name > table.sql
```

So, I just had to send him this `table.sql` file somehow, and he had to execute that SQL file in his database, which can be easily done with the `pg` command:

```
$ psql \
  -h localhost \
  -p 5432 \
  -U user \
  database-name \
  -f table.sql
```

That's it. A quick and useful tip that I have used many times and will probably use many more.
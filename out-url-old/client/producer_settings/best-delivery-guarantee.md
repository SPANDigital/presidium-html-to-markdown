---
title: "Best delivery guarantee¶"
weight: 2
---

**TLDR: These recommended defaults offer good throughput, with the best**
**guarantees we have that your data will make it to the cluster and that**
**your latest writes will be protected (except in the case of a**
**multi-broker failure).**

To make things easy to setup and have durability as a priority we’ve
created a set of defaults you can use. Follow the instructions on
[Default Configuration Injection]({{< ref "client/default_config_injection" >}}) to
get started with these.

Configuration:

- `acks=-1` (ALL)
- `max.block.ms=Long.MAX_VALUE`
- `buffer.memory=<large number of bytes>`
- `retries=Integer.MAX_VALUE`
- `max.in.flight.requests.per.connection=1`

These recommended configurations provide two guarantees

1. **Durability Guarantees**The main parameter that impacts durability is the value of
   `acks` configuration, which controls how many brokers will
   write your record before ack’ing your request. So to achieve the
   highest durability `acks` must be set to all (which is
   `-1`) this means the broker will wait for all in-sync
   replicas before ack’ing your records. Once an ack is received, you know
   that your record has made it to RAM on multiple brokers. It hasn’t been
   written to disk yet, so this means your data is safe only to a point.
   With our topic-level defaults of replication factor = 4 and
   `min.insync.replicas=2`, your availability is protected in
   the presence of two broker failures. However, in the case of combined
   failures (like a full power outage in a data hall), your 2 in-sync
   replicas could both die, and data loss could happen.Please keep in mind that latency will be the trade-off for durability as
   producer latency is the time between producer sending data to the
   ack’ing of the produced records so setting the `acks=-1`
   will add to this time.**How to size your buffer.memory?**The producer will keep buffering incoming requests until your buffer
   fills up. Once full, further produce requests will block, (with the
   settings given). Therefore, if you don’t want the produce requests to
   block, then your buffer needs to be big enough to store requests until
   any transitory issue is resolved. How much and how long is acceptable is
   generally use-case specific, but being able to handle a few minutes to a
   few hours is the norm. It is also not uncommon to set
   `max.block.ms` to the max acceptable time your use-case
   allows. When there is a network blip the produce will buffer your
   message until connectivity recovers. When a broker restarts the producer
   will buffer your message until it can send messages. When a broker fails
   your previously ack’ed messages are safe. When multiple brokers fail
   some previously ack’ed messages will be lost. (See above for details of
   how many broker failures your data is resilient to).
2. **Ordering Guarantees**By increasing your `retries` the producer doesn’t give up
   on the first issue it encounters. Though be warned, without reducing
   `max.in.flight.requests.per.connection` to one your messages
   can, and will, get sent out of order within a single partition. This may
   or may not be an issue for your use case. Of course, dropping
   `max.in.flight.requests.per.connection` also affects
   throughput.


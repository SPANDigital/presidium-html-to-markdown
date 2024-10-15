---
title: "Producer Properties¶"
weight: 7
---

You are **strongly urged** to set the following property for producers:

```
acks=-1

```

This minimizes the chance of data loss as well as the need to perform
unclean leader elections on the brokers. If your use case warrants, you
should also consider setting the following producer properties that also
serve to minimize data loss:

```
retries=2147483647
max.in.flight.requests.per.connection=1
max.block.ms=2147483647

```


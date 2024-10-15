---
title: "What’s happening?¶"
weight: 2
---

As part of our work to unify metric reporting for all of our clusters we are updating our metric reporting system for
Mosaic. This involves changing our broker metrics in small but important ways.

This new metric reporting aims to:

- Unify metrics & dashboards across “ACI” and “AMP” clusters
- Fix stale metrics being reported in some circumstances
- Remove legacy terms such as `piekafka`
- Better align with metric naming conventions in Open Source ecosystem
- Improve metric consistency & quality

Additionally, we are simplifying the way we share consumer lag metrics going forward. These metrics will now be shared
on a per-[Internal Link: customer group](../concepts.html#customer-group) basis, rather than
per-[Internal Link: client identity](../concepts.html#client-identities).

**Impacted:**

- Kafka broker metrics for “ACI” clusters: metrics in the `kafka-broker-*` Mosaic namespaces
- Consumer lag metrics: metrics starting with `kafka_consumergroup_` or in the `kafka-consumer-lag-*` namespaces

**NOT impacted:**

- Synthetic monitoring metrics: metrics in the `synthetic-monitor-*` namespaces
- Metrics for the ex-AMP clusters (these have already been updated to match these conventions), including AMP lag
  metrics


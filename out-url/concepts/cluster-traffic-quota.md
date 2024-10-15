---
title: "Cluster traffic quota¶"
weight: 6
---

ACI Kafka provides a high traffic quota for all multi-tenant clusters.
The quota is enforced on all [Internal Link: client identities](#client-identities) on a per-broker basis to safeguard multi-tenant clusters.
By default, the quota is 35% of the broker bandwidth capacity.

For example if an ACI Kafka cluster has a default produce and consume quota of 315 MB/s and 315 MB/s, respectively, then each client identity will be throttled on a broker when it exceeds 315 MB/s _produce_ or 315 MB/s _consume_ from one broker.

To learn more about the Apache Kafka quota management design, see [External Link: KIP-13 - Quotas](https://cwiki.apache.org/confluence/display/KAFKA/KIP-13+-+Quotas).


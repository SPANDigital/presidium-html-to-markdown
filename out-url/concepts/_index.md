---
title: "Concepts¶"
weight: 1
---

This section aims to provide basic information about _customer groups_, _client identities_, _namespaces_, _topics_, _topic accesses_, and _topic quotas_ related to Apple Cloud Infrastructure (ACI) Kafka services and offerings.

ACI Kafka provides an abstraction of clusters for the following reasons:

- To provide you with topics and client identities in a managed environment.
  We take care of the size of the clusters and the configuration of the underlying brokers.
- Abstraction enables ACI Kafka to make operational changes without requiring you to modify your configurations.
- Provides multi-tenant clusters, allowing ACI Kafka to better utilize existing capacity.


---
title: "Notable changes¶"
weight: 2
---

The ACI Kafka 0.14 release includes new features available in Apache
Kafka 1.1.

**Summary of tent-pole changes**

- Apache Kafka 1.1 includes a number of scalability improvements for
  supporting clusters with large number of partitions
- Consumer lag metric name has been corrected to roll up partition
  ids - KIP-225 authored and implemented by Charly from ACI Kafka team
- Kaffe configuration injection is now enabled by default to ease
  customer configuration and prepare for the envelope deprecation
  transition
- ACI Kafka Portal now supports full management of topics and clients,
  including create, update, and delete actions

**ACI Kafka**

- Kaffe configuration injection is now enabled by default
- Kaffe heart beating has been removed to increase scalability of use
  cases with large number of connected clients
- Inspector now provides Kafka for All VIPs on the cluster page, and
  Zookeeper connect string has been removed for security reasons
- Kafka Portal now supports deletion of topics, clients, namespaces,
  and customer groups
- Kafka Portal now supports updating contact information in customer
  groups

**Apache Kafka**

- Significant Controller improvements (much faster and session
  expiration edge cases fixed)
- Data balancing across log directories (JBOD)
- More efficient replication when the number of partitions is large
- Dynamic Broker Configs
- KIP-145 - Expose Record Headers in Kafka Connect
- KIP-223 - Add per-topic min lead and per-partition lead metrics to
  KafkaConsumer
- KIP-225 - Use tags for consumer “records.lag” metrics
- KIP-227 - Introduce Incremental FetchRequests to Increase Partition
  Scalability
- Complete [External Link: Apache Kafka 1.1.1 Release
  Notes](https://archive.apache.org/dist/kafka/1.1.1/RELEASE_NOTES.html)
  and [External Link: upgrade
  notes](https://kafka.apache.org/11/documentation.html#upgrade)


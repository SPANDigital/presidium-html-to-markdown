---
title: "Past¶"
weight: 4
---

Below are the developments delivered in prior years.

### 2020 [¶](#id1 "Link to this heading")

- **SilkRoad Build Out** Delivering ACI Kafka for use in Silk Road, in
  preparation for Golden Gate.
- **Mirrors for SR** Allow replication of topics from SilkRoad to
  AODCs - limited availability
- **SLOs** Integrate with [slo.apple.com]({{%baseurl%}}//files/slo.apple.com) for
  Kafka to provide visibility of our platform reliability
- **Improved Data Rebalancing** added the ability to rebalance data
  between brokers to improve disk utilization

### 2019 [¶](#id2 "Link to this heading")

- JDK11 Compatibility
- Public API available to all for self-service entity management
- Integration with Topic Key Management Service Front End service
  (Iris) to eliminate C\* client dependency
- APIs and plan to facilitate the deprecation of the PIE Envelope in
  favor of Kafka Message Headers
- Utilization reports for storage, CPU, and other Kafka and system
  resources
- Kubernetes support to manage Kafka cluster running on K8s for SR and
  GG

### 2018 [¶](#id3 "Link to this heading")

- Production cluster available for Phase 1 and 2
- Entity management is fully done with Portal and is backed by PIE
  Workflow.
- Users can use portal for most actions necessary
- Published SLOs about service reliability and reviewing them monthly.
- Kafka ACLs enforcement with support for prefixed principle and
  resource identifiers
- Configuration injection with Kaffe to enable easier setup for user
- Kafka for All to allow polyglot use of ACI Kafka
- Upgrade to Kafka 1.1
- Inspector release to allow customers to troubleshoot problems and
  see their consumer lag.
- Quota accounting in place which will enable us to enforce it in the
  future.
- Limited access preview to mirroring solution based on Kafka connect.

### 2017 [¶](#id4 "Link to this heading")

- First Shared QA cluster available for Phase 1 customers in June
- Begin placing production hardware orders for dedicated high-volume
  clusters
- Early Availability for the ACI Kafka client
- New Shared QA clusters added for Phase 2 customers
- End-to-end encryption available to customers
- First Shared Production cluster available for Phase 1 customers
- First dedicated production clusters live

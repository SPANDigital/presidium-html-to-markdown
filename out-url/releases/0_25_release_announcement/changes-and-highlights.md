---
title: "Changes and highlights¶"
weight: 3
---

We’d like to highlight the following improvements and bug fixes since 0.24 (Kafka 3.2):

- Common client changes:
  - [External Link: KAFKA-13793](https://issues.apache.org/jira/browse/KAFKA-13793) \- Add validators for serialization and deserialization related configuration
    - Tighter checks for correct configuration
  - [External Link: KAFKA-13848](https://issues.apache.org/jira/browse/KAFKA-13848) \- Clients remain connected after SASL re-authentication fails
  - [External Link: KAFKA-13911](https://issues.apache.org/jira/browse/KAFKA-13911) \- Rate is calculated as `NaN` for minimum config values
- Affecting the Producer only:
  - [External Link: KAFKA-10888](https://issues.apache.org/jira/browse/KAFKA-10888) \- **Sticky partition leads to uneven product msg, resulting in abnormal delays in some partitions**
    - Producers since 2.4 (kafka-pie-client since 0.17) can create a partition skew when a broker is slow. See [External Link: KIP-794](https://cwiki.apache.org/confluence/display/KAFKA/KIP-794%3A+Strictly+Uniform+Sticky+Partitioner) for more details.
  - [External Link: KAFKA-13967](https://issues.apache.org/jira/browse/KAFKA-13967) \- Guarantees for producer callbacks on transaction commit should be documented
  - [External Link: KAFKA-13834](https://issues.apache.org/jira/browse/KAFKA-13834) \- batch drain for nodes might have starving issue
  - [External Link: KAFKA-14156](https://issues.apache.org/jira/browse/KAFKA-14156) \- Built-in partitioner may create suboptimal batches with large linger.ms
  - [External Link: KAFKA-14055](https://issues.apache.org/jira/browse/KAFKA-14055) \- Transaction markers may be lost during cleaning if data keys conflict with marker keys
  - [External Link: KAFKA-14020](https://issues.apache.org/jira/browse/KAFKA-14020) \- Performance regression in Producer
- Affecting the Consumer only
  - [External Link: KAFKA-13917](https://issues.apache.org/jira/browse/KAFKA-13917) \- **Avoid calling `lookupCoordinator()` in tight loop**
    - Fixes a bug that caused a spike in CPU on the Consumer when one of the brokers crashes. Affects Consumers since at least 2.8 (kafka-pie-client 0.21).
  - [External Link: KAFKA-14013](https://issues.apache.org/jira/browse/KAFKA-14013) \- Limit the length of the `reason` field sent on the wire
  - [External Link: KAFKA-13778](https://issues.apache.org/jira/browse/KAFKA-13778) \- Fetch from follower should never run the preferred read replica selection
  - [External Link: KAFKA-13777](https://issues.apache.org/jira/browse/KAFKA-13777) \- Fix `FetchResponse#responseData`: Assignment of lazy-initialized members should be the last step with double-checked locking
  - [External Link: KAFKA-13791](https://issues.apache.org/jira/browse/KAFKA-13791) \- Fix `FetchResponse#fetchData` and `forgottenTopics`: Assignment of lazy-initialized members should be the last step with double-checked locking
  - [External Link: KAFKA-14024](https://issues.apache.org/jira/browse/KAFKA-14024) \- Consumer stuck during cooperative rebalance for Commit offset in `onJoinPrepare`
  - [External Link: KAFKA-14201](https://issues.apache.org/jira/browse/KAFKA-14201) \- Consumer should not send group instance ID if committing with empty member ID
  - [External Link: KAFKA-14208](https://issues.apache.org/jira/browse/KAFKA-14208) \- `KafkaConsumer#commitAsync` throws unexpected `WakeupException`
  - [External Link: KAFKA-14194](https://issues.apache.org/jira/browse/KAFKA-14194) \- NPE in `Cluster.nodeIfOnline`
  - [External Link: KAFKA-13998](https://issues.apache.org/jira/browse/KAFKA-13998) \- `JoinGroupRequestData.reason` can be too large
  - [External Link: KAFKA-14078](https://issues.apache.org/jira/browse/KAFKA-14078) \- Replica fetches to follower should return `NOT_LEADER` error
  - [External Link: KAFKA-13891](https://issues.apache.org/jira/browse/KAFKA-13891) \- sync group failed with `rebalanceInProgress` error cause rebalance many rounds in cooperative

For a full list of changes please consult the [External Link: release notes for Kafka 3.3.0](https://archive.apache.org/dist/kafka/3.3.0/RELEASE_NOTES.html) and [External Link: 3.3.1](https://archive.apache.org/dist/kafka/3.3.1/RELEASE_NOTES.html).
Or, for a general overview on what’s new in Kafka 3.3 you might be interested in [External Link: this post](https://blogs.apache.org/kafka/entry/what-rsquo-s-new-in) in the Apache Kafka
blog.

For more information on the ACI Kafka service offering, including
features, caveats, and project road map, visit our [Internal Link: overview section](../overview.html).

You can [Internal Link: reach out to the ACI Kafka team](../contact.html) with
questions on this release or general inquiries related to ACI Kafka.

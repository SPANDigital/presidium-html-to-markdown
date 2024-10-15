---
title: "Changes and highlights¶"
weight: 3
---

We’d like to highlight the following improvements and bug fixes since 0.25 (Kafka 3.3):

- Common client changes:
  - [External Link: KAFKA-10360](https://issues.apache.org/jira/browse/KAFKA-10360) \- Disabling JmxReporter registration
- Affecting the Producer only:
  - [External Link: KAFKA-4852](https://issues.apache.org/jira/browse/KAFKA-4852) \- ByteBufferSerializer not compatible with offsets
  - [External Link: KAFKA-14097](https://issues.apache.org/jira/browse/KAFKA-14097) \- Separate configuration for producer ID expiry
  - [External Link: KAFKA-13559](https://issues.apache.org/jira/browse/KAFKA-13559) \- The broker’s ProduceResponse may be delayed for 300ms
- Affecting the Consumer only
  - [External Link: KAFKA-13715](https://issues.apache.org/jira/browse/KAFKA-13715) \- Add “generation” field into consumer protocol

For a full list of changes please consult the [External Link: release notes for Kafka 3.4.0](https://archive.apache.org/dist/kafka/3.4.0/RELEASE_NOTES.html).
Or, for a general overview on what’s new in Kafka 3.4 you might be interested in [External Link: this post](https://blogs.apache.org/kafka/entry/what-s-new-in-apache9) in the Apache Kafka
blog.

For more information on the ACI Kafka service offering, including
features, caveats, and project road map, visit our [Internal Link: overview section](../overview.html).

You can [Internal Link: reach out to the ACI Kafka team](../contact.html) with
questions on this release or general inquiries related to ACI Kafka.

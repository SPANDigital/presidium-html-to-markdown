---
title: "What’s new in this release?¶"
weight: 3
---

This release aims to provide relevant changes for ASE Kafka customers,
including Apache Kafka development that focuses on:

- **(Early Access) JBOD support in KRaft** – includes changes to allow using
  multiple disks in Kraft mode.
- **(Early Access) Next Generation of Consumer Rebalance Protocol** \- includes
  changes that simplifies the consumer group protocol and moves the complexity
  away from the client to the group coordinator in the broker.

### PieEnvelope[Internal Link: ¶](\#pieenvelope)

Warning

Consuming and deserializing messages in the old envelope format is no longer
supported.

We deprecated the envelope format in 2019. In 2021, we stopped
supporting messages using the envelope format in version 0.20.

If you are still producing records in the envelope format using a producer
older than version 0.20:

- Ensure your topics no longer contain messages in the old format before upgrading consumers to the 0.29 version.
  - Record retention is one week unless explicitly configured otherwise on the topic.
- Set the configuration in those producers as shown below.

```
pie.queue.serializer.use.payload.envelope=false
pie.queue.envelope.serializer.use.payload.envelope=false

```

Refer to the documentation [Internal Link: `PieEnvelope` deprecation](../client/envelope.html) for details.

### JBOD support in Kraft[Internal Link: ¶](\#jbod-support-in-kraft)

Kafka replaces Zookeeper with a custom Raft implementation.
Over the past four years, most of the development in Apache Kafka
has shifted away from features and bug-fixing to
rewriting and migrating into the new system.

This release introduces early access support for running Kafka clusters
with multiple disks per broker (JBOD) in KRaft mode. To this date, this has been Apple’s
biggest contribution upstream and is crucial for ASE Kafka’s
service offering, as we use JBOD in our clusters. Our contributions spanned
from proposing the KIP, to implementing and testing it. We continue to test and
experiment with this feature, to prepare for conversion to KRaft mode.
Please refer to [External Link: KIP-858](https://cwiki.apache.org/confluence/display/KAFKA/KIP-858%3A+Handle+JBOD+broker+disk+failure+in+KRaft) for further information.

**This work does not impact existing ASE Kafka use**
and we will continue to communicate relevant changes in the future.

### Next Generation of Consumer Rebalance Protocol[Internal Link: ¶](\#next-generation-of-consumer-rebalance-protocol)

Kafka consumer relies on a group-wide synchronization barrier, which means a
single misbehaving consumer can potentially cause disruption due to rebalancing
of the entire consumer group. In addition, the existing protocol has gotten too
complex over the years.

This change introduces early access support for a new group membership and
rebalance protocol which aims to be truly incremental and cooperative. Ideally,
it means a consumer should not be impacted by a rebalance if its assignment
hasn’t changed. Please refer to [External Link: KIP-848](https://cwiki.apache.org/confluence/display/KAFKA/KIP-848%3A+The+Next+Generation+of+the+Consumer+Rebalance+Protocol) for further information.

### Other highlights[Internal Link: ¶](\#other-highlights)

- [External Link: KIP-714](https://cwiki.apache.org/confluence/display/KAFKA/KIP-714%3A+Client+metrics+and+observability) Provides more visibility to operators about clients connecting to a Kafka cluster
  by adding support for client-level metrics via a standardized OpenTelemetry interface.
- [External Link: KIP-951](https://cwiki.apache.org/confluence/display/KAFKA/KIP-951%3A+Leader+discovery+optimisations+for+the+client) Reduces the time for a client to find partition leader.
- [External Link: KIP-975](https://cwiki.apache.org/confluence/display/KAFKA/KIP-975%3A+Docker+Image+for+Apache+Kafka) Introduces official Docker images based on upstream Apache Kafka.
- [External Link: KIP-1013](https://cwiki.apache.org/confluence/pages/viewpage.action?pageId=284789510) Marks Java 11 as deprecated for Kafka brokers. Kafka clients and Kafka Connect continue to support Java 11.

For a complete list of changes, please consult the [External Link: release notes for Kafka 3.7](https://downloads.apache.org/kafka/3.7.0/RELEASE_NOTES.html).

For a general overview of what’s new in Kafka 3.7, read [External Link: the 3.7 announcement](https://kafka.apache.org/blog#apache_kafka_370_release_announcement) in the Apache Kafka
blog.

* * *

Visit the [Internal Link: overview section](../overview.html) for more information about
our ASE Kafka service offering, including features, concepts, and project road map.

For questions about this release or general inquiries related to ASE Kafka, [Internal Link: contact the ASE Kafka team](../contact.html).

---
title: "What’s new in this release?¶"
weight: 3
---

This release aims to provide relevant changes for ACI Kafka customers,
including Apache Kafka development that focuses on:

- **KRaft** – includes non-user-impacting changes relating to KRaft.
- **Tiered Storage** – includes changes to provide optional support for tiered storage.
- Refactoring the project structure.

### PieEnvelope [¶](#pieenvelope "Link to this heading")

Warning

Consuming and deserializing messages in the old envelope format is no longer
supported.

We deprecated the envelope format in 2019. In 2021, we stopped
supporting messages using the envelope format in version 0.20.

If you are still producing records in the envelope format using a producer
older than version 0.20:

- Set the configuration in those producers as shown below.
- Ensure your Topics no longer contain messages in the old format before upgrading consumers to the 0.28 version.
  - Record retention is one week unless explicitly configured otherwise on the Topic.

```
pie.queue.serializer.use.payload.envelope=false
pie.queue.envelope.serializer.use.payload.envelope=false

```

Refer to the documentation [`PieEnvelope` deprecation]({{< ref "client/envelope" >}}) for details.

### KRaft [¶](#kraft "Link to this heading")

Kafka replaces Zookeeper with a custom Raft implementation.
Over the past four years, most of the development in Apache Kafka
has shifted away from features and bug-fixing to
rewriting and migrating into the new system.

ACI Kafka has been following these changes. We’ve start testing and
preparing for the conversion to KRaft mode Kafka, including experimentally running a KRaft cluster
in our QA environment and engaging with the open-source project to close
the gaps on bugs and missing features.

We’re also preparing our services to support Apache Kafka’s
transition to KRaft mode. **This work does not impact existing ACI Kafka use**
and we will continue to communicate relevant changes in the future.

### Tiered storage [¶](#tiered-storage "Link to this heading")

Tiered storage is a popular and anticipated feature that allows
a pluggable second tier of storage to extend retention capacity.
Kafka version 3.6 introduces these features as early access.
At this time, however, we DO NOT recommend using it in production.

While we recognize this is an appealing feature, we are not
currently offering it with our services due to the
high-priority KRaft work and limited resources.

If this or any other feature is important for your use case, please
[contact]({{< ref "contact" >}}) us.

### Other highlights [¶](#other-highlights "Link to this heading")

- [KIP-890](https://cwiki.apache.org/confluence/display/KAFKA/KIP-890%3A+Transactions+Server-Side+Defense) transactions in Kafka have been a prominent source of problems. This
  change aims to prevent hanging transaction issues.
- [KIP-863]({{%baseurl%}}/https://cwiki.apache.org/confluence/pages/viewpage.action?pageId=225152035) improves memory allocation and performance during deserialization.

For a complete list of changes, please consult the [release notes for Kafka 3.6](https://downloads.apache.org/kafka/3.6.0/RELEASE_NOTES.html).

For a general overview of what’s new in Kafka 3.6, read [this post](https://kafka.apache.org/blog#apache_kafka_360_release_announcement) in the Apache Kafka
blog.

* * *

Visit the [overview section]({{< ref "overview" >}}) for more information about
our ACI Kafka service offering, including features, caveats, and project road map.

For questions about this release or general inquiries related to ACI Kafka, [contact the ACI Kafka team]({{< ref "contact" >}}).

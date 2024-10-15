---
title: "Upgrade Instructions¶"
weight: 2
---

The ACI Kafka 0.21 client is available on
[Artifactory](https://artifacts.apple.com/pie-release-local/com/apple/pie/queue/kafka-pie-client).
If you are using an older version of the ACI Kafka client, please upgrade to the latest version.

**Removal of deprecated `PieEnvelope` and related objects**

Support for consuming messages in the envelope format is still available.
This continues to guarantee interoperability with older clients.
However, producing messages in this old format is no longer possible.

The classes `PieEnvelope`, `Envelope[Des,S]erializer` and `Envelope[Des,S]erializerHelper` are no longer available.
Instead of producing or consuming `PieEnvelope<T>` messages, the client should instead just produce/consume `T`.

Please make sure you have fully migrated away from using `PieEnvelope` and the envelope wire format.
For more details, please refer to the documentation on the [`PieEnvelope` deprecation]({{< ref "client/envelope" >}}).

**Notable upstream changes**

- Kafka is replacing its dependency on ZooKeeper with its own metadata quorum. This version includes early,
  non production-ready access to the new feature. Currently there are no expected changes for applications using ACI Kafka.
- Topic IDs are slowly being introduced as part of [KIP-516](https://cwiki.apache.org/confluence/display/KAFKA/KIP-516%3A+Topic+Identifiers). At this point, there are no expected changes for applications using ACI Kafka.
- Bugfix [KAFKA-10470](https://issues.apache.org/jira/browse/KAFKA-10470) \- zstd decompression with small batches is slow and causes excessive GC
- Bugfix [KAFKA-12152](https://issues.apache.org/jira/browse/KAFKA-12152) \- Idempotent Producer does not reset the sequence number of partitions without in-flight batches
- Bugfix [KAFKA-12193](https://issues.apache.org/jira/browse/KAFKA-12193) \- Re-resolve IPs when a client is disconnected

For additional changes and more details please consult the notable changes for version
[2.8]({{%baseurl%}}/https://kafka.apache.org/documentation.html#upgrade_280_notable) in the Apache Kafka documentation or the
[release notes for Kafka 2.8]({{%baseurl%}}/https://kafka.apache.org/downloads.html#2.8.0).

For more information on the ACI Kafka service offering, including
features, caveats, and project road map, visit our [overview\
section]({{< ref "overview" >}}).

You can [reach out to the ACI Kafka team]({{< ref "contact" >}}) with
questions on this release or general inquiries related to ACI Kafka.

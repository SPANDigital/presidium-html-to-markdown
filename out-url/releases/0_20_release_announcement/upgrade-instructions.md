---
title: "Upgrade Instructions¶"
weight: 2
---

The ACI Kafka 0.20 client is available on
[External Link: Artifactory](https://artifacts.apple.com/pie-release-local/com/apple/pie/queue/kafka-pie-client)
and can be configured as a [Internal Link: Gradle
Dependency](../client/quick_start.html#gradle-dependency).
If you are using older versions of ACI Kafka, please upgrade to the
latest version.

**Removal of deprecated `PieEnvelope` methods**

The methods to use the `PieEnvelope` directly [External Link: have been removed](https://github.pie.apple.com/pie/kafka-tools/commit/5a45d094f).
However, the envelope format is still supported on the wire.
For more details, please refer to the documentation on the [Internal Link: `PieEnvelope` deprecation](../client/envelope.html).

**Other changes**

- Bugfix [External Link: KAFKA-12152](https://issues.apache.org/jira/browse/KAFKA-12152) \- Idempotent Producer does not reset the sequence number of partitions without in-flight batches
- Bugfix [External Link: KAFKA-12193](https://issues.apache.org/jira/browse/KAFKA-12193) \- Re-resolve IPs when a client is disconnected
- The `ConsoleConsumer` will now print record offsets and headers, see [External Link: KIP-431](https://cwiki.apache.org/confluence/display/KAFKA/KIP-431%3A+Support+of+printing+additional+ConsumerRecord+fields+in+DefaultMessageFormatter)

For additional changes and more details please consult the notable changes for version
[External Link: 2.7](https://kafka.apache.org/documentation.html#upgrade_270_notable) in the Apache Kafka documentation.

For more information on the ACI Kafka service offering, including
features, caveats, and project road map, visit our [Internal Link: overview
section](../overview.html).

You can [Internal Link: reach out to the ACI Kafka team](../contact.html) with
questions on this release or general inquiries related to ACI Kafka.

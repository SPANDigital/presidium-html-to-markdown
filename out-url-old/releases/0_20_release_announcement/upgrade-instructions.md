---
title: "Upgrade Instructions¶"
weight: 2
---

The ACI Kafka 0.20 client is available on
[Artifactory](https://artifacts.apple.com/pie-release-local/com/apple/pie/queue/kafka-pie-client)
and can be configured as a [Gradle\
Dependency]({{%baseurl%}}/../client/quick_start.html#gradle-dependency).
If you are using older versions of ACI Kafka, please upgrade to the
latest version.

**Removal of deprecated `PieEnvelope` methods**

The methods to use the `PieEnvelope` directly [have been removed](https://github.pie.apple.com/pie/kafka-tools/commit/5a45d094f).
However, the envelope format is still supported on the wire.
For more details, please refer to the documentation on the [`PieEnvelope` deprecation]({{< ref "client/envelope" >}}).

**Other changes**

- Bugfix [KAFKA-12152](https://issues.apache.org/jira/browse/KAFKA-12152) \- Idempotent Producer does not reset the sequence number of partitions without in-flight batches
- Bugfix [KAFKA-12193](https://issues.apache.org/jira/browse/KAFKA-12193) \- Re-resolve IPs when a client is disconnected
- The `ConsoleConsumer` will now print record offsets and headers, see [KIP-431](https://cwiki.apache.org/confluence/display/KAFKA/KIP-431%3A+Support+of+printing+additional+ConsumerRecord+fields+in+DefaultMessageFormatter)

For additional changes and more details please consult the notable changes for version
[2.7]({{%baseurl%}}/https://kafka.apache.org/documentation.html#upgrade_270_notable) in the Apache Kafka documentation.

For more information on the ACI Kafka service offering, including
features, caveats, and project road map, visit our [overview\
section]({{< ref "overview" >}}).

You can [reach out to the ACI Kafka team]({{< ref "contact" >}}) with
questions on this release or general inquiries related to ACI Kafka.

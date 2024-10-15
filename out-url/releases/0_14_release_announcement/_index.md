---
title: "ACI Kafka 0.14 Release Announcement¶"
weight: 1
---

The ACI Kafka team is excited to announce the release of ACI Kafka 0.14.
This release is based on Apache Kafka 1.1 and contains a number of new
features, fixes, and improvements.

The ACI Kafka 0.14 client is [External Link: available on
Artifactory](https://artifacts.apple.com/pie-release-local/com/apple/pie/queue/kafka-pie-client/0.14.3/kafka-pie-client-0.14.3.jar)
and can be configured as a [Internal Link: Gradle
Dependency](../client/quick_start.html#gradle-dependency).
If you are using older versions of ACI Kafka 0.14, please upgrade to the
latest version.

Server binaries are now running ACI Kafka 0.14 and broker protocol
versions are now configured to 1.1 in all environments.

In the next ACI Kafka release 0.15, we will be deprecating the
[External Link: createPieProducer](https://github.pie.apple.com/pie/kafka-tools/blob/develop/kafka-pie-client/src/main/java/com/apple/pie/queue/kafka/client/KafkaClientUtils.java#L50)
and
[External Link: createPieConsumer](https://github.pie.apple.com/pie/kafka-tools/blob/develop/kafka-pie-client/src/main/java/com/apple/pie/queue/kafka/client/KafkaClientUtils.java#L37)
methods in KafkaClientUtils in favor of APIs that do not expose the
PieEnvelope as a value type. More information on PIE Envelope
Deprecation can be found in
[External Link: QTIP-28](https://github.pie.apple.com/pie/kafka-tools/blob/develop/improvement-proposals/028-envelope-deprecation.md).

The 0.15 release will also include Iris, a new front-end service for e2e
crypto key distribution. Iris removes the C\* dependency for the ACI
Kafka client and more information can be found in
[External Link: QTIP-29](https://github.pie.apple.com/pie/kafka-tools/blob/develop/improvement-proposals/029-iris-service.md).

For the latest documentation, including an overview of ACI Kafka and the
ACI Kafka Client, please see visit our
[External Link: documentation](https://docs.aci.apple.com/kafka).


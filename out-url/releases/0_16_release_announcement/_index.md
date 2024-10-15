---
title: "ACI Kafka 0.16 release announcement¶"
weight: 1
---

The ACI Kafka team is excited to announce the release of ACI Kafka 0.16.
This is the first ACI Kafka release based on Apache Kafka 2.1.1, a major
upgrade that contains a number of improvements and new features. Please
see below for notable changes.

For the latest documentation, including an overview of ACI Kafka and the
ACI Kafka Client, please see our
[External Link: documentation](https://docs.aci.apple.com/kafka).

**Upgrade Instructions**

The ACI Kafka 0.16.1 client is available on
[External Link: Artifactory](https://artifacts.apple.com/pie-release-local/com/apple/pie/queue/kafka-pie-client/0.16.1)
and can be configured as a [Internal Link: Gradle
Dependency](../client/quick_start.html#gradle-dependency).
If you are using older versions of ACI Kafka, please upgrade to the
latest version, 0.16.1.

**Notable changes**

For clients not using Kaffe configuration injection, the client
configuration property ssl.endpoint.identification.algorithm needs to be
set to the empty string to enable access to the ACI Kafka clusters. This
is handled automatically when clients use Kaffe configuration injection.

The default value for request.timeout.ms has been changed from 5 minutes
to 30 seconds.

Kafka now support zstandard compression. See
[External Link: KIP-110](https://cwiki.apache.org/confluence/display/KAFKA/KIP-110%3A+Add+Codec+for+ZStandard+Compression))
for more information.


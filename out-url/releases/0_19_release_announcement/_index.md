---
title: "ACI Kafka 0.19 release announcement¶"
weight: 1
---

The ACI Kafka team is excited to announce the release of ACI Kafka 0.19.
This release is based on our [External Link: internal
build](https://github.pie.apple.com/pie/apache-kafka) of Apache Kafka
2.6.1 and contains a number of new features, fixes, and improvements.

For the latest documentation, including an overview of ACI Kafka and the
ACI Kafka Client, please see our
[External Link: documentation](https://docs.aci.apple.com/kafka).

**Upgrade Instructions**

The ACI Kafka 0.19 client is available on
[External Link: Artifactory](https://artifacts.apple.com/pie-release-local/com/apple/pie/queue/kafka-pie-client)
and can be configured as a [Internal Link: Gradle
Dependency](../client/quick_start.html#gradle-dependency).
If you are using older versions of ACI Kafka, please upgrade to the
latest version.

**Scala version upgrade**

This release uses Scala version 2.13.

**Enabling TLSv1.3**

TLSv1.3 is the default for Kafka 2.6.1. The client and server will
negotiate TLSv1.3 if both support it and fallback to TLSv1.2 otherwise.
See
[External Link: KIP-573](https://cwiki.apache.org/confluence/display/KAFKA/KIP-573%3A+Enable+TLSv1.3+by+default)
for more details.

ACI Kafka brokers support TLSv1.3 in addition to TLSv1.2.
Kaffe no longer injects the default `ssl.enabled.protocols=TLSv1.2`.
Instead, the default value is `ssl.enabled.protocols=TLSv1.3,TLSv1.2`.
ACI Kafka brokers accept only TLSv1.3 and TLSv1.2.

**DNS Lookup**

The default value for the `client.dns.lookup` configuration has been
changed from default to `use_all_dns_ips`. If a hostname resolves to
multiple IP addresses, clients and brokers will now attempt to connect
to each IP in sequence until the connection is successfully established.
See
[External Link: KIP-602](https://cwiki.apache.org/confluence/display/KAFKA/KIP-602%3A+Change+default+value+for+client.dns.lookup)
for more details.

If you are not setting this manually setting this property, Kaffe is
already injecting this new default even if you are using an older client
version.

**Dropped support for live objects in configuration**

It is no longer possible to use live objects in configuration, as [External Link: one
of the private
patches](https://github.pie.apple.com/pie/apache-kafka/commit/6b1471e788b2577059bcc67753f9435e6fd74905)
maintained in our internal build has been dropped.

**Other notable changes**

- When `RebalanceProtocol#COOPERATIVE` is used, `Consumer#poll` can
  still return data while it is in the middle of a rebalance for those
  partitions still owned by the consumer; in addition
  `Consumer#commitSync` now may throw a non-fatal
  `RebalanceInProgressException` to notify users of such an event, in
  order to distinguish from the fatal `CommitFailedException` and
  allow users to complete the ongoing rebalance and then reattempt
  committing offsets for those still-owned partitions.
- Added a new Serde type Void to represent null keys or null values
  from input topic.
- Improved exactly-once semantics by adding a pending offset fencing
  mechanism and stronger transactional commit consistency check, which
  greatly simplifies the implementation of a scalable exactly-once
  application. Check out
  [External Link: KIP-447](https://cwiki.apache.org/confluence/display/KAFKA/KIP-447%3A+Producer+scalability+for+exactly+once+semantics)
  for the full details.

For additional changes and more details please consult the notable
changes for
[External Link: 2.5](https://kafka.apache.org/documentation.html#upgrade_250_notable) and
[External Link: 2.6](https://kafka.apache.org/documentation.html#upgrade_260_notable) in
the Apache Kafka documentation.

For more information on the ACI Kafka service offering, including
features, caveats, and project road map, visit our [Internal Link: overview
section](../overview.html).

You can [Internal Link: reach out to the ACI Kafka team](../contact.html) with
questions on this release or general inquiries related to ACI Kafka.

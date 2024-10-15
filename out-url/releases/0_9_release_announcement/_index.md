---
title: "ACI Kafka 0.9 Release Announcement¶"
weight: 1
---

The ACI Kafka team is excited to announce the release of ACI Kafka 0.9.
This release is based on Apache Kafka 1.0.1 and contains a number of new
features, fixes, and improvements. We’ve also made invisible
operational changes that allow us to deliver faster, like deploying
customer-facing services (e.g. [Internal Link: Kaffe](../client/kaffe.html) and [Internal Link: Kafka
Portal](../self-service/portal.html)) on
[External Link: Compute](https://docs.aci.apple.com/compute/) once or more a week.

The ACI Kafka 0.9.4 client is [External Link: available on
Artifactory](https://artifacts.apple.com/pie-release-local/com/apple/pie/queue/kafka-pie-client/0.9.4/kafka-pie-client-0.9.4.jar)
and can be configured as a [Internal Link: Gradle
Dependency](../client/quick_start.html#gradle-dependency).
If you are using older versions of ACI Kafka 0.9, please upgrade to
0.9.4.

Server binaries are now running ACI Kafka 0.9 and broker protocol
versions are now configured to `1.0` in all environments, but log format
upgrades have not completed in production. Some new features in Apache
Kafka 1.0 and 0.11 (e.g. Exactly Once Semantics (EOS)) will not be
available in IF-PROD Namespaces until this roll out completes this month
(June 2018).

In the next ACI Kafka release, we will fast-forward to version 0.12 to
reduce ambiguity with Apache Kafka versions.

For the latest documentation, including an overview of ACI Kafka and the
ACI Kafka Client, please see visit our
[External Link: documentation](https://docs.aci.apple.com/kafka).


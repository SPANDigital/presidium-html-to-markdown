---
title: "Upgrade Instructions¶"
weight: 2
---

The ACI Kafka 0.15.8 client is available on
[External Link: Artifactory](https://artifacts.apple.com/pie-release-local/com/apple/pie/queue/kafka-pie-client/0.15.8)
and can be configured as a [Internal Link: Gradle
Dependency](../client/quick_start.html#gradle-dependency).
If you are using older versions of ACI Kafka 0.15, please upgrade to the
latest version, 0.15.8.

If you are using end-to-end crypto, you will need to add additional
configurations for the Iris connect string if you are not already using
[Internal Link: Kaffe configuration injection](../client/kaffe.html). Other
configuration previously required for e2e crypto should also be removed.
For more details, [Internal Link: see
here](../client/security_settings.html#payload-encryption-and-decryption-settings).


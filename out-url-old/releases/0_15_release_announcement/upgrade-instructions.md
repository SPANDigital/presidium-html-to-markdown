---
title: "Upgrade Instructions¶"
weight: 2
---

The ACI Kafka 0.15.8 client is available on
[Artifactory]({{%baseurl%}}/https://artifacts.apple.com/pie-release-local/com/apple/pie/queue/kafka-pie-client/0.15.8)
and can be configured as a [Gradle\
Dependency]({{%baseurl%}}/../client/quick_start.html#gradle-dependency).
If you are using older versions of ACI Kafka 0.15, please upgrade to the
latest version, 0.15.8.

If you are using end-to-end crypto, you will need to add additional
configurations for the Iris connect string if you are not already using
[Kaffe configuration injection]({{< ref "client/kaffe" >}}). Other
configuration previously required for e2e crypto should also be removed.
For more details, [see\
here]({{%baseurl%}}/../client/security_settings.html#payload-encryption-and-decryption-settings).


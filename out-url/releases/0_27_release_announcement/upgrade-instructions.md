---
title: "Upgrade Instructions¶"
weight: 2
---

The ACI Kafka 0.27 client is available on [External Link: Artifactory](https://artifacts.apple.com/pie-release-local/com/apple/pie/queue/kafka-pie-client).
If you are using an older version of the ACI Kafka client, please upgrade to
the [Internal Link: current recommended version](index.html).

### PieEnvelope[Internal Link: ¶](\#pieenvelope)

Although this release brings no changes in the envelope deprecation process, we would like to remind you to migrate your application away from the envelope if you haven’t done so yet.
Support for consuming messages in the envelope format is still available in this version.
This continues to guarantee interoperability with older clients producing messages in the envelope format.
However, as with the previous release, producing messages in this old format is no longer possible.
Please make sure you have fully migrated away from the envelope wire format.
For more details, please refer to the documentation on the [Internal Link: `PieEnvelope` deprecation](../client/envelope.html).

Warning

This will be the last release that still supports consuming messages in the envelope wire format.
The envelope format has been deprecated in 2019, and support for producing messages in this
format was removed in 2021 in version 0.20.


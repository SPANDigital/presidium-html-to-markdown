---
title: "Upgrade Instructions¶"
weight: 2
---

The ACI Kafka 0.25 client is available on [External Link: Artifactory](https://artifacts.apple.com/pie-release-local/com/apple/pie/queue/kafka-pie-client).
If you are using an older version of the ACI Kafka client, please upgrade to
the [Internal Link: current recommended version](index.html).

### PieEnvelope[Internal Link: ¶](\#pieenvelope)

Although this release brings no changes in the envelope deprecation process, we would like to remind you to migrate your application away from the envelope if you haven’t done so yet.
Support for consuming messages in the envelope format is still available in this version.
This continues to guarantee interoperability with older clients producing messages in the envelope format.
However, as with the previous release, producing messages in this old format is no longer possible.
Please make sure you have fully migrated away from using `PieEnvelope` and the envelope wire format.
For more details, please refer to the documentation on the [Internal Link: `PieEnvelope` deprecation](../client/envelope.html).


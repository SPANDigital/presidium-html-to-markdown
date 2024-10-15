---
title: "Upgrade Instructions¶"
weight: 2
---

The ACI Kafka 0.23 client is available on [External Link: Artifactory](https://artifacts.apple.com/pie-release-local/com/apple/pie/queue/kafka-pie-client).
If you are using an older version of the ACI Kafka client, please upgrade to the latest version.

The default value for request timeout - `pie.queue.kaffe.request.timeout.ms` \- has been increased from 2000 to 10000 milliseconds.

Although this release brings no changes in the envelope deprecation process, we would like to remind you to migrate your application away from the envelope if you haven’t done so yet.
Support for consuming messages in the envelope format is still available in this version.
This continues to guarantee interoperability with older clients producing messages in the envelope format.
However, as with the previous release, producing messages in this old format is no longer possible.

Please make sure you have fully migrated away from using `PieEnvelope` and the envelope wire format.
For more details, please refer to the documentation on the [Internal Link: `PieEnvelope` deprecation](../client/envelope.html).

We don’t have any upstream changes to highlight on this particular release.
For details please consult the [External Link: release notes for Kafka 3.1](https://archive.apache.org/dist/kafka/3.1.0/RELEASE_NOTES.html).
Or, for general overview on what’s new in Kafka 3.1 you might be interested in [External Link: this post](https://blogs.apache.org/kafka/entry/what-s-new-in-apache7) in the Apache Kafka
blog.

For more information on the ACI Kafka service offering, including
features, caveats, and project road map, visit our [Internal Link: overview
section](../overview.html).

You can [Internal Link: reach out to the ACI Kafka team](../contact.html) with
questions on this release or general inquiries related to ACI Kafka.

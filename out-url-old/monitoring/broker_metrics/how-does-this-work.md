---
title: "How does this work?¶"
weight: 3
---

ACI Kafka reports metrics to Mosaic in multiple ways:

1. Every Kafka broker runs a Mosaic reporter, publishing metrics that are tagged with `topic` or `client` back to Mosaic.
2. We run an internal fork [lightbend/kafka-lag-exporter](https://github.com/lightbend/kafka-lag-exporter).
3. We run an internally developed synthetic monitoring tool, publishing metrics that are tagged with `usecase` ( _i.e._ cluster) to Mosaic.

Whenever applicable, all metrics are tagged with `namespace`, `topic` and `client` are published to Mosaic.

Moreover, ACI Kafka uses [Mosaic Sharing Policies](https://telemetry.g.apple.com/docs/sharing-policies) to share metrics with
a [customer groups]({{%baseurl%}}/../concepts.html#customer-group) Apple Directory group.

Note

Please make sure your [customer groups]({{%baseurl%}}/../concepts.html#customer-group) Apple Directory groups are up-to-date. They must be configured as follows:

- Group/contact email is enabled
- Members visibility should be `Shown`
- Membership eligibility should be `Only those with Apple email addresses`
- Other Apple Directory groups can become members of this group


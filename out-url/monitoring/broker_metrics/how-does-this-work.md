---
title: "How does this work?¶"
weight: 3
---

ASE Kafka reports metrics to Mosaic in multiple ways:

1. Every Kafka broker runs a reporter, publishing metrics that are tagged with `topic` or `user` back to Mosaic.
2. We run an internal fork [External Link: lightbend/kafka-lag-exporter](https://github.com/lightbend/kafka-lag-exporter).
3. We run an internally developed synthetic monitoring tool, publishing metrics that are tagged with `usecase` ( _i.e._ cluster) to Mosaic.

Whenever applicable, all metrics are tagged with `namespace`, `topic` and `client` are published to Mosaic.


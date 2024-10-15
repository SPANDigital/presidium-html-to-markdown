---
title: "Kafka Hubble Yammer Reporter¶"
weight: 1
---

The hubble yammer reporter only extends _PieHubbleProperties_ with a new
stats collection for yammer metrics.

We plan in the future to report other system stats (bandwidth use, disk
usage…). Configuration is the same as the
[Internal Link: kafka-hubble-reporter](kafka_hubble_reporter.html).

To start using it you only need to specify in your broker properties
file:

```
metric.reporters=com.apple.pie.queue.kafka.hubble.HubbleMetricsReporter
pie.queue.config.class=com.apple.pie.queue.kafka.hubble.PieBrokerHubbleProperties

```

scope in yammer are treated like tags in KafkaMetrics so the
NameResolver will behave the same way for both types of stats.


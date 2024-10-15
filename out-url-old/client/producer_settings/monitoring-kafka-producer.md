---
title: "Monitoring Your Kafka Producer¶"
weight: 5
---

We recommend that you monitor your producer activities and to make this
easier we provide [kafka hubble\
reporter]({{< ref "monitoring" >}}) The
following metrics are recommended to track in the producer

| Metric | Description |
| --- | --- |
| `batch-size-avg` and `batch-size-max` | monitor bytes per batch |
| `record-queue-time-avg` | how long it takes to fill a batch |
| `producer-metrics-buffer-available-bytes` | checking available memory. |
| `producer-topic-metrics-byte-rate` and `producer-topic-metrics-byte-total` | to monitor the average and maximum number of bytes sent for a topic. |
| `producer-metrics-outgoing-byte-rate` and `producer-metrics-outgoing-byte-total` | monitor the average and maximum number of bytes sent for all brokers. |
| `producer-metrics-produce-throttle-time-avg` and `producer-metrics-produce-throttle-time-max` | monitor<br>average and maximum time in milliseconds a request was throttled by<br>a broker. |


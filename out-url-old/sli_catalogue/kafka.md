---
title: "Kafka¶"
weight: 2
---

Kafka is the core of the ACI Kafka service. We measure the availability
of Kafka for produce and consume requests.

### Produce Availability [¶](#produce-availability "Link to this heading")

We measure produce availability in two ways:

- **ISR Correlation** Kafka natively understands _availability to_
  _receive produce requests_ as being those requests made whilst the
  topic has at least the minimum number of in-sync replicas (aka
  MinISR). Therefore, we count the number of produce requests made when
  the number of ISRs is below MinISR and classify these as failed.
  Produce requests made when the number of ISRs is equal to or above
  MinISR are classified as succeeding. Availability, therefore, is the
  ratio of failed to successful produce requests.
- **Failed vs. Errored Requests** we also measure availability as the
  ratio of produce requests which result in an error to those that do
  not.

### Consume Availability [¶](#consume-availability "Link to this heading")

- **Failed vs. Errored Requests** we measure consume availability as
  the ratio of consume requests which result in an error to those that
  do not.

### Kafka Consume Latency [¶](#kafka-consume-latency "Link to this heading")

We measure consume latency as the Broker’s TotalTimeMs, which is the
total time for a consumer fetch to complete.

### Kafka Produce Latency [¶](#kafka-produce-latency "Link to this heading")

We measure produce latency as the Broker’s LocalTimeMs, which is the
time the request is being processed by the leader locally.


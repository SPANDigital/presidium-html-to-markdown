---
title: "Kafka Client OpenTelemetry integration¶"
weight: 1
---

The Java Kafka client (producer or consumer) can be instrumented with OpenTelemetry for monitoring. The open source
OpenTelemetry project provides [External Link: instrumentation](https://github.com/open-telemetry/opentelemetry-java-instrumentation/tree/v2.4.0/instrumentation/kafka/kafka-clients/kafka-clients-2.6/library)
for metrics and tracing support.

Once your client is instrumented, metrics can be exported to Mosaic.


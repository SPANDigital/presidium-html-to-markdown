---
title: "Configuring OpenTelemetry for Mosaic¶"
weight: 4
---

Follow the [Mosaic documentation](https://telemetry.apple.com/docs/collection/opentelemetry).

If using the Java agent or autoconfigured SDK, you’ll need at least the following set of config to export metrics to
the [Mosaic OpenTelemetry Metrics Gateway](https://telemetry.apple.com/blog/mosaic-otlp-metrics-gateway).

```
otel.metrics.exporter: otlp
otel.exporter.otlp.endpoint: https://mosaic-ingestion-gateway.telemetry.g.apple.com:24184
otel.exporter.otlp.protocol: grpc
otel.exporter.otlp.certificate: /certs-dir/trusted-root.pem
otel.exporter.otlp.client.key: /certs-dir/application.key
otel.exporter.otlp.client.certificate: /certs-dir/application.crt
otel.exporter.otlp.headers: X-MOSAIC-WORKSPACE=$WORKSPACE,X-MOSAIC-NAMESPACE=$NAMESPACE

```

---
title: "Broker metrics¶"
weight: 1
---

ASE Kafka customers can access metrics for their namespaces and clients on [External Link: Mosaic](https://telemetry.g.apple.com) using Grafana or [External Link: PromQL](https://prometheus.io/docs/prometheus/latest/querying/basics/).
ASE Kafka publishes metrics to Mosaic, including:

- The health of Kafka topics
- The real usage of resources
- Client throttling (if enabled)
- Consumer lag
- Produce and consume experience (through [Internal Link: synthetic monitoring](#synthetic-metrics))

Note

**This tool is provided as-is, without a Production SLA. SLA and SLO will be announced in the future.**[Internal Link: Contact us](../contact.html) if you would like to access your metrics.

Note

Please keep in mind that, at the moment, this does not replace the Hubble dashboards for ASE Kafka clusters your team has access to.


---
title: "Broker metrics¶"
weight: 1
---

ACI Kafka customers can access metrics for their namespaces and clients on [Mosaic]({{%baseurl%}}//files/telemetry.g.apple.com) using Grafana or [PromQL](https://prometheus.io/docs/prometheus/latest/querying/basics/).
ACI Kafka publishes metrics to Mosaic, including:

- The health of Kafka topics
- The real usage of resources
- Client throttling (if enabled)
- Consumer lag
- Produce and consume experience (through [synthetic monitoring](#synthetic-metrics))

To set up Mosaic alerts against these metrics, customers are required to own a dedicated Mosaic workspace.
Please follow the [Can I set up alerts?]({{%baseurl%}}/../client/lag_metrics.html#how-do-i-set-up-alerts) section on how to set up alerts based on these metrics.

Note

**This tool is provided as-is, without a Production SLA. SLA and SLO will be announced in the future.** [Contact us]({{< ref "contact" >}}) if you would like to access your metrics.


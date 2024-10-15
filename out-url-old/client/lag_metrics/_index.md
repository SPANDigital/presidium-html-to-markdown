---
title: "Consumer Group’s lag metrics¶"
weight: 1
---

ACI Kafka customers can access lag metrics for their eligible consumer group on [Mosaic]({{%baseurl%}}//files/telemetry.g.apple.com) using Grafana or [PromQL](https://prometheus.io/docs/prometheus/latest/querying/basics/).
ACI Kafka produces metrics to Mosaic, and doesn’t grantee the ability to setup alerts as this require that customer own a dedicated Mosaic’s workspace.

Please follow the [Can I set up alerts?](#how-do-i-set-up-alerts) section on how to set up alerts based on these metrics.

Note

**This tool is provided as-is, without a Production SLA. SLA and SLO will be announced in the future.** [Contact us]({{< ref "contact" >}}) if you would like to access your metrics.


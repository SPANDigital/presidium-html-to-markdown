---
title: "Checking cluster health¶"
weight: 3
---

You can check if the Kafka cluster is generally healthy using the
[Synthetic Monitor dashboard](https://aci-kafka.grafana.telemetry.g.apple.com/d/d421ddc0-a18f-467f-93ae-3b71132999cf/synthetic-monitor).

Note

If you can’t access the dashboard please self-join [aci-kafka-announce Apple Directory group](adir://joingroup/670759).
It might take some time for the Grafana access to sync.

Once you select your environment and cluster name you should be able to see round trip latency, broker coverage and failure counts.


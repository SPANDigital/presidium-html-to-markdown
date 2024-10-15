---
title: "What Kafka resources can I monitor?¶"
weight: 2
---

- [Namespaces]({{%baseurl%}}/../concepts.html#namespace)
- [Topic partitions]({{%baseurl%}}/../concepts.html#topic)
- Client metrics include throttling due to [bandwidth throttling]({{%baseurl%}}/../concepts.html#cluster-traffic-quota)
- [Consumer lag metrics]({{< ref "client/lag_metrics" >}})
- [Synthetic metrics](#synthetic-metrics) ( _dedicated clusters only_)
  - Produce experience
  - Consume experience
  - Latency

Note

If you are a customer with a dedicated ACI Kafka cluster, we can share all Kafka metrics available on Mosaic with your team.
Please keep in mind that, at the moment, this does not replace the Hubble dashboards for ACI Kafka clusters your team has access to.


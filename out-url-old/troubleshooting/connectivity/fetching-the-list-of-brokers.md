---
title: "Fetching the list of brokers¶"
weight: 4
---

You can get the [list of monitored brokers from the Synthetic Monitor dashboard](https://aci-kafka.grafana.telemetry.g.apple.com/d/d421ddc0-a18f-467f-93ae-3b71132999cf/synthetic-monitor).

Alternatively you can get the list of brokers for a given cluster using our [Kaffe gRPC API]({{%baseurl%}}/../client/kaffe.html#kaffe-endpoints):

```
grpcurl -proto kaffe.proto \
  -d '{"proto_info": {"version": "0.3"}, "resource_id": {"type": 2, "id": "<cluster-id>"}}' \
  kaffe-prod.kafka.g.apple.com:443 protocol.Kaffe/WatchResource \
  | jq -r '.cluster.brokers'

```


---
title: "Generating commands to check network connectivity¶"
weight: 5
---

You can use the Kaffe cluster response to generate netcat ( `nc`) commands, which you can then execute on the hosts running your Kafka client:

```
grpcurl -proto kaffe.proto \
  -d '{"proto_info": {"version": "0.3"}, "resource_id": {"type": 2, "id": "<cluster-id>"}}' \
  kaffe-prod.kafka.g.apple.com:443 protocol.Kaffe/WatchResource \
  | jq -r '.cluster.brokers | .[] | "nc -vz \(.host) \(.endpoint.SASL_SSL)"'

```

Warning

When checking connectivity to the brokers make sure you’re running the command from the same host, pod or instance that is logging the errors or timeouts.

As network issues can be intermittent you should try running the command in a loop for a couple of minutes.


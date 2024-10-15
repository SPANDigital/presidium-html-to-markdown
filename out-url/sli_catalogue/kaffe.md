---
title: "Kaffe¶"
weight: 4
---

Kaffe supports the service discovery mechanism in our ACI Kafka client. Its
availability is important for clients, because they query Kaffe during startup
to obtain a list of Kafka brokers. We probe two endpoints (heartbeat and watch
resource) of the Kaffe service from three AODC Kube regions ( `us-east-1a`,
`us-west-1a` and `us-west-2a`) using
[External Link: Prober](https://prober.apple.com/workspace/aci-kafka) against its GSLB
endpoint. The watch resource endpoint is queried for client, cluster and
namespace resources, each for existing and non-existing resources.

### Availability[Internal Link: ¶](\#availability)

The availability is defined as the number of failed requests over the number of
total requests, expressed as a percentage: `1 - (failures / total_requests)`

### Latency[Internal Link: ¶](\#latency)

The latency is defined as the number of slow requests over the number of total
requests, expressed as a percentage: `1 - (slow_requests / total_requests)`. In
the case of Kaffe, a request is considered to be slow if it took more than
3000ms.


---
title: "Iris¶"
weight: 5
---

Iris is a service which serves public keys for message recipients and encrypted
private keys to allow asymmetric encryption of messages. Its availability is
important for clients as they fetch keys at initialization time and periodically
thereafter. We probe three endpoints (health, public and private key) of the
Iris service from three AODC Kube regions ( `us-east-1a`, `us-west-1a` and
`us-west-2a`) using [Prober](https://prober.apple.com/workspace/aci-kafka)
against its GSLB endpoint. The public and private key endpoints are each queried
for existing and non-existing topics.

### Availability [¶](#id1 "Link to this heading")

The availability is defined as the number of failed requests over the number of
total requests, expressed as a percentage: `1 - (failures / total_requests)`

### Latency [¶](#id2 "Link to this heading")

The latency is defined as the number of slow requests over the number of total
requests, expressed as a percentage: `1 - (slow_requests / total_requests)`. In
the case of Iris, a request is considered to be slow if it took more than 3000ms.


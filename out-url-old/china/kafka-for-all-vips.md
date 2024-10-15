---
title: "Kafka For All VIPs¶"
weight: 6
---

If you use Kafka For All, you will need to know the Kubernetes headless
service name for your cluster. These follow the pattern:

```
<aci-kafka-cluster-name>.pie-queue-prod-local.svc.kube.<kube-cluster>.k8s.cloud.silu.net:9152

```

In Kubernetes, headless services return a multi-A record result for DNS.
They function (that is, resolving to valid DNS responses) as long as a
single Pod is alive and healthy in the cluster. Kafka then uses this to
bootstrap connectivity to the cluster, where further discovery is
managed using the Kafka binary protocol.

Note

You will also need to configure the trust root in your client.

### Verifying Connectivity [¶](#verifying-connectivity "Link to this heading")

You can verify DNS and Connectivity using the examples below
(substituting the Kafka and Kube cluster values):

```
dig <aci-kafka-cluster-name>.pie-queue-prod-local.svc.kube.<kube-cluster>.k8s.cloud.silu.net
nc -vz  <aci-kafka-cluster-name>.pie-queue-prod-local.svc.kube.<kube-cluster>.k8s.cloud.silu.net 9152

```


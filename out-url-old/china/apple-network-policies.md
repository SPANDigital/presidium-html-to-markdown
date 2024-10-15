---
title: "Apple Network Policies (ANPs)¶"
weight: 7
---

If you are deployed within Kubernetes, you will need us to grant access
from your namespace to reach our clusters. This is typically done during
on-boarding, but if it was missed or additional policies are required
then please file a radar on `ACI Queue | Kafka Support`, including the
namespaces you require access from.

If you do not provide or specify a Kube cluster name, we will default to
`cluster: ALL` in the ANP.

Note

If you are deployed in Kube, you will also need to configure your applications ANP to permit traffic _outbound_.
Our namespace for all clusters is `pie-queue-prod-local.kube`, and we recommend `cluster: ALL`.


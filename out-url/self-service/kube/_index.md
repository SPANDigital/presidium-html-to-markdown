---
title: "Kubernetes CRDs¶"
weight: 1
---

ACI Kafka _topics_ and _client identities_ can be managed using Kubernetes through
[External Link: custom resource definitions (CRDs)](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/#customresourcedefinitions) such as `KafkaTopic`, `KafkaClientIdentity` and `KafkaCustomerGroupBinding`.

It can be easier to self-service ACI Kafka entities via Kubernetes using an API ( `kube-api`) and/or CLI ( `kubectl`) which you may already be familiar with.
It may also be useful to define and manage specifications for your topics and client identities along with workloads or other infrastructure.


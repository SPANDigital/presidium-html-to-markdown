---
title: "Self Service¶"
weight: 8
---

For all clusters in China, we strongly encourage the use of Kubernetes
CRDs for self service. As part of your on-boarding, we will configure a
single namespace for you to manage all of our entities (currently Topics
and Client Identities).

Note

Self service is defined at the Customer Group level, where ACI Kafka authorization is modeled.
This means you can only use a single Kube namespace to manage all entities within that Customer Group.
For example, if you have ACI Kafka Namespaces in both cn-east-1a and cn-east-2a,
and self service configured for your Kube Namespace in cn-east-1a,
**you will only be able to manage Kafka topics using your cn-east-1a namespace**.

Self service is currently available in the following Kube Clusters in
China:

- cn-east-1a
- cn-east-2a

For more detailed information, check the [Kubernetes\
CRD]({{< ref "self-service/kube" >}}) section of our self service
documentation.

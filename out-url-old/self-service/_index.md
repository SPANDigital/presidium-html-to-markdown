---
title: "Self Service¶"
weight: 1
---

Apple Cloud Infrastructure (ACI) Kafka **Self Service** provides you with autonomy by offering multiple ways to manage your entities.
Before starting, we recommend becoming familiar with the ACI Kafka [Concepts]({{< ref "concepts" >}}), including clusters, customer groups, client identities, namespaces, and topics.

You can manage your ACI Kafka resources using one of the following methods:

- [Kubernetes CRDs]({{< ref "self-service/kube" >}})
- [ACI Kafka CLI]({{< ref "self-service/kafka_cli" >}})
- [Public API]({{< ref "self-service/public_api" >}})
- [Inspector]({{< ref "self-service/inspector" >}})
- [Broker metrics]({{< ref "monitoring/broker_metrics" >}})

**[Kubernetes Custom Resource Definitions (CRDs)]({{< ref "self-service/kube" >}})**
— Allows you to manage ACI Kafka resources using `kubectl`. This self-service option is for customers who can access a Kube namespace in [one of the supported clusters]({{%baseurl%}}/kube.html#availability).

**[ACI Kafka CLI]({{< ref "self-service/kafka_cli" >}})**
— Lets you modify your entities using a command-line interface built on top of the [public API]({{< ref "self-service/public_api" >}}).

**[Public API]({{< ref "self-service/public_api" >}})**
— Allows you to interact with ACI Kafka entities with `curl` or an `HTTP` client of choice.

**[Inspector]({{< ref "self-service/inspector" >}})**
— Allows you to see your topic’s health, its offsets and active consumer groups.

Note

Inspector is being replaced by another metrics based solution. For more details check [Broker metrics]({{< ref "monitoring/broker_metrics" >}})

**[Broker metrics]({{< ref "monitoring/broker_metrics" >}})**
— Allows you to see metrics regarding your topic partitions health, your producer and consumer activities including consumer lag and your actual usage.

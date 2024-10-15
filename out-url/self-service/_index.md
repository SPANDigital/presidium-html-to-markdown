---
title: "Self Service¶"
weight: 1
---

Apple Cloud Infrastructure (ACI) Kafka **Self Service** provides you with autonomy by offering multiple ways to manage your entities.
Before starting, we recommend becoming familiar with the ACI Kafka [Internal Link: Concepts](../concepts.html), including clusters, customer groups, client identities, namespaces, and topics.

You can manage your ACI Kafka resources using one of the following methods:

- [Internal Link: Kubernetes CRDs](kube.html)
- [Internal Link: ACI Kafka CLI](kafka_cli.html)
- [Internal Link: Public API](public_api.html)
- [Internal Link: Inspector](inspector.html)
- [Internal Link: Broker metrics](../monitoring/broker_metrics.html)

**[Internal Link: Kubernetes Custom Resource Definitions (CRDs)](kube.html)**
— Allows you to manage ACI Kafka resources using `kubectl`. This self-service option is for customers who can access a Kube namespace in [Internal Link: one of the supported clusters](kube.html#availability).

**[Internal Link: ACI Kafka CLI](kafka_cli.html)**
— Lets you modify your entities using a command-line interface built on top of the [Internal Link: public API](public_api.html).

**[Internal Link: Public API](public_api.html)**
— Allows you to interact with ACI Kafka entities with `curl` or an `HTTP` client of choice.

**[Internal Link: Inspector](inspector.html)**
— Allows you to see your topic’s health, its offsets and active consumer groups.

Note

Inspector is being replaced by another metrics based solution. For more details check [Internal Link: Broker metrics](../monitoring/broker_metrics.html)

**[Internal Link: Broker metrics](../monitoring/broker_metrics.html)**
— Allows you to see metrics regarding your topic partitions health, your producer and consumer activities including consumer lag and your actual usage.

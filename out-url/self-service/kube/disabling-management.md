---
title: "Disabling management¶"
weight: 6
---

To disable management and disassociate the Kube namespace with your ACI Kafka customer group, refer to the relevant section in the [Internal Link: Public API](public_api.html#kube) or [Internal Link: ACI Kafka CLI](kafka_cli.html#kube) documentation for instructions.
Once the change has been made you can safely delete the `KafkaCustomerGroupBinding` and other resources without causing any changes to your topics and client identities.

```
$ kubectl describe KafkaClientIdentity <resource-name>

```

Alternatively use `kubectl get -o yaml <resource-name>` to see the full YAML definition and status.


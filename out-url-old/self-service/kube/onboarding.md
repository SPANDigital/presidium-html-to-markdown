---
title: "Onboarding¶"
weight: 3
---

To use Kubernetes to manage ACI Kafka topics and client identities, you must first authorize a Kubernetes namespace.
This should be a namespace which your team owns, and you have access to.
For example, you can create and manage Pods, ConfigMaps, and other resources.

To authorize a Kubernetes namespace for ACI Kafka self-service, [customer groups]({{%baseurl%}}/../concepts.html#customer-group) must first be decorated with the Kubernetes namespace name, as well as the Kubernetes cluster name.
Refer to the relevant section in [Public API]({{%baseurl%}}/public_api.html#kube) or [ACI Kafka CLI]({{%baseurl%}}/kafka_cli.html#kube) for instructions.

Once authorized, a `KafkaCustomerGroupBinding` resource needs to be created in the authorized namespace.
This should happen automatically. If it does not, create a file with a binding definition, i.e., `binding.yaml`:

```
apiVersion: api.kafka.aci.apple.com/v1
kind: KafkaCustomerGroupBinding
metadata:
  name: <object name>
spec:
  # The fully qualified customer group ID in form <customer>.<group>.<env>
  customerGroupID: <customerGroup>

```

Then apply the file against a Kubernetes cluster and namespace:

```
$ kubectl apply -f binding.yaml

```

Once it is created, you want to check if the resource is `Ready`:

```
# Output the full YAML
$ kubectl get kafkacustomergroupbinding <resource-name> -o yaml

# Or get the condition directly
$ kubectl get kafkacustomergroupbinding <resource-name> -o jsonpath='{.status.conditions[?(@.type=="Ready")]}'

```

If the authorization is not set, your output might contain the following:

```
status:
  conditions:
  - lastTransitionTime: "2020-11-04T12:33:25Z"
    message: customer group is not authorized for this namespace
    observedGeneration: 2
    reason: Error
    status: "False"
    type: Ready

```

If the authorization is set correctly, your output should look similar to:

```
status:
  conditions:
  - lastTransitionTime: "2020-11-04T12:42:42Z"
    message: specification is fully reconciled
    observedGeneration: 2
    reason: Ready
    status: "True"
    type: Ready

```

Once the resource is `Ready`, you are fully set up and can start managing ACI Kafka topics and client identities in the same Kubernetes namespace.


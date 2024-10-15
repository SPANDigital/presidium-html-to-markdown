---
title: "Usage¶"
weight: 4
---

The following sections describe how you can interact with the ACI Kafka CRDs, once you have finished [Internal Link: onboarding](#onboarding).

### List topics and client identities[Internal Link: ¶](\#list-topics-and-client-identities)

Creating the aforementioned `KafkaCustomerGroupBinding` will trigger the creation of `KafkaTopic` or `KafkaClientIdentity` resources for every existing ACI Kafka topic and client identity under the customer group identified by the name of the `KafkaCustomerGroupBinding` object.

You can list all `KafkaTopic` and `KafkaClientIdentity` resources:

```
$ kubectl get KafkaTopic,KafkaClientIdentity

```

### Delete a topic or a client identity[Internal Link: ¶](\#delete-a-topic-or-a-client-identity)

Deleting a Kubernetes resource of type `KafkaTopic` or `KafkaClientIdentity` in an authorized namespace will delete the respective ACI Kafka topic or client identity:

```
# Delete a KafkaTopic
$ kubectl delete KafkaTopic <resource-name>

# Delete a KafkaClientIdentity
$ kubectl delete KafkaClientIdentity <resource-name>

```

### Describe a specific topic[Internal Link: ¶](\#describe-a-specific-topic)

Each ACI Kafka topic is represented by a resource of type `KafkaTopic`.
When you create a `KafkaCustomerGroupBinding` for a specific customer group in an authorized Kubernetes namespace, all ACI Kafka topics under that customer group become represented in the same namespace as newly-created resources of type `KafkaTopic`.

```
$ kubectl describe KafkaTopic <resource-name>

```

Alternatively, use `kubectl get -o yaml <resource-name>` to see the full YAML definition and status.

### Create a new topic[Internal Link: ¶](\#create-a-new-topic)

First define the topic in a file, i.e., `topic.yaml`:

```
apiVersion: api.kafka.aci.apple.com/v1
kind: KafkaTopic
metadata:
  name: <object name>
spec:
  # The fully qualified topic ID, in string form "<customer>.<group>.<env>.<namespace>.<topic-name>"
  topicID: "<customer>.<group>.<env>.<namespace>.<topic-name>"

  # A text description for the topic
  description: "An example topic"

  # Topic Resources and quotas for retention and throughput
  resources:
    # Consume capacity. The maximum rate, in bytes per second, of consumed messages from the Topic.
    consumePerSec: 50Mi

    # Unit of parallelism, - the number of partitions in the topic.
    partitions: 30

    # Produce Capacity. The maximum rate, in bytes per second, of produced messages to the Topic.
    producePerSec: 35M

    # Retention capacity. Controls the maximum time before old messages will be discarded.
    # If not defined, no time limit is applied. See https://kafka.apache.org/documentation.html#retention.ms
    retentionTime: 5h0m0s # this is optional

    # The maximum total size in bytes of retained messages in the topic.
    # Controls the maximum size the Topic can grow to before old messages are discarded.
    # This setting divided by the number of Partitions defines the `retention.bytes` property in the Topic.
    # i.e. `retention.bytes` will be `size / partitions`
    # See https://kafka.apache.org/documentation.html#retention.bytes
    size: 10Gi

  # List of client identities that can access the topic
  accesses:

    # Client accesses that can produce to the topic
    producers:
    - client: <fully qualified ID of the client> # The fully qualified client identity ID, in string form
      # e.g. "<customer>.<group>.<env>.<client-identity-name>"
      quota: 31Mi # The access quota in bytes per second. The maximum rate, in bytes per second, at which messages will flow

    consumers:
    - client: <fully qualified ID of the client> # The fully qualified client identity ID, in string form
      # e.g. "<customer>.<group>.<env>.<client-identity-name>"
      quota: 20M # The access quota in bytes per second. The maximum rate, in bytes per second, at which messages will flow

```

After saving the file, you can apply it:

```
$ kubectl apply -f topic.yaml

```

Once the resource is created, you want to check if it is `Ready`:

```
# Output the full YAML
$ kubectl get KafkaTopic <resource-name> -o yaml

# Or get the condition directly
$ kubectl get KafkaTopic <resource-name> -o jsonpath='{.status.conditions[?(@.type=="Ready")]}'

```

Once the resource is `Ready`, the topic has been provisioned and is ready for usage.

### Update a topic[Internal Link: ¶](\#update-a-topic)

One way to update a topic is to update the YAML definition file, i.e., `topic.yaml` and re-apply it:

```
# 1. If you have no copy of the topic.yaml, grab the current version from Kube:
$ $ kubectl get KafkaTopic <topic-name> -o yaml > topic.yaml

# 2. Edit the file in the editor of your choice

# 3. Finally, apply the new version
$ kubectl apply -f topic.yaml

```

You can also edit the definition directly using `kubectl edit KafkaTopic <resource-name>`.

### Create a new client identity[Internal Link: ¶](\#create-a-new-client-identity)

To create a new client identity, you can define the client identity in a YAML file, i.e., `identity.yaml`:

```
apiVersion: api.kafka.aci.apple.com/v1
kind: KafkaClientIdentity
metadata:
  name: <object name>
spec:
  # The fully qualified client identity ID, in string form
  # "<customer>.<group>.<env>.<client-identity-name>"
  clientIdentityID: "<customer>.<group>.<env>.<client-identity-name>"

  # A PEM encoded, PKCS#8 RSA public key.
  # A key pair can be generated with:
  #   openssl genrsa 2048 | openssl pkcs8 -topk8 -nocrypt -outform pem > private_key.pem
  #   openssl rsa -pubout -in private_key.pem -out public_key.pem
  publicKey: |
    -----BEGIN PUBLIC KEY-----
    MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0Ep9WbZL+EiObXM/TGCW
    eJVzZ0o7E6BWXcG5JMAPAySqBFTMlK3VtS7z2SuYWCnn1hGvMi28+dnDbpmmMh4+
    K4xR2YXMLnLOJvh0zRfB8z4wqoUuW+ZKXFANwq6dheYhyf7Zu62PmNwAJPLkWt/S
    dPREPEgCmnijWv3RpD1zWEQIx6FnSofTHDC8S+2unPt0myoaEekCa1YfQTzLrZjz
    JAsTo2X6dGHH8WUn8JEE1JrFnNQofAETms91KJq+lzFjVphtMzJbFa/uIgNUML6P
    V4DMZhpPxqX+AMDC+CCjz0a/rcvQFRcpmh9S6C9p/ivtf6UB8J3SPmFchMIK7Le+
    uQIDAQAB
    -----END PUBLIC KEY-----

```

And apply the definition:

```
$ kubectl apply -f identity.yaml

```

Once the resource is created, you want to wait for it to become `Ready`:

```
# Output the full YAML
$ kubectl get KafkaClientIdentity <resource-name> -o yaml

# Or get the condition directly
$ kubectl get KafkaClientIdentity <resource-name> -o jsonpath='{.status.conditions[?(@.type=="Ready")]}'

```

Once the resource is `Ready`, the client identity has been provisioned and is ready for usage.

### Describe a specific client identity[Internal Link: ¶](\#describe-a-specific-client-identity)

Each ACI Kafka client identity is represented by a resource of type `KafkaClientIdentity`.
When you create a `KafkaCustomerGroupBinding` for a specific customer group in an authorized Kubernetes namespace, all the ACI Kafka client identities under that customer group become represented in the same namespace as newly created resources of type `KafkaClientIdentity`.


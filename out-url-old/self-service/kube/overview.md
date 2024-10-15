---
title: "Overview¶"
weight: 2
---

### Availability [¶](#availability "Link to this heading")

ACI Kafka self-service via Kubernetes is available in the following clusters:

- `us-west-1a`
- `us-west-2a`
- `us-west-3a`
- `us-east-1a`\*
- `cn-east-1a`
- `cn-east-2a`

You can [reach out to the team]({{< ref "contact" >}}) if you would like to use a Kubernetes cluster that isn’t part of this list.

\* **QA only**: Available for QA (if1) resources only.

### Exploring ACI Kafka CRDs with `kubectl explain` [¶](#exploring-aci-kafka-crds-with-kubectl-explain "Link to this heading")

Our CRDs use a complete OpenAPI specification, and as such we support using `kubectl explain` to help you understand the resources and what each field should contain.

As such, you can explain the root `KafkaTopic` entity:

```
$ kubectl explain KafkaTopic
KIND:     KafkaTopic
VERSION:  api.kafka.aci.apple.com/v1

DESCRIPTION:
     KafkaTopic is the Schema for the kafkatopics API

FIELDS:
   apiVersion <string>
     APIVersion defines the versioned schema of this representation of an
     object. Servers should convert recognized schemas to the latest internal
     value, and may reject unrecognized values. More info:
     https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources

   kind   <string>
     Kind is a string value representing the REST resource this object
     represents. Servers may infer this from the endpoint the client submits
     requests to. Cannot be updated. In CamelCase. More info:
     https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds

   metadata   <Object>
     Standard object's metadata. More info:
     https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata

   spec   <Object>
     KafkaTopicSpec defines the desired state of KafkaTopic

   status <Object>
     KafkaTopicStatus defines the observed state of KafkaTopic

```

It is also possible to drill down, and explain the sub-resources:

```
$ kubectl explain KafkaTopic.spec
KIND:     KafkaTopic
VERSION:  api.kafka.aci.apple.com/v1

RESOURCE: spec <Object>

DESCRIPTION:
     KafkaTopicSpec defines the desired state of KafkaTopic

FIELDS:
   accesses   <Object>
     List of client identities that can access the topic

   compacted  <boolean>
     See Apache Kafka topic compaction.
     https://learning.oreilly.com/library/view/kafka-the-definitive/9781491936153/ch05.html

   description    <string> -required-
     A text description for the topic

   resources  <Object> -required-
     Topic Resources and quotas for retention and throughput

   topicID    <string> -required-
     The fully qualified topic ID, in string form
     "<customer>.<group>.<env>.<namespace>.<topic-name>"

```


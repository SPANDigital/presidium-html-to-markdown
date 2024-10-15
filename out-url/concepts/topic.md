---
title: "Topic¶"
weight: 5
---

ACI Kafka _topics_ are used to manage Kafka topics and include:

- Number of partitions
- Set of configuration parameters
- Produce and consume quotas

ACI Kafka clusters use a replication factor of four as it provides the best operational simplicity.

### Topic name[Internal Link: ¶](\#topic-name)

Topic names have a maximum length of 95 characters and can contain letters, numbers, and `_` or `-` characters.

Topics will appear on the cluster with their fully qualified name, which is prefixed by the namespace.
For example _topic name_ `topic123` in the namespace `icloud.ck-content.prod.dc1` appears on a cluster as: `icloud.ck-content.prod.dc1.topic123`

The ACI Kafka Java client provides a class to create fully qualified topic names for producers and consumers:

```
String namespaceId = "icloud.ck-content.prod.dc1";
String topicName = "topic123";

PieTopicName pieTopic = new PieTopicName(PieNamespaceName.of(namespaceId), topicName);
String fqTopicName = pieTopic.getFullName();

```

### Topic retention[Internal Link: ¶](\#topic-retention)

ACI Kafka topics require configuring retention size to allow for easier capacity planning of Kafka clusters.

If **time-based retention** _(optional)_ and **sized-based retention** _(required)_ are configured, data is deleted by brokers when the first retention rule is met.

Size-based retention is configured per partition.
Computationally, a topic with 100 partitions at 10 GB size-based retention will consume 4 TB of storage: 100 partitions x 10 GB x replication factor of 4.

For example, if a topic with one partition is configured with a size-based retention of 1 GB and a 24-hour time-based retention limit, only the latest 1 GB of data is retained.
All data older than 24 hours is deleted.

From the [External Link: Kafka Topic-Level Configuration Documentation](https://kafka.apache.org/documentation.html#topicconfigs):

**retention.bytes**

> This configuration controls the maximum size a partition (which
> consists of log segments) can grow to before we will discard old log
> segments to free up space if we are using the “delete” retention
> policy. By default there is no size limit only a time limit. Since
> this limit is enforced at the partition level, multiply it by the
> number of partitions to compute the topic retention in bytes.

**retention.ms**

> (default: 21 days) This configuration controls the maximum time we will retain a log
> before we will discard old log segments to free up space if we are
> using the default “delete” retention policy. This represents an SLA on how
> soon consumers must read their data.

Note

To ensure reliable cluster operations, we limit `retention.bytes` to `<= 250GB`.
We also limit the number of partitions to 512 when using self-service API.

Please contact us if your use case requires a higher limit.

### Topic quotas[Internal Link: ¶](\#topic-quotas)

You must configure all topic quotas for expected bytes in and bytes out when using ACI Kafka.
These topic-level quotas are used to budget the cluster and are not enforced on the actual produce or consume usage.


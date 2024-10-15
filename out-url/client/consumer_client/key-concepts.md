---
title: "Key Concepts¶"
weight: 2
---

### Consumer Group (a.k.a `group.id` config)[Internal Link: ¶](\#consumer-group-a-k-a-group-id-config)

A consumer group is a set of instances of the same application that needs all the messages from one or more topics.
The topic partitions consumed by this consumer are divided evenly among instances in the given group.
Adding instances to an existing consumer group can help to scale the consumption and processing of messages from the given topics.

Please keep in mind, while Kafka tries to divide the partitions’ load evenly between instances, it may fail in the following cases:

1. Consuming with an uneven ratio of instances to partitions.
   For example, when consuming using two instances from a topic with five partitions, one instance will process three partitions, whereas the other one will process 2.
2. Consuming from skewed partitions, when one partition has more data than the other. This may happen with compacted topics with a large number of partitions and a small number of keys.

![](http://client/../_images/consumer_group.png)

You need to be familiar with the following two concepts related to groups:

#### 1\. Group Coordinator[Internal Link: ¶](\#group-coordinator)

The Kafka server (broker) designated to oversee the group and group activities.
The coordinator is elected automatically by the cluster when the consumer starts consuming for the first time.

#### 2\. Group Member[Internal Link: ¶](\#group-member)

Each instance in a consumer group is called _**Member**_ and gets a unique member id when it joins.

Note

Number of members in a consumer group should not exceed the total number of consumed partitions.
As this will waste resources as no partitions will be assigned to any extra member and will have negative impact on
the process of dividing the partitions.

For more details on Consumer Group please check the [External Link: Kafka Consumer documentation](https://docs.confluent.io/platform/current/clients/consumer.html#consumer-groups).

### Poll[Internal Link: ¶](\#poll)

Poll is a method that a Kafka consumer calls to retrieve records from all assigned partitions in one list.

```
try {
     while
         }(true) {
       ConsumerRecords<String, String> records = consumer.poll(Duration.ofMillis(pollTimeoutMs));
         processRecords(records);
    } finally {
       consumer.close();
    }
 }

```

### Offset Commit[Internal Link: ¶](\#offset-commit)

Each consumer’s instance must inform the Group Coordinator of its current consumption position (the last processed record offset).
This process needs to happen as soon as the consumer processes the records.
Hence, if a consumer’s instance stops and comes back later, it restarts from the last committed position.

The default behavior of Consumer client is to automatically commit offset of any received records within the last 5 seconds
in the background even if the consumer didn’t finish processing the records, which may lead to data loss.
It’s advised to disable this behavior by configuring the client with `enable.auto.commit=false` and commit offset manually
as soon as you finish processing your records.

```
try {
    while(true) {
        ConsumerRecords<String, String> records = consumer.poll(Duration.ofMillis(pollTimeoutMs));
          processRecords(records);
          if (successfullyFinishedProcessing()) {
              consumer.commitSync();
        }
     } finally {
        consumer.close();
    }
}

```

### Lag[Internal Link: ¶](\#lag)

The delta between the last produced message to a topic partition and the last consumer committed offset for the same partition.

### Heartbeat[Internal Link: ¶](\#heartbeat)

A sign of the consumer’s instance liveness.
The group coordinator expects each instance to send a heartbeat every defined period before the current session expires or considers the instance dead.

Heartbeats are sent to the cluster every `heartbeat.interval.ms` by a background thread.
The window of the current session is expired every `session.timeout.ms`.
Check the [Internal Link: configuration section](#key-configurations) for more details on how to set the value of `heartbeat.interval.ms`.

Heartbeat response is a sign of liveness of the Group Coordinator, so if the group does not receive a heartbeat from the coordinator,
it will assume it is dead and automatically ask for a new coordinator.

### Rebalance[Internal Link: ¶](\#rebalance)

A rebalance occurs when one (or multiple) consumer instances join or leave the group. This redivides the total partitions onto all instances now remaining in the group.
While the rebalance process is makes it safe for an instance to join or leave the Consumer Group, it can be expensive for large ones.
They can even be undesirable as they disturb consumption.

### Summary[Internal Link: ¶](\#summary)

Kafka consumer applications are scaled by having multiple instances in the Consumer Group, which get assigned to a subset of partitions.
The consumer is backed up by the coordinator that manages the group’s state.
The following diagram summarizes the key stages in the consumer lifecycle, from creating the client instances to shutting down the client.

![](http://client/../_images/consumer_lifecycle.png)

For more details on the consumer concept, please check [External Link: Chapter 4 - Kafka: The Definitive Guide, 2nd Edition](https://learning.oreilly.com/library/view/kafka-the-definitive/9781492043072/ch04.html).


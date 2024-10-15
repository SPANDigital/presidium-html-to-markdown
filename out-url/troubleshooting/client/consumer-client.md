---
title: "Consumer client¶"
weight: 4
---

Consumer clients ingest data from Kafka in order to process the data independently.
Consumers can face challenges due to some misconfigurations, unexpected infrastructure issue, or unexpected load.
Common challenges include:

### Consumer group is experiencing too many rebalancing events[Internal Link: ¶](\#consumer-group-is-experiencing-too-many-rebalancing-events)

#### Description[Internal Link: ¶](\#id7)

Consumer group divides the number of partitions among consumers every time a consumer joins or leaves the consumer group;
This process is known as [Internal Link: Rebalancing](../client/consumer_client.html#rebalance).
This is critical for the scalability and availability of the consumer group.
However, in the default consumer group protocol this event can disturb the consumer when it gets triggered too many times unnecessarily.
This happens because consumer instances need to _stop-the-world_ in order for the rebalance to finish.

#### Impact[Internal Link: ¶](\#id8)

Data processing will be delayed or paused everytime a consumer group encounters a rebalancing event.

#### Causes and Solutions[Internal Link: ¶](\#id9)

1. A consumer group contains more instances than the total number of topic partitions.
   This can increase the chance of triggering join/leave events from the stand-by instance, which trigger a rebalance.

_Solution:_

Make sure the number of instances doesn’t exceed the number of [Internal Link: topic partitions](../self-service/public_api.html#fetch-single-topic).

1. Most of the causes for this issue can be traced back to how long it took the consumer to process a record or when a consumer sent its last [Internal Link: heartbeat](../client/consumer_client.html#heartbeat) signal
   , for example:
   1. The consumer did not finish processing records within `max.poll.interval.ms`.
      This may happen if the topic has an unexpected large record that took longer than expected to process.
   2. Something caused the record processing to take longer, or the consumer has a complicated and slow record processor.
      For example writing to slow database or encounters a fatal error during processing and not consumption.
   3. The consumer did not send a heartbeat within `session.timeout.ms`.
      This can be caused by network issues.

_Solution:_

- Make sure to [Internal Link: test the connectivity](connectivity.html) with the brokers.
- A simple solution is increasing `max.poll.interval.ms`(and/or decreasing `max.poll.records`/ `max.partition.fetch.bytes`)
- Consider using _Static Membership_.
  More details can be found in [External Link: Kafka Rebalance Protocol for the Cloud: Static Membership](https://www.confluent.co.uk/blog/kafka-rebalance-protocol-static-membership/).

### Consumer is experiencing large Lag[Internal Link: ¶](\#consumer-is-experiencing-large-lag)

#### Description[Internal Link: ¶](\#id10)

A subset of consumer group or the whole consumer group is slow and the offset of the last consumed record from Kafka is far
This is known as [Internal Link: lag](../client/consumer_client.html#lag).

#### Impact[Internal Link: ¶](\#id11)

This means the consumer will be slow to catch up with the producer which might impact users experience especially for real-time applications.

#### Causes and Solutions[Internal Link: ¶](\#id12)

1. The consumer might be [Internal Link: lagging](../client/consumer_client.html#lag) on all partitions due to problem with committing the offset, or the fetch rate is slower than the produce rate.

_Solution:_

- Commit as soon as possible, and make sure that [Internal Link: committing offsets](../client/consumer_client.html#offset-commit) isn’t failing.
- Increase `fetch.min.bytes` to match the produce rate.
- If increasing `fetch.min.bytes` did not help, then tuning `max.poll.records` and `max.partition.fetch.bytes` can assist with reducing lags.
- Scale up the consumer group to the total number of topic partitions.

1. Consumer group is lagging only on a few partitions.
   This issue may be an indication that the topic(s) has uneven partitions.
   This creates an imbalanced consumer group, despite the fact that the number of partitions are equally distributed.

_Solution:_

- Scale up the consumer group horizontally by adding more instances to have one consumer assigned to one partition.
- Scale up the group vertically by considering implementing an advanced [Internal Link: multithreaded consumer](../client/consumer_settings.html#multi-threaded-consumer).
- Contact the producing side team asking them to fix the imbalanced partitions at the core.
- If the client is consuming from multiple topics, make sure they’re all balanced and have the same number of partitions.

### Consumer failed to commit offsets[Internal Link: ¶](\#consumer-failed-to-commit-offsets)

#### Description[Internal Link: ¶](\#id13)

Consumer may experience offset management issues, such as:

- committing offsets too frequently before finishing processing the records.
- not committing offsets at all, which could happen if the processing of records is stuck, or if the commit offsets is failing.
- committing incorrect offsets, which happens when consumed records have been removed due to hitting the topic [Internal Link: `retention limit`](../concepts.html#topic-retention) on the broker.
  In this case, the client will usually get an `Offset out of range` error.

#### Impact[Internal Link: ¶](\#id14)

Committed offset is what indicates to Kafka the next record to send to the consumer on the next poll.
If the client failed to commit the offset, the next time the `poll` function is called Kafka will return the old records again.
This leads to duplicate processing and lag build up.

#### Causes and Solutions[Internal Link: ¶](\#id15)

The cause of this might vary, and depends on the flow of the client data processing.

_Solution:_

Make sure client is applying the correct offset management technique that fits their usage, depending on the client needs.

- If using `commitSync()` or `commitAsync() ` make sure they don’t fail.
  If the client is subscribing to too many topic partitions, make sure the commit message doesn’t exceed `message.max.bytes` (most clusters are limited to 3MB).
- If the client has enabled `enabled.auto.commit` then make sure `auto.commit.interval.ms` is set to an appropriate value.
- Consider setting `auto.offset.reset` to an appropriate value (which is `earliest` or the `latest` offset) that fits the client’s needs.
  This config defines the behavior of the consumer when there is no committed position (which occurs when the group is first initialized), or when an offset is out of range.
- If the client wants to handle the `offset out of range` error manually, consider utilizing the `seek()` method to reset the consumer’s location inside a partition.


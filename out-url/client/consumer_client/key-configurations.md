---
title: "Key Configurations¶"
weight: 3
---

Here we list the key configuration that impact the consumption behavior.

| Configuration | Description | Default | Should I use the default value? |
| --- | --- | --- | --- |
| `fetch.min.bytes` | The default is 1 byte. Fetch requests will be completed as soon as a single byte is available to transfer to the client. This increases CPU usage for both the broker and consumers. | 1 | No. Increase the value for higher throughput. |
| `fetch.max.wait.ms` | This controls how long Kafka will wait, or the maximum latency on the consumer side for a fetch request to complete, before the sent data is available to a consumer | 500 | Maybe. Adjust the value for low volume topics or for a low latency requirement. |
| `max.partition.fetch.bytes` | This defines the record batch size returned to the consumer per partition. It is limited by the `message.max.bytes` configuration at the broker level. Therefore, updating it to values larger than `message.max.bytes` has no effect. | 1048588 | Yes. |
| `max.poll.records` | Controls the maximum number of records returned from a single poll request. Based on other configurations defined above, this can be increased to have a better throughput. | 500 | Maybe. Increasing the value assist with lag. |
| `max.poll.interval.ms` | This is a key property that defines the health of a consumer in a consumer group. This can effect how long Kafka will consider a client instance to be idle before reassign partitions to other instances in the consumer group. This is essential if processing polled records is taking more than `max.poll.interval.ms` or the topic has a few records that are larger than the others. | 300000 | Maybe. Increasing the value may reduce unnecessary rebalancing. |
| `session.timeout.ms` | The timeout used to detect client failures by the coordinator. Similar to `max.poll.interval.ms`, you may consider tuning this if your record processing is slow or the topic contains a few records larger than the others. Check the [Internal Link: Heartbeat](#heartbeat) concept for more details. | 45000 | Maybe. Increasing the value may reduce unnecessary rebalancing. |
| `heartbeat.interval.ms` | The expected time between [Internal Link: heartbeats](#heartbeat) to the consumer coordinator. This value must be less than `session.timeout.ms`, but typically should be set no higher than one third of that value. | 3000 | Maybe, if you need to adjust rebalance time. |
| `enable.auto.commit` | [Internal Link: Offset commits](#offset-commit) will happened periodically in the background. | `true` | No, consumer should disable this and commit manually after process the message. |

For more details about the consumer configuration check the [External Link: upstream Consumer Configs](https://docs.confluent.io/platform/current/installation/configuration/consumer-configs.html).

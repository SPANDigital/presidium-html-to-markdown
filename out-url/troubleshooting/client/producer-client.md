---
title: "Producer client¶"
weight: 5
---

Producer client determines which partition within a topic a record should be assigned to, and then sends the record to this partition.
Sometime the producer will face challenges due to some misconfigurations, unexpected infrastructure issue, or unexpected load.

Common challenges include:

### Producer failed to send record[Internal Link: ¶](\#producer-failed-to-send-record)

#### Description[Internal Link: ¶](\#id16)

When sending records, the producer client might face a temporary network failure, or a metadata mismatch, usually caused by changes in the topic partition leadership.

#### Impact[Internal Link: ¶](\#id17)

If the producer can’t send the record to Kafka it might cause an outage for the application, impact the consumption side, impact user experience, and in the worst case, data loss.

#### Causes and Solutions[Internal Link: ¶](\#id18)

1. Network issues: Kafka clients are designed intentionally to treat network errors as retryable as it’s most likely temporary.

_Solution:_

- Make sure to [Internal Link: test the connectivity](connectivity.html) with the brokers.
- Check if you need any extra network ACLs or configuration to connect to your Kafka brokers.
- If network connectivity isn’t the problem, then make sure `retires` is not set to a low value.
  Please follow our [Internal Link: recommendation for best delivery guarantee](../client/producer_settings.html#best-delivery-guarantee)

1. No authorization to produce to these topics: If the client isn’t authorized to produce to the topic then sending will fail with an authorization error.

_Solution:_

- Make sure [Internal Link: topic access](../concepts.html#client-identity-topic-access) and authorization are set up correctly.

1. Partition leadership isn’t correct: Kafka moves the leadership of partitions for multiple reasons, maintaining high availability.
   However sometimes this make a small distribution to the producer as the metadata on the producer side is out-of-date.

_Solution:_

- This is a retryable error, the producer will retry sending the record, next time to the correct leader.

### Missing data[Internal Link: ¶](\#missing-data)

#### Description[Internal Link: ¶](\#id19)

Data has been sent to one Kafka broker, but hasn’t been correctly written to all of a topic partition’s replicas.

#### Impact[Internal Link: ¶](\#id20)

The impact of data loss varies depend on the sensitivity of the data to the user’s experience.

#### Causes and Solutions[Internal Link: ¶](\#id21)

1. The client set `acks` to a value other than `ALL` which lowers the delivery guarantees.

_Solution:_

- Make sure to set `acks` to `ALL` which will force the producer to consider a write successful, only when all in-sync replicas receive the record, not just the leader.
  Please follow our [Internal Link: recommendation for best delivery guarantee](../client/producer_settings.html#best-delivery-guarantee)

### Throughput or latency issues[Internal Link: ¶](\#throughput-or-latency-issues)

#### Description[Internal Link: ¶](\#id22)

The producer can’t maintain the necessary message throughput, causing latency, or impacting the end-to-end latency for the use-case.
Usually this is due to misconfigurations at the client side, the broker being very busy, or the producer hitting a limit either on the Kafka broker side or the application’s deployment side.

#### Impact[Internal Link: ¶](\#id23)

The impact on end users will be reduced producer efficiency. This will vary from one use-case to another.

#### Causes and Solutions[Internal Link: ¶](\#id24)

1. Hit a limit: Producer might hit a limit on the Kafka broker side, such as a quota, or application’s deployment side.
   This usually happens with large batches, which can be detected using metrics on the client side, including:
   - `producer-metrics-produce-throttle-time-avg` which monitors the average time in milliseconds a request was throttled by the broker.
     This means you hit one of the quotas on the cluster side.
   - `producer-metrics-buffer-available-bytes` which shows the available memory. Once the buffer is full, further produce requests will be blocked.

_Solution:_

- Monitor producer performance parameters such as message rate, latency, throttling and buffer utilization.
- If `producer-metrics-buffer-available-bytes` metrics are low consider increasing `buffer.memory`. If this offers no improvement, then consider lowering `batch.size` and enable compression.

1. Unoptimized batching behavior: when it comes to throughput vs latency, batching is the first thing to look at.

_Solution:_
There is no one recipe that fits all, however here is list of things to consider:

- set `batch.size` and `linger.ms` to appropriate values that fit the clients requirements.
- Enable compression by setting `compression.type=zstd`
  More details can be found [Internal Link: here](../client/producer_settings.html#optimize-for-latency-vs-throughput)

1. If none of the these solutions offered improvement, please consider:

- Checking our recommendations on how to optimize for Latency vs Throughput [Internal Link: here](../client/producer_settings.html#optimize-for-latency-vs-throughput)
- Using asynchronous `send` with callbacks to increase speed by allowing the producer to keep delivering messages while waiting for acknowledgements.

---
title: "2. Partitioner¶"
weight: 4
---

The Partitioner decides which partition of the topic a record should be written to.
For example, suppose the record has a key (a.k.a compacted topic). In that case, a partition will be chosen based on a unique hash code for the key.
If a key is not presented Kafka will try to allocate the record into some partition in a round-robin fashion to achieve an even balance between partition sizes.
The Partitioner will do nothing if the Producer specified a partition in `ProducerRecord`.
This generally is not recommended for most applications as it will interfere with the partitioner and may lead to skewed partitions.

### 3\. Record accumulator[Internal Link: ¶](\#record-accumulator)

At this step, the Producer will group records into batches per partition based on `batch.size` configuration.
It’s important to remember the `send` method will return at this point and doesn’t need to wait for records to be actually
sent out as the Producer will keep accumulating the records till it reaches the desired batch size, or it reached `linger.ms`
which is the time to wait before sending data to brokers.

Remember that the buffer is stored locally, so depending on the producer configuration, the buffer may reach `buffer.memory`. In such a situation, the Producer will be blocked from invoking `send` until the producer frees buffer space by sending batches to the broker. To see the recommended configuration for `buffer.memory`, check our [Internal Link: recommendations](producer_settings.html#best-delivery-guarantee).

Note

By default, send is non-blocking, so if the Producer wants to block until message sent to the brokers,
it will need to invoke `send(record).get()` . By using `get()` the Producer is forced to wait for a reply from Kafka.
This method will throw an exception if the record is not sent successfully to Kafka.

### 4\. Compression[Internal Link: ¶](\#compression)

The stage at which the batch records will get compressed before it gets sent to the broker.

By default, compression is not enabled, but it’s recommended to enable it specially when producing to a multi-tenant Kafka cluster (like ACI Kafka).
All tenant should reduce their request size, network bandwidth usage as well as reducing their disk footprint.

The configuration `compression.type` can be set to `snappy`, `gzip`, `lz4`, or `zstd`.

| Type | CPU usage | Bandwidth usage | Compression ratio | Compression Speed |
| --- | --- | --- | --- | --- |
| `zstd` (recommended) | Moderate | Medium | Medium | Moderate |
| `snappy` | Moderate | Medium | Medium | Moderate |
| `gzip` | High | Low | High | Slow |
| `lz4` | Low | High | Low | Fast |

Check our recommendations for [Internal Link: compression](producer_settings.html#optimize-for-latency-vs-throughput).

### 5\. Sender Thread[Internal Link: ¶](\#sender-thread)

The sender thread is a background thread that groups the batches from all partitions by broker to open one connection with the broker.
The thread will send batches whenever `batch.size` or `linger.ms` is met first.

The thread will keep retrying to send the batches till it succeeds, or it reaches the retries configuration.

For more details about the Producer, check [External Link: Chapter 3 - Kafka: The Definitive Guide, 2nd Edition](https://learning.oreilly.com/library/view/kafka-the-definitive/9781491936153/ch03.html)


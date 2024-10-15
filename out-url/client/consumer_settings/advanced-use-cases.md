---
title: "Advanced use-cases¶"
weight: 6
---

### Exactly-Once consumer[Internal Link: ¶](\#exactly-once-consumer)

There are two sides for Exactly-Once semantics (EOS) in Kafka:

1. Avoiding duplication during data writing. This can be achieved by enabling the [Internal Link: Idempotent Producer](producer_settings.html#idempotent-producer-and-transactional-producer)).
2. Avoiding duplication during data processing. Even if the producer is idempotent the consumer could consume records twice if it did crash before committing the offset of these records.

A Kafka Consumer Client is unable to handle the second case. However, one popular pattern to achieve EOS consumer is to build a record handler that stores an id for the processed record and skip it if it is consumed again:

1. Keep track of the processed record in another DB (either store the records offset or use a UUID from the record value) even if your consumer didn’t commit the offset back to Kafka. As best practice, store this information in a key-value store, where the key is a unique identifier for the record and the value is the instance that handled the record.
2. Set up the consumer to skip the record that existed in DB.
3. Add metrics to count how many duplicates the client skipped.

### Multi-threaded consumer[Internal Link: ¶](\#multi-threaded-consumer)

**Usecases for multi-thread consumer**

- Scaling up vertically by distributing the load on threads on the same JVM instead of multiple nodes.
- Speeding up processing on the same JVM.

**Patterns**

There are few options in the open-source community. The following two are the most popular ones:

1. One thread per consumer clientThe application creates N threads to represent N consumer clients and connects to Kafka in this pattern. Each thread will be assigned to a subset of partitions.
2. Multi-Thread consumerIn this pattern, the application has a single main thread that contains one consumer client polling from N partitions.
   Once the consumer polls the records, they can be sent to the threads pool where they will be processed in parallel.

![](http://client/../_images/consumer_multi_threading.png)

Note

Multi-threading isn’t supported by the Kafka consumer. If you wish to achieve this functionality, you will need to either implement it or check with the open source community. The following links should provide a starting point:

- [External Link: Multi-Threaded Message Consumption with the Apache Kafka Consumer](https://www.confluent.io/blog/kafka-consumer-multi-threaded-messaging/)
- [External Link: Introducing the Confluent Parallel Consumer](https://www.confluent.io/blog/introducing-confluent-parallel-message-processing-client/)

You could also use an OpenSource Kafka consumer that supports parallel processing. Please contact the legal team before using any open source clients.

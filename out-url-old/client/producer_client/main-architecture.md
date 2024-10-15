---
title: "Main architecture¶"
weight: 2
---

![]({{%baseurl%}}//images/producer_arch.png)

The diagram above shows the producer flow when the application invokes `kafkaProducer.send(record)`.
As shown in the diagram, the record needs to go through five key stages/threads at the client-side before the record reaches the broker.
Understanding these stages is critical to tune and design any Kafka Producer.


---
title: "Exactly-Once¶"
weight: 7
---

EOS means the Producer will deliver all messages without any duplications.
This semantic can be achieved in Kafka by switching the Producer to an **Idempotent Producer**;
this can be achieved by configuring `enable.idempotence=true`, `acks=all` and `retries > 0`.
For more about configuring idempotent Producer, check [recommended configurations]({{%baseurl%}}/producer_settings.html#idempotent-producer-and-transactional-producer).

**Idempotent Producer** adds two fields to each record it sends

- `ProducerID`
- Sequence Number (Epoch)

So if an **Idempotent Producer** tries to send the same record with the same epoch, the broker will ignore it and send an acknowledgment response back to the Producer.

![]({{%baseurl%}}//images/producer-idempotent.png)

For more details on Idempotent Producer check [EOS blog](https://www.confluent.io/blog/exactly-once-semantics-are-possible-heres-how-apache-kafka-does-it/)).


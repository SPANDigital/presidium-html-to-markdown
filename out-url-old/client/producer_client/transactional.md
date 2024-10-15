---
title: "Transactional¶"
weight: 8
---

Transactions enable atomic writes to multiple Kafka topics and partitions where all the messages included in the transaction will be successfully written,
or none of them will be. This producer mode is known as **Transactional Producer**;
the first step to switching a Kafka Producer to a Transactional producer is to set `transactional.id` to a unique value.

The following steps are how a producer can send a transaction.

1. Initialize a transaction:This step is where the producer registers a transaction using `transactional.id`.
   As part of this step, Kafka cluster will elect a Kafka server (broker) to be the **Transactional Coordinator**.
   The coordinator will oversee the transactions state and the health of the producers.Keep in mind that each producer instance needs to have its unique `transactional.id`,
   check [recommended configurations]({{%baseurl%}}/producer_settings.html#idempotent-producer-and-transactional-producer) for how to set `transactional.id`.
2. Every time a producer needs to send a transaction, it must inform the coordinator that a new transaction starts shortly.
3. Send messages to multiple topics/partitions.
4. Commit/Abort Transaction:Once all messages get written successfully, the producer must inform the coordinator that the transaction finished successfully by committing the transactions.
   If something goes wrong, the producer must inform the coordinator that this transaction failed by sending an abort request.The transactional coordinator will add a marker to all the records as either committed or aborted in these transactions.
   Consumers use this marker to consume only the successful transactions and skip aborted transactions.

![]({{%baseurl%}}//images/transactional_producer.png)

For more details on [Transactions\
API](https://www.confluent.io/blog/transactions-apache-kafka/).


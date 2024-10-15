---
title: "Verifiable Producer/Consumer¶"
weight: 1
---

We have taken the Kafka `VerifiableProducer` and `VerifiableConsumer`
and extended them to provide richer, more complete data that can be used
for auditing. The idea is that we’ll be able to use them to verify
correctness while doing destructive system testing, among other uses.
There is an [improvement\
proposal]({{%baseurl%}}/https://github.pie.apple.com/pie/kafka-tools/blob/develop/improvement-proposals/001-verifiable-producer-consumer.md)
that describes the design.

This section describes a little bit of how to run the verifiable
producer and consumer.


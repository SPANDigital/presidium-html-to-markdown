---
title: "ACI Kafka Java client interoperability (ACI Kafka 0.15+)¶"
weight: 9
---

Prior to version 0.15 of the ACI Kafka client, all messages were
serialized inside of the [Internal Link: PieEnvelope](envelope.html). Now that Kafka
has added record headers, there is a plan to deprecate `PieEnvelope`,
detailed in [Internal Link: PieEnvelope
Deprecation](envelope.html#pieenvelope-deprecation-and-migration-details)
section of our docs.

Non-Java clients that don’t ever need to interoperate with Java clients
by interacting with the same topics don’t need to do anything special.

Non-Java producers that have Java-based consumers should produce a
record header called `pie.fmt` with an associated single-element byte
array value of `0x00` to indicate that the message is not serialized
using the envelope. This assumes the Java-based consumers are on
kafka-pie-client 0.15+. If the non-Java producer cannot be made to
include the `pie.fmt` header, Java consumers can use
`KafkaClientUtils.createProducer()` and set the consumer property
`pie.queue.deserializer.assume.payload.envelope=false` to ensure that
they do not try to deserialize an envelope.

Java producers that need to support non-Java consumers should upgrade to
0.15+, create a producer using `KafkaClientUtils.createProducer()`, and
set the producer property
`pie.queue.serializer.use.payload.envelope=false` to ensure that
messages are not serialized inside of an envelope.

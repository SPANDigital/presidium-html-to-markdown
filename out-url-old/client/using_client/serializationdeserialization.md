---
title: "Serialization/Deserialization¶"
weight: 4
---

In order to support the PIE envelope and end-to-end encryption, we use
our own value serializer and deserializer. In order to achieve this, we
replace the user-specified _value.serializer_ and
_value.deserializer_ configurations with our own:
`c.a.p.q.k.client.envelope.internal.Envelope{Deser,Ser}ializer`.
The user-specified serializer/deserializer will still be used in the
serialization/deserialization chain, and they will be propagated with
the `pie.queue.envelope.payload{deser,ser}ialize`{
properties.


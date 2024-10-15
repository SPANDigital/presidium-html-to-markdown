---
title: "1. Serializer¶"
weight: 3
---

Kafka writes data as bytes, so when a producer prepares a record to be sent, it first needs to serialize the record from an object into bytes.
This step is controlled by `key.serializer` and `value.serializer`. Both of these are vital config for any producer client.

Note

Serializer is also used to apply encryption and add more info into record headers. Both features are provided by ACI Kafka,
for more details check [PieEnvelope]({{< ref "client/envelope" >}}) and [encryption]({{< ref "client/security_settings" >}}).


---
title: "Migration steps¶"
weight: 2
---

In short, to make sure you aren’t affected, this is what you need to
do:

- Use kafka-pie-client version 0.15 or newer - [check the current version recommended for production]({{%baseurl%}}/../releases/index.html#recommended-release).
- Use `KafkaClientUtils.createConsumer` instead of
  `KafkaClientUtils.createPieConsumer`
- Use `KafkaClientUtils.createProducer` instead of
  `KafkaClientUtils.createPieProducer`
- If you are including `DeserializerConfigurationInterceptor` in
  `config.interceptors.classes` replace it with
  `PieDeserializerConfigurationInterceptor`
- If you are including `SerializerConfigurationInterceptor` in
  `config.interceptors.classes` replace it with
  `PieSerializerConfigurationInterceptor`
- If you are consuming messages that were produced with a version of
  kafka-pie-client earlier than 0.15 set
  `pie.queue.deserializer.assume.payload.envelope` to true
- If you are producing messages that will be consumed with a version
  of kafka-pie-client earlier than 0.15 set
  `pie.queue.serializer.use.payload.envelope` to true
- If you using PieEnvelope extensions get/set `pie.ext.X.Y` record
  headers for compatibility, or move completely to using other new
  headers

### PieEnvelope Deprecation and Migration details [¶](#pieenvelope-deprecation-and-migration-details "Link to this heading")

The ACI Kafka team has deprecated both the `PieEnvelope` Java object and
the [old binary format](#old-binary-format).

Historically, the ACI Kafka customers were recommended to construct
clients by using these two methods on KafkaClientUtils:

```
public static <K, V> KafkaConsumer<K, PieEnvelope<V>> createPieConsumer(Map<String, Object> config)
public static <K, V> KafkaProducer<K, PieEnvelope<V>> createPieProducer(Map<String, Object> config)

```

In 0.15, those methods were deprecated and two alternative methods -
which do not use the PieEnvelope - were introduced:

```
public static <K, V> KafkaConsumer<K, V> createConsumer(Map<String, Object> config)
public static <K, V> KafkaProducer<K, V> createProducer(Map<String, Object> config)

```

All these variants support both the [new binary\
format](#new-binary-format-introduced-in-kafka-pie-client-0-15) and the
[old binary format](#old-binary-format). The [pie.fmt record\
header](#two-binary-formats) is used to determine how the message is
deserialized. The old methods were retrofitted in 0.15 to support the
new format.

The old producer - created with `createPieProducer` \- chooses the
appropriate message format based on the value of the
`pie.queue.envelope.serializer.use.payload.envelope` property, which
defaults to true, using the [old binary format](#old-binary-format).

The new producer - created with `createProducer` \- chooses the
appropriate message format based on the value of the
`pie.queue.envelope.serializer.use.payload.envelope` property, which
defaults to false, using the [new binary\
format](#new-binary-format-introduced-in-kafka-pie-client-0-15).

Both old and new consumers examine the `pie.fmt` header to determine the
serialized format. When no [pie.fmt record header](#two-binary-formats)
header record is present in a message:

> - The old consumer - created with `createPieConsumer` \- uses the
> `pie.queue.envelope.deserializer.assume.payload.envelope`
> property, which defaults to true and indicates that the [old\
> binary format](#old-binary-format) is assumed.
> - The new consumer - created with `createConsumer` \- uses the
> `pie.queue.deserializer.assume.payload.envelope` property, which
> defaults to false and indicates that the [new binary\
> format](#new-binary-format-introduced-in-kafka-pie-client-0-15) is
> assumed.

Up until March 2021, Kaffe was overriding the default value from `false` to `true` on these properties, to minimize
disruption:

- `pie.queue.serializer.use.payload.envelope`
- `pie.queue.deserializer.assume.payload.envelope`

But these overrides were removed, making the default be `false` for both settings.

If you are consuming messages from or producing messages to clients
running a kafka-pie-client version earlier than 0.15 you need to configure these values
to be `true` to ensure compatibility.

When using the new methods, the `Map<ExtensionKey, ByteBufferOrValue>`
in PieEnvelope are available as a collection of record headers. The
names of these headers are `pie.ext.X.Y`. `X` is the [organization\
id](#organization-id) and `Y` is the [extension id](#extension-id).

A more detailed discussion of this topic can be found in
[QTIP-028]({{%baseurl%}}/https://github.pie.apple.com/pie/kafka-tools/blob/develop/improvement-proposals/028-envelope-deprecation.md).

### Envelope motivation [¶](#envelope-motivation "Link to this heading")

Kafka itself uses an envelope which contains: message key, message
payload, message timestamp, and the compression codec. In Apache Kafka
0.11, record headers were added to this list, but before that, it was
not generally extensible. The ACI Kafka envelope was introduced to
provide general extensibility before the advent of Kafka record headers,
and it utilizes the message payload portion of the Kafka envelope in
order to achieve this.

The envelope provided by ACI Kafka strives to be extensible and enables
users to add their own extensions to records. Extensions can
conceptually be compared to HTTP headers, with the exception that
extensions are single valued. The extensions can be set either by the
producing code, by registering producer interceptors, or by custom
serializers.

On the consumer side, ACI Kafka only deserializes extensions for which
deserializers are registered. This enables users to only deserialize the
extensions they are interested in and avoid adding unnecessary
dependencies.

Now that Kafka has introduced record headers, the envelope is no longer
required to achieve this functionality.

### Two binary formats [¶](#two-binary-formats "Link to this heading")

To support envelope deprecation, there are now two binary formats that
support the envelope. These are distinguished by the value of the
`pie.fmt` record header, which was introduced in kafka-pie-client
version 0.15. The header’s byte array values are as follows:

| Binary format | Header value |
| --- | --- |
| Old | \[0x01\] |
| New | \[0x00\] |

### New binary format (introduced in kafka-pie-client 0.15) [¶](#new-binary-format-introduced-in-kafka-pie-client-0-15 "Link to this heading")

The new binary format leverages Kafka’s native support for record headers.
The PIE Envelope headers are mapped to Kafka record headers.

The Kafka record payload corresponds to the `payload` entry from the
legacy format, as it would for any regular Kafka serializer.

Any envelope extensions present are serialized as Kafka record headers.
The header name is `pie.ext.<organization key>.<extension id>`, and the
header value corresponds to the `extension value` entry from the legacy
format.

As noted above, records serialized to the new format will include the
`pie.fmt` header with a byte array value of `[0x00]`.

### Old binary format [¶](#old-binary-format "Link to this heading")

Back when Kafka did not support record headers, the old binary format enabled users of
ACI Kafka to add headers to messages by wrapping the payload in a well defined envelope
with a header section. This means the payload will be wrapped inside an extra binary structure.

The legacy PieEnvelope binary format uses the standard [Kafka\
serialization]({{%baseurl%}}/https://github.pie.apple.com/pie/kafka-tools/blob/develop/pie-envelope/src/main/java/com/apple/pie/queue/kafka/envelope/EnvelopeDeserializer.java#L40-L74)
format. This enables us to leverage a tested, low-overhead and
dependency-free format.

It is versioned, thus enabling us to completely change the schema in the
future without breaking backwards compatibility.

It is made of an array of envelope extensions and a series of bytes for
the serialized payload. Each extension is made of an extension key and a
payload (the content of the extension). The extension key is made of 2
16-bit integers, the organization\_id and the extension\_id.

Here’s a format definition:

```
envelope magic number : 1 byte (0x7C)
envelope version      : 1 byte
extension count       : 4 bytes
extensions:           : N envelope_extension
payload length        : 4 bytes
payload               : V bytes

```

An envelope\_extension is defined as:

```
organization key: 2 bytes
extension id    : 2 bytes
extension length: 4 bytes
extension value : N bytes

```

For more information have a look at the class
[EnvelopeProtocol]({{%baseurl%}}/https://github.pie.apple.com/pie/kafka-tools/blob/develop/pie-envelope/src/main/java/com/apple/pie/queue/kafka/envelope/internal/EnvelopeProtocol.java).

As noted above, records serialized to the old format using
kafka-pie-client 0.15 or newer will include the `pie.fmt` header with a
byte array value of `[0x01]`.

### Organization id [¶](#organization-id "Link to this heading")

The link to the organization id is usually given out by ACI Kafka.

The current organization ids are:

- 0 ACI Kafka
- 1 AMP XMB
- 2 PIE Insights
- 3 PIE SpaaS
- 4 AMP Content Infrastructure
- 5 AMP Data
  - Extension 1: AMP Data Personalization
- 6 Apple Pay

Rather than creating new organizations, new use cases are encouraged to
use Kafka record headers in place of envelope extensions.

### Extension id [¶](#extension-id "Link to this heading")

Extension ids need only be unique within an organization, so it is up to
the organization to avoid conflicts.

### Using envelope extensions in the client [¶](#using-envelope-extensions-in-the-client "Link to this heading")

To be able to use an envelope extension you will need to give the client
a subclass of
[com.apple.pie.queue.kafka.envelope.extensions.ExtensionSerDeserializer]({{%baseurl%}}/https://github.pie.apple.com/pie/kafka-tools/blob/develop/pie-envelope/src/main/java/com/apple/pie/queue/kafka/envelope/extensions/ExtensionSerDeserializer.java).

To do so, a new configuration parameter was introduced in the clients,
`pie.queue.envelope.extension.serde.classes`, which takes a list of
extension serde. There is a handy utility function that you can use to
add extension serde classes to your configuration:
`EnvelopeConfig.addExtensionSerdesToConfig`.

Anywhere you access the envelope you can use
[getExtension(ExtensionKey)]({{%baseurl%}}/https://github.pie.apple.com/pie/kafka-tools/blob/develop/pie-envelope/src/main/java/com/apple/pie/queue/kafka/envelope/PieEnvelope.java)
to get access to an envelope extension.

For efficiency you can avoid deserializing too many extensions by only
providing the ones needed to your consumer. The other extensions will
still be accessible in their serialized form, thus enabling forwarding
them without incurring the deserialization cost.

### Chaining Serialization [¶](#chaining-serialization "Link to this heading")

To serialize objects you’ll need to go through multiple steps:

- Serialize the payload into bytes
- Encrypt the payload (if enabled)
- Serialize the envelope of the payload into bytes

To deserialize you need to do these action in the opposite order.

This is summarized is the following schema:

![]({{%baseurl%}}//images/serializer_chain.png)

We walk through the linked list created to execute the serializers in
the right order. Deserialization works in the opposite order.

Note

`SerializerConfigurationInterceptor` and
`DeserializerConfigurationInterceptor` add PieEnvelope
serializers/deserializers to your client on startup.
`PieSerializerConfigurationInterceptor` and
`PieDeserializerConfigurationInterceptor` add non-PieEnvelope
serializers/deserializers to your client on startup. Just set
`value.serializer`, `key.serializer`, `value.deserializer`, and
`key.deserializer`, just like you normally would.

### 0-copy optimization (optional) [¶](#copy-optimization-optional "Link to this heading")

If you don’t use encryption, you can further improve your consumer’s
memory consumption by moving the deserialization of your payload to your
client code. This configuration is not recommended for most use cases so
please evaluate carefully if you really need to use it!

To use this 0-copy optimization, encryption and decryption must be disabled, and
the standard Kafka `ByteBuffer` serializer and deserializer may be used.

```
# producer:
import org.apache.kafka.common.serialization.ByteBufferSerializer;
producerConfig.put(EncryptionSerializerConfigDef.ENCRYPT_MESSAGES_CONFIG, false);
producerConfig.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, ByteBufferSerializer.class);
# consumer:
import org.apache.kafka.common.serialization.ByteBufferDeserializer;
consumerConfig.put(DecryptionDeserializerConfigDef.TRY_TO_DECRYPT_MESSAGES_CONFIG, false);
consumerConfig.put(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, ByteBufferDeserializer.class);

```

### Advice on writing good extensions [¶](#advice-on-writing-good-extensions "Link to this heading")

- An extension should be self contained and for a specific purpose. (A
  specific interceptor shouldn’t need to access multiple extensions.)
- An extension should be versioned and backward compatible. In other
  words, remember that you might have old data in your queue forever
  (especially with compacted topics).

Note that Kafka record headers are now recommended in place of
extensions.

### Existing extensions [¶](#existing-extensions "Link to this heading")

We keep here links to extensions which already exist and those being
worked on. Do not hesitate to share yours.

Available:

- [avro\
  serialization](https://github.pie.apple.com/pie/kafka-extensions/tree/develop/avro-serializer)
  with
  [schema-store](https://github.pie.apple.com/amp-content-infra/schema-store)
  integration: Apache Avro serializer and deserializer that integrates
  your application with the Apple
  [schema-store](https://github.pie.apple.com/amp-content-infra/schema-store),
  supporting both compile-time and run-time schemas.
- [proto\
  serialization](https://github.pie.apple.com/pie/kafka-extensions/tree/develop/proto-2-serializer)
  with
  [schema-store](https://github.pie.apple.com/amp-content-infra/schema-store)
  integration. Google Protocol Buffers serializer and deserializer
  that integrates your application with the Apple
  [schema-store](https://github.pie.apple.com/amp-content-infra/schema-store).
- [message\
  auditing](https://github.pie.apple.com/pie/kafka-extensions/tree/develop/client-auditors).
  A set of interceptors and auditors to allow your application to
  submit auditing information tracking the Kafka records it has
  produced and consumed.

### Kafka-extensions repository [¶](#kafka-extensions-repository "Link to this heading")

[Kafka extensions](https://github.pie.apple.com/pie/kafka-extensions) is
a shared repository where we keep a set of interceptors and other tools
to enable clients to do more. Most of the extensions mentioned
previously are in kafka-extensions.

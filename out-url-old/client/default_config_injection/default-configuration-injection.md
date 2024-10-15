---
title: "What will Default Configuration Injection do?¶"
weight: 3
---

Default configuration injection will inject a set of recommended
properties so that it is easier to get started with the ACI Kafka
Client.

**Default configuration injection doesn’t override values for any**
**properties already set. You can still override any of the default**
**properties by setting those properties in your client configuration.**

Since producer and consumer clients require different kind of
configurations, we inject different default properties depending on
whether the client is a producer or a consumer client.

The default properties we currently inject for producer clients are:

```
security.protocol=SASL_SSL
sasl.mechanism=PLAIN
acks=-1
connections.max.idle.ms=180000

```

If encryption is enabled, we also inject these additional properties:

```
# For ACI Kafka client version >= 0.15:
  pie.queue.crypto.iris.connect = <value is environment specific>

# For ACI Kafka client version < 0.15:
  pie.queue.crypto.datastore.namespace = <value is environment specific>
  pie.queue.crypto.signature.verification.key = <value is environment specific>

```

The following producer properties are only injected if the
`pie.queue.kaffe.config.injection` property is set to `true`, even with
ACI Kafka Client version 0.14 and higher. These properties offer a
**high reliability guarantee** and good throughput. **Note: these**
**producer properties are included for compatibility with pre-0.14**
**versions of the ACI Kafka Client. If they are desired, please set them**
**explicitly, as their defaulting may be removed in the future.** :

```
retries=2147483647
max.in.flight.requests.per.connection=1
max.block.ms=2147483647

```

The default properties we inject for consumer clients are:

```
security.protocol=SASL_SSL
sasl.mechanism=PLAIN
pie.queue.metrics.blacklist.patterns=re:consumer-fetch-manager-metrics-.*\.records-lag.*
connections.max.idle.ms=180000

```

If decryption is enabled, we also inject these additional properties:

```
# For ACI Kafka client version >= 0.15:
  pie.queue.crypto.iris.connect = <value is environment specific>

# For ACI Kafka client version < 0.15:
  pie.queue.crypto.datastore.namespace = <value is environment specific>
  pie.queue.crypto.consumer.id = <same as the pie.queue.kaffe.client.id property>

```

---
title: "ACI Kafka 0.18 release announcement¶"
weight: 1
---

The ACI Kafka team is excited to announce the release of ACI Kafka 0.18.
This release is based on Apache Kafka 2.4.2 and contains a number of new
features, fixes, and improvements.

For the latest documentation, including an overview of ACI Kafka and the
ACI Kafka Client, please see our
[documentation](https://docs.aci.apple.com/kafka).

**Upgrade Instructions**

The ACI Kafka 0.18.2 client is available on
[Artifactory]({{%baseurl%}}/https://artifacts.apple.com/pie-release-local/com/apple/pie/queue/kafka-pie-client/0.18.2)
and can be configured as a [Gradle\
Dependency]({{%baseurl%}}/../client/quick_start.html#gradle-dependency).
If you are using older versions of ACI Kafka, please upgrade to the
latest version, 0.18.2.

**Dropped support for Java 1.8**

This is first release of ACI Kafka that is not compatible with JDK8. Use
of JDK8 is
[deprecated]({{%baseurl%}}/https://java.apple.com/applejdk.html#AppleJDKVersions), if
are still using JDK8 you should move to JDK11. If you really need JDK8
compatibility please use the latest
[0.17]({{%baseurl%}}/https://artifacts.apple.com/pie-release-local/com/apple/pie/queue/kafka-pie-client/0.17.6)
release.

**Notes for Kafka users that does not use the factory methods in**
**KafkaClientUtils**

For users not using `KafkaClientUtils.createConsumer()` and
`KafkaClientUtils.createProducer()` there are some configuration changes
related to TLS configuration that are needed for this new release to
work. Previously, ACI Kafka made a change to upstream kafka to enable
`ssl.truststore.location` values that would start with the string
classpath: This enabled a convenient way of shipping the Apple Corporate
cert bundle without needing to manage it in a separate file. This
functionality has been replaced by some new extensions to Apache Kafka.
To have the ACI Kafka client be configured to trust the default Apple
Corporate CA certificates, use the following configuration:

```
security.providers=com.apple.pie.queue.kafka.client.CaCertPemKeyStoreProviderCreator,com.apple.pie.queue.kafka.client.CorpCertKeyStoreProviderCreator
ssl.truststore.type=AppleCorpCert
ssl.truststore.location=/dev/null
ssl.truststore.password=ignored

```

To use another PEM formatted cert bundle, use the following settings:

```
security.providers=com.apple.pie.queue.kafka.client.CaCertPemKeyStoreProviderCreator,com.apple.pie.queue.kafka.client.CorpCertKeyStoreProviderCreator
ssl.truststore.type=PEM
ssl.truststore.location=/path/to/pem/encoded/ca/cert/bundle.pem
ssl.truststore.password=ignored

```

As stated above, when using the `KafkaClientUtils` static methods, or
the `KaffeConfigurationInterceptor` directly, these configuration
changes will be handled transparently.

When specifying `config.interceptor.classes`, order is now more
important. You should ideally specify in order:

- Consumer or Producer ConfigurationInterceptor (e.g.
  `com.apple.pie.queue.kafka.client.configinterceptors.ConsumerConfigurationInterceptor`)
- Kaffe interceptor
  `com.apple.pie.queue.kafka.client.configinterceptors.KaffeConfigurationInterceptor`
- The Deserializer or Serializer, matching Producer or Consumer:
  `com.apple.pie.queue.kafka.client.configinterceptors.PieDeserializerConfigurationInterceptor`

as well as any additional interceptors you may require/have written
applied ideally after the Kaffe interceptor (unless functionally they
are required before).

**Other notable changes**

- Reduced rebalances, more stability in Consumer Groups with [Static\
  Membership]({{%baseurl%}}/https://kafka.apache.org/25/documentation.html#static_membership)
  and incremental cooperative rebalancing to the clients’ group
  protocol.
- The default consumer group id has been changed from the empty string
  (“”) to null. Consumers who use the new default group id will not
  be able to subscribe to topics, and fetch or commit offsets. The
  empty string as consumer group id is deprecated but will be
  supported until a future major release. Old clients that rely on the
  empty string group id will now have to explicitly provide it as part
  of their consumer config. For more information see
  [KIP-289](https://cwiki.apache.org/confluence/display/KAFKA/KIP-289%3A+Improve+the+default+group+id+behavior+in+KafkaConsumer).
- A new INVALID\_RECORD error for producer callers to be better
  informed about the root cause why their sent records were failed.
- The blocking `KafkaConsumer#committed` methods have been extended to
  allow a list of partitions as input parameters rather than a single
  partition. It enables fewer request/response iterations between
  clients and brokers fetching for the committed offsets for the
  consumer group. The old overloaded functions are deprecated and we
  would recommend users to make their code changes to leverage the new
  methods (details can be found in
  [KIP-520](https://cwiki.apache.org/confluence/display/KAFKA/KIP-520%3A+Add+overloaded+Consumer%23committed+for+batching+partitions)).
- Listing committed methods with
  `KafkaConsumer#committed(Set<TopicPartition>)` can now return null
  values, for topic partitions that don’t have offsets yet. The old
  behavior was to omit these topic partitions.
- Records for specific topic with null keys and no assigned partition
  will be sent to the same partition until the batch is ready to be
  sent, using a sticky partitioning strategy in the
  DefaultPartitioner. When a new batch is created, a new partition is
  chosen. This decreases latency to produce, but it may result in
  uneven distribution of records across partitions in edge cases.
  Generally users will not be impacted, but this difference may be
  noticeable in tests and other situations producing records for a
  very short amount of time.

For additional changes and more details please consult the notable
changes for
[2.2]({{%baseurl%}}/https://kafka.apache.org/documentation.html#upgrade_220_notable),
[2.3]({{%baseurl%}}/https://kafka.apache.org/documentation.html#upgrade_230_notable) and
[2.4]({{%baseurl%}}/https://kafka.apache.org/documentation.html#upgrade_240_notable) in
the Apache Kafka documentation.


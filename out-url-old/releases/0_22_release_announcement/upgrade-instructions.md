---
title: "Upgrade Instructions¶"
weight: 2
---

The ACI Kafka 0.22 client is available on [Artifactory](https://artifacts.apple.com/pie-release-local/com/apple/pie/queue/kafka-pie-client).
If you are using an older version of the ACI Kafka client, please upgrade to the latest version.

This release brings no changes in the envelope deprecation process since the previous release.
Support for consuming messages in the envelope format is still available.
This continues to guarantee interoperability with older clients producing messages in the envelope format.
However, and same as the previous release, producing messages in this old format is no longer possible.

Please make sure you have fully migrated away from using `PieEnvelope` and the envelope wire format.
For more details, please refer to the documentation on the [`PieEnvelope` deprecation]({{< ref "client/envelope" >}}).

**Notable upstream changes**

- The default value for `acks` is set to `-1` (all). See [KIP-697](https://cwiki.apache.org/confluence/display/KAFKA/KIP-679%3A+Producer+will+enable+the+strongest+delivery+guarantee+by+default) for more details.
- The default value for `enable.idempotence` is set to `true` for ACI Kafka versions >= 0.22.13 if
  `acks=-1`, `max.in.flight.requests.per.connection <= 5` and `retries > 0`. See [KIP-697](https://cwiki.apache.org/confluence/display/KAFKA/KIP-679%3A+Producer+will+enable+the+strongest+delivery+guarantee+by+default) and
  [KAFKA-13598](https://issues.apache.org/jira/browse/KAFKA-13598) for more details.
- The default value for the consumer configuration session.timeout.ms was increased from 10s to 45s.
  See [KIP-735](https://cwiki.apache.org/confluence/display/KAFKA/KIP-735%3A+Increase+default+consumer+session+timeout) for more details.
- The `Producer#sendOffsetsToTransaction(Map offsets, String consumerGroupId)` method has been deprecated.
  Please `useProducer#sendOffsetsToTransaction(Map offsets, ConsumerGroupMetadata metadata)` instead, where the
  `ConsumerGroupMetadata` can be retrieved via `KafkaConsumer#groupMetadata()` for stronger semantics.
  See [KIP-732](https://cwiki.apache.org/confluence/x/zJONCg) for more details.
- A number of deprecated classes, methods and configurations have been removed:
  - The config value `default` for the client config `client.dns.lookup` has been removed.
    In the unlikely event that you set this config explicitly, we recommend leaving the config unset
    ( `use_all_dns_ips` is used by default).
  - The `ExtendedDeserializer` and `ExtendedSerializer` classes have been removed.
    Please use `Deserializer` and `Serializer` instead.
  - The `close(long, TimeUnit)` method was removed from the producer, consumer and admin client.
    Please use `close(Duration)`.
  - The `ConsumerConfig.addDeserializerToConfig` and `ProducerConfig.addSerializerToConfig` methods were removed.
    These methods were not intended to be public API and there is no replacement.
  - The default `partition.assignment.strategy` is changed to “\[ `RangeAssignor`, `CooperativeStickyAssignor`\]”,
    which will use the `RangeAssignor` by default, but allows upgrading to the `CooperativeStickyAssignor` with just a
    single rolling bounce that removes the RangeAssignor from the list. Please check the client upgrade path guide
    [here](https://cwiki.apache.org/confluence/display/KAFKA/KIP-429:+Kafka+Consumer+Incremental+Rebalance+Protocol#KIP429:KafkaConsumerIncrementalRebalanceProtocol-Consumer).
    for more detail.
  - The `org.apache.kafka.clients.consumer.internals.PartitionAssignor` class has been removed.
    Please use `org.apache.kafka.clients.consumer.ConsumerPartitionAssignor` instead.

For general overview on what’s new in Kafka 3.0.0 you might be interested in [this post](https://blogs.apache.org/kafka/entry/what-s-new-in-apache6) in the Apache Kafka
blog.

For additional changes and more details please consult the notable changes for version
[3.0]({{%baseurl%}}/https://kafka.apache.org/documentation.html#upgrade_300_notable) in the Apache Kafka documentation or the [release notes for Kafka 3.0]({{%baseurl%}}/https://kafka.apache.org/downloads.html#3.0.0).

For more information on the ACI Kafka service offering, including
features, caveats, and project road map, visit our [overview\
section]({{< ref "overview" >}}).

You can [reach out to the ACI Kafka team]({{< ref "contact" >}}) with
questions on this release or general inquiries related to ACI Kafka.

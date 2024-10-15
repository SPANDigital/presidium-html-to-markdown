---
title: "Notable changes¶"
weight: 2
---

The ACI Kafka 0.9 release includes new features available in Apache
Kafka 0.11 and 1.0. We’ve also added a new feature that makes
configuring clients even easier. See [Default Configuration\
Injection]({{< ref "client/default_config_injection" >}}) to learn more.

**Summary of changes**

- New Apache Kafka 0.11 and 1.0 features: Exactly Once Semantics
  (EOS), Record Headers, JBOD failure resiliency
- Kaffe client can now be configured to inject default configuration
  for producer and consumer clients, and messages are logged if a
  client version is deprecated or has critical fixes available
- Deprecation of old Consumer client
- Customer-facing and internal services are now deployed on Compute
  and receive weekly upgrades in all environments

**ACI Kafka**

- [QTIP-021 - Kaffe Safe Defaults\
  Injection]({{%baseurl%}}/https://github.pie.apple.com/pie/kafka-tools/blob/develop/improvement-proposals/021-kaffe-safe-defaults-injection.md)
- [QTIP-022 - Client Upgrade\
  Notification]({{%baseurl%}}/https://github.pie.apple.com/pie/kafka-tools/blob/develop/improvement-proposals/022-client-upgrade-notification.md)
- Critical fix for [rdar://problem/40723727](rdar://problem/40723727) Kafka client can spawn
  threads when refreshing keys

**Apache Kafka**

- [KIP-82 - Add Record\
  Headers](https://cwiki.apache.org/confluence/display/KAFKA/KIP-82+-+Add+Record+Headers)
- [KIP-98 - Exactly Once Delivery and Transactional\
  Messaging](https://cwiki.apache.org/confluence/display/KAFKA/KIP-98+-+Exactly+Once+Delivery+and+Transactional+Messaging)
- [KIP-109 - Old Consumer\
  Deprecation](https://cwiki.apache.org/confluence/display/KAFKA/KIP-109%3A+Old+Consumer+Deprecation)
- [KIP-119 - Drop Support for Scala 2.10 in Kafka\
  0.11]({{%baseurl%}}/https://cwiki.apache.org/confluence/display/KAFKA/KIP-119%3A+Drop+Support+for+Scala+2.10+in+Kafka+0.11)
- [KIP-152 - Improve diagnostics for SASL authentication\
  failures](https://cwiki.apache.org/confluence/display/KAFKA/KIP-152+-+Improve+diagnostics+for+SASL+authentication+failures)
- Producer configs `block.on.buffer.full`, `metadata.fetch.timeout.ms`
  and `timeout.ms` have been removed. They were initially deprecated
  in Kafka 0.9.0.0.
- `ProducerRecord` and `ConsumerRecord` expose the new Headers API via
  Headers headers() method call.
- `ExtendedSerializer` and `ExtendedDeserializer` interfaces are
  introduced to support serialization and deserialization for headers.
  Headers will be ignored if the configured serializer and
  deserializer are not the above classes.
- A new config, `group.initial.rebalance.delay.ms`, was introduced.
  This config specifies the time, in milliseconds, that the
  `GroupCoordinator` will delay the initial consumer rebalance. The
  rebalance will be further delayed by the value of
  `group.initial.rebalance.delay.ms` as new members join the group, up
  to a maximum of `max.poll.interval.ms`. The default value for this
  is 3 seconds. During development and testing it might be desirable
  to set this to 0 in order to not delay test execution time.
- For offset commit failures in the Java consumer’s commitAsync APIs,
  we no longer expose the underlying cause when instances of
  `RetriableCommitFailedException` are passed to the commit callback.
  See [KAFKA-5052](https://issues.apache.org/jira/browse/KAFKA-5052)
  for more detail.
- Authentication failures are now reported to clients as one of the
  subclasses of `AuthenticationException`. No retries will be
  performed if a client connection fails authentication.
- When using an `Authorizer` and a user doesn’t have required
  permissions on a topic, the broker will return
  `TOPIC_AUTHORIZATION_FAILED` errors to requests irrespective of
  topic existence on broker. If the user have required permissions and
  the topic doesn’t exists, then the `UNKNOWN_TOPIC_OR_PARTITION`
  error code will be returned.
- Other notable changes in [Apache Kafka 0.11\
  Changes]({{%baseurl%}}/https://kafka.apache.org/0110/documentation.html#upgrade_1100_notable)
  and [1.0\
  Changes]({{%baseurl%}}/https://kafka.apache.org/11/documentation.html#upgrade)
- Full release notes [Apache Kafka 0.11 Release\
  Notes]({{%baseurl%}}/https://kafka.apache.org/downloads.html#0.11.0.0)
  and [1.0 Release\
  Notes]({{%baseurl%}}/https://kafka.apache.org/downloads.html#1.0.0)

**ACI Kafka Invisible Updates**

- [QTIP-023 - Getting to Continuous\
  Delivery]({{%baseurl%}}/https://github.pie.apple.com/pie/kafka-tools/blob/develop/improvement-proposals/023-continuous-delivery.md)
- rdar://problem/36478792 \[RTLF\] Migrate Kaffe to PIE


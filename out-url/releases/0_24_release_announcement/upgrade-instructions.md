---
title: "Upgrade Instructions¶"
weight: 2
---

The ACI Kafka 0.24 client is available on [External Link: Artifactory](https://artifacts.apple.com/pie-release-local/com/apple/pie/queue/kafka-pie-client).
If you are using an older version of the ACI Kafka client, please upgrade to the latest version.

Although this release brings no changes in the envelope deprecation process, we would like to remind you to migrate your application away from the envelope if you haven’t done so yet.
Support for consuming messages in the envelope format is still available in this version.
This continues to guarantee interoperability with older clients producing messages in the envelope format.
However, as with the previous release, producing messages in this old format is no longer possible.

Please make sure you have fully migrated away from using `PieEnvelope` and the envelope wire format.
For more details, please refer to the documentation on the [Internal Link: `PieEnvelope` deprecation](../client/envelope.html).

This release includes a number of improvements:

- [External Link: KIP-704](https://cwiki.apache.org/confluence/display/KAFKA/KIP-704%3A+Send+a+hint+to+the+partition+leader+to+recover+the+partition) should allow for smoother leadership failovers by having the controller
  indicate to new partition leaders whether they’re being elected uncleanly.
- [External Link: KAFKA-7572](https://issues.apache.org/jira/browse/KAFKA-7572) Producer should not send requests with negative partition id
- [External Link: KAFKA-9279](https://issues.apache.org/jira/browse/KAFKA-9279) Silent data loss in Kafka producer
- [External Link: KAFKA-13418](https://issues.apache.org/jira/browse/KAFKA-13418) Brokers disconnect intermittently with TLS1.3 - We have backported this fix to our internal builds of 3.0 and 3.1
- [External Link: KAFKA-13310](https://issues.apache.org/jira/browse/KAFKA-13310) KafkaConsumer cannot jump out of the poll method, and the consumer is blocked in the ConsumerCoordinator method maybeAutoCommitOffsetsSync(Timer timer). CPU and traffic on broker side increases sharply
- [External Link: KAFKA-12841](https://issues.apache.org/jira/browse/KAFKA-12841) NPE from the provided metadata in client callback in case of ApiException
- [External Link: KAFKA-12256](https://issues.apache.org/jira/browse/KAFKA-12256) auto commit causes delays due to retryable UNKNOWN\_TOPIC\_OR\_PARTITION
- [External Link: KAFKA-13435](https://issues.apache.org/jira/browse/KAFKA-13435) Static membership protocol should let the leader skip assignment (KIP-814)
- [External Link: KAFKA-13782](https://issues.apache.org/jira/browse/KAFKA-13782) Producer may fail to add the correct partition to transaction
- [External Link: KAFKA-13794](https://issues.apache.org/jira/browse/KAFKA-13794) Producer batch lost silently in TransactionManager

For a full list of changes please consult the [External Link: release notes for Kafka 3.2](https://archive.apache.org/dist/kafka/3.2.0/RELEASE_NOTES.html).
Or, for a general overview on what’s new in Kafka 3.2 you might be interested in [External Link: this post](https://blogs.apache.org/kafka/entry/what-s-new-in-apache8) in the Apache Kafka
blog.

For more information on the ACI Kafka service offering, including
features, caveats, and project road map, visit our [Internal Link: overview section](../overview.html).

You can [Internal Link: reach out to the ACI Kafka team](../contact.html) with
questions on this release or general inquiries related to ACI Kafka.

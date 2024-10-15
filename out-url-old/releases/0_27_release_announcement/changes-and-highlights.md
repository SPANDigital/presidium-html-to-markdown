---
title: "Changes and highlights¶"
weight: 3
---

Improvements and bug fixes since 0.26 (Kafka 3.4):

- [KIP-866](https://cwiki.apache.org/confluence/display/KAFKA/KIP-866+ZooKeeper+to+KRaft+Migration) early access feature to migrate ZK clusters to KRaft, which ACI Kafka will start testing soon.
- [KIP-887](https://cwiki.apache.org/confluence/display/KAFKA/KIP-887%3A+Add+ConfigProvider+to+make+use+of+environment+variables) introduces a new `ConfigProvider` implementation, `EnvVarConfigProvider`,
  which allows retrieving configuration from environment variables.
- [KIP-903](https://cwiki.apache.org/confluence/display/KAFKA/KIP-903%3A+Replicas+with+stale+broker+epoch+should+not+be+allowed+to+join+the+ISR): Replicas with stale broker epoch should not be allowed to join the ISR.
  Addresses a limitation of the replication protocol which could lead to data loss when
  a broker fails while another one had an unclean shutdown.
- [KIP-881](https://cwiki.apache.org/confluence/display/KAFKA/KIP-881%3A+Rack-aware+Partition+Assignment+for+Kafka+Consumers) Rack-aware Partition Assignment for Kafka Consumers — note however that this feature
  is not supported by ACI Kafka, as our rack configuration is currently optimized for availability.

For a full list of changes please consult the [release notes for Kafka 3.5](https://archive.apache.org/dist/kafka/3.5.0/RELEASE_NOTES.html).
Or, for a general overview on what’s new in Kafka 3.5 you might be interested in [this post](https://kafka.apache.org/blog#apache_kafka_350_release_announcement) in the Apache Kafka
blog.

For more information on the ACI Kafka service offering, including
features, caveats, and project road map, visit our [overview section]({{< ref "overview" >}}).

You can [reach out to the ACI Kafka team]({{< ref "contact" >}}) with
questions on this release or general inquiries related to ACI Kafka.

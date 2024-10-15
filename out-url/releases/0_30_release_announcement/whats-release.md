---
title: "What’s new in this release?¶"
weight: 3
---

For a complete list of changes, please consult the [External Link: release notes for Kafka 3.8](https://downloads.apache.org/kafka/3.8.0/RELEASE_NOTES.html).

For a general overview of what’s new in Kafka 3.8, read [External Link: the 3.8 announcement](https://kafka.apache.org/blog#apache_kafka_380_release_announcement) in the Apache Kafka
blog.

These are some of the highlights in the changes included in this release.

### Specifying the compression level[Internal Link: ¶](\#specifying-the-compression-level)

From this version on, it is possible to specify the compression level for
gzip, lz4 and zstd compression types using the configuration property
`compression.[gzip|lz4|zstd].level`. See [External Link: KIP-390](https://cwiki.apache.org/confluence/display/KAFKA/KIP-390%3A+Support+Compression+Level) for more details.

### Client rebootstrap[Internal Link: ¶](\#client-rebootstrap)

From this version onwards, Kafka clients are able to repeat the bootstrap
process when updating metadata if none of the known brokers are available.
See [External Link: KIP-899](https://cwiki.apache.org/confluence/display/KAFKA/KIP-899%3A+Allow+producer+and+consumer+clients+to+rebootstrap) for more details.

### JBOD support in Kraft[Internal Link: ¶](\#jbod-support-in-kraft)

JBOD in KRaft is no longer considered an early access feature.
This feature will enable us to keep ASE Kafka up to date, through the
KRaft transition for any clusters still supported by Bare Metal.
See [External Link: KIP-858](https://cwiki.apache.org/confluence/display/KAFKA/KIP-858%3A+Handle+JBOD+broker+disk+failure+in+KRaft) for more details.

### Next Generation of Consumer Rebalance Protocol[Internal Link: ¶](\#next-generation-of-consumer-rebalance-protocol)

This feature is available as preview in this version. ASE Kafka recommendations
and documentation on how to adopt this feature will be made available soon.
For now, please refer to [External Link: KIP-848](https://cwiki.apache.org/confluence/display/KAFKA/KIP-848%3A+The+Next+Generation+of+the+Consumer+Rebalance+Protocol) for further information.

* * *

Visit the [Internal Link: overview section](../overview.html) for more information about
our ASE Kafka service offering, including features, concepts, and project road map.

For questions about this release or general inquiries related to ASE Kafka, [Internal Link: contact the ASE Kafka team](../contact.html).

---
title: "What are available metrics and labels?¶"
weight: 5
---

### Topic usage [¶](#topic-usage "Link to this heading")

| Metric name | Description | Labels |
| --- | --- | --- |
| `kafka_log_Log_LogEndOffset` | The offset of the last stored message on disk. | `_wc_`, `_ns_`, `namespace`, `topic_name`, `partition`, `broker` |
| `kafka_log_Log_LogStartOffset` | The offset of the earliest stored message on disk. | `_wc_`, `_ns_`, `namespace`, `topic_name`, `partition`, `broker` |
| `kafka_log_Log_Size` | The partition’s size stored on each broker in bytes, including replicas. | `_wc_`, `_ns_`, `namespace`, `topic_name`, `partition`, `broker` |
| `kafka_cluster_Partition_ReplicasCount` | Total number of replicas (default: 4). | `_wc_`, `_ns_`, `namespace`, `topic_name`, `partition`, `broker` |

### Topic health [¶](#topic-health "Link to this heading")

| Metric name | Description | Labels |
| --- | --- | --- |
| `kafka_cluster_Partition_InSyncReplicasCount` | Number of in-sync replicas. If the value of this metrics is equal to `4`, then your partition is healthy. | `_wc_`, `_ns_`, `namespace`, `topic_name`, `partition`, `broker` |
| `kafka_cluster_Partition_UnderReplicated` | If this value isn’t zero, then there are replicas of this partition that struggle to catchup. Kafka will attempt to recover from this state, and it is expected to observe replicas falling out of sync. | `_wc_`, `_ns_`, `namespace`, `topic_name`, `partition`, `broker` |
| `kafka_cluster_Partition_AtMinIsr` | If the value isn’t zero, then only 2 partition replicas are in-sync. The partition is at risk of being marked as offline if another replica falls out of sync. | `_wc_`, `_ns_`, `namespace`, `topic_name`, `partition`, `broker` |
| `kafka_cluster_Partition_UnderMinIsr` | If the value isn’t zero, your partition is offline and all produce and consume requests to it will fail. | `_wc_`, `_ns_`, `namespace`, `topic_name`, `partition`, `broker` |

### Produce activities [¶](#produce-activities "Link to this heading")

| Metric name | Description | Labels |
| --- | --- | --- |
| `kafka_server_BrokerTopicMetrics_MessagesInPerSec` | Aggregate of incoming message rate per topic. | `_wc_`, `_ns_`, `namespace`, `topic_name, ` broker\` |
| `kafka_server_BrokerTopicMetrics_BytesInPerSec` | Aggregate of incoming bytes rate per topic. | `_wc_`, `_ns_`, `namespace`, `topic_name`, `broker` |
| `kafka_server_BrokerTopicMetrics_TotalProduceRequestsPerSec` | Produce request rate per topic. | `_wc_`, `_ns_`, `namespace`, `topic_name`, `broker` |
| `kafka_server_BrokerTopicMetrics_BytesRejectedPerSec` | Rejected produce request rate per topic. | `_wc_`, `_ns_`, `namespace`, `topic_name`, `broker` |
| `kafka_server_BrokerTopicMetrics_NoKeyCompactedTopicRecordsPerSec` | Message validation failure rate due to no key specified for compacted topic. | `_wc_`, `_ns_`, `namespace`, `topic_name`, `broker` |
| `Produce_throttle_time` | Produce bandwidth quota metrics per user on each broker. | `_wc_`, `_ns_`, `customer_user`, `broker` |
| `Produce_byte_rate` | Produce rate per user on each broker. | `_wc_`, `_ns_`, `customer_user`, `broker` |

### Consume activities [¶](#consume-activities "Link to this heading")

| Metric name | Description | Labels |
| --- | --- | --- |
| `kafka_server_BrokerTopicMetrics_BytesOutPerSec` | Aggregate of outgoing bytes rate per topic. | `_wc_`, `_ns_`, `namespace`, `topic_name`, `broker` |
| `kafka_server_BrokerTopicMetrics_TotalFetchRequestsPerSec` | Consume request rate per topic. | `_wc_`, `_ns_`, `namespace`, `topic_name`, `broker` |
| `kafka_server_BrokerTopicMetrics_FailedFetchRequestsPerSec` | Failed consume request rate per topic. | `_wc_`, `_ns_`, `namespace`, `topic_name`, `broker` |
| `Fetch_byte_rate` | Consume rate per user on each broker. | `_wc_`, `_ns_`, `customer_user`, `broker` |
| `Fetch_throttle_time` | Consume bandwidth quota metrics per user on each broker. | `_wc_`, `_ns_`, `customer_user`, `broker` |

### Synthetic Metrics [¶](#synthetic-metrics "Link to this heading")

A set of metrics emitted by a synthetic monitoring tool we have developed which continually produces and consumes from the cluster in small bursts, capturing and recording produce or consume failures, as well as the maximum round trip latency.

| Metric name | Description | Labels |
| --- | --- | --- |
| `kafka_synthetic_monitor_round_trip_record_produce_failures_total` | Count of records which were created but failed to be sent by the producer. | `_wc_`, `_ns_`, `namespace`, `topic_name`, `partition`, `usecase` |
| `kafka_synthetic_monitor_round_trip_record_consume_failures_total` | Count of records which were expected to be consumed by the consumer but were not. | `_wc_`, `_ns_`, `namespace`, `topic_name`, `partition`, `usecase` |
| `kafka_synthetic_monitor_round_trip_latency_ms_max` | The maximum time (in a moving 30 second window) taken to send a record to the topic and subsequently consume the same record. | `_wc_`, `_ns_`, `namespace`, `topic_name`, `partition`, `usecase` |
| `kafka_synthetic_monitor_monitoring_failures_total` | Count of failures which occurred when producing and consuming from the synthetic monitoring topic. | `_wc_`, `_ns_`, `usecase` |
| `kafka_synthetic_monitor_preparation_failures_total` | Count of failures which occurred when preparing the the synthetic monitoring topic for use. | `_wc_`, `_ns_`, `usecase` |
| `kafka_synthetic_monitor_runtime_failures_total` | Count of unexpected failures which occurred while monitoring this cluster. | `_wc_`, `_ns_`, `usecase` |

### Others [¶](#others "Link to this heading")

| Metric name | Description | Labels |
| --- | --- | --- |
| `Request_request_time` | Request rate per user on each broker. | `_wc_`, `_ns_`, `customer_user`, `broker` |
| `Request_throttle_time` | Request quota metric per user on each broker. | `_wc_`, `_ns_`, `customer_user`, `broker` |

### Lag metrics [¶](#lag-metrics "Link to this heading")

Details for lag metrics are found [in the dedicated page]({{%baseurl%}}/../client/lag_metrics.html#what-are-available-metrics-and-labels).

### Labels [¶](#labels "Link to this heading")

Each metric may include one of the following labels:

- `_wc_` : Mosaic workspace. This is a constant value set to `aci-kafka` and cannot be changed.
- `_ns_`: Mosaic namespace. The value for `_ns_` for the above metrics is `kafka-broker-<cluster-name-with-hyphens>-piekafka`
- `namespace`: ACI Kafka [namespace]({{%baseurl%}}/../concepts.html#namespace).
- `topic_name`: ACI Kafka topic name. This is **not** the fully qualified name.
- `partition`: The Kafka partition number.
- `customer_user` : ACI Kafka [client identity]({{%baseurl%}}/../concepts.html#client-identities).
- `usecase`: ACI Kafka cluster name.

Note

If you have a dedicated cluster, you may have access to more metrics to check the clusters overall health.


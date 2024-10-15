---
title: "What are available metrics and labels?¶"
weight: 5
---

### Topic/user throughput metrics[Internal Link: ¶](\#topic-user-throughput-metrics)

Warning

When querying `BrokerTopicMetrics` you **must** include a filter for at least one of the `topic`, `customer_group`,
`namespace`, or `topic_name` labels. A filter of an empty value (e.g. `topic=""`) queries the total for **all** topics
in the cluster. Any per-topic query must include a non-empty filter or an inverse filter ( `topic!=""`).

#### Produce activities[Internal Link: ¶](\#produce-activities)

| Metric name | Description | Labels |
| --- | --- | --- |
| `kafka_server_BrokerTopicMetrics_MessagesIn` | Total count of incoming messages per topic. | `topic`, `customer_group`, `namespace`, `topic_name` |
| `kafka_server_BrokerTopicMetrics_BytesIn` | Total count of incoming bytes per topic. | `topic`, `customer_group`, `namespace`, `topic_name` |
| `kafka_server_BrokerTopicMetrics_TotalProduceRequests` | Produce request count per topic. | `topic`, `customer_group`, `namespace`, `topic_name` |
| `kafka_server_BrokerTopicMetrics_BytesRejected` | Rejected produce byte count per topic. | `topic`, `customer_group`, `namespace`, `topic_name` |
| `kafka_server_BrokerTopicMetrics_NoKeyCompactedTopicRecords` | Message validation failure count due to no key specified for compacted topic. | `topic`, `customer_group`, `namespace`, `topic_name` |
| `kafka_server_Produce_byte_rate` | Produce rate per user on each broker. | `user` |
| `kafka_server_Produce_throttle_time` | Amount of time produce requests were throttled on each broker (milliseconds). | `user` |

#### Consume activities[Internal Link: ¶](\#consume-activities)

| Metric name | Description | Labels |
| --- | --- | --- |
| `kafka_server_BrokerTopicMetrics_BytesOut` | Total count of outgoing bytes per topic. | `topic`, `customer_group`, `namespace`, `topic_name` |
| `kafka_server_BrokerTopicMetrics_TotalFetchRequests` | Consume request count per topic. | `topic`, `customer_group`, `namespace`, `topic_name` |
| `kafka_server_BrokerTopicMetrics_FailedFetchRequests` | Failed consume request count per topic. | `topic`, `customer_group`, `namespace`, `topic_name` |
| `kafka_server_Fetch_byte_rate` | Consume rate per user on each broker. | `user` |
| `kafka_server_Fetch_throttle_time` | Amount of time consume requests were throttled on each broker (milliseconds). | `user` |

### Partition metrics[Internal Link: ¶](\#partition-metrics)

| Metric name | Description | Labels |
| --- | --- | --- |
| `kafka_log_Log_LogEndOffset` | The offset of the last stored message on disk. | `topic`, `customer_group`, `namespace`, `topic_name`, `partition` |
| `kafka_log_Log_LogStartOffset` | The offset of the earliest stored message on disk. | `topic`, `customer_group`, `namespace`, `topic_name`, `partition` |
| `kafka_log_Log_Size` | The partition’s size stored on each broker in bytes, including replicas. | `topic`, `customer_group`, `namespace`, `topic_name`, `partition` |
| `kafka_cluster_Partition_InSyncReplicasCount` | Number of in-sync replicas. | `topic`, `customer_group`, `namespace`, `topic_name`, `partition` |

### Synthetic Metrics[Internal Link: ¶](\#synthetic-metrics)

A set of metrics emitted by a synthetic monitoring tool we have developed which continually produces and consumes from the cluster in small bursts, capturing and recording produce or consume failures, as well as the maximum round trip latency.

| Metric name | Description | Labels |
| --- | --- | --- |
| `kafka_synthetic_monitor_round_trip_record_produce_failures_total` | Count of records which were created but failed to be sent by the producer. | `namespace`, `topic_name`, `partition`, `usecase` |
| `kafka_synthetic_monitor_round_trip_record_consume_failures_total` | Count of records which were expected to be consumed by the consumer but were not. | `namespace`, `topic_name`, `partition`, `usecase` |
| `kafka_synthetic_monitor_round_trip_latency_ms_max` | The maximum time (in a moving 30 second window) taken to send a record to the topic and subsequently consume the same record. | `namespace`, `topic_name`, `partition`, `usecase` |
| `kafka_synthetic_monitor_monitoring_failures_total` | Count of failures which occurred when producing and consuming from the synthetic monitoring topic. | `usecase` |
| `kafka_synthetic_monitor_preparation_failures_total` | Count of failures which occurred when preparing the the synthetic monitoring topic for use. | `usecase` |
| `kafka_synthetic_monitor_runtime_failures_total` | Count of unexpected failures which occurred while monitoring this cluster. | `usecase` |

### Others[Internal Link: ¶](\#others)

| Metric name | Description | Labels |
| --- | --- | --- |
| `kafka_server_Request_request_time` | Request rate per user on each broker. | `user` |
| `kafka_server_Request_throttle_time` | Request quota metric per user on each broker. | `user` |

### Lag metrics[Internal Link: ¶](\#lag-metrics)

Details for lag metrics are found [Internal Link: in the dedicated page](../client/lag_metrics.html#what-are-available-metrics-and-labels).

### Labels[Internal Link: ¶](\#labels)

**Labels on all broker metrics**

- `_ws_` : Mosaic workspace. This is a constant value set to `aci-kafka` and cannot be changed.
- `_ns_`: Mosaic namespace. The value for `_ns_` for the above metrics is `kafka-cluster-<cluster-name>`
- `broker_id`: The Kafka broker’s ID.
- `cluster`: Apple Kafka cluster name.

**Labels on per-topic/partition metrics**

- `customer_group`: Apple Kafka \[customer group\]({{< ref “../../concepts#customer-group” >}}).
- `namespace`: Apple Kafka \[namespace\]({{< ref “../../concepts#namespace” >}}).
- `topic_name`: Apple Kafka topic name. The topic name within the namespace.
- `topic`: The fully-qualified topic name including the namespace name.
- `partition`: The Kafka partition number.
- `user` : Apple Kafka \[client identity\]({{< ref “../../concepts#client-identities” >}}).

**Labels on Synthetic Metrics**

- `_ns_`: Mosaic namespace. One of `synthetic-monitor-prod` or `synthetic-monitor-if1`, depending on the environment.
- `usecase`: Apple Kafka cluster name.


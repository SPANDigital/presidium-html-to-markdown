---
title: "What changes?¶"
weight: 4
---

### Broker metrics[Internal Link: ¶](\#broker-metrics)

There are several changes to metric naming. Although all the same metrics remain available, most have changed name
and/or the set of labels used.

1. **Mosaic namespace**: each cluster’s broker metrics now live under the `kafka-cluster-${cluster}` namespace,
   whereas previously this was `kafka-broker-${cluster-in-kebab-case}-piekafka`.
2. **Unprefixed metric names**: all Kafka metrics are now named starting with the prefix `kafka_`. In all cases where an
   unprefixed name was used, those names are now prefixed with `kafka_server_`.
3. **Cluster labels**: all clusters are now identified via the labels `cluster`, `environment`, and `datacenter`. The
   `application_name` label has been replaced by `cluster`.
4. **Broker labels**: the label for broker ID is now `broker_id` instead of `broker`. The `instance_id` label, which
   contained the Carnival instance ID, has been removed.
5. **Metrics ending in `_global`**: the `_global` suffix is no longer used.
6. **Total of all topics metrics**: for metrics with a `topic` label, the sum of all topics for that metric can be
   queried using the `topic=""` filter. Previously, these totals were available in `_global`-suffixed metrics.
7. **Counters for rates**: most per-second rate metrics (i.e. gauges) are now also available as total counters for use
   with PromQL’s `rate()` and `increase()` functions. Where a metric name ended in `PerSec`, the total counter is
   available under the same name without the `PerSec` suffix. In other cases, `_total` is appended to the name. If you
   are unsure of a metric’s type you can check the `_type_` label.
8. **Additional topic labels**: topic metrics now include a `customer_group` label, enabling simplified sharing
   policies. The `topic` label now includes the fully-qualified topic name which can be useful to cross-correlate with
   client-side metrics.
9. **Simplified client ID/user labels**: the `client`, `customer_client`, and `customer_user` labels have been removed.
   The `client_id` or `user` labels should be used instead.

Warning

_6\. **Total of all topics metrics**_: This is especially important to note for `kafka_server_BrokerTopicMetrics_*`
metrics. In these cases you **must** update your query filters to include either `topic=""` (for the total of all
topics) _or_ `topic!=""` (for topic-specific metrics). Not doing so may result in double-counting of metrics.

Note

_7\. **Counters for rates**_: note that the PromQL `rate()` function (and similar) need _at least_ 2 data points per
vector. Broker metrics are reported every 30s. This means that rates must be calculated over at least a 1-minute period
(ideally, more for safety). Currently, Mosaic’s minimum interval in Grafana defaults to 10s, so even when using
[External Link: `[$__rate_interval]`](https://grafana.com/blog/2020/09/28/new-in-grafana-7.2-__rate_interval-for-prometheus-rate-queries-that-just-work/)
you will see discontinuous graphs by default. We recommend setting the `Min interval` value for Grafana panels to `1m`
in combination with `$__rate_interval`.

The table below provides an example of each change:

|  | Change | Before | After |
| --- | --- | --- | --- |
| 1 | Mosaic namespace | `kafka-broker-icloud-shared-prod-8-piekafka` | `kafka-cluster-icloud_shared_prod_8` |
| 2 | Unprefixed metric names | `Fetch_byte_rate` | `kafka_server_Fetch_byte_rate` |
| 3 | Cluster labels | `{application_name="icloud_shared_prod_8-piekafka", datacenter="MR"}` | `{cluster="icloud_shared_prod_8", datacenter="mr", environment="prod"}` |
| 4 | Broker labels | `{broker="1001", instance_id="10101001", host_name="mr37p01if-infs06050101.mr.if.apple.com"}` | `{broker_id="1001", host_name="mr37p01if-infs06050101.mr.if.apple.com"}` |
| 5 | Metrics ending in `_global` | `kafka_server_ReplicaManager_UnderReplicatedPartitions_global` | `kafka_server_ReplicaManager_UnderReplicatedPartitions` |
| 6 | Total of all topics metrics | `kafka_server_BrokerTopicMetrics_BytesInPerSec_global` | `kafka_server_BrokerTopicMetrics_BytesInPerSec{topic=""}` |
| 7 | Counters for rates | _Not available_ | `kafka_server_BrokerTopicMetrics_BytesIn` |
| 8 | Additional topic labels | `{namespace="icloud.family.prod.mr", topic_name="prod-MR-current"}` | `{namespace="icloud.family.prod.mr", customer_group="icloud.family.prod", topic_name="prod-MR-current", topic="icloud.family.prod.mr.prod-MR-current"}` |
| 9 | Simplified client ID/user labels | `{client="unknown", customer_user="icloud.edu.swu-prod.figaro-producer"}` | `{user="icloud.edu.swu-prod.figaro-producer"}` |

### Consumer lag metrics[Internal Link: ¶](\#consumer-lag-metrics)

Consumer lag metrics will now be reported with one of two additional labels:

- `client_customer_group`: The customer group of the consumer group’s client.
- `customer_group`: The customer group of the topic being consumed.

Going forward, we will be sharing metrics with policies that use these labels. Because of the way that Mosaic sharing
policies work, it is necessary to include these labels in any queries that use these metrics. For now, we will retain
the existing sharing policies so that existing queries continue to work, but we recommend updating your queries with
the new labels.

Queries that used the `client` label should now use the `client_customer_group` label, for example:
label:

```
--- a.promql
+++ b.promql
@@ -1,6 +1,6 @@
 kafka_consumergroup_group_lag{
     _ws_="aci-kafka",
     _ns_="kafka-consumer-lag-prod",
-    client="ist.columbo.prod.columboProductionClient",
+    client_customer_group="ist.columbo.prod",
     consumer_group="ist.columbo.prod.columboProductionClient.elastic-consumer"
 }

```

We will now also share lag metrics based on the customer group of the **topic**. This means you will be able to monitor
consumer groups that consume from your topics but may be part of a separate customer group. For example, the following
query queries the lag for **all** consumers of a specific topic:

```
kafka_consumergroup_group_lag{
   _ws_="aci-kafka",
   _ns_="kafka-consumer-lag-prod",
   customer_group="ist.columbo.prod",
   namespace="ist.columbo.prod.st",
   topic_name="session-crash"
}

```

These sharing policies will only work for consumer groups that follow our
[Internal Link: client naming convention](../client/consumer_settings.html#consumer-group-and-aci-kafka-client).

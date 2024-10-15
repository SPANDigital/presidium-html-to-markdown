---
title: "Key configurations¶"
weight: 2
---

### Default ACI Kafka configuration injection [¶](#default-aci-kafka-configuration-injection "Link to this heading")

The Kafka team provides a set of defaults to make things easy to set up, while prioritizing for durability.
Follow the instructions described in [Default Configuration Injection]({{< ref "client/default_config_injection" >}}) to get started.

### Consumer Group and ACI Kafka Client [¶](#consumer-group-and-aci-kafka-client "Link to this heading")

Any ACI Kafka client with its [client identity topic access]({{%baseurl%}}/../concepts.html#client-identity-topic-access) set to `CONSUME`
is authorized to create a [consumer group]({{%baseurl%}}/consumer_client.html#consumer-group-a-k-a-group-id-config) by prefixing the `group.id`
property with their [ACI Kafka client identity]({{%baseurl%}}/../concepts.html#client-identities):

```
group.id = "<aci_kafka_customer_group>.<aci_kafka_client_name>.<suffix_to_identify_group>"
kaffe.client.id = "<aci_kafka_customer_group>.<aci_kafka_client_name>"
client.id = "<aci_kafka_customer_group>.<aci_kafka_client_name>"

```

#### Example [¶](#example "Link to this heading")

A user `aci.queue.qa.user1` will have access to all consumer groups prefixed with
`aci.queue.qa.user1.`. Therefore, `aci.queue.qa.user1.consumer1` is a valid consumer group
name.

Note

Make sure to follow the convention before starting using ACI Kafka client version >=0.8. **The default behavior in >=0.8**
**is to enforce the consumer group convention, i.e., the consumer will throw an exception during initialization if the consumer**
**group does not follow the convention.** We can help move a consumers’ offset from one consumer group to another if they do not follow the convention.
This requires a few minutes of downtime on the consumer clients and coordination with the ACI Kafka team.

If the consumer can afford downtime and change the consumer group. In that case, we may allow an exception if there is a strong reason.
If the consumer qualifies for the exception, we can assist in disabling the checks.

If you are using ACI Kafka client with third-party apps like Spark you have to use the [custom ConfigurationInterceptor]({{%baseurl%}}/https://github.pie.apple.com/pie/kafka-tools/blob/develop/kafka-pie-client/src/main/java/com/apple/pie/queue/kafka/client/configinterceptors/ConsumerGroupIdConfigurationInterceptor.java). Spark manipulates the consumer group internally and requires this interceptor as a workaround.
`kafka-pie-client>=0.13 include ` com.apple.pie.queue.kafka.client.configinterceptors.ConsumerGroupIdConfigurationInterceptor\`.

### Take control of offset commits [¶](#take-control-of-offset-commits "Link to this heading")

The Kafka Consumer client [commits the offset]({{%baseurl%}}/consumer_client.html#offset-commit) automatically every 5 seconds.
This behavior may lead to data loss if the offset is committed before the client finishes processing the records.
Therefore, we recommend setting `enabled.auto.commit=false` to disable the automatic offset commits.

### Optimize for throughput or latency [¶](#optimize-for-throughput-or-latency "Link to this heading")

Usually clients need to choose between throughput or latency, however, we recommend a middle ground:

- Increase `fetch.min.bytes` for throughput, allowing the consumer to fetch batches
- Keep `fetch.max.wait.ms` to default

If consumer has a high throughput requirements, increasing `fetch.min.bytes` may help. However, if the consumer has a low latency requirement or consumes from a topic with a low volume, adjusting `fetch.max.wait.ms` should help.


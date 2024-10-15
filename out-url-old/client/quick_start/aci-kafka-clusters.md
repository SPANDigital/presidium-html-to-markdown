---
title: "Java Properties for connecting to ACI Kafka Clusters¶"
weight: 5
---

**Please be aware that Kafka clusters are NOT reachable from your local**
**desktop environment.**

Note

You can use our [Sidecar\
Solution]({{%baseurl%}}/../dev-and-ci.html#local-environments) for integrating the
ACI Kafka Client with a local Apache Kafka cluster. See the 0.9.x
[Dev Mode Flag]({{%baseurl%}}/../dev-and-ci.html#local-environments) feature for
details on how to minimize configuration differences between local
dev/testing and production.

ACI Kafka client 0.8.x provides built-in cluster discovery and secrets
injection, and starting in 0.9, we will begin injecting safe defaults
for producers and consumers, along with additional secret information
needed for [end-to-end\
crypto]({{%baseurl%}}/security_settings.html#payload-encryption-and-decryption-settings),
if enabled. More on Kaffe and client configuration and secure secrets
injection can be found in the [Kaffe documentation]({{< ref "client/kaffe" >}}).

Here’s an example of the minimum properties needed for configuring your
Kafka clients to connect to ACI Kafka clusters in IF1 or IF-PROD:

```
String kaffeConnect = ...     // see Kaffe Connect Endpoints above
String kaffeClientId = ...    // i.e. {customer}.{group}.{env}.{client}
String kaffeNamespaceId = ... // i.e. {customer}.{group}.{env}.{namespace}

// If not using PIE secrets, provide the file path of your client's private key
String pieSecretsKey = ...    // your secret key if using PIE Secrets
String kaffePrivateKeyPath = String.format("%s/%s",
                                           System.getenv("PLATFORM_SECRETS_PATH"), pieSecretsKey);

String topicName = "topic-123";
PieTopicName pieTopic = new PieTopicName(PieNamespaceName.of(kaffeNamespaceId), topicName);

HashMap<String, Object> configs = new HashMap<>();
configs.put("pie.queue.kaffe.connect", kaffeConnect);
configs.put("pie.queue.kaffe.client.id", kaffeClientId);
configs.put("pie.queue.kaffe.namespace.id", kaffeNamespaceId);
configs.put("pie.queue.kaffe.client.private.key.location", kaffePrivateKeyPath);

// Default configuration is available in ACI Kafka client 0.9 and higher.
// See below to learn more about this feature.
configs.put("pie.queue.kaffe.config.injection", true);

```

**The following example is a demo, don’t use it on production** An
example KafkaProducer can be instantiated with the following additional
properties:

```
configs.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class);
configs.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, StringSerializer.class);
try (KafkaProducer<String, String> producer = KafkaClientUtils.createProducer(configs)) {
   producer.send(new ProducerRecord<>(pieTopic.getFullName(), "key", "Hello World"));
}

```

**The following example is a demo, don’t use it on production** An
example KafkaConsumer can be instantiated with the following additional
properties:

```
long pollTimeoutMs = 30000;

// For more information on group id naming convention,
// see https://docs.aci.apple.com/kafka/client/consumer_settings.html#consumer-group-convention-and-enforcement
String groupId = String.format("%s.%s", kaffeClientId, "consumer-group-123");

configs.put(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class);
configs.put(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class);
configs.put(ConsumerConfig.GROUP_ID_CONFIG, groupId);
try (KafkaConsumer<String, String> consumer = KafkaClientUtils.createConsumer(configs)) {
  consumer.subscribe(Arrays.asList(pieTopic.getFullName()));

  while(true) {
    ConsumerRecords<String, String> records = consumer.poll(Duration.ofMillis(pollTimeoutMs));
    for (TopicPartition partition : records.partitions()) {
      List<ConsumerRecord<String, String>> partitionRecords = records.records(partition);
      for (ConsumerRecord<String, String> record : partitionRecords) {
        System.out.println(
          String.format("Fetched record from topic='%s', partition='%d', key='%s', value='%s'",
          record.topic(), record.partition(), record.key(), record.value()));
      }
      System.out.println(String.format("Updated offset for topic='%s' partition='%d' to offset='%d'",
        partition.topic(), partition.partition(), nextOffset));
    }
  }
}

```

The following properties are needed for producers and consumers not
configured for [Default Configuration\
Injection]({{< ref "client/default_config_injection" >}}) using Kaffe
(optional in ACI Kafka Client 0.9 through 0.13, or enabled by default in
0.14 and higher):

```
configs.put("security.protocol", "SASL_SSL");
configs.put("sasl.mechanism", "PLAIN");

```

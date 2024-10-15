---
title: "Using the client with PIE Spark¶"
weight: 6
---

You shouldn’t have problems using our client with PIE Spark.

You will need to add the spark dependencies for streaming and exclude
the original client:

```
compile ("org.apache.spark:spark-streaming-kafka-0-10_2.11:${sparkVersion}") {
    exclude group: 'org.apache.kafka', module: 'kafka-clients'
}
// This avoids conflicts as the platform also provides a version of spark streaming
compileOnly "org.apache.spark:spark-streaming_2.11:${sparkVersion}"

compile "com.apple.pie.queue:kafka-pie-client:${kafkaClientVersion}"

```

You can then build a direct stream quite easily:

```
PieNamespaceName namespace = new PieNamespaceName("<customer>.<group>.<env>.<namespace>")
PieTopicName topic = new PieTopicName(namespace, "<topic>");
Map<String, Object> kafkaParams = new HashMap<>();
kafkaParams.put("config.interceptors.classes", "com.apple.pie.queue.kafka.client.configinterceptors.ConsumerConfigurationInterceptor,com.apple.pie.queue.kafka.client.configinterceptors.KaffeConfigurationInterceptor,com.apple.pie.queue.kafka.client.configinterceptors.PieDeserializerConfigurationInterceptor");
// Add all your other properties (notably secret key provider and kaffe connect, namespace and clientId...).

JavaStreamingContext streamingContext = new JavaStreamingContext(new SparkConf(), new Duration(TimeUnit.MINUTES.toMillis(10)));
JavaInputDStream<ConsumerRecord<String, String>> stream = KafkaUtils.createDirectStream(
      streamingContext,
      LocationStrategies.PreferConsistent(),
      ConsumerStrategies.<String, String>Subscribe(Collections.singleton(topic.getFullName()), kafkaParams)
);

// Do what you want with your stream

```

You will also need to use a specific interceptor for consumers as Spark
prefixes consumer groups. Instructions for this setup can be found in
our [consumer settings\
docs]({{%baseurl%}}/consumer_settings.html#consumer-group-and-aci-kafka-client).

Warning

The order of values in `config.interceptors.classes` matters (especially as of version 0.18)!
Incorrect orders may result in values that later interceptors depend on not being present.

For producers, you will need to adjust the config interceptors:

```
kafkaParams.put("config.interceptors.classes", "com.apple.pie.queue.kafka.client.configinterceptors.ProducerConfigurationInterceptor,com.apple.pie.queue.kafka.client.configinterceptors.KaffeConfigurationInterceptor,com.apple.pie.queue.kafka.client.configinterceptors.PieSerializerConfigurationInterceptor");

```

If you want to know more about see the [PIE Spark\
documentation](https://docs.aci.apple.com/spark/).

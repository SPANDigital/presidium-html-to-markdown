---
title: "ConfigurationInterceptors¶"
weight: 3
---

We added configuration interceptors in Kafka, for the moment this change
only exists in our fork but we will push it upstream in the future.

A ConfigurationInterceptor is a simple function which we apply to the
configuration before creating the Kafka client.

This enables us to do 2 things:

- Set the encryption and envelope serializers/deserializers.
- Call Kaffe to resolve the cluster and add safe defaults.

If you can modify the code that initializes Kafka clients you should use
the 2 methods to create a client in _KafkaClientUtils_ they
will set the correct configuration interceptors.

If you don’t have access to the code (for example in Spark), you should
add this property to your Kafka properties.

Consumer:

```
config.interceptors.classes=com.apple.pie.queue.kafka.client.configinterceptors.DeserializerConfigurationInterceptor,com.apple.pie.queue.kafka.client.configinterceptors.KaffeConsumerConfigurationInterceptor

```

Producer:

```
config.interceptors.classes=com.apple.pie.queue.kafka.client.configinterceptors.SerializerConfigurationInterceptor,com.apple.pie.queue.kafka.client.configinterceptors.KaffeProducerConfigurationInterceptor

```

For the Kaffe interceptor you will need to set a Kaffe url, a client id,
and a namespace. See the [Kaffe]({{< ref "client/kaffe" >}}) documentation for
details.

When interacting with the client, you will need to pass in your topic
names as fully qualified names.

You can do this with code by using:

```
PieNamespaceName pieNamespaceName = new PieNamespaceName(namespaceName)
String fullyQualifiedName = pieNamespaceName.maybeExpand(topicName)

```

In the case where you are not using namespaces (e.g. in development),
you can simply not set a namespace.


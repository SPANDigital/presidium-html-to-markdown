---
title: "Configuring your client with Kaffe¶"
weight: 3
---

Kaffe configuration is meant to be easy to plug into your existing
client.

The [External Link: Kaffe client
library](https://github.pie.apple.com/pie/kafka-tools/tree/develop/kaffe-client)
is hooked into the kafka-pie-client thanks to a ClusterResolver
implementation.

To be able to start using Kaffe, all you need to do is to set a
configuration interceptor, point to the right Kaffe URL (without a trailing slash `/`), and set your
clientId and the namespace you want to use:

```
# Producers:
config.interceptors.classes=com.apple.pie.queue.kafka.client.configinterceptors.KaffeProducerConfigurationInterceptor
# Consumers:
config.interceptors.classes=com.apple.pie.queue.kafka.client.configinterceptors.KaffeConsumerConfigurationInterceptor
# Both producers and consumers:
pie.queue.kaffe.connect=<kaffe-url>
pie.queue.kaffe.client.id=<client-id-given-out-by-PIE>
pie.queue.kaffe.namespace.id=<namespace-id-given-out-by-PIE>

```

To make it easier we also provide
`KafkaClientUtils.createPieConsumer(config)` and
`KafkaClientUtils.createPieProducer(config)` which set the
right configuration interceptors.


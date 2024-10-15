---
title: "In a nutshell¶"
weight: 2
---

While we highly recommend that you read the rest of this documentation,
here’s a short summary of what you need to know:

- Each client instance is scoped to a namespace, so you won’t be able
  to consume from or produce to multiple namespaces with a single
  client instance.
- Interceptors expose the envelope so you can read from or write to
  it.
- You can use _KafkaClientUtils.createProducer_ and
  _KafkaClientUtils.createConsumer_ to add the
  configuration interceptor.

In a ACI Kafka environment, you must [configure Kaffe]({{< ref "client/kaffe" >}}) for
cluster discovery and [configure security]({{< ref "client/security_settings" >}}).

In ACI Kafka you will need to use fully qualified topic names (with the
namespace as a prefix). You can rebuild a topic name by using the class
PieTopicName.


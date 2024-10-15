---
title: "What will Kaffe do?¶"
weight: 5
---

### When you instantiate a consumer or a producer[Internal Link: ¶](\#when-you-instantiate-a-consumer-or-a-producer)

- The Kaffe client will propagate the properties passed to the
  consumer or producer constructor, as well as some extra metadata, to
  the Kaffe service.
- The Kaffe client will establish a Session with a UUID (we will call
  this a sessionId). A session lasts for as long as the producer or
  consumer is running, and the sessionId is unique.
- The Kaffe service will send down possibly modified properties for
  the client to use.

Note

- You may have sensitive information (e.g. passwords, secrets…) which shouldn’t be sent to Kaffe.
You can hide these by using a
[External Link: ConfigFilter](https://github.pie.apple.com/pie/kafka-tools/blob/develop/kafka-pie-client/src/main/java/com/apple/pie/queue/kafka/client/kaffe/configfilter/ConfigFilter.java)
and then specifying its path in the producer/consumer property
`pie.queue.kaffe.config.filter.class`. The property
`pie.queue.kaffe.config.filter.class` takes in a list of
`ConfigFilter` classes. You can also use the default
filters list which contains a version configuration filter and a regex
obfuscation filter. The version filter will enable us to know the
Kafka and client JAR versions for troubleshooting. The regex filter
will use regexes passed in the property:
`com.apple.pie.queue.kaffe.config.filter.simple.regex` (the default
regex suppresses properties that contain _secret_, _password_, or
_jaas_).
- We only log properties on our side in order to provide more efficient
troubleshooting. Kaffe does not yet modify properties passed in. In
the future, Kaffe will set properties to correct defaults, but this is
not yet implemented.

### When you start consuming or producing to a new topic[Internal Link: ¶](\#when-you-start-consuming-or-producing-to-a-new-topic)

- The client validates that the topic belongs to the configured
  namespace (fully qualified topic names begin with the namespace
  name)
- The information of the topic contains the location of the topic and
  a cluster, Kafka then calls Kaffe to know the bootstrap string for
  this cluster
- The client creates an instance of the producer/consumer for this
  cluster.

Kaffe will select the appropriate listener/port based on the client
configuration for `security.protocol`. If there is no option for the
selected protocol, an exception will be thrown by the ClusterResolver
(see javadoc for detailed information).

The cluster lookup is cached and is done only once per cluster.


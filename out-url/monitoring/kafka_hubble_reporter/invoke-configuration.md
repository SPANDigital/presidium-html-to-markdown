---
title: "Invoke your own configuration¶"
weight: 6
---

By default configuration is extracted from your client or broker
configuration. You might have another way to provide configuration
(framework base, hard coded…). You can easily integrate with your
custom provider by implementing
_com.apple.pie.queue.kafka.hubble.HubbleProperties_. You will then need
to specify in your configuration _pie.queue.config.class_ to the path of
this class. By default it uses
_com.apple.pie.queue.kafka.hubble.PieHubbleProperties_. But for the
broker we extend this class to use the Kafka broker id as the instance
id in _com.apple.pie.queue.kafka.hubble.PieBrokerHubbleProperties_. We
then add to all our broker properties:

```
pie.queue.config.class=com.apple.pie.queue.kafka.hubble.PieBrokerHubbleProperties

```

Another existing _HubbleProperties_ implementation is
_com.apple.pie.queue.kafka.hubble.PassThroughProperties_ which will pass
in all your configuration down to hubble.


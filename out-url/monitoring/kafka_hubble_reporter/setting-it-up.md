---
title: "Setting it up¶"
weight: 3
---

The reporter is configured just as a regular
_o.a.k.common.metrics.KafkaMetricReporter_ reporters by providing your
client and broker properties. You will need to add to your client
configuration our metric reporter by adding:

```
metric.reporters=com.apple.pie.queue.kafka.hubble.HubbleMetricsReporter

```

If your application doesn’t already initialize the HubbleAgent you will
want to add to your configuration:

```
pie.queue.hubble.agent.initialize=true

```

If this property is set, you will need to set the configuration of
hubble. We will be read from your HubbleProperties the path to publish
stats and other general hubble properties.

By default configuration name are:

> - _hubble.config.file_ A configuration file sent directly to hubble.
> - _pie.queue.applicationName_ or _ISApplicationName_ used for the
> application name part of the hubble path.
> - _pie.queue.datacenterName_ used for the data center name to report
> to.
> - _pie.queue.podName_ or _partition_ the pod name part of the hubble
> path.
> - _pie.queue.hostName_ the hostname part of the hubble path.
> - _pie.queue.instanceId_ or _ISInstanceId_ the instance id part of
> the hubble path.
> - _pie.queue.metrics.prefix_ the prefix to add in front of all the
> stats.

Other useful configuration used for hubble are :

```
hubble.publisher.publish=true
hubble.agent.enable.percentile=true

# The publishing frequency in seconds.
hubble.publisher.frequency=60

hubble.event.includeAll=true
hubble.event.names=*
hubble.properties.canOverideFromFileAtRuntime=false

hubble.channel.tcp.proto.include=true
hubble.publisher.tcp.server.url=<url of the hubble to report to>

```

Please refer to the [External Link: hubble agent
documentation](https://github.pie.apple.com/Hubble/hubble-agent-docs)
for more details.

For pie queue deployment we keep all these in a common
_hubble.properties_ that we point to in our configuration by setting
_hubble.config.file_.


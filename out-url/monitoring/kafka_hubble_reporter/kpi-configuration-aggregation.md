---
title: "KPI configuration and aggregation¶"
weight: 4
---

Hubble has a limitation in the number of KPIs which can be emitted. It
also doesn’t enable you to easily roll things up together on something
other than the machine placement.

In Kafka all stats are emitted with a scope or a set of tags (taking the
shape of key/values). For example _kafka.log-Log-LogEndOffset_ is
emitted for each topic and partition. Because you can have many topics
and partitions and don’t want to create alerts for each tuple you’ll
want to roll things up together.

**By default all tags are rolled up together** and emit stats like
_kafka.log-Log-LogEndOffset(partition:aggregated)(topic:aggregated)_.

A non exhaustive list of possible tags are:

- partition (the partition number of the topic).
- topic (the topic name).
- client-id (the client.id property passed in to the instance).

Because some tags for some values may be more important you can
selectively enable them with:

```
pie.queue.metrics.nameResolver[<tagName>]=<regularExpression>

```

If a regular expression is invalid an error message will be logged and
we’ll simply ignore it.

This will push 2 stats for the previous event
_kafka.log-Log-LogEndOffset(partition:aggregated)(topic:aggregated)_ and
_kafka.log-Log-LogEndOffset(partition:aggregated)(topic:)_. If a tag has a many values where a subset of the topics are
selectively enabled, the other topics will be aggregated with under the
value _nonIncluded_.

Here is an example of selectively enabling topic metrics:

```
pie.queue.metrics.nameResolver[topic]=foo.*

```

This will enable all the KPIs which have a key ‘topic’ which starts
with _foo_. So if we take the event _compression rate_ and the topics
foo, foo2, bar and baz 4 KPIs will be emitted:

```
producer-topic-metrics-compression-rate(client-id:aggregated)(topic:aggregated)
producer-topic-metrics-compression-rate(client-id:aggregated)(topic:foo)
producer-topic-metrics-compression-rate(client-id:aggregated)(topic:foo2)
producer-topic-metrics-compression-rate(client-id:aggregated)(topic:nonIncluded)

```

in _aggregated_ we’ll have all topics, in _foo_ and _foo2_ we’ll just
have respectively _foo_ and _foo2_ and in _nonIncluded_ _bar_ and
_baz_.

For aggregated values such as averages and percentiles rolling up is not
straight forward we therefore compute the weighted average of them,
unfortunately this has limited mathematical significance and we
recommend using the maximum.

If you want to know more about the existing KPIs see the [External Link: public
documentation](https://kafka.apache.org/documentation.html#monitoring).


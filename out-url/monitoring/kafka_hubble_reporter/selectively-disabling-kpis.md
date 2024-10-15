---
title: "Selectively Disabling KPIs¶"
weight: 5
---

If you want to selectively disable some KPIs you can use the property:
_pie.queue.metrics.excluded.patterns_. Disabling takes priority over
enable, so when a KPI is selectively disabled, any explicit enabling of
that KPI will have no effect.

**Note: We marked _pie.queue.metrics.blacklist.patterns_ for deprecation.**

It takes a comma separated list of patterns to be applied to the root
name of the KPI (doesn’t take into account tags used for aggregation).

There are 2 ways to selectively disable KPIs:

- String matching (Using String.contains()), to do so simply give a
  portion of the KPI name as a parameter
- regular expression matching (Using java Pattern), to do so prefix
  your pattern by _re:_.

You can escape commas (.i.e: _,_) character with a backslash. If a
regular expression is invalid an error message will be logged and we’ll
simply ignore it.

Here’s an example of a valid configuration:

```
pie.queue.metrics.excluded.patterns=producer-topic,topic-metric[0-9]{2,4}

```

Note:

This feature was added because of
[External Link: KAFKA-5890](https://issues.apache.org/jira/browse/KAFKA-5890). You can
use this configuration to ignore all the extra KPIs:

```
pie.queue.metrics.excluded.patterns=re:consumer-fetch-manager-metrics-.*\.records-lag.*

```


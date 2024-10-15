---
title: "Migrating to the client¶"
weight: 5
---

Because it requires using the envelope, however, there will be problems
if you already have items in a topic as our consumer will try to parse
an envelope which does not exist.

If you introduce a Kafka client by transitive dependency you should
exclude it. You can do so in gradle this way:

```
compile(group: 'org.apache.spark', name: 'spark-streaming-kafka-0-10_2.11', version: '...') {
    exclude group: 'org.apache.kafka', module: 'kafka-clients'
}

```


---
title: "Gradle Dependency¶"
weight: 2
---

The client is a simple java dependency available in
[Artifactory](https://artifacts.apple.com/pie-release-local/com/apple/pie/queue/kafka-pie-client/)
at these coordinates:
`com.apple.pie.queue:kafka-pie-client:{version}`.
Check the current [version recommended for production]({{%baseurl%}}/../releases/index.html#recommended-release).

Add the ACI Kafka client dependency by first selecting your
`{version}` and then adding a dependency:

```
dependencies {
    compile "com.apple.pie.queue:kafka-pie-client:{version}"
}

```

If you have a transitive dependency on Apache Kafka, you’ll need to
exclude it from your dependencies.

Here’s an example of excluding the transitive dependency for Apache
Spark:

```
compile(group: 'org.apache.spark', name: 'spark-streaming-kafka-0-10_2.11', version: '...') {
    exclude group: 'org.apache.kafka', module: 'kafka-clients'
}

```


---
title: "Verifiable Consumer¶"
weight: 3
---

The consumer also has a convenience script. Here is a sample command:

```
deployment/scripts/kafka-verifiable-consumer.sh --broker-list `docker-machine ip`:9092 --topic mytopic --group-id mygroup

```

This will start a consumer reading from the `mytopic` topic using the
`mygroup` group.

Launching the consumer via Carnival is similar to the producer:

```
java ... -cp client-tools.jar org.apache.kafka.tools.VerifiableConsumer ...

```

After the class name, the arguments are the same as for the command-line
script.

---
title: "Verifiable Producer¶"
weight: 2
---

There is a convenience script for running the producer in the
development environment. In order to run it, a topic to produce to must
already exist. Here is a sample command:

```
deployment/scripts/kafka-verifiable-producer.sh --broker-list `docker-machine ip`:9092 --topic mytopic --throughput 100

```

This will start a producer sending 100 messages per second to the
`mytopic` topic. If no throughput argument is specified, throughput
simply will not be throttled. Execution will continue until the process
is given a termination signal. Alternatively, the `max-messages`
argument can be specified, in which case the producer will exit after
sending the given number of messages.

When running via Carnival, it is recommended to execute the Java class
directly. The producer and consumer are packaged as a rio application,
but there is no main class. So, you will need to invoke like this:

```
java ... -cp client-tools.jar org.apache.kafka.tools.VerifiableProducer ...

```

After the class name, the arguments are the same as for the command-line
script.


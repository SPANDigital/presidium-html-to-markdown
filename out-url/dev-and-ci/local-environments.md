---
title: "Local Environments¶"
weight: 3
---

Our Rio Sidecar can be started locally in order to experiment and begin
the process of integrating with Kafka.

Under the hood, this sidecar is backed by containers, therefore, either
[External Link: Vessel](https://vessel.apple.com/) or Docker are prerequisites for running
it locally. This has been tested with Vessel and Docker for Mac.

In addition to the above, we recommend using the Rio Sidecar binary. This
binary provides parity with a Rio pipeline while offering a convenient
interface to manage the lifecycle of your local environment.

Please ensure you’re on macOS >= `13.0`, have Rio Sidecar
binary >= `2.2.3` installed and have Vessel’s
[External Link: compat mode](https://vessel.apple.com/usage/docker-compatibility/compat-mode)
set up if you’re using sidecar with Vessel.

Start by creating the `aci-kafka.yml` file inside your `sidecars` folder
with the following content:

```
schemaVersion: 2.0

image: docker://docker-upstream.apple.com/apache/kafka:<oss-kafka-version>
host: PIE_KAFKA_HOST
ports:
  - name: KAFKA_PORT
    default: 9092

commands:
  - >
      HOST=$([ "$PIE_KAFKA_HOST" = "0.0.0.0" ] && echo "127.0.0.1" || echo "$PIE_KAFKA_HOST")
      KAFKA_NODE_ID=1
      KAFKA_PROCESS_ROLES='broker,controller'
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP='CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT'
      KAFKA_ADVERTISED_LISTENERS="PLAINTEXT_HOST://$HOST:$KAFKA_PORT,PLAINTEXT://$HOST:19092"
      KAFKA_CONTROLLER_QUORUM_VOTERS="1@$HOST:29093"
      KAFKA_LISTENERS="CONTROLLER://:29093,PLAINTEXT_HOST://:$KAFKA_PORT,PLAINTEXT://:19092"
      KAFKA_INTER_BROKER_LISTENER_NAME='PLAINTEXT'
      KAFKA_CONTROLLER_LISTENER_NAMES='CONTROLLER'
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1
      KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS=0
      KAFKA_TRANSACTION_STATE_LOG_MIN_ISR=1
      KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR=1
      /etc/kafka/docker/run

```

Please refer to [Internal Link: version mapping](releases/index.html#apache-kafka-version-mapping)
and substitute the Apache Kafka version in place for `<oss-kafka-version>`.

At the root of your project, use `sidecar start`. You’ll need to set
some environment variables for your Kafka client, load them with
`eval $(sidecar env)`.

If you’re running integration tests against the sidecar from IntelliJ,
run the following command instead:

```
sidecar env | sed 's/export //' | pbcopy

```

This will copy the environment configuration into your clipboard. Then
in IntelliJ, under **Run** \> **Edit Configurations…** >
**Templates**, select the template you’re interested in and paste the
environment variables from your clipboard.

Make sure to check if your Kafka client configuration meets the
requirements to connect to your new local environment as described in
[Internal Link: Client Configuration](#client-configuration).


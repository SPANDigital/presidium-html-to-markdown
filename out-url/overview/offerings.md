---
title: "Offerings¶"
weight: 3
---

ACI Kafka manages Kafka, including procuring, validating, deploying, operating, monitoring hardware, and the following services:

- A self-service API that allows you to manage topics, namespaces, client identities, and so on. See the [Internal Link: Concepts](concepts.html) section for details.
- A [Internal Link: sidecar](dev-and-ci.html#local-environments) for local development.
- An extended set of [Internal Link: client APIs](client/index.html) for cluster discovery [Internal Link: using Kaffe](client/kaffe.html).
- Client configuration management to enable [Internal Link: Security settings](client/security_settings.html).

ACI Kafka offers limited support for architecture and configuration reviews. It also supports various users, from large mission-critical use cases to small proof-of-concepts.

See [Internal Link: ACI Kafka caveats and limitations](#aci-kafka-caveats-and-limitations) for using ACI Kafka.

### Topics-as-a-Service[Internal Link: ¶](\#topics-as-a-service)

ACI Kafka allows you to manage Kafka topics without requiring operating the cluster itself. The ACI Kafka team provides the following support:

- **Operations:**
  The ACI Kafka team manages hardware build-outs, server configuration, monitoring the health of hosts, brokers, and [External Link: Apache ZooKeeper](https://zookeeper.apache.org/) instances. They also manage operational tasks like host replacement and partition rebalancing.
- **Scalability:**
  The ACI Kafka team uses the data you provide for capacity planning, so your clients can produce and consume gigabytes (GBs) per second. The team’s in-house operation’s tooling allows quick responses to cluster, hardware, and data center issues.
- **Reliability:**
  ACI Kafka implements Kafka access control lists (ACLs) to prevent others from reading or writing to your topics. The team uses Traffic Quota Management on multi-tenant clusters to enable high quota on a cluster level, enabling client rate-limiting so your topics stay available, replicated, and healthy. Data is replicated four times, allowing for continuous maintenance and improving overall uptime.
- **Security:**
  All broker instances and clients authenticate and communicate over transport layer security (TLS). In addition to end-to-end transport security, the [Internal Link: ACI Kafka client](#aci-kafka-client) offers end-to-end data security for topics with sensitive or confidential data. The hardware and infrastructure are patched continuously for security issues without action required.

### ACI Kafka client[Internal Link: ¶](\#aci-kafka-client)

[Internal Link: **ACI Kafka client**](client/index.html) is a set of patches on top of the open-source Apache Kafka client. The client is written in Java and is only available for services and applications written in the Java language or running on top of the [External Link: Java virtual machine (JVM)](https://en.wikipedia.org/wiki/Java_virtual_machine). It is API-compatible, allowing a safe drop-in replacement on your existing code base. You can also use the client to connect to clusters not managed by the ACI Kafka team.

The ACI Kafka client includes:

- An Apache Kafka client that provides [Internal Link: cluster discovery](client/kaffe.html) and [Internal Link: Telemetry integration](monitoring/index.html).
- Default property injection and automatic, secure delivery of credentials and secrets allow you to connect securely to a cluster
  with only three properties: **client**, **namespace**, and **secret**.
- [Internal Link: Kafka Sidecar](dev-and-ci.html#local-environments) for developing
  locally with ACI Kafka.
- Critical bug fixes are proactively back-ported when necessary. Open-source software releases may be delayed when releasing versions
  with these patches.
- The ability to report your configuration to the ACI Kafka team, enabling the team to efficiently troubleshoot issues and plan
  migrations more easily.

### ACI Kafka with third-party clients[Internal Link: ¶](\#aci-kafka-with-third-party-clients)

In some situations, the ACI Kafka client may not be suitable for your use cases. For example:

- If you are not using a Java language.
- If you don’t directly control the code that calls Kafka, such as third-party plugins or libraries.

For [Internal Link: third-party (non-Java client) use cases](client/kafka_for_all.html), you can use an ACI Kafka connect string and Simple Authentication and Security Layer/Secure Sockets Layer (SASL/SSL).

### Self-service[Internal Link: ¶](\#self-service)

The ACI Kafka [Internal Link: **Self-service**](self-service/index.html) offering aims to provide you with autonomy. Below are ways to help you manage your entities:

- Automate the management of ACI Kafka resources using a secured REST API.
- Enable authorized team members and users to control or view Kafka resources.
- Manage Topics, Clients, and Access Control configuration.
- View and inspect the health of your topics and consumer groups.
- Modify who can access your entities.
- Automate resource management using a command-line interface (CLI) tool.

### ACI Kafka caveats and limitations[Internal Link: ¶](\#aci-kafka-caveats-and-limitations)

Below are the specified conditions when using ACI Kafka:

- The ACI Kafka latest [Internal Link: release](releases/index.html) is based on Apache Kafka releases. The wire protocol
  negotiation is configured to support Apache
  Kafka clients from version 0.10.2 and later.
- [External Link: Kafka Streams](https://kafka.apache.org/documentation/streams/) are not officially supported.
- End-to-end encryption may be enabled in the ACI Kafka client for certain use cases.
- Size-based topic retention is required. An optional time-based retention can also be enforced.
- Client version greater or equal to 0.16 is required for [External Link: Java Development Kit 11 (JDK11)](https://java.apple.com/downloads/applejdk-11.html) support.
- Direct access to ZooKeeper is not allowed. Broker IPs or hostnames are provided through [Internal Link: Kaffe](client/kaffe.html).
- All Topic, Client, and Access changes are made using the ACI Kafka CLI or the ACI Kafka API.
  - Using the Kafka Admin Client is not supported.
- All multi-tenant clusters contain a default traffic quota per [Internal Link: client identity](concepts.html#client-identities) per broker that a single tenant should not exceed.
  The default quota is 35% of the cluster’s average traffic capacity. This quota acts as safeguard to prevent one tenant from using all of the cluster’s resources.
- ACI Kafka limits the number of connections allowed in the broker at any time. The default is 65K per broker.


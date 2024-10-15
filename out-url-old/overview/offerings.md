---
title: "Offerings¶"
weight: 3
---

ACI Kafka manages Kafka, including procuring, validating, deploying, operating, monitoring hardware, and the following services:

- A self-service API that allows you to manage topics, namespaces, client identities, and so on. See the [Concepts]({{< ref "concepts" >}}) section for details.
- A [sidecar]({{%baseurl%}}/dev-and-ci.html#local-environments) for local development.
- An extended set of [client APIs]({{< ref "client" >}}) for cluster discovery [using Kaffe]({{< ref "client/kaffe" >}}).
- Client configuration management to enable [Security settings]({{< ref "client/security_settings" >}}).

ACI Kafka offers limited support for architecture and configuration reviews. It also supports various users, from large mission-critical use cases to small proof-of-concepts.

See [ACI Kafka caveats and limitations](#aci-kafka-caveats-and-limitations) for using ACI Kafka.

### Topics-as-a-Service [¶](#topics-as-a-service "Link to this heading")

ACI Kafka allows you to manage Kafka topics without requiring operating the cluster itself. The ACI Kafka team provides the following support:

- **Operations:**
  The ACI Kafka team manages hardware build-outs, server configuration, monitoring the health of hosts, brokers, and [Apache ZooKeeper](https://zookeeper.apache.org/) instances. They also manage operational tasks like host replacement and partition rebalancing.
- **Scalability:**
  The ACI Kafka team uses the data you provide for capacity planning, so your clients can produce and consume gigabytes (GBs) per second. The team’s in-house operation’s tooling allows quick responses to cluster, hardware, and data center issues.
- **Reliability:**
  ACI Kafka implements Kafka access control lists (ACLs) to prevent others from reading or writing to your topics. The team uses Traffic Quota Management on multi-tenant clusters to enable high quota on a cluster level, enabling client rate-limiting so your topics stay available, replicated, and healthy. Data is replicated four times, allowing for continuous maintenance and improving overall uptime.
- **Security:**
  All broker instances and clients authenticate and communicate over transport layer security (TLS). In addition to end-to-end transport security, the [ACI Kafka client](#aci-kafka-client) offers end-to-end data security for topics with sensitive or confidential data. The hardware and infrastructure are patched continuously for security issues without action required.

### ACI Kafka client [¶](#aci-kafka-client "Link to this heading")

[**ACI Kafka client**]({{< ref "client" >}}) is a set of patches on top of the open-source Apache Kafka client. The client is written in Java and is only available for services and applications written in the Java language or running on top of the [Java virtual machine (JVM)](https://en.wikipedia.org/wiki/Java_virtual_machine). It is API-compatible, allowing a safe drop-in replacement on your existing code base. You can also use the client to connect to clusters not managed by the ACI Kafka team.

The ACI Kafka client includes:

- An Apache Kafka client that provides [cluster discovery]({{< ref "client/kaffe" >}}) and [Telemetry integration]({{< ref "monitoring" >}}).
- Default property injection and automatic, secure delivery of credentials and secrets allow you to connect securely to a cluster
  with only three properties: **client**, **namespace**, and **secret**.
- [Kafka Sidecar]({{%baseurl%}}/dev-and-ci.html#local-environments) for developing
  locally with ACI Kafka.
- Critical bug fixes are proactively back-ported when necessary. Open-source software releases may be delayed when releasing versions
  with these patches.
- The ability to report your configuration to the ACI Kafka team, enabling the team to efficiently troubleshoot issues and plan
  migrations more easily.

### ACI Kafka with third-party clients [¶](#aci-kafka-with-third-party-clients "Link to this heading")

In some situations, the ACI Kafka client may not be suitable for your use cases. For example:

- If you are not using a Java language.
- If you don’t directly control the code that calls Kafka, such as third-party plugins or libraries.

For [third-party (non-Java client) use cases]({{< ref "client/kafka_for_all" >}}), you can use an ACI Kafka connect string and Simple Authentication and Security Layer/Secure Sockets Layer (SASL/SSL).

### Self-service [¶](#self-service "Link to this heading")

The ACI Kafka [**Self-service**]({{< ref "self-service" >}}) offering aims to provide you with autonomy. Below are ways to help you manage your entities:

- Automate the management of ACI Kafka resources using a secured REST API.
- Enable authorized team members and users to control or view Kafka resources.
- Manage Topics, Clients, and Access Control configuration.
- View and inspect the health of your topics and consumer groups.
- Modify who can access your entities.
- Automate resource management using a command-line interface (CLI) tool.

### ACI Kafka caveats and limitations [¶](#aci-kafka-caveats-and-limitations "Link to this heading")

Below are the specified conditions when using ACI Kafka:

- The ACI Kafka latest [release]({{< ref "releases" >}}) is based on Apache Kafka releases. The wire protocol
  negotiation is configured to support Apache
  Kafka clients from version 0.10.2 and later.
- [Kafka Streams](https://kafka.apache.org/documentation/streams/) are not officially supported.
- End-to-end encryption may be enabled in the ACI Kafka client for certain use cases.
- Size-based topic retention is required. An optional time-based retention can also be enforced.
- Client version greater or equal to 0.16 is required for [Java Development Kit 11 (JDK11)](https://java.apple.com/downloads/applejdk-11.html) support.
- Direct access to ZooKeeper is not allowed. Broker IPs or hostnames are provided through [Kaffe]({{< ref "client/kaffe" >}}).
- All Topic, Client, and Access changes are made using the ACI Kafka CLI or the ACI Kafka API.
  - Using the Kafka Admin Client is not supported.
- All multi-tenant clusters contain a default traffic quota per [client identity]({{%baseurl%}}/concepts.html#client-identities) per broker that a single tenant should not exceed.
  The default quota is 35% of the cluster’s average traffic capacity. This quota acts as safeguard to prevent one tenant from using all of the cluster’s resources.
- ACI Kafka limits the number of connections allowed in the broker at any time. The default is 65K per broker.


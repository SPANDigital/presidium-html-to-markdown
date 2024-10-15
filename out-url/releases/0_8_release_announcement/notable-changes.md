---
title: "Notable changes¶"
weight: 2
---

### Client Update[Internal Link: ¶](\#client-update)

In 0.8 the ACI Kafka Client was updated to better integrate into other
services and frameworks (e.g. MirrorMaker, Spark, etc.). You can read
more about these changes in the [External Link: New ACI Kafka Client
QTIP](https://github.pie.apple.com/pie/kafka-tools/blob/develop/improvement-proposals/019-new-client.md).

Summary of changes –

- _PieKafkaConsumer_ and _PieKafkaProducer_
  are deprecated in 0.8 and will be removed in 0.9
- _ConfigurationInterceptor_ is a new interface to allow
  for injecting configuration via Kaffe or other custom cluster
  discovery mechanism. See our docs for more information on
  [Internal Link: ConfigurationInterceptors](../client/using_client.html#configurationinterceptors).
- ACI Kafka Client now shades internal dependencies, and we now
  recommend you remove transitive inclusion of Apache Kafka from your
  dependencies. See [Internal Link: our
  instructions](../client/using_client.html#migrating-to-the-client)
  for excluding Apache Kafka from your transitive dependencies.

### Kafka ACLs & Consumer group id convention[Internal Link: ¶](\#kafka-acls-consumer-group-id-convention)

In 0.8 we will begin enforcing server-side ACLs, and the client now
requires consumer group names (i.e. _group.id_) to conform
to a new [Internal Link: naming
convention](../client/consumer_settings.html#consumer-group-and-aci-kafka-client).
We’ve begun reaching out to teams to coordinate migrating offsets for
changing consumer groups.

In our next release (0.9), brokers will begin enforcing consumer group
ACLs, so please plan to upgrade clients to 0.8 over the next few weeks.

Server-side ACL enforcement should be invisible to teams using ACI
Kafka, but if you experience any issues or if you have any questions on
the migration, please [Internal Link: contact us](../contact.html).

### Secure delivery of Kafka secrets[Internal Link: ¶](\#secure-delivery-of-kafka-secrets)

The ACI Kafka client now automatically receives and injects secrets for
your SASL and end-to-end encryption configuration. Teams already using
ACI Kafka were asked to provide key pairs for their clients, and now the
private key file can be directly configured to allow Kaffe to
automatically inject secrets.

See [Internal Link: our
docs](../client/client_secrets_distribution.html#configuring-your-client-to-enable-kaffe-configuration-injection)
for how to configure your client with private key files, and details on
how to implement your own private key provider.

### Kafka Sidecar[Internal Link: ¶](\#kafka-sidecar)

Local integration and testing with ACI Kafka is now possible with new
[Internal Link: Sidecar](../dev-and-ci.html#local-environments) support. Sidecars can
be used to bring-up Kafka and Zookeeper locally, and in future releases,
Kafka Sidecar will include tools for local topic and client management.
Later this spring, we’ll be adding support for ACI Kafka features like
Kaffe cluster discovery, configuration injection, and end-to-end data
security.

[Internal Link: Get started](../dev-and-ci.html#local-environments) with Kafka Sidecar,
and discover other Sidecars in [External Link: Rio’s Sidecar community
repo](https://github.pie.apple.com/pie/community-rio-sidecar-specs).

### Kafka Portal[Internal Link: ¶](\#kafka-portal)

Kafka Portal is a new interface for managing your ACI Kafka Resources.
Portal is powered by the new Kafka Public API which will be made
available to teams in a future release.

When teams on-board to ACI Kafka, the Queue team will create Customer
Groups and Namespaces with authorization set such that your team and
delegate teams can manage and/or view topics, clients, and access
control.

You can read more about Customer Groups, Namespaces, Group Access
Control, etc. in the [Internal Link: Concepts
overview](../concepts.html). We’ll be
reaching out to teams over the next few weeks to share links and further
instructions for accessing Kafka Portal.


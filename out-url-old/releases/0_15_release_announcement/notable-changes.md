---
title: "Notable changes¶"
weight: 3
---

The ACI Kafka 0.15 release includes new features available in the ACI
Kafka client.

**Summary of tent-pole changes**

- PIE Envelope Deprecation and Migration Plan
- Iris: a new front-end service for fetching keys for end-to-end (e2e)
  crypto
- Kafka CLI and Python APIs: Python tool and APIs for automating the
  management of ACI Kafka resources
- OJDK11 support for running in a Java 11 runtime

**ACI Kafka**

- [QTIP-028 - PIE Envelope\
  Deprecation]({{%baseurl%}}/https://github.pie.apple.com/pie/kafka-tools/blob/develop/improvement-proposals/028-envelope-deprecation.md),
  also see [ACI Kafka Docs]({{< ref "client/envelope" >}})
- [QTIP-029 - Iris Service for Scalable and Secure Key\
  Distribution]({{%baseurl%}}/https://github.pie.apple.com/pie/kafka-tools/blob/develop/improvement-proposals/029-iris-service.md)
- [ACI Kafka CLI](https://packages.apple.com/pypi/pie-kafka-cli)

**PIE Envelope Deprecation**

The PIE Envelope was introduced as a way to encapsulate record-level
metadata via custom extensions embedded within the record payload, but
since Apache Kafka 0.11, record headers are now supported as a first
class concept. In this release, we introduce new factory APIs for
constructing your producers and consumers, and these clients are capable
of writing envelope information either as PIE Envelope payloads or using
Apache Kafka record headers. Using the new APIs now allows teams to use
the value type of their records directly in the Apache Kafka APIs, no
longer exposing a PieEnvelope Java object and abstracting away the
underlying data format.

We plan to seamlessly migrate users from the PIE Envelope format to the
record header format using server-side flags, without requiring any
manual user intervention.

More details can be found in the [PIE Envelope Deprecation and Migration\
page]({{%baseurl%}}/../client/envelope.html#pieenvelope-deprecation-and-migration-details),
and for teams using Kafka for All, [ACI Kafka Java client\
interoperability (ACI Kafka\
0.14+)]({{%baseurl%}}/../client/kafka_for_all.html#aci-kafka-java-client-interoperability-aci-kafka-0-15)
provides an explanation of the implications for non-ACI Kafka clients.

**Iris: Front-end service for e2e crypto keys**

For users depending on the [e2e crypto\
feature]({{%baseurl%}}/../client/security_settings.html#payload-encryption-and-decryption-settings),
there is now a front-end service for key distribution, eliminating the
Cassandra dependency and direct connections to the backend data store.

If you’re currently using e2e crypto, you will need to add additional
configurations for the Iris connect string if you are not already using
[Kaffe configuration injection]({{< ref "client/kaffe" >}}). Other
configuration previously required for e2e crypto should also be removed.
For more details, [see\
here]({{%baseurl%}}/../client/security_settings.html#payload-encryption-and-decryption-settings).

**Kafka CLI**

Kafka CLI provides a way for customers to manage their own entities
programmatically. You can perform all the same operations on the CLI
that you can on ACI Kafka Portal, including creating and updating
topics, clients, and client-topic accesses. Customer groups and
namespaces are still created by the ACI Kafka team, and if you’re
interested in onboarding, please [contact\
us]({{%baseurl%}}/../contact.html#on-boarding-or-expansion).

This CLI is based on a Python library that can be integrated with
existing automation. You can learn more at [Apple\
PyPi](https://packages.apple.com/pypi/pie-kafka-cli).

**OpenJDK11 Support**

OpenJDK 11 is now supported for production use cases. Customers can now
run ACI Kafka clients 0.15 and higher in Java 11 environments.

**Reminders**

In case you missed it, since ACI Kafka 0.14.1, Apple’s corporate
certificate chain is now included in the ACI Kafka client artifact. This
supports [IS&T’s transition to the corporate certificate\
chain](https://istweb.apple.com/security-certificates/ssl-server-certificates-faq).
Upgrading to 0.15 obviates the need to install this certificate chain on
hosts connecting to ACI Kafka when all services eventually transition
away from GeoTrust-based certificates to Apple Corporate issued
certificates.

We also now support a [GSLB-enabled\
endpoint]({{%baseurl%}}/../client/kaffe.html#kaffe-endpoints)
for connecting to Kaffe in production. This is the recommended way for
connecting to Kaffe as it provides higher availability than configuring
DC-specific Kaffe endpoints.


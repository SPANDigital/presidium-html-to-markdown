---
title: "Kafka for All (non-Java clients)¶"
weight: 1
---

We recommend the [ACI Kafka client]({{< ref "client/quick_start" >}}) for everyone
using Java or other JVM languages.

The Java client utilizes Kaffe for configuration injection, but the
Kaffe client hasn’t been ported to other languages. For customers using
non-JVM languages, this guide provides an overview of the configuration
required to enable a third-party client.

In order to support non-Java clients the ACI Kafka team has created ways
to let you:

1. Connect to ACI Kafka clusters by providing a VIP(s) for configuring bootstrap servers
2. Fetch credentials (username and password) for authenticating to ACI Kafka clusters


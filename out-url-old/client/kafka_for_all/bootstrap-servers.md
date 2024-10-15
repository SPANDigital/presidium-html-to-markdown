---
title: "Bootstrap Servers¶"
weight: 5
---

Note

As there are currently no SLOs for the public API, it is not advised to rely on it to fetch the ACI Kafka bootstrap servers.

For Production purposes, store the list of bootstrap servers in your application config instead.

The `bootstrap.servers` property should be set to the `bootstrap-servers` property in the state of the Namespace.
You can fetch the Namespace state using the [CLI]({{%baseurl%}}/../self-service/kafka_cli.html#fetch-single-namespace) or the [API]({{%baseurl%}}/../self-service/public_api.html#fetch-single-namespace).


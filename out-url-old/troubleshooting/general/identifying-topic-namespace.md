---
title: "Identifying your topic, namespace and cluster¶"
weight: 2
---

You will usually find your topic name in the logs or client metrics that first indicated a problem.
Your ACI Kafka topic name will be prefixed with your [customer group and namespace]({{%baseurl%}}/../concepts.html#namespace).
For example `icloud.ck-content.prod.dc1.topic123` topic belongs to the `icloud.ck-content.prod.dc1` namespace (which itself belongs to the `icloud.ck-content.prod` customer group).

### Finding your ACI Kafka cluster name [¶](#finding-your-aci-kafka-cluster-name "Link to this heading")

Once you know your customer group and namespace name you can find the name of the ACI Kafka cluster your client is connecting to by fetching your Namespace information.
The cluster name will be listed in the `cluster-id` field of your ACI Kafka namespace.

#### Using the CLI or the API [¶](#using-the-cli-or-the-api "Link to this heading")

You can use [the CLI]({{%baseurl%}}/../self-service/kafka_cli.html#fetch-single-namespace) or [the Public API]({{%baseurl%}}/../self-service/public_api.html#fetch-single-namespace) if you’re familiar with them.

#### Using the API in the browser [¶](#using-the-api-in-the-browser "Link to this heading")

You can view your Namespace info in your browser by filling the customer group (e.g. `icloud.ck-content.prod`) and namespace name (e.g. `dc1`) in the following URL:

[https://ui-prod.aci-kafka.apple.com/api/public/v2/groups/(group-id)/namespaces/(namespace)](https://ui-prod.aci-kafka.apple.com/api/public/v2/groups/(group-id)/namespaces/(namespace))

#### From the bootstrap.servers [¶](#from-the-bootstrap-servers "Link to this heading")

If you’re configuring `bootstrap.servers` in your Kafka client configuration it should contain the name of the cluster:

```
# GSLB bootstrap endpoints
bootstrap.servers=a.<cluster-name-with-dashes>.clusters.kafka.g.apple.com:65506,...

```

```
# Legacy bootstrap endpoints
bootstrap.servers=<cluster-name-with-dashes>-piekafka.<dc>.if.apple.com:65506

```

Replace any `-` characters with `_` to get the cluster name.


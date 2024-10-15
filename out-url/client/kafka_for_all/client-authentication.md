---
title: "Client Authentication¶"
weight: 6
---

Each ACI Kafka _Client_ is associated with a SASL
username/password combination that can be used to connect to clusters
the client has been granted access to. ACI Kafka clusters are configured
to use a SASL\_SSL listener, and clients must be configured to connect to
it. The properties that are required are:

```
security.protocol=SASL_SSL
sasl.mechanism=PLAIN
sasl.jaas.config=org.apache.kafka.common.security.plain.PlainLoginModule required username="<username>" password="<password>";

```

In this case `<username>` is the fully qualified name of your ACI Kafka
`Client`, e.g., [Internal Link: `<identityRegex>`](../self-service/kafka_cli.html#client-identities)), and `<password>` can be obtained with help from
the CLI or API by retrieving the Client resource and running the
following command using the _encryptedPassword_ value:

```
echo "<encryptedPassword value>" | base64 --decode | openssl rsautl -decrypt -inkey private_key.pem

```


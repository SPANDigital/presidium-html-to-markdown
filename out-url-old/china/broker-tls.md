---
title: "Broker TLS¶"
weight: 5
---

As mentioned above, the brokers present TLS certificates rooted from the
Identity root. Using ACI Kafka client 0.16.4 or greater, you can simply
set the following properties, because the ACI Kafka client extends
upstream to provide support for PEM loading:

```
ssl.truststore.type=PEM
ssl.truststore.location=<path-to-certificate-volume>/trusted-root.pem

```

If you are using the vanilla Apache Kafka client, you will either need
to switch to [our drop in client]({{< ref "client/using_client" >}}) or
configure your own JKS truststore.

Note

When building a truststore and importing the trusted root from a PEM file,
you must ensure that this PEM file has only a _single certificate_ present.
We’ve observed that when a JKS is created with multiple certificates from a single PEM file,
the resulting truststores are “valid” with respect to `keytool`,
yet **fail to correctly trust either root during usage in a JVM application**.

An example of creating your own JKS, using the `trusted-root.pem` file
from the Identity mount:

```
openssl x509 -outform der -in trusted-root.pem -out trusted-root.der
keytool -import -trustcacerts -alias <unique-name> -file trusted-root.der -keystore cacerts.jks -storepass changeit

```

and using the above, with the corresponding Kafka properties:

```
ssl.truststore.location=/location/of/your/truststore.jks
ssl.truststore.password=changeit

```


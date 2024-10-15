---
title: "TLS and Authentication Settings (required)¶"
weight: 2
---

The ASE Kafka client communicates with brokers over TLS, and depending
on the platform where the brokers are running, brokers will use
certificates signed different Root CAs. The client can be configured
to trust the broker’s certificate with the following configuration:

```
ssl.truststore.type=PEM
ssl.truststore.location=path/to/trusted-roots.pem

```

The `trusted-root.pem` file should include is available [Internal Link: here](../_static/trusted-roots.pem).

In some cases the JVM may trust the broker’s certificate by default, but the [External Link: Apple JVM recommendation](https://java.apple.com/faq.html#q-my-application-needs-root-certificate-xxx-why-is-it-not-included-in-applejdk) is “not to rely upon the default trust store that comes with AppleJDK for your production use cases, and instead pull in a trust store that is appropriate for your environment and expected/trusted peers”.

ASE Kafka expects plain-text SASL authentication over this encrypted
connection.

### SASL Settings[Internal Link: ¶](\#sasl-settings)

For most environments, the following properties are sufficient to enable
SASL/PLAIN-based authentication over TLS for producer and consumer
clients:

```
security.protocol=SASL_SSL
sasl.mechanism=PLAIN
# Keep this property secret as it includes a password
sasl.jaas.config = org.apache.kafka.common.security.plain.PlainLoginModule required username="<username>" password="<password>";

```

**If you have enabled**[Internal Link: Kaffe configuration
injection](default_config_injection.html), **then you don’t require**
**to set these properties.**

**If you are using Apache client version 2.0+, then you need to include**
**this extra property:**

```
ssl.endpoint.identification.algorithm=   # an empty string

```


---
title: "Client Configuration¶"
weight: 5
---

You may need to adjust some client-side configurations for Kafka in
order for things to work. In the production scenario, you don’t need to
set these properties, as [Kaffe]({{< ref "client/kaffe" >}}) will handle these
for you.

### Enable Development Mode [¶](#enable-development-mode "Link to this heading")

You only need to add the following property in order to connect to the
sidecar.

```
pie.queue.kafka.dev.mode.enabled=true

```

We populate the `bootstrap.servers` string automatically for you using
the environment variables provided by the sidecar. Make sure these are
set when working locally. When working on a Rio pipeline, these will be
set for you as part of starting the sidecar in a pipeline.

### Security Protocols [¶](#security-protocols "Link to this heading")

The sidecar definition above only supports the PLAINTEXT protocol.
This differs from QA and Production environments where we mandate the use
of `SASL_SSL`.

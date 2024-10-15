---
title: "Continuous Integration¶"
weight: 4
---

When running Sidecars in Rio, you are required to read the connection
information for Kafka from the Sidecar provided environment variables.
This is because Rio controls running the Sidecar docker container, and
with dynamic port assignment.

Setting dev mode (covered below) will mean that the ACI Kafka client
will automatically read these environment variables, and populate the
`bootstrap.servers` for you.


---
title: "Configuring your client to enable Default Configuration Injection¶"
weight: 2
---

In ACI Kafka Client version 0.14 and higher, this functionality is
enabled by default (with one exception for producer properties, noted
below).

In ACI Kafka Client versions 0.9 through 0.13, you can enable this
feature by setting the following property in your client configuration:

```
pie.queue.kaffe.config.injection=true

```


---
title: "Kaffe¶"
weight: 4
---

The URL for Kaffe Connect is:

```
pie.queue.kaffe.connect = https://kafka-kaffe-prod.cn.g.silu.net:14684

```

In China, you will always need to set the following Kaffe properties:

```
pie.queue.kaffe.tls.trusted.certs.path=<path-to-certificate-volume>/trusted-root.pem
pie.queue.kaffe.tls.hostname.verification=false

```

Disabling hostname verification is required to use the VIP until all the
required SANs are present. We will update this documentation once that
has been finished.

Note

Access to the Kaffe VIP requires a new entitlement in China, that wasn’t required in SilkRoad.
You need to request an outgoing entitlement for “GCBD VIP Networks”.

The entitlement required is self add, using the `ent` CLI:

```
ent request --network-out "GCBD VIP Networks" <your-namespace-name>.kube

```


---
title: "SDR Identity Certificates¶"
weight: 3
---

Where possible, ACI Kafka utilizes the [External Link: SDR Identity
certificates](https://kube.apple.com/docs/guides/network-security-and-traffic-control/#network-security-and-traffic-control)
provided on top of ACI Kube. You can find details of how to mount the
per-application certificates under [External Link: Appendix Application
Certificates](https://kube.apple.com/docs/guides/network-security-and-traffic-control/#appendix-application-certificates).

In China, these are rooted from a different CA to US AODCs.

The above certificates are used for at least:

- ACI Kafka brokers, when handling TLS connectivity
- Kaffe gRPC service


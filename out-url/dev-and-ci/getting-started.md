---
title: "Getting Started¶"
weight: 2
---

The ACI Kafka team suggests using a [External Link: Rio
Sidecar](https://docs.aci.apple.com/rio/build-features/sidecars.html),
providing customers with supported tooling for managing
[Internal Link: local](#local-environments), ephemeral,
and lightweight ACI Kafka environments.

The provided Rio Sidecar can also be used as part of a
[Internal Link: CI process](#continuous-integration). Both
scenarios are covered by this section.

The Rio Sidecar is published at the same cadence as our general
releases, using the same versioning semantics as ACI Kafka itself.

### Supported Features[Internal Link: ¶](\#supported-features)

We suggest using the upstream [External Link: docker image](https://hub.docker.com/r/apache/kafka)
for the sidecar. That means some ACI Kafka features
like [Internal Link: Kaffe](client/kaffe.html) or the Public API are not available.

If your use case requires accessing any of these features, consider
[Internal Link: applying](onboarding.html) for a qualification namespace.


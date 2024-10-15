---
title: "Using the client¶"
weight: 1
---

The client is a simple java dependency available in
[External Link: artifactory](https://docs.aci.apple.com/rio/infrastructure/artifactory.html)
at these coordinates:
`com.apple.pie.queue:kafka-pie-client:{version}`.
Check the current [Internal Link: version recommended for production](../releases/index.html#recommended-release).

All items produced and consumed are wrapped in a common
[Internal Link: envelope](envelope.html).

You can extend the envelope yourself to add your own fields as explained
in the documentation of the [Internal Link: envelope](envelope.html). We urge users
to use the standard extension points rather than re-wrapping messages in
a nested envelope.

The javadoc for the client is pretty complete; this document will only
present high-level concepts.


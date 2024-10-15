---
title: "Using the client¶"
weight: 1
---

The client is a simple java dependency available in
[artifactory](https://docs.aci.apple.com/rio/infrastructure/artifactory.html)
at these coordinates:
`com.apple.pie.queue:kafka-pie-client:{version}`.
Check the current [version recommended for production]({{%baseurl%}}/../releases/index.html#recommended-release).

All items produced and consumed are wrapped in a common
[envelope]({{< ref "client/envelope" >}}).

You can extend the envelope yourself to add your own fields as explained
in the documentation of the [envelope]({{< ref "client/envelope" >}}). We urge users
to use the standard extension points rather than re-wrapping messages in
a nested envelope.

The javadoc for the client is pretty complete; this document will only
present high-level concepts.


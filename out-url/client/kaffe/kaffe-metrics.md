---
title: "Kaffe metrics¶"
weight: 7
---

Both Kaffe client and server are able to publish their own Metrics. The server uses the raw Hubble agent, whereas the
client uses a regular reporter so that you can retrieve these metrics in other systems as well.

KPIs need to be improved, but they are all described in
[External Link: KaffeClientEvents](https://github.pie.apple.com/pie/kafka-tools/blob/develop/kaffe-client/src/main/java/com/apple/pie/queue/kaffe/KaffeClientEvents.java)
and
[External Link: KaffeServerEvents](https://github.pie.apple.com/pie/kafka-tools/blob/develop/kaffe-service/src/main/java/com/apple/pie/queue/kaffe/server/KaffeServerEvents.java).


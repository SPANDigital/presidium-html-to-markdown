---
title: "Consumer Properties¶"
weight: 8
---

The consumer group id ( `<groupId>` property) should always have the
client id as the prefix. The purpose of this requirement is to support
consumer group ACLs.

For instance, when you try to initiate a consumer with a client id
`pie.queue.test.client1`, your `<groupId>` property should be in the
format of `pie.queue.test.client1.xxx`.

More details about the consumer group convention can be found in the
[Internal Link: Recommended Consumer Settings](consumer_settings.html)


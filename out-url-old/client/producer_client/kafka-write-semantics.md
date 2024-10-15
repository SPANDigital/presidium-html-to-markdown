---
title: "Kafka Write semantics¶"
weight: 5
---

### At-Least-Once [¶](#at-least-once "Link to this heading")

By default, Kafka Producer guarantees at-least-once delivery with `retries` set to 2147483647 by default.
This value allows the Producer never to drop a message but might duplicate the occasional message.


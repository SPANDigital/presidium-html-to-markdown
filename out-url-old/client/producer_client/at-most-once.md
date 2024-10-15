---
title: "At-Most-Once¶"
weight: 6
---

At-most-once writes means that the Producer will never retry to write message twice; instead, it will skip writing the message on failure.
This semantic can be achieved if `retries` is set to 0.


---
title: "Tracing a session¶"
weight: 6
---

One of the big advantages of Kaffe is that it enables us to know how
your clients are configured and what they are up to.

Whenever you are investigating a problem and reaching out to us, please
retrieve the sessionId which is logged on startup in a message:

```
Client started sessionId='{}' client='{}'"

```

Once we have this we can look at our state tracker (currently Splunk)
and see your client version, the different things you’ve been asking
Kaffe for, and the properties you are using.

This session tracing also enables us to know which clients are running
which version and be proactive in the case a serious bug is discovered.


---
title: "General troubleshooting guide¶"
weight: 1
---

When investigating timeouts, delays or other issues related to Kafka, the first step is to try and identify the impact and the source of the problem.

There are three common categories of problems users run into:

1. Client-side and application issues
2. Network or ACL problems
3. Issues with the Kafka cluster

Useful questions to try to answer in order to identify the cause include:

1. Is the issue happening on all service/application instances or a subset / single instance?
   A single instance showing errors could indicate a hardware issue with the host for example.
2. Have there been any changes recently (like a new deployment or a configuration change) that correlate with the time you’ve started seeing issues?
3. What is the impact on your application?
   A lot of [Internal Link: client errors](client.html#common-errors) are retryable and don’t always indicate an immediate problem.


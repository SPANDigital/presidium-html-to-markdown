---
title: "Inspector¶"
weight: 1
---

Note

ACI Kafka offers a metrics based solution to share topic partition health and client activity (including consumer lag and actual usage).
For more details check [broker metrics]({{< ref "monitoring/broker_metrics" >}}).

Inspector is a tool created by Apple Media Group (AMP) and known by the
name kafka-webapi. It enables you to see the health of topics, their
offsets and active consumer groups. This tool is extensively used by our
operations team, and we provide access to customers to provide some
visibility to your namespace.

This tool is provided as-is, without a Production SLA. We provide a best
effort support model for Inspector.

You can view inspector at the following URI’s;

- For Prod: [https://ui-prod.aci-kafka.apple.com/inspector/](https://ui-prod.aci-kafka.apple.com/inspector/)
- For IF1: [https://ui-if1.aci-kafka.apple.com/inspector/](https://ui-if1.aci-kafka.apple.com/inspector/)

You should then be able to navigate to the cluster of choice, which will
present you the following view:

![]({{%baseurl%}}//images/inspector.png)

You’ll be able to see all topics in all namespaces (for which you have
at least visibility access) that are running on this physical cluster.

As this is mainly an operational tool, a lot of the information present
here may not be useful to you. However, you should be able to see the
health of a topic and if you are writing and reading from it.

**Note: The lag of consumers is reported only if you use the built in**
**Kafka consumer offset committing, if you don’t use this you won’t be**
**able to track backlog.**

---
title: "Common Client Problems¶"
weight: 1
---

Most client problems can be fixed at the client side with some basic troubleshooting and config tuning.
Please use this page during a client incident before contacting the Kafka team.
If your problem isn’t listed here, or you’ve tried all the listed solutions, then contact us with the details of the issue including solutions tried and the output of each one.

- [Internal Link: Common errors](#common-errors)
- [Internal Link: All Clients](#all-clients)
  - [Internal Link: Is this a Kafka’s exception?](#how-to-find-the-exact-exception-from-kafka)
  - [Internal Link: Is my cluster healthy?](#is-my-cluster-healthy)
  - [Internal Link: Too many logs saying `Broker may not be available`](#too-many-logs-saying-broker-may-not-be-available)
- [Internal Link: Consumer Client](#consumer-client)
  - [Internal Link: Too many rebalancing events](#consumer-group-is-experiencing-too-many-rebalancing-events)
  - [Internal Link: Large Lag](#consumer-is-experiencing-large-lag)
  - [Internal Link: Can’t commit offsets](#consumer-failed-to-commit-offsets)
- [Internal Link: Producer Client](#producer-client)
  - [Internal Link: Can’t send records](#producer-failed-to-send-record)
  - [Internal Link: Missing data](#missing-data)
  - [Internal Link: Throughput or latency issues](#throughput-or-latency-issues)


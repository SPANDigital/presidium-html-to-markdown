---
title: "Common Client Problems¶"
weight: 1
---

Most client problems can be fixed at the client side with some basic troubleshooting and config tuning.
Please use this page during a client incident before contacting the Kafka team.
If your problem isn’t listed here, or you’ve tried all the listed solutions, then contact us with the details of the issue including solutions tried and the output of each one.

- [Common errors](#common-errors)
- [All Clients](#all-clients)
  - [Is this a Kafka’s exception?](#how-to-find-the-exact-exception-from-kafka)
  - [Is my cluster healthy?](#is-my-cluster-healthy)
  - [Too many logs saying `Broker may not be available`](#too-many-logs-saying-broker-may-not-be-available)
- [Consumer Client](#consumer-client)
  - [Too many rebalancing events](#consumer-group-is-experiencing-too-many-rebalancing-events)
  - [Large Lag](#consumer-is-experiencing-large-lag)
  - [Can’t commit offsets](#consumer-failed-to-commit-offsets)
- [Producer Client](#producer-client)
  - [Can’t send records](#producer-failed-to-send-record)
  - [Missing data](#missing-data)
  - [Throughput or latency issues](#throughput-or-latency-issues)


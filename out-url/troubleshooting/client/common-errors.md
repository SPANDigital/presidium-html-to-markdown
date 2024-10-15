---
title: "Common Errors¶"
weight: 2
---

This section lists some of the common messages from the Kafka client you might see in your application logs.

### Connection to node 0000 could not be established[Internal Link: ¶](\#connection-to-node-0000-could-not-be-established)

This warning is frequently logged by the client in response to normal cluster maintenance operations.
It does not always suggest an issue by itself.
If you’re seeing an impact on your service, please refer to the instructions in [Internal Link: Too many logs saying Broker may not be available](#too-many-logs-saying-broker-may-not-be-available).

### Disconnecting from node 0000 due to … timeout[Internal Link: ¶](\#disconnecting-from-node-0000-due-to-timeout)

This warning is frequently logged by the client in response to normal cluster maintenance operations.
It does not always suggest an issue by itself.
If you’re seeing an impact on your service, please refer to the instructions in [Internal Link: Too many logs saying Broker may not be available](#too-many-logs-saying-broker-may-not-be-available).

### Timeout of 60000ms expired before …[Internal Link: ¶](\#timeout-of-60000ms-expired-before)

Very large timeouts (>30s) often indicate a network problem.
A good start is to check [Internal Link: you have connectivity to the brokers from the affected application instance](connectivity.html) and that the [Internal Link: cluster is healthy](general.html#checking-cluster-health).

Similar messages to look for:

- `Failed to update metadata after 60000 ms.`
- Other instances of `org.apache.kafka.common.errors.TimeoutException`

### Not authorized to access topics[Internal Link: ¶](\#not-authorized-to-access-topics)

Make sure you’re using the [Internal Link: fully qualified topic name](../concepts.html#topic-name) and [Internal Link: the topic allows access for your client identity](../concepts.html#client-identity-topic-access).
You should also check that [Internal Link: `bootstrap.servers` configuration](../client/kafka_for_all.html#bootstrap-servers) is matching the correct cluster for the topic’s namespace.

Similar messages to look for:

- Other instances of `org.apache.kafka.common.errors.TopicAuthorizationException`

### Member … sending LeaveGroup request to coordinator …[Internal Link: ¶](\#member-sending-leavegroup-request-to-coordinator)

Please see the section on [Internal Link: consumer group experiencing too many rebalancing events](#consumer-group-is-experiencing-too-many-rebalancing-events)

Similar messages to look for:

- `... due to consumer poll timeout has expired`


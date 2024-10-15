---
title: "All Clients¶"
weight: 3
---

Some problems are common between all Kafka clients.
Usually these are due to network, bootstrapping, collecting metadata, authentication and authorization issues.
Here are some of the most common ones:

### How to find the exact exception from Kafka? [¶](#how-to-find-the-exact-exception-from-kafka "Link to this heading")

#### Description [¶](#description "Link to this heading")

Sometimes a client gets interrupted and throws customised error messages or exceptions that mask the root cause.
Which makes it harder to determines if this is Kafka related issue or related to client’s business logic.

#### Impact [¶](#impact "Link to this heading")

Client is interrupted and it is taking longer to track down the root cause.

#### Causes and Solutions [¶](#causes-and-solutions "Link to this heading")

As mentioned above the cause of this confusion is down to the fact that client’s logic is masking Kafka errors.

_Solution:_

- Trace the error back to Kafka’s logs and exceptions. These will be prefixed with:
  - `org.apache.kafka`: In this case; check the rest of this page to see if these errors/logs are listed here in the common client’s issues and follow the solutions.
  - `com.apple.pie.queue`: If client has errors from `com.apple.pie.queue` package please check the connectivity with [Kaffe endpoints]({{%baseurl%}}/../client/kaffe.html#kaffe-endpoints) using `nc` bash command.
- If no fatal Kafka errors were found then check if the client has any other errors with other systems and address these first.

If you still suspecting Kafka but don’t have enough evidences please check section [Is my cluster healthy?](#is-my-cluster-healthy).

### Is my cluster healthy? [¶](#is-my-cluster-healthy "Link to this heading")

#### Description [¶](#id1 "Link to this heading")

Kafka logs are a bit verbose and chatty which makes Kafka look like the first suspect during a client incident.
This may lead the client’s owner to question if the cluster is healthy or if they’re the only one experiencing a problem.

#### Impact [¶](#id2 "Link to this heading")

The impact of client interruptions varies by use case.
However, in most instances, the cluster is healthy, and excessive logging from Kafka clients can divert the client’s owner from addressing the actual issue, leading to wasted time.

#### Causes and Solutions [¶](#id3 "Link to this heading")

Excessive logging in Kafka might be due to the log.level setting.
However, even at the lowest log level, Kafka still generates numerous informational logs.
Most Kafka [errors are informative](#too-many-logs-saying-broker-may-not-be-available), and the client can recover from them without any intervention.

_Solution:_

- [Check that the Kafka cluster is healthy]({{%baseurl%}}/general.html#checking-cluster-health).
- Produce and consume experience are monitored through [Synthetic Metrics]({{%baseurl%}}/../monitoring/broker_metrics.html#synthetic-metrics).

### Too many logs saying `Broker may not be available` [¶](#too-many-logs-saying-broker-may-not-be-available "Link to this heading")

#### Description [¶](#id4 "Link to this heading")

Kafka clients connect to Kafka brokers to poll or produce data, or collect metadata from the Kafka cluster.
Sometimes the connection to a few brokers is unavailable due to a network issue, or a planned cluster bounce (explicit maintenance by the Kafka team).
In this case, the client will log `Connection to node 0000 could not be established. Broker may not be available.`, or `Disconnecting from node 0000 due to request timeout`.
These log messages are to provide information only and should not be considered an error.

#### Impact [¶](#id5 "Link to this heading")

Both temporary and permanent connectivity issues will log these informative messages without actually telling the client if this a fatal error or not.
This alone doesn’t help the applications owner to detect a connectivity issue.
The client owner must check other metrics, logs, and run basic network troubleshooting, before making the conclusion there’s a connectivity problem.

#### Causes and Solutions [¶](#id6 "Link to this heading")

- Make sure to [test the connectivity]({{< ref "troubleshooting/connectivity" >}}) with the brokers before investigating any causes.
- [Check that the Kafka cluster is healthy]({{%baseurl%}}/general.html#checking-cluster-health).
- Check client metrics, such as:
- `producer-metrics-record-retry-rate` which shows how many retry requests the producer sent.
  If this is not trending down immediately it means the connectivity issue is permanent.
- `record-send-rate` which shows the average number of records sent per second.
  If this is not showing the expected produce rate per second, then this might be a permanent connectivity issue.
- `consumer-fetch-manager-metrics-bytes-consumed-rate` which shows average number of bytes consumed per second for a topic.
  If this is not showing the expected consume rate per second, then this might suggest a permanent connectivity issue.

If this basic troubleshooting indicates the connectivity is:

1. a temporary issue, meaning that the client is recovering on the next retry of the consumer [poll]({{%baseurl%}}/../client/consumer_client.html#poll) function, or a producer successfully [retried sending the batch]({{%baseurl%}}/https://docs.aci.apple.com/kafka/client/producer_client.html#sender-thread).
   Or that the [cluster is healthy]({{%baseurl%}}/general.html#checking-cluster-health).

_Solution:_

- There is no need to act on this informative log; Kafka clients are designed intentionally to treat most of these errors as retryable.
  Usually, another replica will become a leader for the partition and handle the requests.

1. a permanent issue, meaning that the client isn’t recovering on the retry.

_Solution:_

- Check if you need any extra network ACLs or configuration to connect to Kafka brokers.
- If this is a producer client, make sure `retries` config is set to a higher number so the client will not give up on the batch before the temporary network error resolves itself.
  Check our [recommendation for best delivery guarantee]({{%baseurl%}}/../client/producer_settings.html#best-delivery-guarantee) for more details regarding how to choose an appropriate value.
- If this is a consumer, make sure you are consuming in a loop as consumer doesn’t have in-built retry logic.
- Make sure that `bootstrap.servers` config is to the correct cluster.
- Increase `session.timeout.ms` to an appropriate value if the client is deployed in different DC than the Kafka cluster’s DC. For example client is connecting from USA based DC to Asia based DC.


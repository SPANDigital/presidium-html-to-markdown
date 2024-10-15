---
title: "Common problems or errors you may face¶"
weight: 5
---

### Why my lag doesn’t go down?[Internal Link: ¶](\#why-my-lag-doesn-t-go-down)

Possible causes could be:

- Double check you are filtering out inactive consumer instances from your query/alert by adding `consumer_id!="unknown"`.
- If this isn’t working check the history for the last few days to see if your consumer group ever reached zero lag (or near zero lag).
  If not, check [Internal Link: the details of large lag](../troubleshooting/client.html#consumer-is-experiencing-large-lag) for more details on how to fix this.
- If you are monitoring lag in time, you could be misled if the producer is replaying old records with old the timestamp.

### Why my lag is zero?[Internal Link: ¶](\#why-my-lag-is-zero)

Your lag will be zero when you consume everything in the topic.
However, if you believe your lag shouldn’t be zero then check if the producer is still publishing data.

More common problems with broker metrics can be found [Internal Link: here](../monitoring/broker_metrics.html#common-problems-or-errors-you-may-face).

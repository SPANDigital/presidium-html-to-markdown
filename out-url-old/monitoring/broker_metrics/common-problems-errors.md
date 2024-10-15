---
title: "Common problems or errors you may face¶"
weight: 9
---

### Why am I getting a 403 error when accessing the Grafana dashboards, or setting up alerts? [¶](#why-am-i-getting-a-403-error-when-accessing-the-grafana-dashboards-or-setting-up-alerts "Link to this heading")

The possible causes could be:

1. You are trying to access the metrics before Mosaic applied the sharing policy in their system.
   After the ACI Kafka team shared the metrics with your team, it can take up to an hour for the changes to be reflected.
2. You don’t have access to your team’s Apple Directory groups. Please double-check that you have the right access.
3. You are trying to use regular expressions in your query. Mosaic doesn’t allow tenants to query shared metrics using regular expressions.
   You will need to use the exact label names. These are `_ns_`, `namespace`, or `customer_user` or `client` for lag metrics.

If you still have issues contact the Mosaic team on Slack in [`#help-aci-mosaic`](https://a1391190.slack.com/messages/CHNFM578E)
with your sharing policies IDs. You can list your sharing policies IDs using the (sharing-policies endpoint)\[https://docs.aci.apple.com/service\_apis/mosaic-service.html#security-apis-sharing-policies-get\] of the Mosaic API.

### Why am I seeing small gaps in my metrics? [¶](#why-am-i-seeing-small-gaps-in-my-metrics "Link to this heading")

This might happen when the ACI Kafka team is bouncing clusters and auxiliary services.
Ideally you shouldn’t see more than 5 minute gaps in the metrics.
If the gap is larger than 5 minutes, [please contact ACI Kafka after checking the production incident Slack channel to see if there is any active incident]({{< ref "contact" >}}).

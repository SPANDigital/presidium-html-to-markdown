---
title: "Common problems or errors you may face¶"
weight: 8
---

### Why am I getting a 403 error when accessing the Grafana dashboards, or setting up alerts?[Internal Link: ¶](\#why-am-i-getting-a-403-error-when-accessing-the-grafana-dashboards-or-setting-up-alerts)

The possible causes could be:

1. You are trying to access the metrics before Mosaic applied the sharing policy in their system.
   After the ASE Kafka team shared the metrics with your team, it can take up to an hour for the changes to be reflected.
2. You don’t have access to your team’s Apple Directory groups. Please double-check that you have the right access.

If you still have issues contact the Mosaic team on Slack in [External Link: `#help-ase-mosaic`](https://a1391190.slack.com/messages/CHNFM578E)
with your sharing policies IDs. You can list your sharing policies IDs using the (sharing-policies endpoint)\[https://docs.aci.apple.com/service\_apis/mosaic-service.html#security-apis-sharing-policies-get\] of the Mosaic API.

### Why am I seeing small gaps in my metrics?[Internal Link: ¶](\#why-am-i-seeing-small-gaps-in-my-metrics)

This might happen when the ASE Kafka team is bouncing clusters and auxiliary services.
Ideally you shouldn’t see more than 5 minute gaps in the metrics.
If the gap is larger than 5 minutes, [Internal Link: please contact ASE Kafka after checking the production incident Slack channel to see if there is any active incident](../contact.html).

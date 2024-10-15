---
title: "Paging for Urgent Support¶"
weight: 6
---

For **urgent production, customer impacting issues** and outages you can
page us.

ACI Kafka does not use **Help Central** for receiving customer support requests.
Any tickets you raise there will not receive a timely response.

Note

Please be aware that we have a pager rotation which includes
colleagues in the London timezone (UTC/GMT), so a page during Cupertino
office hours may wake up a member of the team in London. Please be
considerate, and if your issue does not require immediate attention
please start with some of the other channels like Slack or email.

If you have [followed the troubleshooting guide]({{< ref "troubleshooting/general" >}}) and determined that you need to page us during an incident - it
is required to provide the following information so that we can assist
you effectively:

1. The fully qualified Topic(s), or Namespace(s) (e.g.
   `aci.queue.prod.pv-namespace.events-topic`) for which you have
   observed an issue
2. A brief explanation of the impact (e.g. service outage, data loss,
   severe performance degradation, etc)
3. Some metrics or logs that show the observed the change in
   behavior/issue you are experiencing
4. What actions you have already taken to remediate the issue (if
   applicable)
5. A brief description of how your use case uses Kafka

The paging email for Queue support is
**BCC**: [aci-queue@applepie.pagerduty.com]({{%baseurl%}}//files/mailto:?bcc=aci-queue%40applepie.pagerduty.com).

A few things to note when using this alert:

1. It will create a new PagerDuty ticket, so please use BCC: for that
   address, otherwise every email reply will create a new PagerDuty
   ticket.
2. It should be used for genuine emergencies only. Normally an issue in
   a QA environment is handled during business hours only. If your
   issue can wait for a few hours then consider another channel.

Regarding a company shutdown like the Thanksgiving or Christmas breaks,
there are often change moratoriums in place so that everyone can get a
much deserved break.

If a group is planning on working during that time and will need
additional support, please have a valid business reason for doing so,
and make prior arrangements with the SRE team(s) that may be impacted so
that we’re in the best position to address any support needs that may
arise.


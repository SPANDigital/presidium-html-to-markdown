---
title: "Mirrors¶"
weight: 3
---

Mirrors is a new service - still in controlled availability - which
replicates messages from one topic to another, typically between
clusters in different physical locations. Mirrors offer an “at-least-once”
message delivery guarantee, we aim to minimize the number of duplicates, although in
some circumstances duplicates can still occur.

### Mirroring Latency [¶](#mirroring-latency "Link to this heading")

We measure mirror latency at the mirror cluster, as the number of
messages that we have processed that have a timestamp older than 5
minutes at the point of mirroring. Please note that latencies will
increase depending on network distance.

As a note, we have two SLOs for Mirrors, one for clusters that have
China as a Source / Destination, and the other for all other
production Mirror clusters. We currently offer a lower SLO for China
as we have observed a higher chance for network failures, due to the increase in
network physical distances.

We plan to measure latency for Mirrors as being the time taken for a
message to be read from the source topic and written to the destination
topic.

### Mirroring Availability [¶](#mirroring-availability "Link to this heading")

We measure mirror availability in two ways:

- **Attempts vs Exceptions** We get exceptions if we encounter any
  terminal issues - this is one of the ways that it is possible to
  loose data when being mirrored.
- **Attempts vs Attempts - Successes** we also measure availability as
  the ratio of attempts to attempts that didn’t end up as being
  recorded as a success.


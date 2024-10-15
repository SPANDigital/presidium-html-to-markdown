---
title: "Customer group¶"
weight: 2
---

_Customer groups_ provide a mechanism to group client identities and namespaces.
A customer group consists of a _customer name_, a _group name_, and a logical _environment_ (as defined by the group owner).

Customer names are determined by the ACI Kafka team and indicate high-level ownership.
Customer names have a maximum length of 15 characters and can contain letters, numbers, and `_` or `-` characters.

Group names indicate a team-level or use-case-level ownership.
Customer names have a maximum length of 31 characters and can contain letters, numbers, and `_` or `-` characters.

You can request one or more customer groups based on the number of separate environments your team or project needs.
Environment names have a maximum length of 15 characters and can contain letters, numbers, and `_` or `-` characters.

Below are examples of fully qualified customer group identifiers:

- `aci.queue.prod`
- `icloud.messaging.qa1`

### Customer group access control[Internal Link: ¶](\#customer-group-access-control)

_Customer groups_ define access controls for namespaces, topics, and client identities owned by the group.
Access permissions are managed through the _Modification Groups_ and _Visibility Groups_ as defined below.

**Modification Groups**: Users with membership in any of the modification Apple Directory groups can create and edit ACI Kafka resources like topics, client identities, and topic accesses.

**Visibility Groups**: Users with membership in any of the visibility Apple Directory groups have read-only access to the ACI Kafka resources.

### Customer group contact information[Internal Link: ¶](\#customer-group-contact-information)

_Customer group contact information_ should contain your team’s support and contact details as specified below.
It should include valid Apple Directory groups that are email-enabled and can accept emails from anyone in Apple.

**Contact Group**: Used to send automated reports and ACI Kafka announcements.

**Pager Contact Group**: Used to page your team (for example using **PagerDuty**) if required to resolve an active incident.

**Alert Contact Group**: Used to communicate alerts and important maintenance activities.


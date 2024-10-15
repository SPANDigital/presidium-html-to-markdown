---
title: "Namespace¶"
weight: 4
---

A _namespace_ represents a logical and physical grouping of related Kafka topics.
Each namespace is unique across ACI Kafka and is linked to a customer group.

All topics within a namespace are located on the same Kafka cluster.

Each namespace may have constraints that limit the movement of its topics. For example, a namespace may be restricted to clusters running in a specific datacenter.

Namespace names have a maximum length of 63 characters and can contain letters, numbers, and `_` or `-` characters.

A fully qualified namespace identifier is written as: `customer.group.env.name`, for example:

- `pie.queue.prod.pv`
- `icloud.cloudkit.icloud4.p02`

A namespace is used as a prefix before the topic name, separated by a period (‘.’).

### Namespace quotas [¶](#namespace-quotas "Link to this heading")

ACI Kafka requires configuring namespace quotas for _produce_, _consume_, and _storage_ resources.
These namespace-level quotas act as a budget for all topics within the same namespace as defined below:

**Produce quota**: The rate at which clients will cumulatively _produce_ (also known as _Bandwidth IN_) messages to topics in a namespace.

**Consume quota**: The rate at which clients will cumulatively _consume_ (also known as _Bandwidth OUT_) messages from topics in a namespace.

**Storage quota**: The total allocated _storage_ (also known as _Namespace Storage_) capacity for a namespace defines the upper limit for the sum of storage requirements for all topics within the namespace.

Consume and produce quotas are used for cluster capacity planning and are not enforced by Kafka.

Any operation, like creating new topics or updating existing topics that would result in the namespace exceeding the quotas will be rejected.

When requesting ACI Kafka onboarding, you should know your storage and throughput requirements.
These are captured in the questionnaire in the form of produce, consume, and storage numbers.
Providing these requirements ensures quicker onboarding.


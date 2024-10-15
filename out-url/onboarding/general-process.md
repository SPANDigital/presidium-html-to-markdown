---
title: "General Process¶"
weight: 4
---

The questionnaire shared with you and your team needs to be filled in
its entirety. Details requested in the questionnaire have been described
below. If you still have questions, please reach out to the team member
helping you with the on-boarding process.

### Team Name and Contact Info[Internal Link: ¶](\#team-name-and-contact-info)

This data is used to identify the use-case owner and associated team
throughout the on-boarding process.

- **Group**: Org of team requesting for ACI Kafka service. The group
  (Org) is a drop-down and if you do not see your Org simply type in
  your Org name in the field and it will be reviewed by the
  on-boarding member.
- **Team Name**: Name of the team requesting Kafka service(s). Could
  alternatively use your project name.
- **Kafka DRI**: Name of the person/DRI that is requesting Kafka
  service, usually it is the EPM for the project.
- **Team Email**: Apple Directory email group of the team requesting
  ACI Kafka service.

### Background[Internal Link: ¶](\#background)

- **How do you plan to use Kafka?**: Short description of your
  use-case.
- **Does your team currently run Kafka?**: Select “Yes” if your
  use-case is already running Kafka.
- **What language(s) are your services written in?**: Please mention
  the languages your clients are written in. e.g. Java8 FTW.
- **Is your team planning to use Kafka For All?**: Kafka For All is to
  provide support for non-JVM clients. Select Yes/No based on whether
  you plan on using Kafka For All or not.
- **Is your team currently using or evaluating other services in**
  **PIE?**: Select the services you are currently evaluating or are
  already using.

### Timeline / Go-Live Dates[Internal Link: ¶](\#timeline-go-live-dates)

- **DEV/QA**: Target date for QA service to be available.
- **Production**: Target date for Production service to be available.

Note

You should notify at least one month in advance for Production availability.

### Deployments[Internal Link: ¶](\#deployments)

- **In which DCs do you need Prod capacity?** Please select the
  datacenters you want ACI Kafka service to be available in.
- **In which DCs, work units, or labs will your QA clients be**
  **located?** Please select the datacenters you want ACI Kafka QA
  services to be available in.
- **Does your use-case require replicating data across Kafka**
  **clusters?**: Select “Yes/No”. ACI Kafka Mirror for topic
  replication across namespaces is targeted for Q3 FY 1019.

### Data Security[Internal Link: ¶](\#data-security)

- Will you be storing sensitive data in Kafka? If so, will this data
  contain PII or data requiring PCI compliance?ACI Kafka is only PII compliant and is InfoSec approved in terms of
  data security. However if your use-case requires storing PII or PCI
  compliant data of any form, please make sure you validate your
  use-case with InfoSec.

`Note`: ACI Kafka is **NOT** PCI compliant as of today.

### Data In[Internal Link: ¶](\#data-in)

This section captures information about your bandwidth IN i.e. Produce
Capability. Most of the information requested is self explanatory and an
example has been provided in
[Internal Link: Namespace](concepts.html#namespace)
section and in the questionnaire for your reference.

- For each DC, what is your expected requests/sec or messages/sec for
  production traffic on your go-live date? A year from go-live?
- What do you expect to be your average production rate (MB/s per DC)?
- Peak Production rate (MB/s per DC)?
- What is your average message size in bytes for production traffic on
  your go-live date? A year from go-live?
- How many service instances will be producing data?

### Data Out[Internal Link: ¶](\#data-out)

Like “Data In” this section captures information about your Bandwidth
OUT requirements i.e. Consume Capability. An example has been provided
in [Internal Link: Namespace](concepts.html#namespace)
section and in the questionnaire for your reference.

- How many consumer (groups) do you expect to read from Kafka?
- How many Kafka consumers will you create across your application
  instances?
- What do you expect to be your Average consumption rate i.e. (MB/s
  per DC)?
- Peak consumption rate (MB/s)?

### Data Retention[Internal Link: ¶](\#data-retention)

This section talks about disk storage requirements for holding your data
for a length of time before it would be deleted. Data retention policies
are described in more detail in the
[Internal Link: topic retention](concepts.html#topic-retention)
section.

- Enter how much space do you want to allocate for topic retention (in
  GB)?
- How many topics and partitions do you need? Provide a rough estimate
  (Leave blank if unknown).

### [Internal Link: Customer Groups](concepts.html\#customer-group)[Internal Link: ¶](\#customer-groups)

This section captures _customer group_ details of your use case and will
be used to create your groups with the access details your provide in
form of AD group(s). Please see [Internal Link: Customer
Group](concepts.html#customer-group)
section in our concepts page and ensure that you follow the naming
conventions. [Internal Link: Access
Info](concepts.html#customer-group-access-control)
and [Internal Link: Contact
Info](concepts.html#customer-group-contact-information)
details provided for the group will be applicable only for that group.
If you have requested multiple groups and need to have same access
level, please still explicitly provide same information across all
customer groups.

### [Internal Link: Namespace](concepts.html\#namespace)[Internal Link: ¶](\#namespace)

This section captures details of the Namespace(s) you are requesting
for, including Bandwidth IN & OUT. Please see
[Internal Link: Namespace](concepts.html#namespace)
section for more details about naming conventions and concepts to
provide details. Depending on the numbers you provide your namespace
could be on a dedicated cluster or on a shared cluster, which will be
communicated to you later.

**Note**: Queues are not meant to be uses as datastore, so one should
provide storage requirements accordingly.

```
**Important Information**

```

- The information you provide us with-in the Data In and Data Out
  sections will be used only for budgeting and it’s not enforced.
  Your identity is limited by the cluster default traffic quota.
- We recommend that you have your data long lived in a separate data
  store (i.e Cassandra) and not use Kafka as your source of truth.
- PII compliant use-case(s) requires Infosec approval.
- Production hosts 2012 or newer are necessary for E2E encryption.
- Alerts and Pager Contact Groups needs to be active and valid in
  order to receive alerts and pages, or you run into the risk of
  **NOT** getting notified when your service is interrupted.
- Downtime or data loss **IS** a possibility in the event of power
  loss to an Apple Data Center, or simply because of a bug in Kafka.
- The largest message size we support is 3 MiB before Kafka’s
  compression, messages larger than 3 MiB will be rejected.
- In IF1 (QA) you will be provisioned with Standard quota limits i.e.
  1 MB/s Produce and 2 MB/s Consume 100 GB of storage. If your
  use-case requires more quota, you will need to provide business
  justification.

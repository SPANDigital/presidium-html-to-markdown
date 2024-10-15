---
title: "Usage¶"
weight: 4
---

The example usages for different entities are shown below. Placeholders, surrounded by angle brackets, i.e., `<group>`, are to be replaced with the actual value by the user.

- [Internal Link: Activities](#activities)
- [Internal Link: Groups](#groups)
- [Internal Link: Namespaces](#namespaces)
- [Internal Link: Topics](#topics)
- [Internal Link: Client Identities](#client-identities)
- [Internal Link: Accesses](#accesses)

### Activities[Internal Link: ¶](\#activities)

Operations that create, update or delete an entity, schedule a background activity. These are identified by an activity ID, for example, `2ace75a0-3373-11e9-874c-4d3d40fdba04`. An operation is completed when the background activity finishes.

The CLI will automatically wait for the completion of any activity, but you can also verify the status yourself:

```
$ kafka-cli activity \
  --env prod \
  --fetch \
  --activity-id 2ace75a0-3373-11e9-874c-4d3d40fdba04

```

The above will fetch an activity in the environment ( `<env>`) `prod` with a UUID of `2ace75a0-3373-11e9-874c-4d3d40fdba04`.

### Groups[Internal Link: ¶](\#groups)

The [Internal Link: customer group name](../concepts.html#customer-group) is of the form `<customerName>.<groupName>.<environtmentName>` and given the shorthand “group ID” ( `<groupId>`) or “group” ( `<group>`). Example customer groups would be `icloud.Messaging.QA1` and `pie.queue.prod`.

Note

These entities are currently created by the ACI Kafka team. Customers can read entities they have visibility for, and edit limited fields where modifications permissions allow.

#### List all groups[Internal Link: ¶](\#list-all-groups)

```
$ kafka-cli group \
  --env prod \
  --list

```

#### Fetch single group[Internal Link: ¶](\#fetch-single-group)

```
$ kafka-cli group \
  --env prod \
  --fetch \
  --group-id <group>

```

#### Create new group[Internal Link: ¶](\#create-new-group)

```
$ kafka-cli group \
    --env prod \
    --create \
    --group-id <group> \
    --description "A textual description for the group" \
    --contact <email> \
    --alert-contacts <email> \
    --pager-contacts <email> \
    --modification-groups <directoryGroupName>@group.apple.com \
    --visibility-groups <directoryGroupName>@group.apple.com

```

#### Update group[Internal Link: ¶](\#update-group)

```
$ kafka-cli group \
    --env prod \
    --update \
    --group-id <group> \
    --description "A textual description for the group" \
    --contact <email> \
    --alert-contacts <email> \
    --pager-contacts <email> \
    --modification-groups <directoryGroupName>@group.apple.com \
    --visibility-groups <directoryGroupName>@group.apple.com

```

#### Delete group[Internal Link: ¶](\#delete-group)

```
$ kafka-cli group \
  --env prod \
  --delete \
  --group-id <group>

```

### Namespaces[Internal Link: ¶](\#namespaces)

A [Internal Link: namespace](../concepts.html#namespace) contains a set of topics and is usually denoted by a single word, for example, `<namespace>` (or, to minimize ambiguity, `<kafkaNamespace>`). Here are two examples of fully qualified namespaces ( `<namespaceRegex>`), which are of the form `<group>.<namespace>`, i.e., `<customerName>.<groupName>.<environmentName>.<namespace>`:

- `pie.queue.prod.pv` ( `<namespace>` is `pv`)
- `icloud.cloudkit.icloud4.p02` ( `<namespace>` is `p02`)

Note

These entities are currently created by the ACI Kafka team. Customers can read entities they have visibility for, and edit limited fields where modifications permissions allow.

#### List all namespaces[Internal Link: ¶](\#list-all-namespaces)

```
$ kafka-cli namespace \
  --env prod \
  --list \
  --group-id <group>

```

#### Fetch single namespace[Internal Link: ¶](\#fetch-single-namespace)

Note

As there are currently no SLOs for the public API and hence the CLI, it is not advised to rely on it to fetch the ACI Kafka bootstrap servers.

For Production purposes, use [Internal Link: Kaffe](../client/kaffe.html) or store the list of bootstrap servers in your application config.

```
$ kafka-cli namespace \
  --env prod \
  --fetch \
  --group-id <group> \
  --namespace-name <namespace>

```

#### Fetch namespace capacity[Internal Link: ¶](\#fetch-namespace-capacity)

You can fetch the currently allocated and remaining capacity for a namespace with `--show-capacity`.

Note

You need to install `pie-kafka-cli==0.3.35` or later to use this feature.

```
$ kafka-cli namespace \
  --env prod \
  --fetch \
  --group-id <group> \
  --namespace-name <namespace> \
  --show-capacity

```

#### Create new namespace[Internal Link: ¶](\#create-new-namespace)

```
$ kafka-cli namespace \
  --env prod \
  --create \
  --group-id <group> \
  --namespace-name <namespace> \
  --description "A textual description for the namespace" \
  --cluster-dc "<two letter datacenter code>" \
  --cluster-id "<cluster name, e.g., shared_prod_1>" \
  --produce-bytes-per-sec-capacity 10Mi \
  --consume-bytes-per-sec-capacity 10Mi \
  --storage-capacity-bytes 5Gi

```

#### Update namespace[Internal Link: ¶](\#update-namespace)

```
$ kafka-cli namespace \
  --env prod \
  --update \
  --group-id <group> \
  --namespace-name <namespace> \
  --description "A textual description for the namespace" \
  --cluster-dc "<two letter datacenter code>" \
  --cluster-id "<cluster name, e.g., shared_prod_1>" \
  --produce-bytes-per-sec-capacity 10Mi \
  --consume-bytes-per-sec-capacity 10Mi \
  --storage-capacity-bytes 5Gi

```

#### Delete namespace[Internal Link: ¶](\#delete-namespace)

```
$ kafka-cli namespace \
  --env prod \
  --delete \
  --group-id <group> \
  --namespace-name <namespace>

```

### Topics[Internal Link: ¶](\#topics)

Messages can be produced to and/or consumed from a [Internal Link: topic](../concepts.html#topic), e.g., `<topic>`. Each _fully qualified_ topic ( `<topicRegex>`) is of the form `<namespaceRegex>.<topic>`, i.e., `<customerName>.<groupName>.<environmentName>.<namespace>.<topic>`. Thus, if a fully qualified topic is `amp.bigdata.prod.pv.fraud-activity`, the topic is `fraud-activity`.

#### List all topics[Internal Link: ¶](\#list-all-topics)

```
$ kafka-cli topic \
  --env prod \
  --list \
  --group-id <group> \
  --namespace-name <namespace>

```

#### Fetch single topic[Internal Link: ¶](\#fetch-single-topic)

```
$ kafka-cli topic \
  --env prod \
  --fetch \
  --group-id <group> \
  --namespace-name <namespace> \
  --topic-name <topic>

```

#### Fetch topic capacity[Internal Link: ¶](\#fetch-topic-capacity)

You can fetch the currently allocated capacity for a topic.

Note

You need to install `pie-kafka-cli==0.3.35` or later to use this feature.

```
$ kafka-cli topic \
  --env prod \
  --fetch \
  --group-id <group> \
  --namespace-name <namespace> \
  --topic-name <topic> \
  --show-capacity

```

#### Create new topic[Internal Link: ¶](\#create-new-topic)

```
$ kafka-cli topic \
  --env prod \
  --create \
  --group-id <group> \
  --namespace-name <namespace> \
  --topic-name <topic> \
  --description "A textual description for the topic" \
  --produce-bytes-per-sec-capacity 2Mi \
  --consume-bytes-per-sec-capacity 2Mi \
  --retention-bytes 1000M \
  --num-partitions 5

```

#### Update topic[Internal Link: ¶](\#update-topic)

```
$ kafka-cli topic \
  --env prod \
  --update \
  --group-id <group> \
  --namespace-name <namespace> \
  --topic-name <topic> \
  --description "A textual description for the topic" \
  --produce-bytes-per-sec-capacity 2Mi \
  --consume-bytes-per-sec-capacity 2Mi \
  --retention-bytes 1000M \
  --num-partitions 5

```

#### Delete topic[Internal Link: ¶](\#delete-topic)

```
$ kafka-cli topic \
  --env prod \
  --delete \
  --group-id <group> \
  --namespace-name <namespace> \
  --topic-name <topic>

```

### Client Identities[Internal Link: ¶](\#client-identities)

[Internal Link: Client identities](../concepts.html#client-identities) (e.g., `<identity>`) are required to grant clients access to a topic. An example client with an identity name `publisher` in the `icloud.Messaging.QA1` customer group, would have a fully qualified client identifier ( `<identityRegex>`) of `icloud.Messaging.QA1.publisher`, as fully qualified client identifiers are of the form `<group>.<identity>`, i.e., `<customerName>.<groupName>.<environmentName>.<identity>`.

#### List all identities[Internal Link: ¶](\#list-all-identities)

```
$ kafka-cli identity \
  --env prod \
  --list \
  --group-id <group>

```

#### Fetch single client identity[Internal Link: ¶](\#fetch-single-client-identity)

```
$ kafka-cli identity \
  --env prod \
  --fetch \
  --group-id <group> \
  --identity-name <identity>

```

#### Create new client identity[Internal Link: ¶](\#create-new-client-identity)

Note

Learn more about creating a private/public key pair in our [Internal Link: client secret distribution documentation](../client/client_secrets_distribution.html).

```
$ kafka-cli identity \
  --env prod \
  --create \
  --group-id <group> \
  --identity-name <identity> \
  --public-key "$(cat <path to public key PEM file>)"

```

#### Delete client identity[Internal Link: ¶](\#delete-client-identity)

```
$ kafka-cli identity \
  --env prod \
  --delete \
  --group-id <group> \
  --identity-name <identity>

```

### Accesses[Internal Link: ¶](\#accesses)

A client identity can be assigned [Internal Link: access to a topic](../concepts.html#client-identity-topic-access).

Note

While you might want to provide both produce and consume access to a topic at the same time, it is possible to provide only produce or consume access. To do this, leave out one of the arguments and add the corresponding flag, `--no-consume` or `--no-produce`, when creating or updating an access. You need to install `pie-kafka-cli==0.3.38` or later to use this feature.

#### List all accesses[Internal Link: ¶](\#list-all-accesses)

```
$ kafka-cli access \
  --env prod \
  --list \
  --group-id <group> \
  --namespace-name <namespace> \
  --topic-name <topic>

```

#### Fetch single access[Internal Link: ¶](\#fetch-single-access)

```
$ kafka-cli access \
  --env prod \
  --fetch \
  --group-id <group> \
  --namespace-name <namespace> \
  --topic-name <topic> \
  --identity-group-id <group> \
  --identity-name <identity>

```

#### Create new access[Internal Link: ¶](\#create-new-access)

```
$ kafka-cli access \
  --env prod \
  --create \
  --group-id <group> \
  --namespace-name <namespace> \
  --topic-name <topic> \
  --identity-group-id <group> \
  --identity-name <identity> \
  --produce-bytes-per-sec 1Mi \
  --consume-bytes-per-sec 1Mi

```

#### Update access[Internal Link: ¶](\#update-access)

```
$ kafka-cli access \
  --env prod \
  --update \
  --group-id <group> \
  --namespace-name <namespace> \
  --topic-name <topic> \
  --identity-group-id <group> \
  --identity-name <identity> \
  --produce-bytes-per-sec 1Mi \
  --consume-bytes-per-sec 1Mi

```

#### Delete access[Internal Link: ¶](\#delete-access)

```
$ kafka-cli access \
  --env prod \
  --delete \
  --group-id <group> \
  --namespace-name <namespace> \
  --topic-name <topic> \
  --identity-group-id <group> \
  --identity-name <identity>

```

### Kube[Internal Link: ¶](\#kube)

Note

You need to install `pie-kafka-cli==0.3.45` or later to use this feature.

To authorize a Kubernetes namespace for [Internal Link: Kubernetes CRDs](kube.html) self-service
[Internal Link: customer group](../concepts.html#customer-group) must first be
decorated with the Kubernetes namespace name and the Kubernetes cluster
name:

```
$ kafka-cli group \
    --env prod \
    --update \
    --group-id <group> \
    --description "A textual description for the group" \
    --contact <email> \
    --alert-contacts <email> \
    --pager-contacts <email> \
    --modification-groups <directoryGroupName>@group.apple.com \
    --visibility-groups <directoryGroupName>@group.apple.com \
    --kube-cluster "<kube cluster name>" \
    --kube-namespace "<kube namespace name>"

```

#### Disabling management[Internal Link: ¶](\#disabling-management)

```
$ kafka-cli group \
    --env prod \
    --update \
    --group-id <group> \
    --description "A textual description for the group" \
    --contact <email> \
    --alert-contacts <email> \
    --pager-contacts <email> \
    --modification-groups <directoryGroupName>@group.apple.com \
    --visibility-groups <directoryGroupName>@group.apple.com \
    --no-kube

```

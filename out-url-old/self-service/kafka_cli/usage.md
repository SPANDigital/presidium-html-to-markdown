---
title: "Usage¶"
weight: 4
---

The example usages for different entities are shown below. Placeholders, surrounded by angle brackets, i.e., `<group>`, are to be replaced with the actual value by the user.

- [Activities](#activities)
- [Groups](#groups)
- [Namespaces](#namespaces)
- [Topics](#topics)
- [Client Identities](#client-identities)
- [Accesses](#accesses)

### Activities [¶](#activities "Link to this heading")

Operations that create, update or delete an entity, schedule a background activity. These are identified by an activity ID, for example, `2ace75a0-3373-11e9-874c-4d3d40fdba04`. An operation is completed when the background activity finishes.

The CLI will automatically wait for the completion of any activity, but you can also verify the status yourself:

```
$ kafka-cli activity \
  --env prod \
  --fetch \
  --activity-id 2ace75a0-3373-11e9-874c-4d3d40fdba04

```

The above will fetch an activity in the environment ( `<env>`) `prod` with a UUID of `2ace75a0-3373-11e9-874c-4d3d40fdba04`.

### Groups [¶](#groups "Link to this heading")

The [customer group name]({{%baseurl%}}/../concepts.html#customer-group) is of the form `<customerName>.<groupName>.<environtmentName>` and given the shorthand “group ID” ( `<groupId>`) or “group” ( `<group>`). Example customer groups would be `icloud.Messaging.QA1` and `pie.queue.prod`.

Note

These entities are currently created by the ACI Kafka team. Customers can read entities they have visibility for, and edit limited fields where modifications permissions allow.

#### List all groups [¶](#list-all-groups "Link to this heading")

```
$ kafka-cli group \
  --env prod \
  --list

```

#### Fetch single group [¶](#fetch-single-group "Link to this heading")

```
$ kafka-cli group \
  --env prod \
  --fetch \
  --group-id <group>

```

#### Create new group [¶](#create-new-group "Link to this heading")

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

#### Update group [¶](#update-group "Link to this heading")

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

#### Delete group [¶](#delete-group "Link to this heading")

```
$ kafka-cli group \
  --env prod \
  --delete \
  --group-id <group>

```

### Namespaces [¶](#namespaces "Link to this heading")

A [namespace]({{%baseurl%}}/../concepts.html#namespace) contains a set of topics and is usually denoted by a single word, for example, `<namespace>` (or, to minimize ambiguity, `<kafkaNamespace>`). Here are two examples of fully qualified namespaces ( `<namespaceRegex>`), which are of the form `<group>.<namespace>`, i.e., `<customerName>.<groupName>.<environmentName>.<namespace>`:

- `pie.queue.prod.pv` ( `<namespace>` is `pv`)
- `icloud.cloudkit.icloud4.p02` ( `<namespace>` is `p02`)

Note

These entities are currently created by the ACI Kafka team. Customers can read entities they have visibility for, and edit limited fields where modifications permissions allow.

#### List all namespaces [¶](#list-all-namespaces "Link to this heading")

```
$ kafka-cli namespace \
  --env prod \
  --list \
  --group-id <group>

```

#### Fetch single namespace [¶](#fetch-single-namespace "Link to this heading")

Note

As there are currently no SLOs for the public API and hence the CLI, it is not advised to rely on it to fetch the ACI Kafka bootstrap servers.

For Production purposes, use [Kaffe]({{< ref "client/kaffe" >}}) or store the list of bootstrap servers in your application config.

```
$ kafka-cli namespace \
  --env prod \
  --fetch \
  --group-id <group> \
  --namespace-name <namespace>

```

#### Fetch namespace capacity [¶](#fetch-namespace-capacity "Link to this heading")

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

#### Create new namespace [¶](#create-new-namespace "Link to this heading")

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

#### Update namespace [¶](#update-namespace "Link to this heading")

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

#### Delete namespace [¶](#delete-namespace "Link to this heading")

```
$ kafka-cli namespace \
  --env prod \
  --delete \
  --group-id <group> \
  --namespace-name <namespace>

```

### Topics [¶](#topics "Link to this heading")

Messages can be produced to and/or consumed from a [topic]({{%baseurl%}}/../concepts.html#topic), e.g., `<topic>`. Each _fully qualified_ topic ( `<topicRegex>`) is of the form `<namespaceRegex>.<topic>`, i.e., `<customerName>.<groupName>.<environmentName>.<namespace>.<topic>`. Thus, if a fully qualified topic is `amp.bigdata.prod.pv.fraud-activity`, the topic is `fraud-activity`.

#### List all topics [¶](#list-all-topics "Link to this heading")

```
$ kafka-cli topic \
  --env prod \
  --list \
  --group-id <group> \
  --namespace-name <namespace>

```

#### Fetch single topic [¶](#fetch-single-topic "Link to this heading")

```
$ kafka-cli topic \
  --env prod \
  --fetch \
  --group-id <group> \
  --namespace-name <namespace> \
  --topic-name <topic>

```

#### Fetch topic capacity [¶](#fetch-topic-capacity "Link to this heading")

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

#### Create new topic [¶](#create-new-topic "Link to this heading")

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

#### Update topic [¶](#update-topic "Link to this heading")

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

#### Delete topic [¶](#delete-topic "Link to this heading")

```
$ kafka-cli topic \
  --env prod \
  --delete \
  --group-id <group> \
  --namespace-name <namespace> \
  --topic-name <topic>

```

### Client Identities [¶](#client-identities "Link to this heading")

[Client identities]({{%baseurl%}}/../concepts.html#client-identities) (e.g., `<identity>`) are required to grant clients access to a topic. An example client with an identity name `publisher` in the `icloud.Messaging.QA1` customer group, would have a fully qualified client identifier ( `<identityRegex>`) of `icloud.Messaging.QA1.publisher`, as fully qualified client identifiers are of the form `<group>.<identity>`, i.e., `<customerName>.<groupName>.<environmentName>.<identity>`.

#### List all identities [¶](#list-all-identities "Link to this heading")

```
$ kafka-cli identity \
  --env prod \
  --list \
  --group-id <group>

```

#### Fetch single client identity [¶](#fetch-single-client-identity "Link to this heading")

```
$ kafka-cli identity \
  --env prod \
  --fetch \
  --group-id <group> \
  --identity-name <identity>

```

#### Create new client identity [¶](#create-new-client-identity "Link to this heading")

Note

Learn more about creating a private/public key pair in our [client secret distribution documentation]({{< ref "client/client_secrets_distribution" >}}).

```
$ kafka-cli identity \
  --env prod \
  --create \
  --group-id <group> \
  --identity-name <identity> \
  --public-key "$(cat <path to public key PEM file>)"

```

#### Delete client identity [¶](#delete-client-identity "Link to this heading")

```
$ kafka-cli identity \
  --env prod \
  --delete \
  --group-id <group> \
  --identity-name <identity>

```

### Accesses [¶](#accesses "Link to this heading")

A client identity can be assigned [access to a topic]({{%baseurl%}}/../concepts.html#client-identity-topic-access).

Note

While you might want to provide both produce and consume access to a topic at the same time, it is possible to provide only produce or consume access. To do this, leave out one of the arguments and add the corresponding flag, `--no-consume` or `--no-produce`, when creating or updating an access. You need to install `pie-kafka-cli==0.3.38` or later to use this feature.

#### List all accesses [¶](#list-all-accesses "Link to this heading")

```
$ kafka-cli access \
  --env prod \
  --list \
  --group-id <group> \
  --namespace-name <namespace> \
  --topic-name <topic>

```

#### Fetch single access [¶](#fetch-single-access "Link to this heading")

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

#### Create new access [¶](#create-new-access "Link to this heading")

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

#### Update access [¶](#update-access "Link to this heading")

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

#### Delete access [¶](#delete-access "Link to this heading")

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

### Kube [¶](#kube "Link to this heading")

Note

You need to install `pie-kafka-cli==0.3.45` or later to use this feature.

To authorize a Kubernetes namespace for [Kubernetes CRDs]({{< ref "self-service/kube" >}}) self-service
[customer group]({{%baseurl%}}/../concepts.html#customer-group) must first be
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

#### Disabling management [¶](#disabling-management "Link to this heading")

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

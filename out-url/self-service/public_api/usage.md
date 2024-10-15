---
title: "Usage¶"
weight: 4
---

The following sections provide instructions to configure activities, customer groups, namespaces, topics, client identities, and accesses, including authorizing and unauthorizing a Kubernetes namespace for Kubernetes CRDs using an API.

Note

Many endpoints accept an `?expanded` query parameter that adds more details in the API response.
Endpoints that accept this parameter are explicitly called out.

### Activities[Internal Link: ¶](\#activities)

_Activities_ are background operations that create, update, or delete an entity, and are identified by an activity ID—for example, `2ace75a0-3373-11e9-874c-4d3d40fdba04`.

When the background activity finishes, the operation is complete.

You can check the status of an activity by using the `/activities` endpoint:

```
$ kcurl $BASE_URL/activities/2ace75a0-3373-11e9-874c-4d3d40fdba04

```

### Customer groups[Internal Link: ¶](\#customer-groups)

The _[Internal Link: customer group](../concepts.html#customer-group)_ provides a grouping mechanism for both client identities and namespaces.
The customer group name uses the format: `<customerName>.<groupName>.<environtmentName>`

Note

The ACI Kafka team creates these entities. You can read them if you have visibility and can edit limited fields where modification permissions allow.

Run the commands below to list, fetch, create, update, and delete a new customer group.

#### List all groups[Internal Link: ¶](\#list-all-groups)

The _list all groups_ endpoint accepts the `?expanded` query parameter.

Run the following command to see a list of all groups:

```
$ kcurl $BASE_URL/groups

```

#### Fetch single group[Internal Link: ¶](\#fetch-single-group)

Run the following command to fetch a single group:

```
$ kcurl $BASE_URL/groups/<group>

```

#### Create new group[Internal Link: ¶](\#create-new-group)

Create a new group by running the following command:

```
$ kcurl $BASE_URL/groups \
    -X POST \
    -H 'Content-Type: text/json' \
    -d @- << EOF
  {
    "name": "<group>",
    "entity": {
      "description": "A textual description for the group",
      "contact": [
        "<email>"
      ],
      "alert-contact": [
        "<email>"
      ],
      "pager-contact": [
        "<email>"
      ],
      "modification-groups": [
        "<directoryGroupName>@group.apple.com"
      ],
      "visibility-groups": [
        "<directoryGroupName>@group.apple.com"
      ]
    }
  }
EOF

```

#### Update group[Internal Link: ¶](\#update-group)

Update a group by running the following command:

```
$ kcurl $BASE_URL/groups/<group> \
    -X PUT \
    -H 'Content-Type: text/json' \
    -d @- << EOF
  {
    "entity": {
      "description": "A textual description for the group",
      "contact": [
        "<email>"
      ],
      "alert-contact": [
        "<email>"
      ],
      "pager-contact": [
        "<email>"
      ],
      "modification-groups": [
        "<directoryGroupName>@group.apple.com"
      ],
      "visibility-groups": [
        "<directoryGroupName>@group.apple.com"
      ]
    }
  }
EOF

```

#### Delete group[Internal Link: ¶](\#delete-group)

Delete a group by running the following command:

```
$ kcurl $BASE_URL/groups/<group> -X DELETE

```

### Namespaces[Internal Link: ¶](\#namespaces)

A _[Internal Link: namespace](../concepts.html#namespace)_ contains a logical and physical grouping of related topics.

Note

The ACI Kafka team creates these entities. You can read them if you have visibility and can edit limited fields where modification permissions allow.

Run the following commands to list, fetch, create, update, and delete namespaces.

#### List all namespaces[Internal Link: ¶](\#list-all-namespaces)

The _list all namespaces_ endpoint accepts the `?expanded` query parameter.

Run the following command to see a list of all group namespaces:

```
$ kcurl $BASE_URL/groups/<group>/namespaces

```

#### Fetch single namespace[Internal Link: ¶](\#fetch-single-namespace)

Note

Service level objectives (SLOs) don’t exist for the public API. Avoid using the public API to fetch the ACI Kafka bootstrap servers.
Use [Internal Link: Kaffe](../client/kaffe.html) or store the list of bootstrap servers in your application configuration for production.

Fetch a single namespace in a group by running the following command:

```
$ kcurl $BASE_URL/groups/<group>/namespaces/<namespace>

```

#### Fetch namespace capacity[Internal Link: ¶](\#fetch-namespace-capacity)

Fetch allocated and remaining capacity for a namespace by running the command below:

```
$ kcurl $BASE_URL/groups/<group>/namespaces/<namespace>/capacity

```

#### Create new namespace[Internal Link: ¶](\#create-new-namespace)

To create a new namespace for a group, run the following command:

```
$ kcurl $BASE_URL/groups/<group>/namespaces \
    -X POST \
    -H 'Content-Type: text/json' \
    -d @- << EOF
  {
    "name": "<namespace>",
    "entity": {
      "description": "A textual description for the namespace",
      "dc": "<two letter datacenter code>",
      "cluster-id": "<cluster name, e.g., shared_prod_1.1>",
      "total-capacity": {
        "produce-bytes-per-sec": "10Mi",
        "consume-bytes-per-sec": "10Mi",
        "storage-bytes": "5Gi"
      }
    }
  }
EOF

```

#### Update namespace[Internal Link: ¶](\#update-namespace)

To update a namespace in a group, run the following command:

```
$ kcurl $BASE_URL/groups/<group>/namespaces/<namespace> \
    -X PUT \
    -H 'Content-Type: text/json' \
    -d @- << EOF
  {
    "entity": {
      "description": "A textual description for the namespace",
      "dc": "<two letter datacenter code>",
      "cluster-id": "<cluster name, e.g., shared_prod_1.1>",
      "total-capacity": {
        "produce-bytes-per-sec": "10Mi",
        "consume-bytes-per-sec": "10Mi",
        "storage-bytes": "5Gi"
      }
    }
  }
EOF

```

#### Delete namespace[Internal Link: ¶](\#delete-namespace)

To delete a namespace in a group, run the following command:

```
$ kcurl $BASE_URL/groups/<group>/namespaces/<namespace> -X DELETE

```

Note

_Customer groups_ and _namespaces_ are created by the ACI Kafka team.
You can read entities and edit limited fields if you have visibility permissions.

### Topics[Internal Link: ¶](\#topics)

A [Internal Link: _topic_](../concepts.html#topic) is used to manage Kafka topics.
Messages can be _produced to_ and _consumed from_ a topic.

Run the commands below to list, fetch, create, update, and delete topics.

#### List all topics[Internal Link: ¶](\#list-all-topics)

The _list all topics_ endpoint accepts the `?expanded` query parameter.

Run the following command to see a list of all group topics:

```
$ kcurl $BASE_URL/groups/<group>/namespaces/<namespace>/topics

```

#### Fetch single topic[Internal Link: ¶](\#fetch-single-topic)

Run the command below to fetch a single topic:

```
$ kcurl $BASE_URL/groups/<group>/namespaces/<namespace>/topics/<topic>

```

#### Fetch topic capacity[Internal Link: ¶](\#fetch-topic-capacity)

Fetch the currently allocated capacity for a topic by running the following command:

```
$ kcurl $BASE_URL/groups/<group>/namespaces/<namespace>/topics/<topic>/capacity

```

#### Create new topic[Internal Link: ¶](\#create-new-topic)

The key-value pairs `retention-milliseconds`, `compacted` and `connect-enabled` are optional.

Note

When setting _compacted_, you must also set _connect-enabled_ to ensure the correct _cleanup.policy_ is assigned.

With topic compaction enabled:

> Setting `"connect-enabled" : false` assigns `"cleanup.policy" : "compact,delete"`.
>
> Setting `"connect-enabled" : true` assigns `"cleanup.policy" : "compact"`.

Run the following command to create a new topic:

```
$ kcurl $BASE_URL/groups/<group>/namespaces/<namespace>/topics \
    -X POST \
    -H 'Content-Type: text/json' \
    -d @- << EOF
  {
    "name": "<topic>",
    "entity": {
      "description": "A textual description for the topic",
      "total-bandwidth": {
        "produce-bytes-per-sec": "2Mi",
        "consume-bytes-per-sec": "2Mi"
      },
      "retention-bytes": "1000M",
      "partition-count": 0,
      "retention-milliseconds": "3600000",
      "compacted": false,
      "connect-enabled": false
    }
  }
EOF

```

#### Update topic[Internal Link: ¶](\#update-topic)

The key-value pairs `retention-milliseconds`, `compacted` and `connect-enabled` are optional.

Note

When setting _compacted_, it’s advised to set _connect-enabled_ to ensure the correct _cleanup.policy_ is assigned.

With topic compaction enabled:

> Setting `"connect-enabled" : false` assigns `"cleanup.policy" : "compact,delete"`.
>
> Setting `"connect-enabled" : true` assigns `"cleanup.policy" : "compact"`.

Omitting _connect-enabled_ will preserve the current _cleanup.policy_ for the topic.

Run the following command to update a topic:

```
$ kcurl $BASE_URL/groups/<group>/namespaces/<namespace>/topics/<topic> \
    -X PUT \
    -H 'Content-Type: text/json' \
    -d @- << EOF
  {
    "entity": {
      "description": "A textual description for the topic",
      "total-bandwidth": {
        "produce-bytes-per-sec": "2Mi",
        "consume-bytes-per-sec": "2Mi"
      },
      "retention-bytes": "1000M",
      "partition-count": 0,
      "retention-milliseconds": "3600000",
      "compacted": false,
      "connect-enabled": false
    }
  }
EOF

```

#### Delete topic[Internal Link: ¶](\#delete-topic)

Delete a topic by running the following command:

```
$ kcurl $BASE_URL/groups/<group>/namespaces/<namespace>/topics/<topic> -X DELETE

```

### Client identities[Internal Link: ¶](\#client-identities)

Each [Internal Link: _client identity_](../concepts.html#client-identities), also referred to as a _client_ or _identity_ in this documentation, belongs to a customer group. A client identity typically maps to a specific application or service.
Interactions with ACI Kafka use a client identity, which is required to grant access to a topic.

You can run the commands below to manage your client identities.

#### List all client identities[Internal Link: ¶](\#list-all-client-identities)

The _list all identities_ endpoint accepts the `?expanded` query parameter.

Run the following command to see a list of all group client identities:

```
$ kcurl $BASE_URL/groups/<group>/identities

```

#### Fetch single client identity[Internal Link: ¶](\#fetch-single-client-identity)

Run the following command to fetch a single client identity:

```
$ kcurl $BASE_URL/groups/<group>/identities/<identity>

```

#### Create new client identity[Internal Link: ¶](\#create-new-client-identity)

Create a new client identity by running the following command:

```
$ kcurl $BASE_URL/groups/<group>/identities \
    -X POST \
    -H 'Content-Type: text/json' \
    -d @- << EOF
  {
    "name": "<identity>",
    "entity": {
      "public-key": "$(cat <path to public key PEM file>)"
    }
  }
EOF

```

To learn more about how to create a private and public key pair, see [Internal Link: Client Secrets Distribution](../client/client_secrets_distribution.html).

#### Delete client identity[Internal Link: ¶](\#delete-client-identity)

Delete a client identity by running the following command:

```
$ kcurl $BASE_URL/groups/<group>/identities/<identity> -X DELETE

```

### Accesses[Internal Link: ¶](\#accesses)

You can assign a [Internal Link: _client identity access_ to a _topic_](../concepts.html#client-identity-topic-access).
You can also list, fetch, create, update, and delete access by running the following commands.

#### List all accesses[Internal Link: ¶](\#list-all-accesses)

The _list all accesses_ for a topic endpoint accepts the `?expanded` query parameter.

Run the following command to see access for all group topics:

```
$ kcurl $BASE_URL/groups/<group>/namespaces/<namespace>/topics/<topic>/accesses

```

#### Fetch single access[Internal Link: ¶](\#fetch-single-access)

Run the command below to fetch a single access resource for a topic:

```
$ kcurl $BASE_URL/groups/<group>/namespaces/<namespace>/topics/<topic>/accesses/groups/<group>/identities/<identity>

```

#### Create new access[Internal Link: ¶](\#create-new-access)

To create new access for a topic, run the following command:

```
$ kcurl $BASE_URL/groups/<group>/namespaces/<namespace>/topics/<topic>/accesses \
    -X POST \
    -H 'Content-Type: text/json' \
    -d @- << EOF
  {
    "identity-id": {
      "group": "<group>",
      "identity": "<identity>"
    },
    "entity": {
      "produce": {
        "bytes-per-sec": "1Mi"
      },
      "consume": {
        "bytes-per-sec": "1Mi"
      }
    }
  }
EOF

```

#### Update access[Internal Link: ¶](\#update-access)

To update access for a topic, run the following command:

```
$ kcurl $BASE_URL/groups/<group>/namespaces/<namespace>/topics/<topic>/accesses/groups/<group>/identities/<identity> \
    -X PUT \
    -H 'Content-Type: text/json' \
    -d @- << EOF
  {
    "entity": {
      "produce": {
        "bytes-per-sec": "1Mi"
      },
      "consume": {
        "bytes-per-sec": "1Mi"
      }
    }
  }
EOF

```

#### Delete access[Internal Link: ¶](\#delete-access)

To delete access for a topic, run the command below:

```
$ kcurl $BASE_URL/groups/<group>/namespaces/<namespace>/topics/<topic>/accesses/groups/<group>/identities/<identity> -X DELETE

```

### Kube[Internal Link: ¶](\#kube)

To authorize a Kubernetes namespace for [Internal Link: Kubernetes CRDs](kube.html) self-service, you must set `kubernetes-info` with the names of your Kubernetes `cluster` and `namespace`.

```
$ kcurl $BASE_URL/groups/<group> \
    -X PUT \
    -H 'Content-Type: text/json' \
    -d @- << EOF
  {
    "entity": {
      "description": "A textual description for the group",
      "contact": [
        "<email>"
      ],
      "alert-contact": [
        "<email>"
      ],
      "pager-contact": [
        "<email>"
      ],
      "modification-groups": [
        "<directoryGroupName>@group.apple.com"
      ],
      "visibility-groups": [
        "<directoryGroupName>@group.apple.com"
      ],
      "kubernetes-info": {
        "cluster": "<kube cluster name>",
        "namespace": "<kube namespace name>"
      }
    }
  }
EOF

```

#### Disabling management[Internal Link: ¶](\#disabling-management)

To unauthorize a Kubernetes namespace, run the following command:

```
$ kcurl $BASE_URL/groups/<group> \
    -X PUT \
    -H 'Content-Type: text/json' \
    -d @- << EOF
  {
    "entity": {
      "description": "A textual description for the group",
      "contact": [
        "<email>"
      ],
      "alert-contact": [
        "<email>"
      ],
      "pager-contact": [
        "<email>"
      ],
      "modification-groups": [
        "<directoryGroupName>@group.apple.com"
      ],
      "visibility-groups": [
        "<directoryGroupName>@group.apple.com"
      ],
      "kubernetes-info": null
    }
  }
EOF

```

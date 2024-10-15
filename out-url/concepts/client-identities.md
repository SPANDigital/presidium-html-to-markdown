---
title: "Client identities¶"
weight: 3
---

Each interaction with an ACI Kafka cluster uses a _client identity_.
Each client identity is owned by a customer group and usually maps to a specific application or service.
Client identities can be configured with permissions to _consume from_ or _produce to_ multiple topics as explained in [Internal Link: Client identity topic access](#client-identity-topic-access).

Client identity names have a maximum length of 63 characters and can contain letters, numbers, and `_` or `-` characters.

Client identity name is prefixed by the customer group name.
For example a client identity name `publisher` in the `icloud.Messaging.QA1` customer group uses a fully qualified identifier: `icloud.Messaging.QA1.publisher`

Client identities can’t be changed after creation and contain two fields: name and public key.

To identify your client to [External Link: Kaffe](https://docs.aci.apple.com/kafka/client/kaffe.html#kaffe), `pie.queue.kaffe.client.id` property must be configured in ACI Kafka _producers_ and _consumers_ with the fully-qualified client identity name.

The client identity’s private key file can be configured with the property `pie.queue.kaffe.client.private.key.location`.

To learn more about generating public and private key pairs and the related configuration, see [Internal Link: Client Secrets
Distribution](client/client_secrets_distribution.html#client-secrets-distribution).

### Consumer group[Internal Link: ¶](\#consumer-group)

When subscribing to topics, [External Link: _consumer clients_](https://docs.aci.apple.com/kafka/client/consumer_client.html) configure a [External Link: _consumer group_](https://docs.aci.apple.com/kafka/client/consumer_client.html#consumer-group-a-k-a-group-id-config) using the prefix property: `group.id`

In ACI Kafka, consumer groups must use the Kaffe client id prefix (noted above) to enforce Kafka ACLs and protect against unauthorized access to a topic’s consumer offsets.

For more information, see [Internal Link: Consumer Group Convention](client/consumer_settings.html#consumer-group-and-aci-kafka-client).


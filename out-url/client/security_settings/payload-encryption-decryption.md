---
title: "Payload encryption and decryption settings¶"
weight: 3
---

The new end-to-end encryption design for Kafka allows users to own their encryption keys and adds support for arm64.
The design and implementation has been reviewed by ACS Security.

The design requires an individual group, called “context owner(s)” to “manage” an encryption context.
Managing an encryption context involves:

- Creating the encryption context
- Storing the key associated with the encryption context
- Granting decrypt permission to a consumer
- Revoking consumer’s decryption permissions

### Requirements[Internal Link: ¶](\#requirements)

Please ensure your consumers and producers are configured to use at least `com.apple.pie.queue:kafka-pie-client:0.28.3`.
This is a general availability (GA) release of the ASE Kafka client, based on upstream Apache Kafka 3.6.
It supports both the old and new encryption schemes.

In addition, the context owners must install `kafkacli` from [External Link: Homebrew](https://github.pie.apple.com/homebrew/brew) using:

```
brew update
brew install apple/ase/kafkacli

```

### Setup[Internal Link: ¶](\#setup)

#### Creating an encryption context[Internal Link: ¶](\#creating-an-encryption-context)

The context owner may execute the following command to create an encryption context:

```
kafkacli encryption --env if1 create <context_id>

```

`context_id` must be fully qualified and be of the form `<customerName>.<groupName>.<environtmentName>.{contextName}`.

The owner MUST have write permissions for the customer group `<customerName>.<groupName>.<environtmentName>`.

Please, note down the context master secret key that’s printed in the console and save it securely (i.e., [External Link: Whisper](https://whisper.corp.apple.com/), [External Link: sops](https://github.com/getsops/sops), macOS Keychain or 1Password).

Alternatively, the key can be stored and retrieved automatically using Whisper if the relevant flags are used:

```
--whisper-bucket       whisper bucket to use
--whisper-namespace    whisper namespace to use
--whisper-cert         whisper certificate to use. If empty, defaults to current AppleConnect user.
--whisper-key          whisper key to use. If empty, defaults to current AppleConnect user

```

#### Grant a consumer decryption access[Internal Link: ¶](\#grant-a-consumer-decryption-access)

The context owner may grant a consumer identity decryption access with:

```
kafkacli encryption --env if1 grant-decrypt <context_id> <identity_id>

```

where `context_id` and `identity_id` are fully qualified.

#### Revoke a consumer decryption access[Internal Link: ¶](\#revoke-a-consumer-decryption-access)

The context owner may revoke a consumer identity decryption access with:

```
kafkacli encryption --env if1 revoke-decrypt <context_id> <identity_id>

```

where `context_id` and `identity_id` are fully qualified.

**`kafkacli` will prompt for confirmation before updating the encryption context’s public state.**
**This is done to allow all consumers to be configured with the rolled private states or pick them up automatically from [Internal Link: Kaffe](kaffe.html).**
**You should confirm the revocation only when all consumers are up-to-date with the new epoch.**

#### Configuring producers and consumers[Internal Link: ¶](\#configuring-producers-and-consumers)

Encryption and decryption are disabled by default and can be enabled by
setting the following producer or consumer properties for **ASE Kafka**
**client version 0.28.3+**:

```
# For producers:
aci.queue.encryption.context=<context_id>

# For consumers:
aci.queue.decryption.context=<context_id>
pie.queue.crypto.decrypt.enable=true
# Note that if a topic contains both encrypted (old and new scheme) and unencrypted messages, you will need to set this property to `true` to handle both kind of messages.

```

We **strongly recommend** enabling [Internal Link: Kaffe configuration injection](default_config_injection.html) if enabling encryption or/and decryption.
In GCBD, or if Kaffe configuration injection isn’t enabled, you need these additional properties:

```
# For producers:
aci.queue.encryption.context.public.state=<context_public_state>

# For consumers:
aci.queue.decryption.encrypted.context.private.state=<encrypted_context_private_state>
pie.queue.crypto.consumer.id=<provided by ASE Kafka team>
pie.queue.crypto.consumer.secret.key=<provided by ASE Kafka team and required only to support old encryption scheme> (Keep this secret)

# For both producers and consumers:
pie.queue.crypto.iris.connect=<Iris service connect address - https://kafka-iris-if1.pie.apple.com for if1, https://kafka-iris-prod.pie.apple.com for if-prod>

```

where `context_public_state` and `encrypted_context_private_state` are retrieved from the CLI:

```
kafkacli encryption --env if1 public-state <context_id>
kafkacli encryption --env if1 private-state <context_id> <identity_id>

```

Should you have any questions, do not hesitate to [Internal Link: reach out to us](../contact.html).

#### FAQ[Internal Link: ¶](\#faq)

##### What if producer or consumer is required to use a different encryption context for multiple topics?[Internal Link: ¶](\#what-if-producer-or-consumer-is-required-to-use-a-different-encryption-context-for-multiple-topics)

In this case, you need to provide a comma-separated list of all context ids and topic ids mapping:

```
# For producers:
aci.queue.encryption.context=<context_id_1>,<context_id_2>,<context_id_3>
aci.queue.encryption.context.topic.map=<topic_id_1>:<context_id_1>,<topic_id_2>:<context_id_2>,<topic_id_3>:<context_id_3>,

# For consumers:
aci.queue.decryption.context=<context_id_1>,<context_id_2>,<context_id_3>

```

In GCBD or if Kaffe configuration injection isn’t enabled, you also need these additional properties:

```
# For producers:
aci.queue.encryption.context.public.state=<context_public_state_1>,<context_public_state_2>,<context_public_state_3>

# For consumers:
aci.queue.decryption.encrypted.context.private.state=<encrypted_context_private_state_1>,<encrypted_context_private_state_2>,<encrypted_context_private_state_3>

```

##### I don’t trust ASE Kafka team to persist the encrypted context master secret. Can I persist encrypted context master secret locally on disk?[Internal Link: ¶](\#i-don-t-trust-ase-kafka-team-to-persist-the-encrypted-context-master-secret-can-i-persist-encrypted-context-master-secret-locally-on-disk)

Yes, you can provide `--manager=fs` flag when you interact with `kafkacli` to achieve that.

In addition to responsibility of keeping the encrypted context master secret on disk safely, you also need to provide a `--consumer-public-key-path` for the identity on each `grant-decrypt` command.

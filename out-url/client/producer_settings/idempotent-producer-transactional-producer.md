---
title: "Idempotent Producer and Transactional Producer¶"
weight: 4
---

ACI Kafka supports all types of writing semantic in Kafka including EOS
(a.k.a Idempotent Producer) and transaction writes (a.k.a Transactional
Producer), by default when you grant `PRODUCE` access to a
topic you will be granted all necessary access to perform all type of
writes. And it’s up to you to decide which semantic is right for your
use-case.

Note

Idempotent producer has been the default producer since Apache Kafka 2.8 (ACI Kafka 0.21).

- **When to use Exactly-once over At-Least-Once**If your use-case requires avoiding processing a record multiple
  times, then the producer will need to write only once no matter how
  many times it will retry to write the same record (this is known as
  [Internal Link: Idempotent Producer](producer_client.html#exactly-once)).
  To enable idempotence the only thing you need to do is setting
  `enable.idempotence: true`. Note that by setting this to
  true, the producer client will choose the right values for
  `acks`, `retries` and
  `max.in.flight.requests.per.connection` in order to
  provide EOS.Suitable values are:
  - acks=All
  - retries= <1 to Integer.MAX\_VALUE>
    (default is MAX\_VALUE>
  - max.in.flight.requests.per.connection= <1 to 5> (default
    value depends on the client version, most clients use 5)

Note

The idempotent producer only guarantees EOS within a single session.
If the producer client retries sending the same record in a new session, the record will be written into the Kafka topic, even if the client successfully wrote it in a previous session.

You should consider the transactional producer if you are looking for EOS across sessions.

- **When to use Transactional Producer**You may need to consider [Internal Link: Transactions Producer](producer_client.html#transactional) in
  Kafka if you have a use-case that needs to write a batch of messages
  to multiple partitions on the same cluster where either all messages
  in the batch are visible to any consumer or none of them will be at
  all. If this is your use-case you will need to make sure that any
  consumer fetch messages from any of your partitions need to be
  “transaction-aware consumer”, this can be achieved by setting
  `isolation.level:read_committed` and
  `enable.auto.commit` must be disabled.Keep in mind that using a transactional producer means your
  application is handling more than just the number of your messages,
  as the transaction API needs to write a transaction marker on
  completion and update the transaction state for each partition in
  the batch. All of this will add overhead to your application.

Note

Producers are allowed to set `transactional.id` to any value prefixed by
Any ACI-Kafka client identity with a [Internal Link: client identity topic access](../concepts.html#client-identity-topic-access) set to `PRODUCE`
is authorized to create [Internal Link: transactions](producer_client.html#transactional) with `transactional.id`
prefixed with their [Internal Link: ACI-Kafka client identity](../concepts.html#client-identities).

### Common problems with EOS producers[Internal Link: ¶](\#common-problems-with-eos-producers)

#### Misconfigured or rouge idempotent producers will cause Kafka clusters to run out-of-memory (OOM)[Internal Link: ¶](\#misconfigured-or-rouge-idempotent-producers-will-cause-kafka-clusters-to-run-out-of-memory-oom)

Kafka achieves EOS by keeping an in-memory map of the last produced batch by each producer instance.
Each producer instance is identified by a unique ProducerId (PID).
This map is stored in-memory by Kafka brokers for 7 days.

As the EOS guarantee of idempotent producers are limited to a single session, the client will need to request a new PID on every initialization.
This means any unnecessary initialization of the idempotent producer will put the broker at risk of OOM as it will
increase the size of the map of PID to the last produced batch.

Multiple causes can lead to an incident related to this behavior:

- if a misconfigured client enters a crash loop. The client will keep re-initializing.
- the client is initializing new producer instance on retriable errors instead of letting the current instance retry naturally.
- the client logic could have a bug, where it initialized more idempotent producers than necessary.
- an anti-pattern is used, leading to the creation of too many idempotent producers.

In the case of an incident, the ACI Kafka on-call engineer will ask for the use-cases not depending on EOS to disable idempotence.

**Action**

The producer client must set the following configuration:

- `enable.idempotence=false`
- `max.in.flight.requests.per.connection=1`


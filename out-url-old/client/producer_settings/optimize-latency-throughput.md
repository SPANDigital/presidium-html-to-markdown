---
title: "Optimize for Latency vs Throughput¶"
weight: 3
---

While Kafka tries to balance end-to-end latency with good throughput
some use-cases may value one over other. So for example if your goal is
achieving high durability and high throughput then you will need to be
aware the latency will be your trade-off here. The following sections
are options to consider when you are optimizing your producer
performance.

- **Batching**When it comes to latency and throughput for the producer, there are two
  parameters are particularly important for that

  1. `batch.size` controls how many bytes of data to collect,
     before the producer will send messages to Kafka (the default value is
     16384), so increasing this gives you better throughput.
  2. `linger.ms` controls how many milliseconds the producer will
     wait for `batch.size` to get full, so the producer will send
     data either if batch size reached or if linger delay get triggered (the
     default value is 0).

So you can tune these parameters to make sure your producer is sending
all the time, this how you can get the best throughput possible. As a
rule of Thumb small records (less than 100KB) are a good idea.**What can go wrong with large batches**You can set your batch size as high as possible. However, there are many parameters that may effect that for example
  - Available memory allocated at your application, to tune this
    check “How to size your buffer.memory?” section for details.
  - Depending on the size and load on the cluster resources the time
    the cluster will spend on your requests may vary as any of these
    parameters could add additional delays.
  - If any of the records in the batch is larger than 3MiB the
    request will be rejected as ACI Kafka has a hard limit for
    message size. If you believe you need a higher record size limit
    please contact us and we can discuss.
- **Compression**Bandwidth and Disk storage are the two most expensive commodities. We
  highly recommend compression especially if you are optimizing for
  throughput, the common arguments against compression is CPU usage.
  However, CPU usage overhead is usually negligible. If you don’t want to
  bother running a load test with and without compression we recommend you
  just turn compression on.Here are a few reasons why we think you might be surprised by how cheap compression is:

  - With compression there’s a smaller payload to encrypt after
    compression. Encryption is done both at the TLS and SDR layers.
    Smaller payloads are quicker to encrypt.
  - [zstd](https://facebook.github.io/zstd/) (only supported from
    Kafka 2.1) is beating most other codecs on both ratio and speed.
  - If your payloads are smaller your batches will contain more
    information and your throughput will be higher.

By default, compression is disabled in Kafka until you set a value for
`compression.type`, so to enable it set:
`compression.type=zstd` (or another codec if you have
consumers that can’t support zstd yet).
- **Other Considerations**
  - **acks and high throughput****tl;dr ACI Kafka supports only `acks=ALL`**You commonly read that to get higher throughput or better
    latency in Kafka you should play with acks. This is not
    something we support, what acks give you is knowledge of where
    your data will get dropped:

    - With `acks=all`, the producer will consider the write
      successful when all of the in-sync replicas receive the
      record. This is achieved by the leader broker being smart as
      to when it responds to the request. It’ll send back a
      response once all the in-sync replicas receive the record
      themselves. If any failure prevents the records from being
      replicated across all replicas, the client has a chance to
      act on that.
    - With `acks=1`, the producer will consider the write
      successful when the leader receives the record. The leader
      broker will know to immediately respond the moment it
      receives the record and not wait any longer. A failure in
      the network between the broker and its peers, a failure in
      the broker, or a simple change of leadership can cause the
      data to be lost without the producer ever noticing.
    - With `acks=0`, the producer won’t even wait for a response
      from the broker. It immediately considers the write
      successful the moment the record is sent out. A failure in
      the network, a failure in the broker, or a simple change of
      leadership can cause the data to be lost without the
      producer ever noticing.

It’s always better to know where your data is getting lost
because you’ll need to quantify your loss. You should achieve
almost the same throughput with ACKs=ALL that you’d do with
`ACKs=0` so there’s no good reason to do this.We will therefore not support clients to be producing at
anything but `ACKs=ALL` if you believe you need ACKs!=ALL please
contact us and we can discuss.


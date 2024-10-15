---
title: "Kaffe architecture¶"
weight: 8
---

### Protocol [¶](#protocol "Link to this heading")

Kaffe is a gRPC service. It’s [protocol\
definition]({{%baseurl%}}/https://github.pie.apple.com/pie/kafka-tools/blob/develop/kaffe-client/src/main/proto/kaffe.proto)
is available in `kafka-tools`.

For simplicity, a common header called `ProtoInfo` is always sent. It
contains the version. The version has nothing to do with ACI Kafka
versions and only identifies the variant of the protocol in use.

Passing a clear protocol version with each RPC enables ACI Kafka to
easily determine which versions are still in use and make educated
decisions about deprecating older clients.

Kaffe works with a watch system. Each request contains a `ResourceId` and
`Metadata`. `ResourceId` defines the type (namespace, cluster, client), id,
and the last know version of the entity. `Metadata` is just a dictionary
of extra information. The way Kaffe works is that it will park a request
until it knows of a version higher than the one sent in the request.
This enables clients to never miss an update.

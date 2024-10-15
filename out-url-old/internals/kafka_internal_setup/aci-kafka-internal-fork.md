---
title: "ACI Kafka internal fork¶"
weight: 2
---

### Branches [¶](#branches "Link to this heading")

#### Apache Kafka upstream branches [¶](#apache-kafka-upstream-branches "Link to this heading")

- The upstream Apache Kafka project hosts it’s code on
  [https://github.com/apache/kafka](https://github.com/apache/kafka). Development happens in the
  `trunk` branch with changes being merged into MAJOR.MINOR branches
  such as `2.0` or `2.3`. These branches are then the basis on kafka
  releases such as `2.0.1` or `2.3.0`.
- Our internal Apache Kafka releases are based on such public
  releases, with our own changes on top of them, we then append an
  additional version number to the upstream release number. In other
  words the first internal Apache Kafka release based on version
  `2.1.1` will be named `2.1.1.0`. Subsequent releases will have the
  last segment of the version number incremented by one per release.

#### Internal branches and the mechanics of a release [¶](#internal-branches-and-the-mechanics-of-a-release "Link to this heading")

- Each upstream release, such as `2.1.1` has a develop branch in the
  [https://github.pie.apple.com/pie/apache-kafka](https://github.pie.apple.com/pie/apache-kafka) repository, in
  this case named `develop-2.1.1`.
- Both local development and backports of changes from upstream
  `trunk` are landed onto the develop branch by opening a github PR.
- Once a change has been merged into a branch such as `develop-2.1.1`
  a rio pipeline will test the changes and possibly merge the changes
  into the corresponding release- branch, in this case
  `release-2.1.1`.
- Commits merged onto the `release-` branch that define a new version,
  such as `2.1.1.2` in `gradle.properties` will cause a publish
  pipeline to publish new packages that are made available to clients
  and can be referenced by the `kafkaVersion` property in
  `gradle.properties` in `kafka-tools`.

### Scala version [¶](#scala-version "Link to this heading")

We build our packages with Scala version 2.13.2.


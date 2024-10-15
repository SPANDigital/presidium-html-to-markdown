---
title: "ACI Kafka tools¶"
weight: 6
---

### Branches[Internal Link: ¶](\#id1)

- develop and master branch pair for active development.
- develop branch pulls latest desired release artifacts generated from
  ACI Kafka repository, e.g. version 0.10.0.1.0.
- master branch is used for CD.
- develop-x and release-x branches for each major version
- Once a release has been created, the target version of PIE Apache
  Kafka cannot be changed. A new release must be created in order to
  change the version of PIE Apache Kafka.
- carnival packages from release-x branch are the ones which would
  eventually make it to the production hosts.

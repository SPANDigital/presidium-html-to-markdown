---
title: "Setup of jobs¶"
weight: 3
---

### Jenkins jobs [¶](#jenkins-jobs "Link to this heading")

- We use [rio](https://docs.aci.apple.com/rio/) to automate our test,
  build and deployment tasks.
- Here is an example
  [rio.yml]({{%baseurl%}}/https://github.pie.apple.com/pie/apache-kafka/blob/develop-2.1.1/rio.yml)
  file for branch `develop-2.1.1`.
- You can read more about build templates
  [here](https://docs.aci.apple.com/rio/guide-to-rio/build-and-test/builders/java-gradle.html).
- Here is the rio project for
  [pie-apache-kafka](https://rio.apple.com/projects/pie-apache-kafka).
- Typically, there are `prb`, `merge` and `snapshot`/ `release` jobs
  for each branch.

### Binaries [¶](#binaries "Link to this heading")

- We generate artifacts from ACI Kafka repository and carnival
  packages from ACI Kafka tools repository.
- Artifacts generated from master branches in ACI Kafka repository go
  to
  [pie-snaphot-local](https://artifacts.apple.com/artifactory/webapp/#/artifacts/browse/tree/General/pie-snapshot-local/com/apple/pie/queue)
  artifactory path.
- Artifacts generated from release-numbered branches in ACI Kafka
  repository go to
  [pie-release-local](https://artifacts.apple.com/artifactory/webapp/#/artifacts/browse/tree/General/pie-release-local/com/apple/pie/queue)
  artifactory path. [A specific version is recommended for production use]({{%baseurl%}}/../releases/index.html#recommended-release).


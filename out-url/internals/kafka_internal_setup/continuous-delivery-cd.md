---
title: "Continuous delivery (CD)¶"
weight: 4
---

We have continuous delivery enabled on _develop-0.x_ branches in the `kafka-carnival-conf` repository and the _develop_ branch in the `kafka-tools` repository. Any change that lands in either of these branches will immediate trigger a deployment to our development cluster \[pieinternal\_cicd1 `](https://if1.carnival.apple.com/Carnival/HallOfMirrors.jsp?selectedISAppName=pieinternal_cicd_1-piekafka&searchStyle=Instance) in ` if1.

### How?[Internal Link: ¶](\#how)

- Any change in ACI Kafka and ACI Kafka tools triggers
  [External Link: pie-kafka-tools-master-carnival-1build](https://rio.apple.com/projects/pie-kafka-tools/pipeline-specs/pie-kafka-tools-master-carnival-1build/pipelines)
  job.
- Refer [External Link: trigger
  CD](https://github.pie.apple.com/pie/apache-kafka/blob/develop-0.10/rio.yml#L41)
  and [External Link: CD in
  rio.yml](https://github.pie.apple.com/pie/kafka-tools/blob/develop/rio.yml#L73).
- This job
  [External Link: deploys](https://github.pie.apple.com/pie/kafka-tools/blob/develop/rio.yml#L85)
  carnival packages to Kafka cluster in IF1.


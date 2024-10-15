---
title: "Denali profiles for connecting to Kafka brokers¶"
weight: 6
---

If you wish to connect to ACI Kafka clusters from AWS@Apple or another [External Link: Denali](https://denali.sdi.apple.com)-secured platform, use the following security profiles depending on the environment the Kafka cluster is in:

### Kafka brokers in Apple Owned Data Centers (AODC)[Internal Link: ¶](\#kafka-brokers-in-apple-owned-data-centers-aodc)

Ensure that your Denali Security Group(s) subscribes to these Denali Security Profiles:

- [External Link: `IS&T to ISS-NETs PROD AIS Approved`](https://networkpolicy.sdi.g.apple.com/security-profile?securityProfileId=34b8ebb3-9b28-4b76-94ff-abe252d0adbb). If your application connects to QA (IF1) clusters, replace this profile with the following:
  - [External Link: `IS&T to ISS-QA AIS Approved`](https://networkpolicy.sdi.g.apple.com/security-profile?securityProfileId=29aa8383-aef6-40a4-b833-ebddfe97c25c)
- [External Link: `/ISS/ACI/Kafka/Kubev2`](https://networkpolicy.sdi.g.apple.com/security-profile?securityProfileId=b91a312a-c68c-4245-b26e-d286a4b06495)
- `ACI Kafka - <cluster-name>` – if this profile exists on Denali. You can map your [Internal Link: namespace to a cluster name](../faq/index.html#what-cluster-does-my-aci-kafka-namespace-live-on) using our self-service mechanisms.

You may want to refer to Denali documentation, or #help-denali.

### Kafka brokers in AWS@Apple[Internal Link: ¶](\#kafka-brokers-in-aws-apple)

These Kafka clusters are deployed on isolated VPCs, owned and managed by ACI Kafka. Network access to these clusters is also managed through a Denali Security Profile named [External Link: `ACI Kafka brokers in aci-kafka-prod-0617`](https://networkpolicy.sdi.g.apple.com/security-profile?securityProfileId=39540c4b-e69b-4a01-9622-87df41d3dd43).

### Applications running in Kube[Internal Link: ¶](\#applications-running-in-kube)

If your application is running in Kube, you need to configure an ANP in all Kube namespaces and clusters where you application runs to permit outbound traffic into the Kube namespace `pie-queue-prod-local.kube` in `ALL` Kube clusters:

```
apiVersion: pie.apple.com/v1alpha1
kind: AppleNetworkPolicy
metadata:
  name: <anp-name>
spec:
  namespaces_outgoing:
    - cluster: ALL
      namespace: pie-queue-prod-local.kube
  selector:
    sdr.appname: <selector-to-match-your-application>

```

### Supporting Services (e.g. Kaffe)[Internal Link: ¶](\#supporting-services-e-g-kaffe)

If you use our JVM client, you may need some additional Security Profiles to connect to Kaffe and Iris:

- [External Link: `To Kube VIPs HTTPS`](https://networkpolicy.sdi.g.apple.com/security-profile?securityProfileId=28ce21cc-27c0-4aa8-b496-d880b1914b04)
- [External Link: `PLB Service KUBE V2`](https://networkpolicy.sdi.g.apple.com/security-profile?securityProfileId=263bf533-7fb6-4380-bfab-b6ee2c8bba69)

### What about connecting to QA (IF1)?[Internal Link: ¶](\#what-about-connecting-to-qa-if1)

The QA network in Denali is classified differently and the required profile grants a much wider access.
Customers are advised to subscribe to the [External Link: `IS&T to ISS-QA AIS Approved`](https://networkpolicy.sdi.g.apple.com/security-profile?securityProfileId=29aa8383-aef6-40a4-b833-ebddfe97c25c) Service Profile, which should grant access to all of QA.
Approval may be required (likely in the form of verifying that you are not connecting to QA from production).

---
title: "API¶"
weight: 6
---

### Controller [¶](#controller "Link to this heading")

Controller is part of Kubernetes CRDs based API which allows you to manage ACI Kafka resources via Kube.
API Operator is implemented as a collection of controllers where each controller watches a specific resource type.
Each resource contains a definition of desired state of the world which need to be reconciled by API operator.

#### Reconciliation Latency [¶](#reconciliation-latency "Link to this heading")

We measure the reconciliation latency as the percentage of reconciliations that took less than 10 minutes to complete: `1 - (slow reconciliations / total reconciliations)`

#### Reconciliation Availability [¶](#reconciliation-availability "Link to this heading")

Due to the nature of Kube events it can take several reconciliation cycles before a resource becomes `Ready` and fully reconciled.
Controller infinitely retries on errors, so technically it is always available.
That’s why we measure the reconciliation availability as the percentage of events that took more than 10 minutes to complete: `1 - (slow reconciliations / total reconciliations)`
Please note that we differentiate between user errors caused by invalid input and internal or upstream API errors.
User errors are excluded from the calculation.

### Webhook [¶](#webhook "Link to this heading")

Webhook is part of Kubernetes CRDs based API which allows you to manage ACI Kafka resources via Kube.
API Operator performs a set of validations on each Create/Update/Delete request to Kube API in order to reject bad input from users sooner rather than later.

#### Latency [¶](#id3 "Link to this heading")

We measure the webhook latency as the percentage of requests that took less than 10 seconds to complete: `1 - (slow requests / total requests)`

#### Availability [¶](#id4 "Link to this heading")

We measure the webhook availability as the percentage of requests that didn’t result in an error: `1 - (failed requests / total requests)`
Note: requests that were rejected due to bad user input are not considered as failed and are excluded from the calculation.


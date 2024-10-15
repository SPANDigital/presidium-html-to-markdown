---
title: "API¶"
weight: 6
---

### Controller[Internal Link: ¶](\#controller)

Controller is part of Kubernetes CRDs based API which allows you to manage ACI Kafka resources via Kube.
API Operator is implemented as a collection of controllers where each controller watches a specific resource type.
Each resource contains a definition of desired state of the world which need to be reconciled by API operator.

#### Reconciliation Latency[Internal Link: ¶](\#reconciliation-latency)

We measure the reconciliation latency as the percentage of reconciliations that took less than 10 minutes to complete: `1 - (slow reconciliations / total reconciliations)`

#### Reconciliation Availability[Internal Link: ¶](\#reconciliation-availability)

Due to the nature of Kube events it can take several reconciliation cycles before a resource becomes `Ready` and fully reconciled.
Controller infinitely retries on errors, so technically it is always available.
That’s why we measure the reconciliation availability as the percentage of events that took more than 10 minutes to complete: `1 - (slow reconciliations / total reconciliations)`
Please note that we differentiate between user errors caused by invalid input and internal or upstream API errors.
User errors are excluded from the calculation.

### Webhook[Internal Link: ¶](\#webhook)

Webhook is part of Kubernetes CRDs based API which allows you to manage ACI Kafka resources via Kube.
API Operator performs a set of validations on each Create/Update/Delete request to Kube API in order to reject bad input from users sooner rather than later.

#### Latency[Internal Link: ¶](\#id3)

We measure the webhook latency as the percentage of requests that took less than 10 seconds to complete: `1 - (slow requests / total requests)`

#### Availability[Internal Link: ¶](\#id4)

We measure the webhook availability as the percentage of requests that didn’t result in an error: `1 - (failed requests / total requests)`
Note: requests that were rejected due to bad user input are not considered as failed and are excluded from the calculation.


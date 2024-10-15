---
title: "End user responsibilities to use the platform by data tier¶"
weight: 3
---

| End user responsibilities | Tier 1 | Tier 2 | Tier 3 | Tier 4 |
| --- | --- | --- | --- | --- |
| Data must be **encrypted prior to being stored on the platform**. The **platform** will not provide data encryption functionality. | ✓ | x | x | x |
| Connections between your application/service and the **platform must use TLS**. | ✓ | ✓ | ✓ | x |
| Any public (over the internet) connections between your application/service and the **platform** must use TLS and have authentication. | ✓ | ✓ | ✓ | ✓ |
| This data **must never be logged in the platform logs** without Privacy and Security approval. | ✓ | ✓ | ✓\* | x |
| This data **must not exist in QA environments** using this platform. | ✓ | ✓ | x | x |
| **You are responsible for handling authentication** for your service/application, including restricting access to appropriate users only. | ✓ | ✓ | ✓ | x |
| You **must review access** to your cluster periodically. | ✓ | ✓ | x | x |
| You must comply with **user data privacy requirements** while storing customer data for this platform. | ✓ | ✓ | ✓ | x |
| Data destruction: No services support the feature to destroy the data after the end users’ usage. The end user is responsible **for deleting the data from the platform**. | ✓ | ✓ | ✓ | ✓ |
| Check with NPS the Project Sensitivity (Black/Ultra). | ✓ | ✓ | ✓ | ✓ |

- Only requires privacy approval.

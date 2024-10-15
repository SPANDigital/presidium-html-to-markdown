---
title: "Client identity topic access¶"
weight: 7
---

Access grants a specific client identity permissions to produce to or consume from a given topic.

Topic accesses include a produce or consume quota measured in megabytes.
This quota is only used as a guideline and is not enforced.
A client identity can still produce or consume up to the cluster quota limit as mentioned above.
The per assignment quota will be deprecated in the future.

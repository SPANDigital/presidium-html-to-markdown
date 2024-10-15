---
title: "Supported Data Types¶"
weight: 2
---

The table below lists the [data tiers](https://infosec.apple.com/guidance/policies/all/information-classification/) supported by ACI Kafka, **with some important exceptions (below)**. The minimum
[data handling requirements](https://infosec.apple.com/guidance/policies/all/information-classification/) are also included.

**Important — Exceptions to Data Tier Support**

Data tiers may include regulated data that is not yet supported by ACI Services. It is the responsibility of the team using ACI Services to identify if they are using any of these data types, and take the appropriate steps accordingly.

**Kafka is approved to transmit / process / store Apple Financial Reporting Data or SOX-related Data. Customers requesting SOX compliance must contact the Kafka team to enable the required audits for their cluster.**

For any of the following data types, you must reach out to [ASE Compliance]({{%baseurl%}}//files/mailto:ASE_Compliance%40group.apple.com) before leveraging an ASE Service:

- **Credit Cards:** PCI, Credit or Debit Cards, or any other form of Payment Cards
- **Financial:** Apple Financial Reporting Data or SOX-related Data
- **Health:** Health or HIPAA-related information
- **Privacy:** Data that may be in-scope for Privacy Regulations (e.g. GDPR, CCPA, LGPD)

| ACI Service | Supported Data Tier | Data Handling Requirements |
| --- | --- | --- |
| Hosted | Encrypted | Authenticated | Restricted | Stored | Logged | Destroyed |
| --- | --- | --- | --- | --- | --- | --- |
| Kafka | T1, T2, and T3 | In Production only | Over the internet or other non-Apple networks.<br>All [Tier 1 Data](https://infosec.apple.com/guidance/standards/all/information-handling/) must be encrypted at rest. | With Information Security approved authentication. | To users with a business need. | On production equipment only.<br>This information may not be exported to removable media, including laptops and workstations, without Information Security approval. | Never | Per policy, before disposal. |


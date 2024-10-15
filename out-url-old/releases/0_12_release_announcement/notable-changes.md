---
title: "Notable changes:¶"
weight: 2
---

We are launching Kafka Inspector, which is a read-only, authorized
deployment of the AMP Web UI. This provides visibility to the health of
your topics and offsets while keeping isolation using your customer
groups access controls.

We provide a direct link to Kafka Inspector from the namespace list
view, check out the [portal\
docs]({{< ref "self-service/portal" >}}) for more
information.

We are also launching the **kafka.pie.apple.com** subdomain to simplify
access to the portal and other tools.

You can now access portal using:

- IF1: [kafka.pie.apple.com/if1](https://kafka.pie.apple.com/if1).
- IF-PROD:
  [kafka.pie.apple.com/prod](https://kafka.pie.apple.com/prod).

Note: If you are connecting over VPN, you may need to switch to full
tunnel VPN to resolve these new subdomains; [check out the IS&T\
documentation](https://istweb.apple.com/vpn/ac-firsttime).

**Summary of changes**

- Kafka connect based mirroring in alpha for select customers.
- Start rolling out capacity validation and enforcement (allowing for
  fully self-service namespaces in the future and charge back).
- Migration to PIE Workflow for managing topic and client changes
  (this will allow the rest of the operations to be accessible in the
  future).

Note: Tasks triggered on the old workflow system will not be visible
after the switch. This will only impact auditing for a few weeks.

**Notable Bug fixes**

- [rdar://problem/39546710](https://rdar.apple.com/39546710) PIE
  client may be causing file descriptor leaks.


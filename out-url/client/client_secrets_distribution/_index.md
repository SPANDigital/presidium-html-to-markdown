---
title: "Client Secrets Distribution¶"
weight: 1
---

ACI Kafka Client requires sensitive configurations (passwords, keys
etc.) to enable [Internal Link: Security](security_settings.html).

RSA asymmetric encryption is used to protect credentials which are
generated when provisioning your client.

This document provides instructions on how to generate an RSA
public/private key pair, create client using the public key, and get
going with configuring client secrets.


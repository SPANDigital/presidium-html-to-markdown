---
title: "Generating an RSA public/private key pair¶"
weight: 2
---

- Use the OpenSSL tool to generate a public/private RSA key pair by
  executing the following command


  ```
  openssl genrsa 2048 | openssl pkcs8 -topk8 -nocrypt -outform pem > private_key.pem

  ```

- Extract the public key from the RSA key pair into a file
  `public_key.pem`


  ```
  openssl rsa -pubout -in private_key.pem -out public_key.pem

  ```


**You should keep the generated private key securely and reliably, since**
**if it is lost or compromised we’ll need to recreate the client from**
**scratch (which means downtime).**

Note

It is highly recommended to generate a new public/private key pair
per client, but it is possible to use the same public/private key pair
across multiple clients in the same customer group.


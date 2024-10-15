---
title: "Registering non-Java clients¶"
weight: 4
---

1. Create a Kafka topic using [Internal Link: CLI](../self-service/kafka_cli.html) or [Internal Link: API](../self-service/public_api.html). The documentation of [Internal Link: Kafka CLI](../self-service/kafka_cli.html) explains the concepts behind terms in `<angleBrackets>`.
2. Create a public/private key pair for SASL authentication.A random password is generated for each client identity which is encrypted using the customer’s public key. The customer can then use the private key to retrieve the password.


   ```
   openssl genrsa 2048 | openssl pkcs8 -topk8 -nocrypt -outform pem > private_key.pem

   openssl rsa -pubout -in private_key.pem -out public_key.pem

   ```

3. Create a [Internal Link: client identity](../concepts.html#client-identities) using CLI


   ```
   SASL_USERNAME="<customerName>.<groupName>.<environmentName>.<identity>"`

   kafka-cli identity --create \
   -e <env> --group-id <groupId> --identity-name <identity> \
   --public-key "$(cat ./path-to/public_key.pem)"

   ```



   Note

   `<customerName>.<groupName>.<environmentName>` is the [Internal Link: `<groupId>`](../self-service/kafka_cli.html#client-identities). This means that `<groupName>` is not `<groupId>` and `<environmentName>` is not necessarily `<env>`.

4. Retrieve the `SASL_PASSWORD`:


   ```
   ENCRYPTED_PASSWORD=$(kafka-cli identity --fetch -e <env> --group-id <groupId> --identity-name <identity> | jq '.state["encrypted-password"]' -r)
   SASL_PASSWORD=$(echo $ENCRYPTED_PASSWORD | base64 -d | openssl rsautl -decrypt -inkey ./path-to/private_key.pem)

   ```

5. For the [Internal Link: namespace `<kafkaNamespace>`](../self-service/kafka_cli.html#namespaces), assign topic access to the client identity created above:


   ```
   kafka-cli access --create \
   -e <env> -g <groupId> -t <topic> -n <kafkaNamespace> \
   --identity-group-id <groupId> --identity-name <identity> \
   --produce-bytes-per-sec <produce-rate> --consume-bytes-per-sec <consume-rate>

   ```



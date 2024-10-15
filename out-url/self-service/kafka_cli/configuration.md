---
title: "Configuration¶"
weight: 3
---

Note

IAS JWT token authentication is only available in `pie-kafka-cli>=0.4`.

1. [External Link: Install IAS CLI](https://docs.aci.apple.com/ias/docs/IAS_installation.html)
2. Generate IAS client for ACI Kafka API:


   ```
   ias client create <(echo '{"partition": "prod", "service": "aci-kafka-api", "region": "earth"}')

   ```


   Or using a `client.json` file:


   ```
   echo '{"partition": "prod", "service": "aci-kafka-api", "region": "earth"}' > client.json
   ias client create client.json

   ```


   If prompted for identity pick TouchID (e.g. `mac-...`).
3. Set environment variables using the new client details:


   ```
   export KAFKA_CLI_CLIENT_ID=<client_id from `ias client create` output>
   export KAFKA_CLI_CLIENT_SECRET=<client_secret from `ias client create` output>

   ```



---
title: "Configuration¶"
weight: 3
---

The following two sub-sections outline the required steps to set up the authorization header.
You can also use the [External Link: **kcurl**](https://docs.aci.apple.com/kafka/self-service/public_api.html#kcurl) example if you don’t want to repeat the code for each configuration.

Note

All public API requests must include an authorization header using the Identity Authentication Service (IAS) JSON Web Token (JWT): `Authorization: Bearer <IAS JWT id_token>`

### Setting up IAS client environment variables[Internal Link: ¶](\#setting-up-ias-client-environment-variables)

To generate ID tokens, you must set up your IAS client credentials using the steps below:

1. Open a new terminal session and follow the [External Link: Installing IAS CLI](https://docs.aci.apple.com/ias/docs/IAS_installation.html) instructions to install, update, and verify the command line interface (CLI).
2. To generate an IAS client for the ACI Kafka API, run the following command:


   ```
   ias client create <(echo '{"partition": "prod", "service": "aci-kafka-api", "region": "earth"}')

   ```


   You can also use a `client.json` file to generate an IAS client for the API:


   ```
   echo '{"partition": "prod", "service": "aci-kafka-api", "region": "earth"}' > client.json
   ias client create client.json

   ```


   If you are prompted to choose an identity, select **Touch ID**.
3. Set the environment variables using the new client details by running the two commands below:


   ```
   export KAFKA_CLI_CLIENT_ID=<client_id from ias client create output>
   export KAFKA_CLI_CLIENT_SECRET=<client_secret from ias client create output>

   ```


### Generating a JSON Web Token[Internal Link: ¶](\#generating-a-json-web-token)

You can generate a new ID token for the authorization header using the [External Link: IAS API](https://docs.aci.apple.com/ias/tutorials/client_credentials/index.html):

```
curl -X POST https://authorize.apple.com/oauth2/token \
    -u "${KAFKA_CLI_CLIENT_ID}:${KAFKA_CLI_CLIENT_SECRET}" \
    -d "grant_type=client_credentials" \
    -d "scope=openid corpds:ds:firstName corpds:ds:lastName corpds:ds:username corpds:ds:email" \
    -H "Content-Type: application/x-www-form-urlencoded"

```

Set the `Authorization: Bearer <IAS JWT id_token>` header using the value of `id_token` from the IAS JSON response when making requests to the ACI Kafka API.

If you have trouble running the commands above, see the [External Link: IAS documentation](https://docs.aci.apple.com/ias/docs/creating_an_IAS_application.html).

### kcurl[Internal Link: ¶](\#kcurl)

You can call the API using _kcurl_ and avoid repeating the API client setup code in the sections below by defining the `kcurl` function in advance.
You must install [External Link: `jq`](https://stedolan.github.io/jq/): `brew install jq` for macOS before defining the function.

To define the `kcurl` function, run the following command:

```
export KAFKA_CLI_CLIENT_ID=<client_id from ias client create output>
export KAFKA_CLI_CLIENT_SECRET=<client_secret from ias client create output>
export BASE_URL=https://api-prod.aci-kafka.apple.com/api/public/v2

kcurl() {
  kafka_ias_id_token=$(curl -fsS -X POST https://authorize.apple.com/oauth2/token \
      -u "${KAFKA_CLI_CLIENT_ID}:${KAFKA_CLI_CLIENT_SECRET}" \
      -d "grant_type=client_credentials" \
      -d "scope=openid corpds:ds:firstName corpds:ds:lastName corpds:ds:username corpds:ds:email" \
      -H "Content-Type: application/x-www-form-urlencoded" | jq -r '.id_token')
  curl -s -H "Authorization: Bearer ${kafka_ias_id_token}" "$@"
}

```

The `BASE_URL` for QA and Production use cases is required:

- For QA use cases: `https://api-if1.aci-kafka.apple.com/api/public/v2`
- For Production use cases: `https://api-prod.aci-kafka.apple.com/api/public/v2`


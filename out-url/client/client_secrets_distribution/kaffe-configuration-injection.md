---
title: "Kaffe Configuration Injection¶"
weight: 4
---

### What will Kaffe configuration injection do?[Internal Link: ¶](\#what-will-kaffe-configuration-injection-do)

Kaffe configuration injection’s function is to reduce the number of
manual steps that customers need to configure their clients.

- On client instantiation, we retrieve from Kaffe the encrypted client
  secrets.
- Kaffe Cluster Resolver (in the client) decrypts the credentials with
  the private key that you provided to use for client provisioning.
- Kaffe Cluster Resolver then plugs in these secrets into your
  producer/consumer configuration.

### Configuring your client to enable Kaffe configuration injection[Internal Link: ¶](\#configuring-your-client-to-enable-kaffe-configuration-injection)

By default, the client will read your private key from a PKCS#8
formatted file present on disk. If this doesn’t work for you, you can
provide your private key another way.

1. **File-based private key provisioning (default)**

   - (this first step is only for customers upgrading from a version
     older than 0.8) Transform your private key to the
     required format PKCS#8 using command:


     ```
     openssl pkcs8 -topk8 -inform PEM -outform PEM -nocrypt -in <existing_private_key.pem> -out <private_key.pem>

     ```

   - By default the ACI Kafka client is configured to retrieve your
     client’s private key from the path:
     `${PLATFORM_SECRETS_PATH}/pie_kafka_private_key`.

It’s easy with [External Link: PIE Secrets](https://docs.aci.apple.com/secrets) to add
the secret in the correct place.
2. **Different secret path**You can configure the location of the private key by setting the
   property:


   ```
    pie.queue.kaffe.client.private.key.location=<location of private key>.

   ```


   For example, with a secret called `kafkaPrivateKey.pem`,
   configure
   `pie.queue.kaffe.client.private.key.location=${PLATFORM_SECRETS_PATH}/kafkaPrivateKey.pem`
3. **Custom private key provisioning**If you want to provide your private key through another source such
   as a secret store, implement the `PrivateKeyProvider`
   interface and set the property`pie.queue.kaffe.private.key.provider=<class path to provider>`
4. **Environment variable**We implemented an environment variable and java property secret
   loader. To use it set the properties


   ```
   pie.queue.kaffe.private.key.provider=com.apple.pie.queue.kafka.client.kaffe.privatekeyprovider.EnvPropsPrivateKeyProvider
   pie.queue.kaffe.client.private.key.name=<YOUR_ENV_VAR_NAME>

   ```


Note

You should remove the properties `sasl.jaas.config`, `pie.queue.crypto.consumer.secret.key` and
`pie.queue.crypto.datastore.password` from your client configuration if
you are enabling configuration injection, since Kaffe injects these properties.


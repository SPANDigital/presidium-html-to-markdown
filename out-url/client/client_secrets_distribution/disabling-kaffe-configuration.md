---
title: "Disabling Kaffe Configuration Injection¶"
weight: 5
---

It is highly recommended to use Kaffe configuration injection, but it is
possible to still use the client by disabling it.

The default value for the property `pie.queue.kaffe.config.filter.class`
is a list of configuration filters in value
`KaffeConfig.KAFFE_CONFIG_FILTER_CONFIG_DEFAULT_VALUE`.

If you want disable the configuration injection, remove the
`DecryptionFilter.class` filter from the default value and set that as
the value of this property.

> ```
> Map<String, Object> configs = new HashMap<>(); // Kafka configs
> configs = KaffeConfig.removeDecryptionFilter(configs);
>
> ```

Note

The older versions of ACI Kafka Client (< 0.8) do not include the Kaffe configuration injection feature.

If you are using the old client, follow the instructions in the rest of the document to decrypt your client credentials and add the values to your client configurations.

- You will receive a `Getting Started Guide` email which includes two encrypted secrets:
1. authentication password to the cluster (to be put into `sasl.jaas.conf` property or JAAS configuration file)
2. consumer secret key (needed in consumer clients for decrypting encrypted payloads)
- Use the following command to decode and decrypt the received credentials


```
echo "<encrypted secret>" | base64 --decode | openssl rsautl -decrypt -inkey private_key.pem

```


**You should keep the decrypted credentials securely as they are highly sensitive.**

---
title: "Kafka client configuration¶"
weight: 3
---

If you are using the OpenTelemetry Java agent, Kafka clients should automatically be instrumented and no further code
changes should be necessary.

If you are not using the Java agent, follow the
[guide for configuring the instrumentation](https://github.com/open-telemetry/opentelemetry-java-instrumentation/tree/v2.4.0/instrumentation/kafka/kafka-clients/kafka-clients-2.6/library#usage-metrics),
after adding the instrumentation dependency.

It’s important to set up the OpenTelemetry SDK prior to the creation of any instrumented Kafka clients. For example,
using the autoconfigured SDK:

```
import io.opentelemetry.api.OpenTelemetry;
import io.opentelemetry.instrumentation.kafkaclients.v2_6.KafkaTelemetry;
import io.opentelemetry.sdk.autoconfigure.AutoConfiguredOpenTelemetrySdk;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerConfig;
import org.apache.kafka.common.serialization.StringSerializer;

import com.apple.pie.queue.kafka.client.KafkaClientUtils;

import java.util.HashMap;
import java.util.Map;

class MyApp {
  public static void main(String[] args) {
    // Set up the OpenTelemetry SDK early...
    OpenTelemetry otel = initOtelSdk();
    // ...before anything that uses it
    Map<String, ?> metricConfig = KafkaTelemetry.create(otel).metricConfigProperties();
    KafkaProducer<String, String> producer = createProducer(metricConfig);
    // Produce some data...
  }

  OpenTelemetry initOtelSDK() {
    return AutoConfiguredOpenTelemetrySdk.builder()
      .setResultAsGlobal()
      .build()
      .getOpenTelemetrySdk();
  }

  KafkaProducer<String, String> createProducer(Map<String, ?> extraConfigs) {
    Map<String, ?> configs = new HashMap<>();
    configs.put("pie.queue.kaffe.connect", "$kaffeConnect");
    configs.put("pie.queue.kaffe.client.id", "$kaffeClientId");
    configs.put("pie.queue.kaffe.namespace.id", "$kaffeNamespaceId");
    configs.put("pie.queue.kaffe.client.private.key.location", "$kaffePrivateKeyPath");
    configs.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class);
    configs.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, StringSerializer.class);

    configs.putAll(extraConfigs);

    return KafkaClientUtils.createProducer(configs);
  }
}

```


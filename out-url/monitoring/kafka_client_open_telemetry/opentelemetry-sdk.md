---
title: "OpenTelemetry SDK¶"
weight: 2
---

The first step in setting up the OpenTelemetry instrumentation is configuring the SDK. You have a few choices to do
this:

|  | SDK instantiation | SDK configuration | Instrumentation | Setup guide |
| --- | --- | --- | --- | --- |
| Java agent | Automatic | `otel.` config | Bundled, automatic instrumentation | [External Link: OpenTelemetry docs](https://github.com/open-telemetry/opentelemetry-java-instrumentation?tab=readme-ov-file#getting-started) |
| Autoconfigured SDK | In (minimal) code | `otel.` config + (optionally) code | Extra deps, manual instrumentation | [External Link: OpenTelemetry docs](https://github.com/open-telemetry/opentelemetry-java/tree/v1.38.0/sdk-extensions/autoconfigure) |
| Standard SDK | In code | In code | Extra deps, manual instrumentation | [External Link: Mosaic docs](https://telemetry.apple.com/docs/collection/opentelemetry) |

The auto-configured SDK generally provides a good balance of flexibility with standardized config options, but if you
want to instrument several other 3rd party libraries in your application then the Java agent may be a better choice.


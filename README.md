# <img src="./assets/xk6-opentelemetry-logo.png" alt="xk6-opentelemetry logo" style="height: 32px; width:32px;"/> xk6-opentelemetry

The xk6-opentelemetry project is a [k6 extension](https://k6.io/docs/extensions/guides/what-are-k6-extensions/) that enables k6 users to generate random [OpenTelemetry signals](https://opentelemetry.io/docs/reference/specification/glossary/#signals) (metrics, logs, traces) for testing purposes.

Check the [scripts](./scripts/) directory for examples.

## Features

- Generate ExportLogsServiceRequest
- Generate ExportMetricsServiceRequest
    - Types: gauge, sum, histogram
- Set static resource attributes

## Development

As the testing environment and scripts rely on sending messages to Kafka, the extension [xk6-kafka](https://github.com/mostafa/xk6-kafka) needs to be integrated as well when creating the k6 binary.

Create the binary:
```
git clone git@github.com:thmshmm/xk6-opentelemetry.git && cd xk6-opentelemetry
xk6 build --with xk6-opentelemetry=. --with github.com/mostafa/xk6-kafka
```

## Testing

### Local environment

Start a local test environment including Kafka (Redpanda) and an instance of the [OpenTelemetry Collector](https://opentelemetry.io/docs/collector/) using Docker Compose. The OTel Collector is configured to read messages using the Kafka receiver and write all messages to STDOUT of the container.

Start:
```
cd testing
docker compose -p xk6-opentelemetry up -d
```

Stop:
```
docker compose -p xk6-opentelemetry down
```

## Scripts

- **[test-logs.js](./scripts/test-logs.js)** - Create random log signals.
- **[test-metrics.js](./scripts/test-metrics.js)** - Create random gauge/sum metric signals.
- **[test-metrics-histogram.js](./scripts/test-metrics-histogram.js)** - Create random histogram metric signals.

Execute:
```
./k6 run scripts/<SCRIPT>.js
```

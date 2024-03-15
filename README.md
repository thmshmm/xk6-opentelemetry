# <img src="./assets/xk6-opentelemetry-logo.png" alt="xk6-opentelemetry logo" style="height: 32px; width:32px;"/> xk6-opentelemetry

The xk6-opentelemetry project is a [k6 extension](https://k6.io/docs/extensions/guides/what-are-k6-extensions/) that enables k6 users to generate random [OpenTelemetry signals](https://opentelemetry.io/docs/reference/specification/glossary/#signals) (metrics, logs, traces) for testing purposes.

Check the [examples](./examples/) directory which contains some scripts to get started.

## Features

- Generate
    - ExportLogsServiceRequest
    - ExportMetricsServiceRequest (Types: gauge, sum, histogram)
    - ExportTraceServiceRequest
- Set resource attributes

## Usage

### Docker

The easiest way to get started is to build the provided Docker image and run k6 scripts inside the container.

Build the image:
```
docker build -t xk6-opentelemetry .
```

Run a local k6 script:
```
docker run --rm -i xk6-opentelemetry run - <my-script.js
```

### Local build

As the testing environment and scripts rely on sending messages to Kafka, the extension [xk6-kafka](https://github.com/mostafa/xk6-kafka) needs to be integrated as well when creating the k6 binary.

Create the binary:
```
git clone git@github.com:thmshmm/xk6-opentelemetry.git && cd xk6-opentelemetry
xk6 build --with xk6-opentelemetry=. --with github.com/mostafa/xk6-kafka
```

or use [Task](https://taskfile.dev/):
```
task build-k6
```

## Testing

### Local environment

Start a local test environment including Kafka (Redpanda) and an instance of the [OpenTelemetry Collector](https://opentelemetry.io/docs/collector/) using Docker Compose. The OTel Collector is configured to read messages using the Kafka receiver and write all messages to STDOUT of the container.

Start:
```
task testing-up
```

Stop:
```
task testing-down
```

## Examples

- **[test-logs.js](./examples/test-logs.js)** - Create random log signals.
- **[test-metrics.js](./examples/test-metrics.js)** - Create random gauge/sum metric signals.
- **[test-metrics-histogram.js](./examples/test-metrics-histogram.js)** - Create random histogram metric signals.
- **[test-traces.js](./examples/test-traces.js)** - Create random trace signals.

Execute:
```
./k6 run examples/<SCRIPT>.js
```

Run all examples:
```
task run-examples
```

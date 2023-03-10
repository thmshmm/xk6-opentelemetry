import generator from 'k6/x/opentelemetry';
import {
    Writer,
    SchemaRegistry,
    SCHEMA_TYPE_BYTES,
} from 'k6/x/kafka';
import {
    randomItem,
    randomIntBetween,
} from 'https://jslib.k6.io/k6-utils/1.4.0/index.js';

const brokers = ["localhost:9092"]
const topic = "otel_metrics"
const schemaRegistry = new SchemaRegistry();
const producer = new Writer({
    brokers: brokers,
    topic: topic,
});

const metricNames = [
    "metric1",
    "metric2"
]

const metricAttributes = new Map()
metricAttributes.set("operation", "create")

export default function () {
    var messages = []

    const resourceAttributes = new Map()
    resourceAttributes.set("pod", "test-abc")

    // optionally set static attributes for the resource
    generator.setStaticResourceAttributes(resourceAttributes)

    for (let idx = 0; idx < 5; idx++) {
        messages[idx] = {
            value: schemaRegistry.serialize({
                data: Array.from(generator.exportMetricsServiceRequest({
                    "type": "gauge",
                    "name": "gauge_" + randomItem(metricNames),
                    "attributes": metricAttributes,
                    "data": {
                        "value": randomIntBetween(10, 100)
                    },
                })),
                schemaType: SCHEMA_TYPE_BYTES,
            })
        }
    }

    for (let idx = 0; idx < 5; idx++) {
        messages[idx] = {
            value: schemaRegistry.serialize({
                data: Array.from(generator.exportMetricsServiceRequest({
                    "type": "sum",
                    "name": "sum_" + randomItem(metricNames),
                    "attributes": metricAttributes,
                    "data": {
                        "value": randomIntBetween(10, 100),
                        "isMonotonic": false, // true or false
                        "aggregationTemporality": "delta" // delta or cumulative
                    },
                })),
                schemaType: SCHEMA_TYPE_BYTES,
            })
        }
    }

    producer.produce({ messages: messages })
}

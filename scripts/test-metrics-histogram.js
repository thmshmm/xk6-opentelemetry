import generator from 'k6/x/opentelemetry';
import {
    Writer,
    SchemaRegistry,
    SCHEMA_TYPE_BYTES,
} from 'k6/x/kafka';
import {
    randomIntBetween,
} from 'https://jslib.k6.io/k6-utils/1.4.0/index.js';

const brokers = ["localhost:9092"]
const topic = "otel_metrics"
const schemaRegistry = new SchemaRegistry();
const producer = new Writer({
    brokers: brokers,
    topic: topic,
});

// Example of having one cumulative histogram with 10 values after one execution.
export default function () {
    var messages = []
    var count = 0
    var sum = 0.0
    var min = 100.0
    var max = 0.0
    var bucketCounts = [0, 0, 0] // (-inf,10], (10,50], (50,+inf]

    for (let idx = 0; idx < 10; idx++) {
        var recordedValue = randomIntBetween(0, 100)

        var bucketIdx = getBucketIdx(recordedValue)
        bucketCounts[bucketIdx] = bucketCounts[bucketIdx] + 1

        sum += recordedValue
        min = Math.min(min, recordedValue)
        max = Math.max(max, recordedValue)

        messages[idx] = {
            value: schemaRegistry.serialize({
                data: Array.from(generator.exportMetricsServiceRequest({
                    "type": "histogram",
                    "name": "hist_metric1",
                    "unit": "ms",
                    "data": {
                        "aggregationTemporality": "cumulative",
                        "count": ++count,
                        "sum": sum,
                        "min": min,
                        "max": max,
                        "explicitBounds": [10, 50],
                        "bucketCounts": bucketCounts,
                    },
                })),
                schemaType: SCHEMA_TYPE_BYTES,
            })
        }
    }

    producer.produce({ messages: messages })
}

function getBucketIdx(value) {
    if (value <= 10) {
        return 0
    } else if (value > 10 && value <= 50) {
        return 1
    }

    return 2 // value > 50
}

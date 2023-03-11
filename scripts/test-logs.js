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
const topic = "otel_logs"
const schemaRegistry = new SchemaRegistry();
const producer = new Writer({
    brokers: brokers,
    topic: topic,
});

const logBodies = [
    "this is a log msg",
    "this is another log msg"
]

export default function () {
    var messages = []

    const resourceAttributes = new Map()
    resourceAttributes.set("pod", "test-abc")
    resourceAttributes.set("instance", 1)
    resourceAttributes.set("something", 1.23)
    resourceAttributes.set("isProd", true)

    // optionally set static attributes for the resource
    generator.setStaticResourceAttributes(resourceAttributes)

    for (let idx = 0; idx < 10; idx++) {
        messages[idx] = {
            value: schemaRegistry.serialize({
                data: Array.from(generator.exportLogsServiceRequest({
                    // "attributes": logAttributes, // optionally create log specific attribute map
                    "data": {
                        "body": randomItem(logBodies),
                        "severity": randomIntBetween(0, 24) // severity numbers from 0 to 24
                    },
                })),
                schemaType: SCHEMA_TYPE_BYTES,
            })
        }
    }

    producer.produce({ messages: messages })
}

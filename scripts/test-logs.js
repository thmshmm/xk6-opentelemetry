import generator from 'k6/x/opentelemetry';
import {
    Writer,
    SchemaRegistry,
    SCHEMA_TYPE_BYTES,
} from 'k6/x/kafka';

const brokers = ["localhost:9092"]
const topic = "otel_logs"
const schemaRegistry = new SchemaRegistry();
const producer = new Writer({
    brokers: brokers,
    topic: topic,
});

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
                // use 0 as parameter to create random body words from 5 to 10
                // or create a random number here
                data: Array.from(generator.exportLogsServiceRequest(5)),
                schemaType: SCHEMA_TYPE_BYTES,
            })
        }
    }

    producer.produce({ messages: messages })
}

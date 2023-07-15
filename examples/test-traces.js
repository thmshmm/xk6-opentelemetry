import generator from 'k6/x/opentelemetry';
import {
    Writer,
    SchemaRegistry,
    SCHEMA_TYPE_BYTES,
} from 'k6/x/kafka';

const brokers = ["localhost:9092"]
const topic = "otel_traces"
const schemaRegistry = new SchemaRegistry();
const producer = new Writer({
    brokers: brokers,
    topic: topic,
});

export default function () {
    var messages = []

    const resourceAttributes = new Map()
    resourceAttributes.set("pod", "test-abc")

    // make sure to set a service name for correct visualization in Grafana Tempo
    resourceAttributes.set("service.name", "service-1")

    // optionally set static attributes for the resource
    generator.setStaticResourceAttributes(resourceAttributes)

    for (let idx = 0; idx < 10; idx++) {
        var timeNow = generator.timeNowUnixNano()

        var traceId = generator.newTraceID()

        var span1Id = generator.newSpanID()
        var span1 = {
            "traceId": traceId,
            "spanId": span1Id,
            // root span has no parendSpanId
            "name": "say-hello",
            // "attributes": spanAttributes, // optionally create span specific attribute map
            "startTimeUnixNano": timeNow - 1e8, // 100ms earlier
            "endTimeUnixNano": timeNow,
        }

        var span2Id = generator.newSpanID()
        var span2 = {
            "traceId": traceId,
            "spanId": span2Id,
            "parentSpanId": span1Id,
            "name": "random-name",
            "kind": "client",
            "startTimeUnixNano": timeNow - 9e7,
            "endTimeUnixNano": timeNow - 4e7,
        }

        var span3Id = generator.newSpanID()
        var span3 = {
            "traceId": traceId,
            "spanId": span3Id,
            "parentSpanId": span1Id,
            "name": "enrich",
            "startTimeUnixNano": timeNow - 3e7,
            "endTimeUnixNano": timeNow - 1e7,
        }

        messages[idx] = {
            value: schemaRegistry.serialize({
                data: Array.from(generator.exportTraceServiceRequest({
                    "data": [
                        span1,
                        span2,
                        span3,
                    ],
                })),
                schemaType: SCHEMA_TYPE_BYTES,
            })
        }
    }

    producer.produce({ messages: messages })
}

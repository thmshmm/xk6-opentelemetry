package generator

import (
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	resourcepb "go.opentelemetry.io/proto/otlp/resource/v1"
)

func (g *Generator) SetStaticResourceAttributes(attrs map[string]interface{}) {
	attributes := make([]*commonpb.KeyValue, 0)

	for k, v := range attrs {
		attributes = append(attributes, &commonpb.KeyValue{
			Key:   k,
			Value: ToAnyValue(v),
		})
	}

	g.staticResourceAttributes = attributes
}

func (g *Generator) resource() *resourcepb.Resource {
	var attributes []*commonpb.KeyValue

	if g.staticResourceAttributes != nil {
		attributes = g.staticResourceAttributes
	} else {
		attributes = g.randomResourceAttributes()
	}

	return &resourcepb.Resource{
		Attributes: attributes,
	}
}

func (g *Generator) randomResourceAttributes() []*commonpb.KeyValue {
	return []*commonpb.KeyValue{
		{
			Key: "host",
			Value: &commonpb.AnyValue{
				Value: &commonpb.AnyValue_StringValue{
					StringValue: g.faker.Word() + ".example.com",
				},
			},
		},
	}
}

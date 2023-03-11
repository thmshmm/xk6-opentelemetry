package generator

import (
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	resourcepb "go.opentelemetry.io/proto/otlp/resource/v1"
)

func (g *Generator) SetStaticResourceAttributes(attrs map[string]interface{}) {
	g.staticResourceAttributes = ToAttributes(attrs)
}

func (g *Generator) resource() *resourcepb.Resource {
	var attributes []*commonpb.KeyValue

	if g.staticResourceAttributes != nil {
		attributes = g.staticResourceAttributes
	}

	return &resourcepb.Resource{
		Attributes: attributes,
	}
}

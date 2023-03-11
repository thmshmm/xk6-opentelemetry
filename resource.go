package generator

import (
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	resourcepb "go.opentelemetry.io/proto/otlp/resource/v1"
)

func resource(attrs []*commonpb.KeyValue) *resourcepb.Resource {
	return &resourcepb.Resource{
		Attributes: attrs,
	}
}

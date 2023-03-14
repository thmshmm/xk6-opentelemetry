package generator

import (
	coltracepb "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	tracepb "go.opentelemetry.io/proto/otlp/trace/v1"
)

func ExportTraceServiceRequest(
	resourceAttrs []*commonpb.KeyValue,
	config TraceConfig,
) *coltracepb.ExportTraceServiceRequest {
	return &coltracepb.ExportTraceServiceRequest{
		ResourceSpans: []*tracepb.ResourceSpans{
			ResourceSpans(resourceAttrs, config),
		},
	}
}

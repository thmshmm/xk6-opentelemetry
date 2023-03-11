package generator

import (
	collogspb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	logspb "go.opentelemetry.io/proto/otlp/logs/v1"
)

func ExportLogsServiceRequest(
	resourceAttrs []*commonpb.KeyValue,
	config LogConfig,
) *collogspb.ExportLogsServiceRequest {
	return &collogspb.ExportLogsServiceRequest{
		ResourceLogs: []*logspb.ResourceLogs{
			ResourceLogs(resourceAttrs, config),
		},
	}
}

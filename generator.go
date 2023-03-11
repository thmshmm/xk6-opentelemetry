package generator

import (
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
)

type Generator struct {
	staticResourceAttributes []*commonpb.KeyValue
}

// NewGenerator creates a Generator instance which can create supported OpenTelemetry signals.
func NewGenerator() *Generator {
	return &Generator{}
}

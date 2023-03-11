package generator

import (
	"github.com/sirupsen/logrus"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	"google.golang.org/protobuf/proto"
)

type Generator struct {
	staticResourceAttributes []*commonpb.KeyValue
}

// NewGenerator creates a Generator instance which can create supported OpenTelemetry signals.
func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) SetStaticResourceAttributes(attrs map[string]interface{}) {
	g.staticResourceAttributes = ToAttributes(attrs)
}

func (g *Generator) ExportLogsServiceRequest(config LogConfig) []byte {
	request := ExportLogsServiceRequest(g.staticResourceAttributes, config)

	data, err := proto.Marshal(request)
	if err != nil {
		logrus.Errorf("Failed to marshal ExportLogsServiceRequest: %v", err)

		return nil
	}

	return data
}

func (g *Generator) ExportMetricsServiceRequest(config MetricConfig) []byte {
	request := ExportMetricsServiceRequest(g.staticResourceAttributes, config)

	data, err := proto.Marshal(request)
	if err != nil {
		logrus.Errorf("Failed to marshal ExportMetricsServiceRequest: %v", err)

		return nil
	}

	return data
}

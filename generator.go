package generator

import (
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	"google.golang.org/protobuf/proto"
)

type Generator struct {
	*IDGenerator
	staticResourceAttributes []*commonpb.KeyValue
}

// NewGenerator creates a Generator instance which can create supported OpenTelemetry signals.
func NewGenerator() *Generator {
	return &Generator{
		IDGenerator: &IDGenerator{
			randSource: rand.New(rand.NewSource(time.Now().UnixNano())),
		},
	}
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

func (g *Generator) ExportTraceServiceRequest(config TraceConfig) []byte {
	request := ExportTraceServiceRequest(g.staticResourceAttributes, config)

	data, err := proto.Marshal(request)
	if err != nil {
		logrus.Errorf("Failed to marshal ExportTraceServiceRequest: %v", err)

		return nil
	}

	return data
}

func (g *Generator) TimeNowUnixNano() int64 {
	return time.Now().UnixNano()
}

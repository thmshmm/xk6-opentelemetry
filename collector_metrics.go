package generator

import (
	"github.com/sirupsen/logrus"
	colmetricspb "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	metricspb "go.opentelemetry.io/proto/otlp/metrics/v1"
	"google.golang.org/protobuf/proto"
)

func (g *Generator) ExportMetricsServiceRequest(config MetricConfig) []byte {
	request := colmetricspb.ExportMetricsServiceRequest{
		ResourceMetrics: []*metricspb.ResourceMetrics{
			g.resourceMetrics(config),
		},
	}

	data, err := proto.Marshal(&request)
	if err != nil {
		logrus.Errorf("Failed to marshal ExportMetricsServiceRequest: %v", err)

		return nil
	}

	return data
}

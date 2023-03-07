package generator

import (
	"time"

	"github.com/sirupsen/logrus"
	v1 "go.opentelemetry.io/proto/otlp/common/v1"
	metricspb "go.opentelemetry.io/proto/otlp/metrics/v1"
)

type MetricConfig struct {
	Type       string                 `json:"type"`
	Name       string                 `json:"name"`
	Attributes map[string]interface{} `json:"attributes"`
	Value      interface{}            `json:"value"`
}

func (g *Generator) resourceMetrics(config MetricConfig) *metricspb.ResourceMetrics {
	metric := &metricspb.Metric{
		Name: config.Name,
	}

	switch config.Type {
	case "gauge":
		metric.Data = g.gauge(ToAttributes(config.Attributes), config.Value)
	default:
		logrus.Errorf("Unimplemented metric type %q, use one of [gauge]", config.Type)

		return nil
	}

	return &metricspb.ResourceMetrics{
		Resource: g.resource(),
		ScopeMetrics: []*metricspb.ScopeMetrics{
			{
				Metrics: []*metricspb.Metric{
					metric,
				},
			},
		},
	}
}

func (g *Generator) gauge(attrs []*v1.KeyValue, value interface{}) *metricspb.Metric_Gauge {
	intValue, ok := value.(int64)
	if !ok {
		logrus.Error("Failed to parse gauge value, expected int")

		return nil
	}

	return &metricspb.Metric_Gauge{
		Gauge: &metricspb.Gauge{
			DataPoints: []*metricspb.NumberDataPoint{
				{
					Attributes:   attrs,
					TimeUnixNano: uint64(time.Now().UnixNano()),
					Value: &metricspb.NumberDataPoint_AsInt{
						AsInt: intValue,
					},
				},
			},
		},
	}
}

package generator

import (
	"strings"

	"github.com/sirupsen/logrus"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	metricspb "go.opentelemetry.io/proto/otlp/metrics/v1"
)

type MetricConfig struct {
	Type       string                 `json:"type"`
	Name       string                 `json:"name"`
	Unit       string                 `json:"unit"`
	Attributes map[string]interface{} `json:"attributes"`
	Data       map[string]interface{} `json:"data"`
}

func ResourceMetrics(resourceAttrs []*commonpb.KeyValue, config MetricConfig) *metricspb.ResourceMetrics {
	metric := &metricspb.Metric{
		Name: config.Name,
		Unit: config.Unit,
	}

	attrs := ToAttributes(config.Attributes)

	switch config.Type {
	case "gauge":
		data, err := parseGaugeData(config.Data)
		if err != nil {
			logrus.Error(err)

			return nil
		}

		metric.Data = gauge(attrs, data)
	case "sum":
		data, err := parseSumData(config.Data)
		if err != nil {
			logrus.Error(err)

			return nil
		}

		metric.Data = sum(attrs, data)
	case "histogram":
		data, err := parseHistogramData(config.Data)
		if err != nil {
			logrus.Error(err)

			return nil
		}

		metric.Data = histogram(attrs, data)
	default:
		logrus.Errorf("Unimplemented metric type %q, use one of [gauge, sum]", config.Type)

		return nil
	}

	return &metricspb.ResourceMetrics{
		Resource: resource(resourceAttrs),
		ScopeMetrics: []*metricspb.ScopeMetrics{
			{
				Metrics: []*metricspb.Metric{
					metric,
				},
			},
		},
	}
}

func getAggregationTemporality(temporality string) metricspb.AggregationTemporality {
	switch strings.ToLower(temporality) {
	case "delta":
		return metricspb.AggregationTemporality_AGGREGATION_TEMPORALITY_DELTA
	case "cumulative":
		return metricspb.AggregationTemporality_AGGREGATION_TEMPORALITY_CUMULATIVE
	}

	return metricspb.AggregationTemporality_AGGREGATION_TEMPORALITY_UNSPECIFIED
}

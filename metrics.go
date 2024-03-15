package generator

import (
	"log/slog"
	"strings"

	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	metricspb "go.opentelemetry.io/proto/otlp/metrics/v1"
)

type MetricConfig struct {
	Type       string                 `js:"type"`
	Name       string                 `js:"name"`
	Unit       string                 `js:"unit"`
	Attributes map[string]interface{} `js:"attributes"`
	Data       map[string]interface{} `js:"data"`
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
			slog.Error("Failed to parse gauge data", "error", err)

			return nil
		}

		metric.Data = gauge(attrs, data)
	case "sum":
		data, err := parseSumData(config.Data)
		if err != nil {
			slog.Error("Failed to parse sum data", "error", err)

			return nil
		}

		metric.Data = sum(attrs, data)
	case "histogram":
		data, err := parseHistogramData(config.Data)
		if err != nil {
			slog.Error("Failed to parse histogram data", "error", err)

			return nil
		}

		metric.Data = histogram(attrs, data)
	default:
		slog.Error("Unimplemented metric type %q, use one of [gauge, sum]", "type", config.Type)

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

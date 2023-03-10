package generator

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	v1 "go.opentelemetry.io/proto/otlp/common/v1"
	metricspb "go.opentelemetry.io/proto/otlp/metrics/v1"
)

type MetricConfig struct {
	Type       string                 `json:"type"`
	Name       string                 `json:"name"`
	Attributes map[string]interface{} `json:"attributes"`
	Data       map[string]interface{} `json:"data"`
}

func (g *Generator) resourceMetrics(config MetricConfig) *metricspb.ResourceMetrics {
	metric := &metricspb.Metric{
		Name: config.Name,
	}

	attrs := ToAttributes(config.Attributes)

	switch config.Type {
	case "gauge":
		data, err := parseGaugeData(config.Data)
		if err != nil {
			logrus.Error(err)

			return nil
		}

		metric.Data = g.gauge(attrs, data)
	case "sum":
		data, err := parseSumData(config.Data)
		if err != nil {
			logrus.Error(err)

			return nil
		}

		metric.Data = g.sum(attrs, data)
	default:
		logrus.Errorf("Unimplemented metric type %q, use one of [gauge, sum]", config.Type)

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

type GaugeData struct {
	Value int64 `json:"value"`
}

func parseGaugeData(rawData map[string]interface{}) (*GaugeData, error) {
	dataBytes, err := json.Marshal(rawData)
	if err != nil {
		return nil, fmt.Errorf("failed to process provided gauge data: %w", err)
	}

	var data GaugeData
	err = json.Unmarshal(dataBytes, &data)

	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON gauge data: %w", err)
	}

	return &data, nil
}

func (g *Generator) gauge(attrs []*v1.KeyValue, data *GaugeData) *metricspb.Metric_Gauge {
	return &metricspb.Metric_Gauge{
		Gauge: &metricspb.Gauge{
			DataPoints: []*metricspb.NumberDataPoint{
				{
					Attributes:   attrs,
					TimeUnixNano: uint64(time.Now().UnixNano()),
					Value: &metricspb.NumberDataPoint_AsInt{
						AsInt: data.Value,
					},
				},
			},
		},
	}
}

type SumData struct {
	Value                  int64  `json:"value"`
	IsMonotonic            bool   `json:"isMonotonic"`
	AggregationTemporality string `json:"aggregationTemporality"`
}

func parseSumData(rawData map[string]interface{}) (*SumData, error) {
	dataBytes, err := json.Marshal(rawData)
	if err != nil {
		return nil, fmt.Errorf("failed to process provided sum data: %w", err)
	}

	var data SumData
	err = json.Unmarshal(dataBytes, &data)

	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON sum data: %w", err)
	}

	return &data, nil
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

func (g *Generator) sum(attrs []*v1.KeyValue, data *SumData) *metricspb.Metric_Sum {
	return &metricspb.Metric_Sum{
		Sum: &metricspb.Sum{
			IsMonotonic:            data.IsMonotonic,
			AggregationTemporality: getAggregationTemporality(data.AggregationTemporality),
			DataPoints: []*metricspb.NumberDataPoint{
				{
					Attributes:   attrs,
					TimeUnixNano: uint64(time.Now().UnixNano()),
					Value: &metricspb.NumberDataPoint_AsInt{
						AsInt: data.Value,
					},
				},
			},
		},
	}
}

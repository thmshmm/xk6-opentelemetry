package generator

import (
	"encoding/json"
	"fmt"
	"time"

	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	metricspb "go.opentelemetry.io/proto/otlp/metrics/v1"
)

type HistogramData struct {
	AggregationTemporality string    `json:"aggregationTemporality"`
	Count                  uint64    `json:"count"`
	Sum                    float64   `json:"sum"`
	Min                    float64   `json:"min"`
	Max                    float64   `json:"max"`
	ExplicitBounds         []float64 `json:"explicitBounds"`
	BucketCounts           []uint64  `json:"bucketCounts"`
}

func parseHistogramData(rawData map[string]interface{}) (*HistogramData, error) {
	dataBytes, err := json.Marshal(rawData)
	if err != nil {
		return nil, fmt.Errorf("failed to process provided histogram data: %w", err)
	}

	var data HistogramData
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON histogram data: %w", err)
	}

	return &data, nil
}

func histogram(attrs []*commonpb.KeyValue, data *HistogramData) *metricspb.Metric_Histogram {
	return &metricspb.Metric_Histogram{
		Histogram: &metricspb.Histogram{
			AggregationTemporality: getAggregationTemporality(data.AggregationTemporality),
			DataPoints: []*metricspb.HistogramDataPoint{
				{
					Attributes:     attrs,
					TimeUnixNano:   uint64(time.Now().UnixNano()),
					Count:          data.Count,
					Sum:            &data.Sum,
					Min:            &data.Min,
					Max:            &data.Max,
					ExplicitBounds: data.ExplicitBounds,
					BucketCounts:   data.BucketCounts,
				},
			},
		},
	}
}

package generator

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	metricspb "go.opentelemetry.io/proto/otlp/metrics/v1"
)

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

func sum(attrs []*commonpb.KeyValue, data *SumData) *metricspb.Metric_Sum {
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

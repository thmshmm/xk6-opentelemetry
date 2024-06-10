package generator

import (
	"encoding/json"
	"fmt"
	"time"

	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	metricspb "go.opentelemetry.io/proto/otlp/metrics/v1"
)

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

func gauge(attrs []*commonpb.KeyValue, data *GaugeData) *metricspb.Metric_Gauge {
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

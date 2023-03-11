package generator

import (
	"time"

	v1 "go.opentelemetry.io/proto/otlp/common/v1"
	logspb "go.opentelemetry.io/proto/otlp/logs/v1"
)

type LogConfig struct {
	Attributes map[string]interface{} `json:"attributes"`
	Data       LogData                `json:"data"`
}

type LogData struct {
	Body     string `json:"body"`
	Severity int32  `json:"severity"`
}

func (g *Generator) resourceLogs(config LogConfig) *logspb.ResourceLogs {
	return &logspb.ResourceLogs{
		Resource: g.resource(),
		ScopeLogs: []*logspb.ScopeLogs{
			{
				LogRecords: []*logspb.LogRecord{
					g.logRecord(ToAttributes(config.Attributes), config.Data),
				},
			},
		},
	}
}

func (g *Generator) logRecord(attrs []*v1.KeyValue, data LogData) *logspb.LogRecord {
	return &logspb.LogRecord{
		Attributes:     attrs,
		TimeUnixNano:   uint64(time.Now().UnixNano()),
		SeverityNumber: logspb.SeverityNumber(data.Severity),
		Body: &v1.AnyValue{
			Value: &v1.AnyValue_StringValue{
				StringValue: data.Body,
			},
		},
	}
}

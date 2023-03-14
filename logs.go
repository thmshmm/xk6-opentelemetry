package generator

import (
	"time"

	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	logspb "go.opentelemetry.io/proto/otlp/logs/v1"
)

type LogConfig struct {
	Attributes map[string]interface{} `js:"attributes"`
	Data       LogData                `js:"data"`
}

type LogData struct {
	Body     string `js:"body"`
	Severity int32  `js:"severity"`
}

func ResourceLogs(resourceAttrs []*commonpb.KeyValue, config LogConfig) *logspb.ResourceLogs {
	return &logspb.ResourceLogs{
		Resource: resource(resourceAttrs),
		ScopeLogs: []*logspb.ScopeLogs{
			{
				LogRecords: []*logspb.LogRecord{
					logRecord(ToAttributes(config.Attributes), config.Data),
				},
			},
		},
	}
}

func logRecord(attrs []*commonpb.KeyValue, data LogData) *logspb.LogRecord {
	return &logspb.LogRecord{
		Attributes:     attrs,
		TimeUnixNano:   uint64(time.Now().UnixNano()),
		SeverityNumber: logspb.SeverityNumber(data.Severity),
		Body: &commonpb.AnyValue{
			Value: &commonpb.AnyValue_StringValue{
				StringValue: data.Body,
			},
		},
	}
}

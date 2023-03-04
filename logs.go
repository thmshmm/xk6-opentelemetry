package generator

import (
	"math/rand"
	"time"

	v1 "go.opentelemetry.io/proto/otlp/common/v1"
	logspb "go.opentelemetry.io/proto/otlp/logs/v1"
)

const (
	defaultMinBodyWordCount = 5
	defaultMaxBodyWordCount = 10
	maxSeverityNumber       = 24
)

func (g *Generator) resourceLogs(bodyWordCount int) *logspb.ResourceLogs {
	return &logspb.ResourceLogs{
		Resource: g.resource(),
		ScopeLogs: []*logspb.ScopeLogs{
			{
				LogRecords: []*logspb.LogRecord{
					g.logRecord(bodyWordCount),
				},
			},
		},
	}
}

func (g *Generator) logRecord(bodyWordCount int) *logspb.LogRecord {
	return &logspb.LogRecord{
		TimeUnixNano:   uint64(time.Now().UnixNano()),
		SeverityNumber: g.logSeverityNumber(),
		Body:           g.logBody(bodyWordCount),
	}
}

func (g *Generator) logBody(wordCount int) *v1.AnyValue {
	if wordCount == 0 {
		wordCount = defaultMinBodyWordCount + rand.Intn(defaultMaxBodyWordCount-defaultMinBodyWordCount+1)
	}

	return &v1.AnyValue{
		Value: &v1.AnyValue_StringValue{
			StringValue: g.faker.Sentence(wordCount),
		},
	}
}

// severityNumber returns a random severity number [1, 24].
// See opentelemetry/proto/logs/v1/logs.proto for available severity numbers.
func (g *Generator) logSeverityNumber() logspb.SeverityNumber {
	randSeverityNumber := rand.Int31n(maxSeverityNumber) + 1

	return logspb.SeverityNumber(randSeverityNumber)
}

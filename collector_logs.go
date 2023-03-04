package generator

import (
	"github.com/sirupsen/logrus"
	collogspb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	logspb "go.opentelemetry.io/proto/otlp/logs/v1"
	"google.golang.org/protobuf/proto"
)

func (g *Generator) ExportLogsServiceRequest(bodyWordCount int) []byte {
	exportLogsServiceRequest := collogspb.ExportLogsServiceRequest{
		ResourceLogs: []*logspb.ResourceLogs{
			g.resourceLogs(bodyWordCount),
		},
	}

	data, err := proto.Marshal(&exportLogsServiceRequest)
	if err != nil {
		logrus.Error("Failed to marshal ExportLogsServiceRequest: %w", err)

		return nil
	}

	return data
}

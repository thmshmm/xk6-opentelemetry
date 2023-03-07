package generator

import (
	"github.com/sirupsen/logrus"
	collogspb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	logspb "go.opentelemetry.io/proto/otlp/logs/v1"
	"google.golang.org/protobuf/proto"
)

func (g *Generator) ExportLogsServiceRequest(bodyWordCount int) []byte {
	request := collogspb.ExportLogsServiceRequest{
		ResourceLogs: []*logspb.ResourceLogs{
			g.resourceLogs(bodyWordCount),
		},
	}

	data, err := proto.Marshal(&request)
	if err != nil {
		logrus.Errorf("Failed to marshal ExportLogsServiceRequest: %v", err)

		return nil
	}

	return data
}

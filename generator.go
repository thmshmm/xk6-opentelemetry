package generator

import (
	"github.com/brianvoe/gofakeit/v6"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
)

type Generator struct {
	faker                    *gofakeit.Faker
	staticResourceAttributes []*commonpb.KeyValue
}

// NewGenerator creates a Generator instance which can create supported OpenTelemetry signals.
func NewGenerator() *Generator {
	return &Generator{
		faker: gofakeit.NewCrypto(),
	}
}

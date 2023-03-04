package generator

import (
	"github.com/brianvoe/gofakeit/v6"
)

type Generator struct {
	faker *gofakeit.Faker
}

// NewGenerator creates a Generator instance which can create supported OpenTelemetry signals.
func NewGenerator() *Generator {
	return &Generator{
		faker: gofakeit.NewCrypto(),
	}
}

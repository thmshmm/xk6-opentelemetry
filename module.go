package generator

import (
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/opentelemetry", New())
}

type (
	RootModule     struct{}
	ModuleInstance struct {
		generator *Generator
	}
)

var (
	_ modules.Instance = &ModuleInstance{}
	_ modules.Module   = &RootModule{}
)

// New returns a pointer to a new RootModule instance.
func New() *RootModule {
	return &RootModule{}
}

// NewModuleInstance implements the modules.Module interface returning a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &ModuleInstance{
		generator: NewGenerator(),
	}
}

// Exports implements the modules.Instance interface and returns the exported types for the JS module.
func (mi *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{
		Default: mi.generator,
	}
}

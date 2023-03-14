package generator

import (
	"encoding/hex"
	"math/rand"
	"sync"
)

type IDGenerator struct {
	sync.Mutex
	randSource *rand.Rand
}

func (idGen *IDGenerator) NewTraceID() string {
	idGen.Lock()
	defer idGen.Unlock()

	var id [16]byte

	_, _ = idGen.randSource.Read(id[:])

	return hex.EncodeToString(id[:])
}

func (idGen *IDGenerator) NewSpanID() string {
	idGen.Lock()
	defer idGen.Unlock()

	var id [8]byte

	_, _ = idGen.randSource.Read(id[:])

	return hex.EncodeToString(id[:])
}

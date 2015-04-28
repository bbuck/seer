package nodes

import (
	"github.com/satori/go.uuid"
)

type Node interface {
	ID() uuid.UUID
	Type() NodeType
}

func newId() uuid.UUID {
	return uuid.NewV4()
}

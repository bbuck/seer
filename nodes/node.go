package nodes

import (
	"github.com/satori/go.uuid"
	"github.com/tree-server/trees/neo4j"
)

// Interface defining what it means to be a "Node"
type Node interface {
	ID() uuid.UUID
	Type() NodeType
	ValidRelationships() map[RelationshipType]struct{}
	RelateToNode(Node, RelationsihpType) bool
	Save() bool
}

func IsValidRelationship(n Node, r RelationshipType) bool {
	rels := n.ValidRelationships()
	_, ok := rels[r]

	return ok
}

func newId() uuid.UUID {
	return uuid.NewV4()
}

TODO: finish implementation
func Load(id string) Node {
	uuid := uuid.FromString(id)

}

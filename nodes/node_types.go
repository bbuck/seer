package nodes

import (
	"errors"
)

type NodeType int8

const (
	StartNode NodeType = iota
	RouteNode
	ServerNode
)

// String makes NodeType implement fmt.Stringer
func (n NodeType) String() string {
	switch n {
	case StartNode:
		return "Start"
	case RouteNode:
		return "Route"
	case ServerNode:
		return "Server"
	default:
		return "InvalidNode"
	}
}

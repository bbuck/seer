package nodes

type RelationshipType int8

const (
	// Event
	OnLoad RelationshipType = iota
	OnNodeSaved
	OnRequest
	OnResponse
	OnRender

	// Standard
	Route
	Controller
	Render
)

package nodes

type RelationshipType int8

const (
	// Start
	OnLoadOut RelationshipType = iota
	OnNodeSaved

	// Start, Route
	RouteOut

	// Route
	RouteIn
	ControllerOut

	// Route, Server
	RenderOut

	// Start, Route
	OnRequestOut
	OnResponseOut

	// View, Start
	OnRenderOut

	// Server
	OnRequestIn
	OnResponseIn
	OnRenderIn
	OnLoadIn
	OnNodeSavedIn
	ControllerIn
)

// RelationshipType
func (r RelatinoshipType) String() string {

}

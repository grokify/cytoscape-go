// Package cytoscape provides Go types and utilities for Cytoscape.js graph visualization.
package cytoscape

// Element represents a Cytoscape.js element (node or edge).
type Element struct {
	// Group is "nodes" or "edges".
	Group string `json:"group"`

	// Data contains the element's data properties.
	Data ElementData `json:"data"`

	// Classes are CSS classes applied to the element.
	Classes string `json:"classes,omitempty"`

	// Position is the node's position (nodes only).
	Position *Position `json:"position,omitempty"`
}

// ElementData contains the data properties of an element.
type ElementData struct {
	// ID is the unique identifier (required for nodes).
	ID string `json:"id"`

	// Source is the source node ID (edges only).
	Source string `json:"source,omitempty"`

	// Target is the target node ID (edges only).
	Target string `json:"target,omitempty"`

	// Label is the display label.
	Label string `json:"label,omitempty"`

	// Parent is the parent node ID for compound nodes.
	Parent string `json:"parent,omitempty"`

	// Type is a custom type field for filtering.
	Type string `json:"type,omitempty"`

	// Group is for community/cluster grouping.
	Group string `json:"group,omitempty"`

	// Weight is for edge weight.
	Weight float64 `json:"weight,omitempty"`

	// Extra holds additional custom properties.
	Extra map[string]any `json:"extra,omitempty"`
}

// Position represents x,y coordinates.
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Node creates a node element.
func Node(id, label string) *Element {
	return &Element{
		Group: "nodes",
		Data: ElementData{
			ID:    id,
			Label: label,
		},
	}
}

// NodeWithType creates a node element with a type.
func NodeWithType(id, label, nodeType string) *Element {
	return &Element{
		Group: "nodes",
		Data: ElementData{
			ID:    id,
			Label: label,
			Type:  nodeType,
		},
	}
}

// NodeWithGroup creates a node element with a group/community.
func NodeWithGroup(id, label, nodeType, group string) *Element {
	return &Element{
		Group: "nodes",
		Data: ElementData{
			ID:    id,
			Label: label,
			Type:  nodeType,
			Group: group,
		},
	}
}

// Edge creates an edge element.
func Edge(id, source, target string) *Element {
	return &Element{
		Group: "edges",
		Data: ElementData{
			ID:     id,
			Source: source,
			Target: target,
		},
	}
}

// EdgeWithLabel creates an edge element with a label.
func EdgeWithLabel(id, source, target, label string) *Element {
	return &Element{
		Group: "edges",
		Data: ElementData{
			ID:     id,
			Source: source,
			Target: target,
			Label:  label,
		},
	}
}

// EdgeWithType creates an edge element with a type.
func EdgeWithType(id, source, target, label, edgeType string) *Element {
	return &Element{
		Group: "edges",
		Data: ElementData{
			ID:     id,
			Source: source,
			Target: target,
			Label:  label,
			Type:   edgeType,
		},
	}
}

// SetClass sets the CSS classes for the element.
func (e *Element) SetClass(class string) *Element {
	e.Classes = class
	return e
}

// SetPosition sets the position for a node.
func (e *Element) SetPosition(x, y float64) *Element {
	e.Position = &Position{X: x, Y: y}
	return e
}

// SetParent sets the parent node for compound graphs.
func (e *Element) SetParent(parentID string) *Element {
	e.Data.Parent = parentID
	return e
}

// SetExtra sets additional custom data.
func (e *Element) SetExtra(key string, value any) *Element {
	if e.Data.Extra == nil {
		e.Data.Extra = make(map[string]any)
	}
	e.Data.Extra[key] = value
	return e
}

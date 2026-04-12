package cytoscape

import (
	"encoding/json"
	"fmt"
)

// Graph is a container for Cytoscape.js elements and configuration.
type Graph struct {
	// Elements are the nodes and edges.
	Elements []*Element `json:"elements"`

	// Style is the Cytoscape.js stylesheet.
	Style []StyleRule `json:"-"`

	// Layout configures the graph layout algorithm.
	Layout Layout `json:"-"`

	// Metadata holds graph-level metadata.
	Metadata *GraphMetadata `json:"-"`
}

// GraphMetadata contains information about the graph.
type GraphMetadata struct {
	Title       string
	Description string
	NodeCount   int
	EdgeCount   int
	NodeTypes   map[string]int
	EdgeTypes   map[string]int
	Groups      []string
}

// NewGraph creates an empty graph.
func NewGraph() *Graph {
	return &Graph{
		Elements: make([]*Element, 0),
		Style:    DefaultStyle(),
		Layout:   &CoseLayout{},
		Metadata: &GraphMetadata{
			NodeTypes: make(map[string]int),
			EdgeTypes: make(map[string]int),
		},
	}
}

// AddNode adds a node to the graph.
func (g *Graph) AddNode(n *Element) *Graph {
	g.Elements = append(g.Elements, n)
	if g.Metadata != nil {
		g.Metadata.NodeCount++
		if n.Data.Type != "" {
			g.Metadata.NodeTypes[n.Data.Type]++
		}
	}
	return g
}

// AddEdge adds an edge to the graph.
func (g *Graph) AddEdge(e *Element) *Graph {
	g.Elements = append(g.Elements, e)
	if g.Metadata != nil {
		g.Metadata.EdgeCount++
		if e.Data.Type != "" {
			g.Metadata.EdgeTypes[e.Data.Type]++
		}
	}
	return g
}

// SetLayout sets the layout algorithm.
func (g *Graph) SetLayout(layout Layout) *Graph {
	g.Layout = layout
	return g
}

// SetStyle sets the stylesheet.
func (g *Graph) SetStyle(style []StyleRule) *Graph {
	g.Style = style
	return g
}

// SetTitle sets the graph title.
func (g *Graph) SetTitle(title string) *Graph {
	if g.Metadata == nil {
		g.Metadata = &GraphMetadata{}
	}
	g.Metadata.Title = title
	return g
}

// Nodes returns only the node elements.
func (g *Graph) Nodes() []*Element {
	var nodes []*Element
	for _, e := range g.Elements {
		if e.Group == "nodes" {
			nodes = append(nodes, e)
		}
	}
	return nodes
}

// Edges returns only the edge elements.
func (g *Graph) Edges() []*Element {
	var edges []*Element
	for _, e := range g.Elements {
		if e.Group == "edges" {
			edges = append(edges, e)
		}
	}
	return edges
}

// ElementsJSON returns the elements as JSON for embedding.
func (g *Graph) ElementsJSON() (string, error) {
	data, err := json.Marshal(g.Elements)
	if err != nil {
		return "", fmt.Errorf("marshaling elements: %w", err)
	}
	return string(data), nil
}

// StyleJSON returns the style as JSON for embedding.
func (g *Graph) StyleJSON() (string, error) {
	data, err := json.Marshal(g.Style)
	if err != nil {
		return "", fmt.Errorf("marshaling style: %w", err)
	}
	return string(data), nil
}

// LayoutJSON returns the layout config as JSON for embedding.
func (g *Graph) LayoutJSON() (string, error) {
	data, err := json.Marshal(g.Layout.Config())
	if err != nil {
		return "", fmt.Errorf("marshaling layout: %w", err)
	}
	return string(data), nil
}

// CollectGroups extracts unique groups from nodes.
func (g *Graph) CollectGroups() []string {
	seen := make(map[string]bool)
	var groups []string
	for _, e := range g.Elements {
		if e.Group == "nodes" && e.Data.Group != "" {
			if !seen[e.Data.Group] {
				seen[e.Data.Group] = true
				groups = append(groups, e.Data.Group)
			}
		}
	}
	if g.Metadata != nil {
		g.Metadata.Groups = groups
	}
	return groups
}

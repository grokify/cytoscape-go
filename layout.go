package cytoscape

// Layout is the interface for layout algorithms.
type Layout interface {
	Name() string
	Config() map[string]any
}

// CoseLayout is a force-directed layout good for general graphs.
// See: https://js.cytoscape.org/#layouts/cose
type CoseLayout struct {
	// Animate whether to animate the layout
	Animate bool
	// NodeRepulsion higher values push nodes apart more
	NodeRepulsion int
	// IdealEdgeLength target edge length
	IdealEdgeLength int
	// EdgeElasticity how much edges should stretch
	EdgeElasticity int
	// Gravity pulls nodes toward center
	Gravity float64
	// NumIter max iterations
	NumIter int
	// Padding around the graph
	Padding int
}

func (l *CoseLayout) Name() string { return "cose" }

func (l *CoseLayout) Config() map[string]any {
	cfg := map[string]any{
		"name":    "cose",
		"animate": l.Animate,
		"padding": 30,
	}
	if l.NodeRepulsion > 0 {
		cfg["nodeRepulsion"] = func(node any) int { return l.NodeRepulsion }
	}
	if l.IdealEdgeLength > 0 {
		cfg["idealEdgeLength"] = func(edge any) int { return l.IdealEdgeLength }
	}
	if l.Gravity > 0 {
		cfg["gravity"] = l.Gravity
	}
	if l.NumIter > 0 {
		cfg["numIter"] = l.NumIter
	}
	if l.Padding > 0 {
		cfg["padding"] = l.Padding
	}
	return cfg
}

// DagreLayout is a hierarchical layout good for DAGs and ERDs.
// Requires cytoscape-dagre extension.
// See: https://github.com/cytoscape/cytoscape.js-dagre
type DagreLayout struct {
	// RankDir direction: TB (top-bottom), BT, LR, RL
	RankDir string
	// NodeSep separation between nodes
	NodeSep int
	// RankSep separation between ranks
	RankSep int
	// EdgeSep separation between edges
	EdgeSep int
	// Padding around the graph
	Padding int
}

func (l *DagreLayout) Name() string { return "dagre" }

func (l *DagreLayout) Config() map[string]any {
	cfg := map[string]any{
		"name":    "dagre",
		"rankDir": "TB",
		"nodeSep": 50,
		"rankSep": 80,
		"edgeSep": 30,
		"padding": 30,
	}
	if l.RankDir != "" {
		cfg["rankDir"] = l.RankDir
	}
	if l.NodeSep > 0 {
		cfg["nodeSep"] = l.NodeSep
	}
	if l.RankSep > 0 {
		cfg["rankSep"] = l.RankSep
	}
	if l.EdgeSep > 0 {
		cfg["edgeSep"] = l.EdgeSep
	}
	if l.Padding > 0 {
		cfg["padding"] = l.Padding
	}
	return cfg
}

// GridLayout arranges nodes in a grid.
type GridLayout struct {
	Rows    int
	Cols    int
	Padding int
}

func (l *GridLayout) Name() string { return "grid" }

func (l *GridLayout) Config() map[string]any {
	cfg := map[string]any{
		"name":    "grid",
		"padding": 30,
	}
	if l.Rows > 0 {
		cfg["rows"] = l.Rows
	}
	if l.Cols > 0 {
		cfg["cols"] = l.Cols
	}
	if l.Padding > 0 {
		cfg["padding"] = l.Padding
	}
	return cfg
}

// CircleLayout arranges nodes in a circle.
type CircleLayout struct {
	Padding int
	Radius  int
}

func (l *CircleLayout) Name() string { return "circle" }

func (l *CircleLayout) Config() map[string]any {
	cfg := map[string]any{
		"name":    "circle",
		"padding": 30,
	}
	if l.Radius > 0 {
		cfg["radius"] = l.Radius
	}
	if l.Padding > 0 {
		cfg["padding"] = l.Padding
	}
	return cfg
}

// ConcentricLayout arranges nodes in concentric circles by degree or attribute.
type ConcentricLayout struct {
	Padding        int
	MinNodeSpacing int
	// Concentric function returns the concentric level for a node
	// Higher values = closer to center
	LevelAttr string // Use node data attribute for level
}

func (l *ConcentricLayout) Name() string { return "concentric" }

func (l *ConcentricLayout) Config() map[string]any {
	cfg := map[string]any{
		"name":           "concentric",
		"padding":        30,
		"minNodeSpacing": 50,
	}
	if l.MinNodeSpacing > 0 {
		cfg["minNodeSpacing"] = l.MinNodeSpacing
	}
	if l.Padding > 0 {
		cfg["padding"] = l.Padding
	}
	return cfg
}

// BreadthFirstLayout arranges nodes in a BFS tree.
type BreadthFirstLayout struct {
	Roots    []string // Root node IDs
	Directed bool
	Padding  int
	Circle   bool // Arrange in a circle instead of tree
}

func (l *BreadthFirstLayout) Name() string { return "breadthfirst" }

func (l *BreadthFirstLayout) Config() map[string]any {
	cfg := map[string]any{
		"name":     "breadthfirst",
		"directed": l.Directed,
		"padding":  30,
		"circle":   l.Circle,
	}
	if len(l.Roots) > 0 {
		cfg["roots"] = l.Roots
	}
	if l.Padding > 0 {
		cfg["padding"] = l.Padding
	}
	return cfg
}

// PresetLayout uses positions specified in node data.
type PresetLayout struct {
	Padding int
}

func (l *PresetLayout) Name() string { return "preset" }

func (l *PresetLayout) Config() map[string]any {
	cfg := map[string]any{
		"name":    "preset",
		"padding": 30,
	}
	if l.Padding > 0 {
		cfg["padding"] = l.Padding
	}
	return cfg
}

// ColaLayout is an advanced force-directed layout.
// Requires cytoscape-cola extension.
type ColaLayout struct {
	Animate     bool
	MaxSimTime  int // Max simulation time in ms
	NodeSpacing int
	EdgeLength  int
	Padding     int
	Randomize   bool
	Convergence float64
}

func (l *ColaLayout) Name() string { return "cola" }

func (l *ColaLayout) Config() map[string]any {
	cfg := map[string]any{
		"name":              "cola",
		"animate":           l.Animate,
		"padding":           30,
		"randomize":         false,
		"maxSimulationTime": 2000,
	}
	if l.MaxSimTime > 0 {
		cfg["maxSimulationTime"] = l.MaxSimTime
	}
	if l.NodeSpacing > 0 {
		cfg["nodeSpacing"] = l.NodeSpacing
	}
	if l.EdgeLength > 0 {
		cfg["edgeLength"] = l.EdgeLength
	}
	if l.Padding > 0 {
		cfg["padding"] = l.Padding
	}
	if l.Convergence > 0 {
		cfg["convergenceThreshold"] = l.Convergence
	}
	return cfg
}

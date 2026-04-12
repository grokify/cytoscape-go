package cytoscape

// StyleRule is a Cytoscape.js stylesheet rule.
type StyleRule struct {
	Selector string         `json:"selector"`
	Style    map[string]any `json:"style"`
}

// NewStyleRule creates a new style rule.
func NewStyleRule(selector string) *StyleRule {
	return &StyleRule{
		Selector: selector,
		Style:    make(map[string]any),
	}
}

// Set sets a style property.
func (r *StyleRule) Set(key string, value any) *StyleRule {
	r.Style[key] = value
	return r
}

// DefaultStyle returns a default stylesheet suitable for general graphs.
func DefaultStyle() []StyleRule {
	return []StyleRule{
		{
			Selector: "node",
			Style: map[string]any{
				"background-color":      "#6c757d",
				"label":                 "data(label)",
				"text-valign":           "center",
				"text-halign":           "center",
				"font-family":           "system-ui, -apple-system, sans-serif",
				"font-size":             "12px",
				"color":                 "#ffffff",
				"text-outline-color":    "#6c757d",
				"text-outline-width":    "2px",
				"width":                 "40px",
				"height":                "40px",
				"border-width":          "2px",
				"border-color":          "#495057",
				"text-max-width":        "100px",
				"text-wrap":             "ellipsis",
				"text-overflow-wrap":    "anywhere",
			},
		},
		{
			Selector: "node:selected",
			Style: map[string]any{
				"background-color": "#0d6efd",
				"border-color":     "#0a58ca",
				"border-width":     "3px",
			},
		},
		{
			Selector: "node.highlighted",
			Style: map[string]any{
				"background-color": "#0d6efd",
				"border-color":     "#0a58ca",
			},
		},
		{
			Selector: "edge",
			Style: map[string]any{
				"width":              2,
				"line-color":         "#adb5bd",
				"target-arrow-color": "#adb5bd",
				"target-arrow-shape": "triangle",
				"curve-style":        "bezier",
				"opacity":            0.7,
			},
		},
		{
			Selector: "edge:selected",
			Style: map[string]any{
				"line-color":         "#0d6efd",
				"target-arrow-color": "#0d6efd",
				"width":              3,
				"opacity":            1,
			},
		},
		{
			Selector: "edge.highlighted",
			Style: map[string]any{
				"line-color":         "#0d6efd",
				"target-arrow-color": "#0d6efd",
				"width":              3,
				"opacity":            1,
			},
		},
		{
			Selector: "node.faded",
			Style: map[string]any{
				"opacity": 0.2,
			},
		},
		{
			Selector: "edge.faded",
			Style: map[string]any{
				"opacity": 0.1,
			},
		},
	}
}

// CodeGraphStyle returns a stylesheet optimized for code/call graphs.
func CodeGraphStyle() []StyleRule {
	return []StyleRule{
		// Base node style
		{
			Selector: "node",
			Style: map[string]any{
				"background-color":   "#e9ecef",
				"label":              "data(label)",
				"text-valign":        "center",
				"text-halign":        "center",
				"font-family":        "'SF Mono', Monaco, 'Courier New', monospace",
				"font-size":          "11px",
				"color":              "#212529",
				"width":              "label",
				"height":             "30px",
				"padding":            "10px",
				"shape":              "round-rectangle",
				"border-width":       "1px",
				"border-color":       "#dee2e6",
				"text-max-width":     "150px",
				"text-wrap":          "ellipsis",
			},
		},
		// Node types
		{
			Selector: "node[type='package']",
			Style: map[string]any{
				"background-color": "#cfe2ff",
				"border-color":     "#9ec5fe",
				"shape":            "round-rectangle",
			},
		},
		{
			Selector: "node[type='file']",
			Style: map[string]any{
				"background-color": "#e2e3e5",
				"border-color":     "#c4c8cb",
			},
		},
		{
			Selector: "node[type='function']",
			Style: map[string]any{
				"background-color": "#d1e7dd",
				"border-color":     "#a3cfbb",
			},
		},
		{
			Selector: "node[type='method']",
			Style: map[string]any{
				"background-color": "#cff4fc",
				"border-color":     "#9eeaf9",
			},
		},
		{
			Selector: "node[type='struct']",
			Style: map[string]any{
				"background-color": "#fff3cd",
				"border-color":     "#ffe69c",
			},
		},
		{
			Selector: "node[type='interface']",
			Style: map[string]any{
				"background-color": "#f8d7da",
				"border-color":     "#f1aeb5",
			},
		},
		// Selected/highlighted
		{
			Selector: "node:selected",
			Style: map[string]any{
				"background-color": "#0d6efd",
				"border-color":     "#0a58ca",
				"border-width":     "2px",
				"color":            "#ffffff",
			},
		},
		{
			Selector: "node.highlighted",
			Style: map[string]any{
				"border-color": "#0d6efd",
				"border-width": "2px",
			},
		},
		// Edges
		{
			Selector: "edge",
			Style: map[string]any{
				"width":              1.5,
				"line-color":         "#adb5bd",
				"target-arrow-color": "#adb5bd",
				"target-arrow-shape": "triangle",
				"curve-style":        "bezier",
				"opacity":            0.6,
			},
		},
		{
			Selector: "edge[type='calls']",
			Style: map[string]any{
				"line-color":         "#198754",
				"target-arrow-color": "#198754",
			},
		},
		{
			Selector: "edge[type='imports']",
			Style: map[string]any{
				"line-color":         "#0d6efd",
				"target-arrow-color": "#0d6efd",
				"line-style":         "dashed",
			},
		},
		{
			Selector: "edge[type='implements']",
			Style: map[string]any{
				"line-color":         "#dc3545",
				"target-arrow-color": "#dc3545",
				"line-style":         "dotted",
			},
		},
		{
			Selector: "edge:selected",
			Style: map[string]any{
				"width":   3,
				"opacity": 1,
			},
		},
		{
			Selector: "edge.highlighted",
			Style: map[string]any{
				"width":   2.5,
				"opacity": 1,
			},
		},
		// Faded (for filtering)
		{
			Selector: "node.faded",
			Style: map[string]any{
				"opacity": 0.15,
			},
		},
		{
			Selector: "edge.faded",
			Style: map[string]any{
				"opacity": 0.05,
			},
		},
	}
}

// ERDStyle returns a stylesheet optimized for entity-relationship diagrams.
func ERDStyle() []StyleRule {
	return []StyleRule{
		{
			Selector: "node",
			Style: map[string]any{
				"background-color":   "#f8f9fa",
				"border-color":       "#dee2e6",
				"border-width":       1,
				"label":              "data(label)",
				"text-valign":        "top",
				"text-halign":        "center",
				"text-margin-y":      8,
				"font-family":        "system-ui, -apple-system, sans-serif",
				"font-size":          14,
				"font-weight":        "bold",
				"color":              "#212529",
				"shape":              "round-rectangle",
				"width":              "label",
				"height":             50,
				"padding-top":        "30px",
				"padding-bottom":     "10px",
				"padding-left":       "15px",
				"padding-right":      "15px",
			},
		},
		{
			Selector: "node:selected",
			Style: map[string]any{
				"background-color": "#cfe2ff",
				"border-color":     "#0d6efd",
				"border-width":     2,
			},
		},
		{
			Selector: "edge",
			Style: map[string]any{
				"width":                    2,
				"line-color":               "#6c757d",
				"target-arrow-color":       "#6c757d",
				"target-arrow-shape":       "triangle",
				"curve-style":              "bezier",
				"label":                    "data(label)",
				"font-family":              "system-ui, -apple-system, sans-serif",
				"font-size":                11,
				"color":                    "#495057",
				"text-background-color":    "#ffffff",
				"text-background-opacity":  0.9,
				"text-background-padding":  "3px",
				"text-rotation":            "autorotate",
			},
		},
		{
			Selector: "edge:selected",
			Style: map[string]any{
				"line-color":         "#0d6efd",
				"target-arrow-color": "#0d6efd",
				"width":              3,
			},
		},
	}
}

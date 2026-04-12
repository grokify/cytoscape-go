# cytoscape-go

Go library for generating [Cytoscape.js](https://js.cytoscape.org/) graph visualizations.

[![Go Reference](https://pkg.go.dev/badge/github.com/grokify/cytoscape-go.svg)](https://pkg.go.dev/github.com/grokify/cytoscape-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/grokify/cytoscape-go)](https://goreportcard.com/report/github.com/grokify/cytoscape-go)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Features

- Build Cytoscape.js graphs programmatically in Go
- Generate standalone HTML visualizations with embedded JavaScript
- Multiple layout algorithms (Cose, Dagre, Cola, Circle, Grid, etc.)
- Customizable stylesheets with pre-built themes for code graphs and ERDs
- Interactive features: search, filters, export (PNG/SVG), minimap
- Dark mode support

## Installation

```bash
go get github.com/grokify/cytoscape-go
```

## Quick Start

```go
package main

import (
    "os"

    cytoscape "github.com/grokify/cytoscape-go"
)

func main() {
    // Create a graph
    g := cytoscape.NewGraph()
    g.SetTitle("My Graph")
    g.SetStyle(cytoscape.DefaultStyle())
    g.SetLayout(&cytoscape.DagreLayout{RankDir: "TB"})

    // Add nodes
    g.AddNode(cytoscape.NodeWithType("n1", "Node 1", "type-a"))
    g.AddNode(cytoscape.NodeWithType("n2", "Node 2", "type-b"))
    g.AddNode(cytoscape.NodeWithType("n3", "Node 3", "type-a"))

    // Add edges
    g.AddEdge(cytoscape.EdgeWithType("e1", "n1", "n2", "", "connects"))
    g.AddEdge(cytoscape.EdgeWithType("e2", "n2", "n3", "", "connects"))

    // Generate HTML
    html, _ := g.ToHTML(cytoscape.DefaultHTMLOptions())
    os.WriteFile("graph.html", html, 0644)
}
```

## Graph Building

### Creating Elements

```go
// Simple node
node := cytoscape.Node("id", "Label")

// Node with type (for styling)
node := cytoscape.NodeWithType("id", "Label", "function")

// Node with group/community
node := cytoscape.NodeWithGroup("id", "Label", "function", "auth")

// Simple edge
edge := cytoscape.Edge("edge-id", "source-id", "target-id")

// Edge with label
edge := cytoscape.EdgeWithLabel("e1", "src", "tgt", "calls")

// Edge with type (for styling)
edge := cytoscape.EdgeWithType("e1", "src", "tgt", "label", "calls")
```

### Element Methods

```go
// Set CSS class
node.SetClass("highlighted important")

// Set position (for preset layout)
node.SetPosition(100.0, 200.0)

// Set parent for compound nodes
node.SetParent("parent-id")

// Add custom data
node.SetExtra("file", "main.go")
node.SetExtra("line", 42)
```

## Layouts

### Force-Directed (Cose)

```go
g.SetLayout(&cytoscape.CoseLayout{
    Animate:         true,
    NodeRepulsion:   8000,
    IdealEdgeLength: 100,
    Gravity:         0.25,
})
```

### Hierarchical (Dagre)

Requires cytoscape-dagre extension (auto-loaded in HTML).

```go
g.SetLayout(&cytoscape.DagreLayout{
    RankDir: "TB",  // TB, BT, LR, RL
    NodeSep: 50,
    RankSep: 80,
})
```

### Cola (Advanced Force-Directed)

Requires cytoscape-cola extension (auto-loaded in HTML).

```go
g.SetLayout(&cytoscape.ColaLayout{
    Animate:     true,
    MaxSimTime:  2000,
    NodeSpacing: 30,
})
```

### Other Layouts

```go
// Circle
g.SetLayout(&cytoscape.CircleLayout{})

// Grid
g.SetLayout(&cytoscape.GridLayout{Rows: 3, Cols: 4})

// Concentric
g.SetLayout(&cytoscape.ConcentricLayout{MinNodeSpacing: 50})

// Breadth-First Tree
g.SetLayout(&cytoscape.BreadthFirstLayout{
    Roots:    []string{"root-node"},
    Directed: true,
})

// Preset (manual positions)
g.SetLayout(&cytoscape.PresetLayout{})
```

## Stylesheets

### Pre-built Styles

```go
// General purpose
g.SetStyle(cytoscape.DefaultStyle())

// Optimized for code/call graphs
g.SetStyle(cytoscape.CodeGraphStyle())

// Optimized for ERD diagrams
g.SetStyle(cytoscape.ERDStyle())
```

### Custom Styles

```go
style := []cytoscape.StyleRule{
    {
        Selector: "node",
        Style: map[string]any{
            "background-color": "#6c757d",
            "label":            "data(label)",
            "width":            "40px",
            "height":           "40px",
        },
    },
    {
        Selector: "node[type='important']",
        Style: map[string]any{
            "background-color": "#dc3545",
            "border-width":     "3px",
        },
    },
    {
        Selector: "edge",
        Style: map[string]any{
            "line-color":         "#adb5bd",
            "target-arrow-shape": "triangle",
            "curve-style":        "bezier",
        },
    },
}
g.SetStyle(style)
```

## HTML Generation

### Options

```go
opts := cytoscape.HTMLOptions{
    Title:              "My Graph",
    Description:        "Graph visualization",
    SourceURL:          "https://github.com/example/repo",
    ShowSearch:         true,
    ShowFilters:        true,
    ShowLegend:         true,
    ShowStats:          true,
    ShowExport:         true,
    ShowLayoutSelector: true,
    ShowMinimap:        false,
    DarkMode:           false,
    MaxLabelLength:     30,
    UseDagre:           true,
    UseCola:            true,
}

html, err := g.ToHTML(opts)
```

### Pre-configured Options

```go
// General purpose
opts := cytoscape.DefaultHTMLOptions()

// Code graphs
opts := cytoscape.CodeGraphHTMLOptions()

// ERD diagrams
opts := cytoscape.ERDHTMLOptions()
```

## JSON Export

```go
// Export elements as JSON (Cytoscape format)
json, err := g.ToJSON()
```

## Use Cases

- **Knowledge Graphs**: Visualize code dependencies, call graphs
- **Entity-Relationship Diagrams**: Database schema visualization
- **Network Diagrams**: Infrastructure, API relationships
- **Dependency Graphs**: Package dependencies, build systems

## Projects Using cytoscape-go

- [graphize](https://github.com/plexusone/graphize) - LLM-powered knowledge graphs for Go codebases
- [entscape](https://github.com/grokify/entscape) - Ent.go schema visualization

## License

MIT

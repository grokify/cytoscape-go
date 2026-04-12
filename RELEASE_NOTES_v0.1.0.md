# v0.1.0

**Release Date:** 2026-04-12

Initial release of cytoscape-go, a Go library for generating [Cytoscape.js](https://js.cytoscape.org/) graph visualizations.

## Highlights

- Build Cytoscape.js graphs programmatically in Go
- Generate standalone HTML visualizations with embedded JavaScript
- Multiple layout algorithms with pre-built styles
- Interactive features out of the box

## Features

### Graph Building

Create nodes and edges with typed metadata:

```go
g := cytoscape.NewGraph()
g.SetTitle("My Graph")

// Add nodes with types for styling
g.AddNode(cytoscape.NodeWithType("n1", "Node 1", "function"))
g.AddNode(cytoscape.NodeWithType("n2", "Node 2", "type"))

// Add edges with labels
g.AddEdge(cytoscape.EdgeWithLabel("e1", "n1", "n2", "calls"))
```

### Layout Algorithms

8 layout algorithms supported:

| Layout | Description | Use Case |
|--------|-------------|----------|
| Cose | Force-directed | General graphs |
| Dagre | Hierarchical | Call graphs, trees |
| Cola | Advanced force-directed | Large graphs |
| Circle | Circular arrangement | Cycles, rings |
| Grid | Grid arrangement | Regular structures |
| Concentric | Concentric circles | Centrality visualization |
| BreadthFirst | Tree layout | Hierarchies |
| Preset | Manual positions | Custom layouts |

### Pre-built Styles

Three style presets for common use cases:

- `DefaultStyle()` - General purpose with neutral colors
- `CodeGraphStyle()` - Optimized for code/call graphs with distinct node types
- `ERDStyle()` - Entity-relationship diagrams with table-like nodes

### HTML Generation

Generate standalone HTML with interactive features:

```go
opts := cytoscape.HTMLOptions{
    Title:              "My Graph",
    ShowSearch:         true,
    ShowFilters:        true,
    ShowLegend:         true,
    ShowExport:         true,
    ShowLayoutSelector: true,
    UseDagre:           true,
}

html, _ := g.ToHTML(opts)
os.WriteFile("graph.html", html, 0644)
```

Interactive features include:

- Search box for finding nodes
- Type filters for nodes and edges
- Export to PNG/SVG
- Layout switching at runtime
- Dark mode support

## Installation

```bash
go get github.com/grokify/cytoscape-go
```

## Projects Using cytoscape-go

- [graphize](https://github.com/plexusone/graphize) - LLM-powered knowledge graphs for Go codebases
- [entscape](https://github.com/grokify/entscape) - Ent.go schema visualization

## Links

- [Documentation](https://pkg.go.dev/github.com/grokify/cytoscape-go)
- [GitHub Repository](https://github.com/grokify/cytoscape-go)
- [Changelog](CHANGELOG.md)

package cytoscape

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
)

//go:embed templates/*.html
var templatesFS embed.FS

// HTMLOptions configures HTML generation.
type HTMLOptions struct {
	// Title is the page title.
	Title string

	// Description is shown in the header.
	Description string

	// SourceURL is a link to the source (optional).
	SourceURL string

	// ShowSearch enables the search box.
	ShowSearch bool

	// ShowFilters enables node/edge type filters.
	ShowFilters bool

	// ShowLegend enables the type legend.
	ShowLegend bool

	// ShowStats enables the stats display.
	ShowStats bool

	// ShowExport enables export buttons (PNG, SVG).
	ShowExport bool

	// ShowLayoutSelector enables layout switching.
	ShowLayoutSelector bool

	// ShowMinimap enables the minimap (for large graphs).
	ShowMinimap bool

	// DarkMode uses dark theme.
	DarkMode bool

	// MaxLabelLength truncates long labels.
	MaxLabelLength int

	// UseDagre loads the dagre extension for hierarchical layouts.
	UseDagre bool

	// UseCola loads the cola extension for advanced force-directed.
	UseCola bool
}

// DefaultHTMLOptions returns sensible defaults.
func DefaultHTMLOptions() HTMLOptions {
	return HTMLOptions{
		Title:              "Graph",
		ShowSearch:         true,
		ShowFilters:        true,
		ShowLegend:         true,
		ShowStats:          true,
		ShowExport:         true,
		ShowLayoutSelector: true,
		MaxLabelLength:     30,
		UseDagre:           true,
	}
}

// CodeGraphHTMLOptions returns options for code/call graphs.
func CodeGraphHTMLOptions() HTMLOptions {
	return HTMLOptions{
		Title:              "Code Graph",
		Description:        "Knowledge Graph",
		ShowSearch:         true,
		ShowFilters:        true,
		ShowLegend:         true,
		ShowStats:          true,
		ShowExport:         true,
		ShowLayoutSelector: true,
		MaxLabelLength:     40,
		UseDagre:           true,
		UseCola:            true,
	}
}

// ERDHTMLOptions returns options for ERD diagrams.
func ERDHTMLOptions() HTMLOptions {
	return HTMLOptions{
		Title:              "Entity-Relationship Diagram",
		ShowSearch:         true,
		ShowFilters:        false,
		ShowLegend:         false,
		ShowStats:          true,
		ShowExport:         true,
		ShowLayoutSelector: true,
		MaxLabelLength:     50,
		UseDagre:           true,
	}
}

// templateData holds data passed to the HTML template.
type templateData struct {
	Title        string
	Description  string
	SourceURL    string
	ElementsJSON template.JS
	StyleJSON    template.JS
	LayoutJSON   template.JS
	Options      HTMLOptions
	NodeTypes    []string
	EdgeTypes    []string
	NodeCount    int
	EdgeCount    int
	Groups       []string
}

// ToHTML generates a standalone HTML file from the graph.
func (g *Graph) ToHTML(opts HTMLOptions) ([]byte, error) {
	// Set defaults
	if opts.Title == "" {
		opts.Title = "Graph"
	}

	// Generate JSON
	elementsJSON, err := g.ElementsJSON()
	if err != nil {
		return nil, err
	}

	styleJSON, err := g.StyleJSON()
	if err != nil {
		return nil, err
	}

	layoutJSON, err := g.LayoutJSON()
	if err != nil {
		return nil, err
	}

	// Collect metadata
	g.CollectGroups()
	nodeTypes := make([]string, 0, len(g.Metadata.NodeTypes))
	for t := range g.Metadata.NodeTypes {
		nodeTypes = append(nodeTypes, t)
	}
	edgeTypes := make([]string, 0, len(g.Metadata.EdgeTypes))
	for t := range g.Metadata.EdgeTypes {
		edgeTypes = append(edgeTypes, t)
	}

	// Check if using dagre layout
	if _, ok := g.Layout.(*DagreLayout); ok {
		opts.UseDagre = true
	}
	if _, ok := g.Layout.(*ColaLayout); ok {
		opts.UseCola = true
	}

	// JSON is generated from internal data structures via json.Marshal,
	// which properly escapes special characters. Not user input.
	data := templateData{
		Title:        opts.Title,
		Description:  opts.Description,
		SourceURL:    opts.SourceURL,
		ElementsJSON: template.JS(elementsJSON), //nolint:gosec // JSON from json.Marshal is safe
		StyleJSON:    template.JS(styleJSON),    //nolint:gosec // JSON from json.Marshal is safe
		LayoutJSON:   template.JS(layoutJSON),   //nolint:gosec // JSON from json.Marshal is safe
		Options:      opts,
		NodeTypes:    nodeTypes,
		EdgeTypes:    edgeTypes,
		NodeCount:    g.Metadata.NodeCount,
		EdgeCount:    g.Metadata.EdgeCount,
		Groups:       g.Metadata.Groups,
	}

	// Parse and execute template
	tmplContent, err := templatesFS.ReadFile("templates/graph.html")
	if err != nil {
		return nil, fmt.Errorf("reading template: %w", err)
	}

	tmpl, err := template.New("graph").Parse(string(tmplContent))
	if err != nil {
		return nil, fmt.Errorf("parsing template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, fmt.Errorf("executing template: %w", err)
	}

	return buf.Bytes(), nil
}

// ToJSON exports the graph as JSON (Cytoscape elements format).
func (g *Graph) ToJSON() ([]byte, error) {
	return json.MarshalIndent(g.Elements, "", "  ")
}

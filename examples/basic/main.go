// Package main demonstrates basic usage of cytoscape-go.
package main

import (
	"log"
	"os"

	cytoscape "github.com/grokify/cytoscape-go"
)

func main() {
	// Create a new graph
	g := cytoscape.NewGraph()
	g.SetTitle("Example Call Graph")

	// Use code graph styling
	g.SetStyle(cytoscape.CodeGraphStyle())
	g.SetLayout(&cytoscape.DagreLayout{
		RankDir: "TB",
		NodeSep: 60,
		RankSep: 100,
	})

	// Add nodes
	g.AddNode(cytoscape.NodeWithType("main", "main()", "function"))
	g.AddNode(cytoscape.NodeWithType("init", "init()", "function"))
	g.AddNode(cytoscape.NodeWithType("server", "Server", "struct"))
	g.AddNode(cytoscape.NodeWithType("start", "Start()", "method"))
	g.AddNode(cytoscape.NodeWithType("stop", "Stop()", "method"))
	g.AddNode(cytoscape.NodeWithType("handler", "Handler", "interface"))

	// Add edges
	g.AddEdge(cytoscape.EdgeWithType("e1", "main", "init", "", "calls"))
	g.AddEdge(cytoscape.EdgeWithType("e2", "main", "server", "", "creates"))
	g.AddEdge(cytoscape.EdgeWithType("e3", "main", "start", "", "calls"))
	g.AddEdge(cytoscape.EdgeWithType("e4", "server", "handler", "", "implements"))
	g.AddEdge(cytoscape.EdgeWithType("e5", "start", "handler", "", "uses"))
	g.AddEdge(cytoscape.EdgeWithType("e6", "stop", "handler", "", "uses"))

	// Generate HTML
	html, err := g.ToHTML(cytoscape.CodeGraphHTMLOptions())
	if err != nil {
		log.Fatalf("Error generating HTML: %v", err)
	}

	// Write to file
	if err := os.WriteFile("graph.html", html, 0600); err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	log.Println("Generated graph.html")
}

package arbitrarynodelabels

import "fmt"

type ID int

// Graph only the ID (the label is important for this example)
type Graph struct {
	Nodes map[ID]struct{}
	Edges map[ID]map[ID]struct{}
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[ID]struct{}),
		Edges: make(map[ID]map[ID]struct{}),
	}
}

func (g *Graph) AddNode(id ID) bool {
	if _, ok := g.Nodes[id]; ok {
		return false
	}

	// Node labels need to be unique
	g.Nodes[id] = struct{}{}
	return true
}

func (g *Graph) AddEdge(from, to ID) {
	// Add the nodes if they don't exist
	_ = g.AddNode(from)
	_ = g.AddNode(to)

	// from -> to
	if _, ok := g.Edges[from]; !ok {
		g.Edges[from] = make(map[ID]struct{})
	}
	g.Edges[from][to] = struct{}{}

	// For undirected graphs add to -> from
	//if _, ok := g.Edges[to]; !ok {
	//	g.Edges[to] = make(map[ID]struct{})
	//}
	//g.Edges[to][from] = struct{}{}
}

func (g *Graph) printFromNode(id ID) {
	if _, ok := g.Edges[id]; !ok {
		fmt.Printf("\nNo edges from %d", id)
	}

	for edgeID := range g.Edges[id] {
		fmt.Printf("\nEdge %d -> %d", id, edgeID)
	}
}

func (g *Graph) print() {
	for from := range g.Nodes {
		g.printFromNode(from)
	}
}

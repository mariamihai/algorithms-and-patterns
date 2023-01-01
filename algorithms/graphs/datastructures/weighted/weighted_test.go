package weighted

import "fmt"

func ExamplePrint() {
	g := NewGraph(5)

	addEdge(g, 0, 4, 2)
	addEdge(g, 1, 2, 1)
	addEdge(g, 2, 3, 2)
	addEdge(g, 2, 4, 3)
	addEdge(g, 4, 2, 4)

	g.print()

	// Output:
	// Edge 0 -> 4 (2)
	// Edge 1 -> 2 (1)
	// Edge 2 -> 3 (2)
	// Edge 2 -> 4 (3)
	// No edges from 3
	// Edge 4 -> 2 (4)
}

func addEdge(g *Graph, from, to, weight int) {
	err := g.AddDirectedEdge(from, to, weight)

	if err != nil {
		fmt.Printf("Could not add the edge %d -> %d", from, to)
	}
}

package unweighted

import "fmt"

type Graph struct {
	NumNodes int
	Edges    [][]int
}

func NewGraph(n int) *Graph {
	return &Graph{
		NumNodes: n,
		Edges:    make([][]int, n),
	}
}

func (g *Graph) AddDirectedEdge(from, to int) error {
	if isInvalidNode(from, g.NumNodes) {
		return fmt.Errorf("invalid from node - %d", from)
	}

	if isInvalidNode(to, g.NumNodes) {
		return fmt.Errorf("invalid to node - %d", to)
	}

	// from -> to
	g.Edges[from] = append(g.Edges[from], to)

	return nil
}

func (g *Graph) AddUndirectedEdge(from, to int) error {
	if isInvalidNode(from, g.NumNodes) {
		return fmt.Errorf("invalid from node - %d", from)
	}

	if isInvalidNode(to, g.NumNodes) {
		return fmt.Errorf("invalid to node - %d", to)
	}

	// from -> to
	g.Edges[from] = append(g.Edges[from], to)

	// to -> from
	g.Edges[to] = append(g.Edges[to], from)

	return nil
}

func isInvalidNode(node, numNodes int) bool {
	return node < 0 || node >= numNodes
}

func (g *Graph) printFromNode(startingNodeKey int) error {
	if isInvalidNode(startingNodeKey, g.NumNodes) {
		return fmt.Errorf("invalid node - %d", startingNodeKey)
	}

	if len(g.Edges[startingNodeKey]) == 0 {
		fmt.Printf("\nNo edges from %d", startingNodeKey)
		return nil
	}

	fmt.Printf("\nEdges from %d:", startingNodeKey)
	for _, key := range g.Edges[startingNodeKey] {
		fmt.Printf(" %d", key)
	}

	return nil
}

func (g *Graph) print() {
	for index := range g.Edges {
		err := g.printFromNode(index)

		if err != nil {
			fmt.Printf("Error encountered for node %d - %v", index, err)
		}
	}
}

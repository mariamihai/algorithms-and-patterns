package weighted

import "fmt"

type Edge struct {
	From   int
	To     int
	Weight int
}

type Graph struct {
	NumNodes int
	Edges    [][]Edge
}

func NewGraph(n int) *Graph {
	return &Graph{
		NumNodes: n,
		Edges:    make([][]Edge, n),
	}
}

func (g *Graph) AddEdge(from, to, weight int) error {
	if isInvalidNode(from, g.NumNodes) {
		return fmt.Errorf("invalid from node - %d", from)
	}

	if isInvalidNode(to, g.NumNodes) {
		return fmt.Errorf("invalid to node - %d", to)
	}

	// from -> to
	g.Edges[from] = append(g.Edges[from], Edge{
		From:   from,
		To:     to,
		Weight: weight,
	})

	// For undirected graphs add to -> from
	//g.Edges[to] = append(g.Edges[to], Edge{
	//	From:   from,
	//	To:     to,
	//	Weight: weight,
	//})

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

	for _, edge := range g.Edges[startingNodeKey] {
		fmt.Printf("\nEdge %d -> %d (%d)", startingNodeKey, edge.To, edge.Weight)
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

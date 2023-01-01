package arbitrarynodelabels

func ExamplePrint() {
	g := NewGraph()
	g.AddNode(3)
	g.AddNode(5)
	g.AddNode(7)

	g.AddEdge(3, 5)
	g.AddEdge(3, 7)
	g.AddEdge(5, 7)

	g.print()

	// (Uses map who doesn't preserve order - example not run)
	// Output:
	// Edge 3 -> 7
	// Edge 3 -> 5
	// Edge 5 -> 7
	// No edges from 7
}

package unweighted

import (
	"fmt"
	"testing"
)

func ExamplePrint() {
	g := NewGraph(5)

	addEdge(g, 0, 4)
	addEdge(g, 1, 2)
	addEdge(g, 2, 3)
	addEdge(g, 2, 4)
	addEdge(g, 4, 2)

	g.print()

	// Output:
	// Edges from 0: 4
	// Edges from 1: 2
	// Edges from 2: 3 4
	// No edges from 3
	// Edges from 4: 2
}

func addEdge(g *Graph, from, to int) {
	err := g.AddDirectedEdge(from, to)

	if err != nil {
		fmt.Printf("Could not add the edge %d -> %d", from, to)
	}
}

func TestGraph_AddEdge(t *testing.T) {
	type args struct {
		from int
		to   int
	}
	tests := []struct {
		name    string
		fields  *Graph
		args    args
		wantErr bool
	}{
		{
			name:   "Valid new edge",
			fields: NewGraph(2),
			args: args{
				from: 0,
				to:   1,
			},
			wantErr: false,
		},
		{
			name:   "Invalid from node",
			fields: NewGraph(2),
			args: args{
				from: 2,
				to:   1,
			},
			wantErr: true,
		},
		{
			name:   "Invalid to node",
			fields: NewGraph(2),
			args: args{
				from: 0,
				to:   2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				NumNodes: tt.fields.NumNodes,
				Edges:    tt.fields.Edges,
			}
			if err := g.AddDirectedEdge(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("AddEdge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

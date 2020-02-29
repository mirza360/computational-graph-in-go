package main

import "fmt"

type graph struct {
	numNodes int
	edges    [][]edge
}

type edge struct {
	from   int
	to     int
	weight int
}

//NewGraph Create graph with n nodes.
func newGraph(n int) *graph {
	grph := &graph{numNodes: n, edges: make([][]edge, n)}
	return grph
}

//AddEdge Add an edge from u to v.
func (g *graph) addEdge(u, v, w int) {
	g.edges[u] = append(g.edges[u], edge{from: u, to: v, weight: w})

	// For undirected graph add edge from v to u.
	// g.Edges[v] = append(g.Edges[v], Edge{From: v, To: u, Weight: w})
}

func (g *graph) adjacentEdgesExample() {
	u := 1 // Example node label.

	fmt.Printf("Printing all edges adjacent to %d\n", u)
	for _, e := range g.edges[u] {
		fmt.Printf("Edge: %d -> %d (%d)\n", e.from, e.to, e.weight)
	}

	fmt.Println("Printing all edges in graph.")
	for _, adjacent := range g.edges {
		for _, e := range adjacent {
			fmt.Printf("Edge: %d -> %d (%d)\n", e.from, e.to, e.weight)
		}
	}
}

func main() {
	gr := newGraph(4)
	gr.addEdge(0, 1, 10)
	gr.addEdge(1, 2, 9)
	gr.addEdge(2, 3, 8)
	gr.addEdge(3, 4, 7)
	//gr.addEdge(4, 5, 6)

	gr.adjacentEdgesExample()

}

package main

func main() {
	g := &Graph2{vertices: make([]*Vertex, 0), edges: make([]*Edge2, 0)}
	for i := 1; i < 6; i++ {
		g.AddVertex(i)
	}

	g.AddWeightedEdge(1, 2, 2)
	g.AddWeightedEdge(1, 3, 1)
	g.AddWeightedEdge(1, 4, 3)
	g.AddWeightedEdge(2, 5, 3)
	g.AddWeightedEdge(3, 4, 2)
	g.AddWeightedEdge(4, 5, 1)

	// display(g.vertices)
	displayNeighbours(g.vertices)
	g.Dijkstra(1, 5)
}
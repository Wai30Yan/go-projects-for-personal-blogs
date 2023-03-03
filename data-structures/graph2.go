package main

import "fmt"

type Graph2 struct {
	vertices []*Vertex
	edges []*Edge2
}

type Vertex struct {
	value    int
	adjacent []*Vertex
}

type Edge2 struct {
	src, dst *Vertex
	weight int
}

func (g *Graph2) AddVertex(value int) {
	found := findVertex(g.vertices, value)
	if found {
		fmt.Printf("Vertex %d already exists.\n", value)
		return
	}
	vertex := &Vertex{value: value, adjacent: make([]*Vertex, 0)}
	g.vertices = append(g.vertices, vertex)
}

func (g *Graph2) AddWeightedEdge(src, dst int, w int) {
	f1 := findVertex(g.vertices, src)
	f2 := findVertex(g.vertices, dst)
	edgeFound, _ := findEdge(src, dst, g.edges)

	if !f1 || !f2 {
		fmt.Printf("Cannot find vertex %d nor %d in the graph.\n", src, dst)
		return
	}

	if edgeFound {
		fmt.Printf("Edge from %d to %d already exists.\n", src, dst)
		return
	}

	s := g.GetVertex(src)
	d := g.GetVertex(dst)

	s.adjacent = append(s.adjacent, d)
	edge := &Edge2{src: s, dst: d, weight: w}
	g.edges = append(g.edges, edge)
}

func (g *Graph2) GetVertex(value int) *Vertex {
	for _, v := range g.vertices {
		if v.value == value {
			return v
		}
	}
	return nil
}

func findVertex(adj []*Vertex, value int) bool {
	for _, v := range adj {
		if v.value == value {
			return true
		}
	}
	return false
}

func findEdge(src, dst int, edges []*Edge2) (bool, *Edge2) {
	for _, v := range edges {
		if v.src.value == src && v.dst.value == dst {
			return true, v
		}
	}
	return false, nil
}

func display(vertices []*Vertex) {
	for k := range vertices {
		fmt.Println("Vertex:", vertices[k].value)
	}
}

func displayNeighbours(neighbours []*Vertex) {
	for _, v := range neighbours {
		if len(v.adjacent) == 0 {
			continue
		}
		fmt.Print("adjacent of ", v.value, " are ")
		for _, vertex := range v.adjacent {
			fmt.Print(vertex.value, " ")
		}
		fmt.Println()
	}
}
package main

type Graph struct {
	adjacentList map[string][]string
	edges []Edge
}

type Edge struct {
	src, dst string
	weight int
}

func (g *Graph) AddVertex(v string) {
	list := g.adjacentList
	list[v] = make([]string, 0)
}

func (g *Graph) AddEdge(src, dst string) {
	list := g.adjacentList
	list[src] = append(list[src], dst)
}

func (g *Graph) FindEdge(src, dst string) bool {
	list := g.adjacentList
	if neighbours, ok := list[src]; ok {
		for _, v := range neighbours {
			if dst == v {
				return true
			}
		}
	}
	return false
}

func (g *Graph) AddEdgeWeight(src, dst string, w int) {
	list := g.adjacentList
	_, ok := list[src]
	if !ok {
		g.AddVertex(src)
	}

	_, ok = list[dst]
	if !ok {
		g.AddVertex(dst)
	}

	found := g.FindEdge(src, dst)

	if !found {
		g.AddEdge(src, dst)
	}

	edge := Edge{src: src, dst: dst, weight: w}
	g.edges = append(g.edges, edge)
}
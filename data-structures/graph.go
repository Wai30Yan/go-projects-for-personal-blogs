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

// func main() {
// 	g := Graph{
// 		adjacentList: make(map[string][]string), 
// 		edges: make([]Edge, 0),
// 	}
// 	v := "A"
// 	for i := 0; i < 6; i++ {
// 		fmt.Println("adding", v)
// 		g.AddVertex(v)
// 		v = string(rune(int(v[0]) + 1))
// 	}

// 	keys := []string{}

// 	for k := range g.adjacentList {
// 		keys = append(keys, k)
// 	}

// 	sort.Strings(keys)

// 	g.AddEdgeWeight("A", "B", 1)
// 	g.AddEdgeWeight("A", "D", 5)
// 	g.AddEdgeWeight("A", "E", 4)
// 	g.AddEdgeWeight("B", "F", 3)
// 	g.AddEdgeWeight("B", "C", 5)
// 	g.AddEdgeWeight("C", "B", 3)
// 	g.AddEdgeWeight("C", "D", 2)
// 	g.AddEdgeWeight("C", "A", 2)
// 	g.AddEdgeWeight("D", "A", 3)
// 	g.AddEdgeWeight("E", "B", 5)
// 	g.AddEdgeWeight("E", "C", 1)
// 	g.AddEdgeWeight("F", "C", 3)

// 	for _, v := range keys {
// 		fmt.Println(v, g.adjacentList[v])
// 	}

// 	for _, v := range g.edges {
// 		fmt.Println(v)
// 	}
// }
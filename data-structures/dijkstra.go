package main

import (
	"fmt"
	"math"
)

func (g *Graph2) Dijkstra(src, dst int) {
	visited := make([]bool, 0)
	dist := make(map[int]int)
	for _, v := range g.vertices {
		dist[v.value] = math.MaxInt
		visited = append(visited, false)
	}
	visited[src-1] = true
	fmt.Println(visited)
	fmt.Println(dist)

	for src != dst {
		total := 0
		cur := g.GetVertex(src)
		for _, v := range cur.adjacent {
			edge := g.GetEdge(cur.value, v.value)
			fmt.Println(edge.src.value, edge.dst.value, edge.weight)
			if dist[edge.dst.value] > edge.weight {
				dist[edge.dst.value] = edge.weight
				visited[edge.dst.value-1] = true
			}
		}
		nearestNode := getMinimun(dist)
		total += dist[nearestNode]
		if src == 4 {
			return
		}
		src = nearestNode
		fmt.Println("updated src", src)
	}
}

func getMinimun(m map[int]int) int {
	min := math.MaxInt
	for _, v := range m {
		if v < min {
			min = v
		}
	}
	for k := range m {
		if m[k] == min {
			// fmt.Println("minimun node", k, m[k])
			return k
		}
	}
	return 0
}
package main

import (
	"fmt"
	"math"
)

func main() {
	edges := []int{5,4,5,4,3,6,-1}

	fmt.Println(closestMeetingNode(edges,0,1))
}

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	dist1, dist2 := make([]int, len(edges)), make([]int, len(edges))
	minDistNode, minDistTillNow := -1, math.MaxInt
	for i := range dist1 {
		dist1[i] = math.MaxInt32
		dist2[i] = math.MaxInt32
	}

	bfs(node1, edges, dist1)
	bfs(node2, edges, dist2)

	fmt.Println(dist1)
	fmt.Println(dist2)

	for currNode := 0; currNode < len(edges); currNode++ {
		if minDistTillNow > max(dist1[currNode], dist2[currNode]) {
			fmt.Println(minDistTillNow, dist1[currNode], dist2[currNode])
			minDistNode = currNode
			minDistTillNow = max(dist1[currNode], dist2[currNode])
		}
	}

	return minDistNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bfs(startNode int, edges, dist []int) {
	queue := []int{startNode}
	visit := make([]bool, len(edges))
	dist[startNode] = 0
	for i := range visit {
		visit[i] = false
	}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		if visit[node] {
			continue
		}

		visit[node] = true
		neighbor := edges[node]

		if neighbor != -1 && !visit[neighbor] {
			dist[neighbor] = 1 + dist[node]
			queue = append(queue, neighbor)
		}
	}
}
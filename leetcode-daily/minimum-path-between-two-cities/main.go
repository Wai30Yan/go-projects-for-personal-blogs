package main

import "math"

func minScore(n int, roads [][]int) int {
	al := make([][][2]int, n+1)
	for _, road := range roads {
		al[road[0]] = append(al[road[0]], [2]int{road[1], road[2]})
		al[road[1]] = append(al[road[1]], [2]int{road[0], road[2]})
	}
	visited := make([]bool, n+1)
	visited[1] = true
	return dfs(al, 1, visited)
}

func dfs(al [][][2]int, city int, visited []bool) int {
	res := math.MaxInt32
	for _, next := range al[city] {
		nextCity, dist := next[0], next[1]
		// Capture the distance to this node. But we only want to visit the node
		// if we haven't already visited.
		res = min(res, dist)
		if !visited[nextCity] {
			visited[nextCity] = true
			res = min(res, dfs(al, nextCity, visited))
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
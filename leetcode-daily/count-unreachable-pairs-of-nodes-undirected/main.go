package main

func countPairs(n int, edges [][]int) int64 {
	adj := make([][]int, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}
	visit := make([]bool, n)
	numberOfPairs := int64(0)
	sizeOfComponent := int64(0)
	remainingNodes := int64(n)
	for i := 0; i < n; i++ {
		if !visit[i] {
			sizeOfComponent = 0
			bfs(adj, i, visit, &sizeOfComponent)
			numberOfPairs += sizeOfComponent * (remainingNodes - sizeOfComponent)
			remainingNodes -= sizeOfComponent
		}
	}
	return numberOfPairs
}

func bfs(adj [][]int, node int, visit []bool, sizeOfComponent *int64) {
	q := make([]int, 0)
	q = append(q, node)
	visit[node] = true
	*sizeOfComponent++
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		for _, v := range adj[node] {
			if !visit[v] {
				q = append(q, v)
				visit[v] = true
				*sizeOfComponent++
			}
		}
	}
}
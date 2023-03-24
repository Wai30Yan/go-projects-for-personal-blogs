package main

func minReorder(n int, connections [][]int) int {
	adList := make(map[int][]int)

	for _, con := range connections {
		adList[con[0]] = append(adList[con[0]], con[1])
		adList[con[1]] = append(adList[con[1]], -con[0])
	}

	count := 0
    visited := make([]bool, n)

	var dfs func(start int)
	dfs = func(start int) {
		visited[start] = true
		for _, nb := range adList[start] {
			if nb < 0 {
				nb = abs(nb)
				if !visited[nb] {
					dfs(nb)
				}
			} else {
				if !visited[nb] {
					count++
					dfs(nb)
				}
			}
		}
	}
	dfs(0)
	return count
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
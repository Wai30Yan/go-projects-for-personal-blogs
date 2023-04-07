package main

func numEnclaves(grid [][]int) int {
    m, n := len(grid), len(grid[0])
	count := 0

	for i := 0; i < m; i++ {
		if grid[i][0] == 1 {
			dfs(i, 0, m, n, grid)
		}
		if grid[i][n-1] == 1 {
			dfs(i, n-1, m, n, grid)
		}
	}
	for j := 1; j < n-1; j++ {
		if grid[0][j] == 1 {
			dfs(0, j, m, n, grid)
		}
		if grid[m-1][j] == 1 {
			dfs(m-1, j, m, n, grid)
		}
	}

	for i := range grid {
		for j := range grid[i] {
			count += grid[i][j]
		}
	}
	return count
}

func dfs(i, j, m, n int, grid [][]int) {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	dfs(i-1, j, m, n, grid)
	dfs(i+1, j, m, n, grid)
	dfs(i, j-1, m, n, grid)
	dfs(i, j+1, m, n, grid)
}
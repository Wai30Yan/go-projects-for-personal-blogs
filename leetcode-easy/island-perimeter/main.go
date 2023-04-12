package main

func islandPerimeter(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	peri := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				dfs(i,j,m,n, &peri, grid)
			}
		}
	}
    return peri
}

func dfs(x, y, m, n int, peri *int, grid [][]int) {
	if x < 0 || y < 0 || x >= m || y >= n || grid[x][y] == 0 {
		*peri += 1
		return
	} 

	if grid[x][y] == -1 {
		return
	}

	grid[x][y] = -1

	dirX := []int{0, 1, 0, -1}
	dirY := []int{1, 0, -1, 0}

	for i := 0; i < 4; i++ {
		r := x + dirX[i]
		c := y + dirY[i]

		dfs(r, c, m, n, peri, grid)
	}

}
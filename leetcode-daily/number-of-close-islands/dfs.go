package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0},
	}

	fmt.Println(closedIslandDFS(grid))

}

func closedIslandDFS(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	ans := 0

	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 && !visited[i][j] && dfs(i,j,m,n,grid,visited) {
				ans++
			}
		}
	}

	return ans
}

func dfs(x, y, m, n int, grid [][]int, visited [][]bool) bool {

	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return false
	}
	
	if grid[x][y] == 1 || visited[x][y] {
		return true
	}

	visited[x][y] = true
	isClosed := true

	dirX := []int{0,1,0,-1}
	dirY := []int{1,0,-1,0}

	for i := 0; i < 4; i++ {
		r := x + dirX[i]
		c := y + dirY[i]

		if !dfs(r, c, m, n, grid, visited) {
			isClosed = false
		}
	}

	return isClosed
}
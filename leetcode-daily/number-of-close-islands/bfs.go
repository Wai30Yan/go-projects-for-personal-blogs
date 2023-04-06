package main

// import "fmt"

// func main() {
// 	grid := [][]int{
// 		{1, 1, 1, 1, 1, 1, 1, 0},
// 		{1, 0, 0, 0, 0, 1, 1, 0},
// 		{1, 0, 1, 0, 1, 1, 1, 0},
// 		{1, 0, 0, 0, 0, 1, 0, 1},
// 		{1, 1, 1, 1, 1, 1, 1, 0},
// 	}

// 	fmt.Println(closedIslandBFS(grid))

// }

func closedIslandBFS(grid [][]int) int {
	// row, column
	m, n := len(grid), len(grid[0])
	ans := 0

	visited := make([][]bool, m)

	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 && !visited[i][j] && bfs(i, j, m, n, grid, visited) {
				ans++
			}
		}
	}
	return ans
}

type Pair struct {
	x, y int
}

func bfs(x, y, m, n int, grid [][]int, visited [][]bool) bool {
	p := Pair{x: x, y: y}
	queue := []Pair{p}
	visited[x][y] = true
	isCosed := true

	dirX := []int{0, 1, 0, -1}
	dirY := []int{1, 0, -1, 0}
	
	for len(queue) != 0 {
		temp := queue[0]
		x = temp.x
		y = temp.y
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			// r & c are row & col
			r := x + dirX[i]
			c := y + dirY[i]

			if r < 0 || r >= m || c < 0 || c >= n {
				// (x, y) is a boundary cell
				isCosed = false
			} else if grid[r][c] == 0 && !visited[r][c] {
				newPair := Pair{x: r, y: c}
				queue = append(queue, newPair)
				visited[r][c] = true
			}
		}
	}
	return isCosed
}
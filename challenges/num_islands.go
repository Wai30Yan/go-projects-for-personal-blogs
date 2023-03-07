package main

func numberOfIslands(grid [][]int) int {
	numIslands := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				numIslands += helperDFS(grid, i, j)
			}
		}
	}
	return numIslands;
}

/*
	this only show 4 islands
*/
func helperDFS(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid) || grid[i][j] == 0 {
		return 0
	} 

	grid[i][j] = 0

	helperDFS(grid, i + 1, j)
	helperDFS(grid, i - 1, j)
	helperDFS(grid, i, j + 1)
	helperDFS(grid, i, j - 1)

	return 1
}


// func main() {
// 	grid := [][]int{
// 		{1, 1, 0, 1, 0},
// 		{1, 1, 0, 0, 1},
// 		{0, 0, 1, 0, 1},
// 		{1, 1, 0, 0, 0},
// 	}

// 	fmt.Println("num of island", numberOfIslands(grid))
// }


/*
	this function shows 5 islands, correct number
*/
// func helperDFS(grid [][]int, i, j int) int {
//     if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 0 {
//         return 0
//     }

//     grid[i][j] = 0

//     helperDFS(grid, i + 1, j)
//     helperDFS(grid, i - 1, j)
//     helperDFS(grid, i, j + 1)
//     helperDFS(grid, i, j - 1)

//     return 1
// }

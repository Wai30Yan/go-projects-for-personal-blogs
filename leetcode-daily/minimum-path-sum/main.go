package main

import "math"

func minPathSum(grid [][]int) int {
	previousRow := make([]int, len(grid[0]))
	for i := range previousRow {
		previousRow[i] = math.MaxInt
	}
	previousRow[len(previousRow)-1] = 0

	for row := len(grid) - 1; row >= 0; row-- {
		currentRow := make([]int, len(grid[0]))

		for col := len(grid[0]) - 1; col >= 0; col-- {
			bottomCell := previousRow[col]
			var rightCell int
			if col == len(grid[0])-1 {
				rightCell = math.MaxInt
			} else {
				rightCell = currentRow[col+1]
			}

			if bottomCell > rightCell {
				currentRow[col] = rightCell + grid[row][col]
			} else {
				currentRow[col] = bottomCell + grid[row][col]
			}

		}
		previousRow = currentRow

	}
	return previousRow[0]
}
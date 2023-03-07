package main

import "fmt"

func knapsack(weight, profit []int, w, items int) {
	tb := make([][]int, 0)

	for i := 0; i <= items; i++ {
		// build the table column from 0 to max-weight (w)
		col := make([]int, w+1)
		tb = append(tb, col)
		for j := 0; j <= w; j++ {
			if i == 0 || j == 0 {
				tb[i][j] = 0
			} else if weight[i] <= j {
				tb[i][j] = max(profit[i]+tb[i-1][j-weight[i]], tb[i-1][j])
			} else {
				tb[i][j] = tb[i-1][j]
			}
		}
		fmt.Println(tb[i])
	}

	i, j := items, w

	for i > 0 && j > 0 {
		// fmt.Println("j is", j, "i is", i)
		if tb[i][j] == tb[i-1][j] {
			// don't buy this item
			fmt.Println(i, "= 0")
			i--
		} else {
			// buy this item
			fmt.Println(i, "= 1, weight[i] =", weight[i])
			j = j - weight[i]
			i--
		}
		fmt.Println("j =", j)
	}
	fmt.Println(i, "= 0")

}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

// func main() {
// 	maxWeight, numOfItems := 8, 4
// 	profit := []int{0,1,2,5,6}
// 	weight := []int{0,2,3,4,5}

// 	knapsack(weight, profit, maxWeight, numOfItems)
// }
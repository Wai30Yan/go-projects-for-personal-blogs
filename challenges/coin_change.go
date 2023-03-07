package main

import "fmt"

func coinChangeDP(amount int, coins []int) {
	combinations := make([]int, amount+1)
	/*  
		important step, as the iteration will start from
		index 1 cuz there is no coin zero,
		so, when there is coin one, it will take the value from indez zero
		which has value 1
	*/
	combinations[0] = 1
	for _, coin := range coins {
		for i := 1; i < len(combinations); i++ {
			// only iterate from the index based on coin value
			if i >= coin {
				combinations[i] += combinations[i - coin]
			}
		}
	}

	fmt.Println("combinations amount is", combinations[len(combinations)-1])
}

// func coinChangeRecur(amount int, coins []int) int {
// 	if amount == 0 {
// 		return 1
// 	}
// 	if amount < 0 {
// 		return 0
// 	}

// 	combos := 0
// 	for i := 0; i < len(coins); i++ {
// 		combos += coinChangeRecur(amount - coins[i])
// 	}
// }
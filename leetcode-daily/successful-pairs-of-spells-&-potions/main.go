package main

import (
	"math"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	m := len(potions) 
	maxPotion := potions[m-1]
	slice := make([]int, len(spells))
	for i, v := range spells {
		minPotion := math.Ceil(float64(success) / float64(v))
		if minPotion > float64(maxPotion) {
			slice[i] = 0
			continue
		}
		index := lowerBound(potions, int(minPotion))
		slice[i] = m - index

	}
	return slice
}

func lowerBound(arr []int, key int) int {
	lo, hi := 0, len(arr)

	for lo < hi {
		mid := lo + (hi-lo) / 2
		if arr[mid] < key {
			lo = mid+1
		} else {
			hi = mid-1
		}
	}

	return lo
}

// func successfulPairs(spells []int, potions []int, success int64) []int {
// 	sort.Ints(potions)
// 	var n, m int = len(spells), len(potions)
// 	res := make([]int, n)

// 	for i := 0; i < n; i++ {
// 		var left int = 0
// 		var right int = m

// 		for left < right {
// 			var mid int = left + (right-left)/2
// 			var currValue int64 = int64(potions[mid] * spells[i])
// 			if currValue >= success {
// 				right = mid
// 			} else {
// 				left = mid + 1
// 			}
// 		}
// 		res[i] = m - left
// 	}

// 	return res
// }
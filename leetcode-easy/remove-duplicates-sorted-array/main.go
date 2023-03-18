package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 3, 3, 3, 3, 4, 4, 5}
	//  output: [1, 2, 3, 4, 5, _, _, _, _, _]
	fmt.Println(removeDuplicates(arr))
}

func removeDuplicates(nums []int) int {
	minIdx := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[minIdx] = nums[i]
			minIdx++
		}
		fmt.Println(nums)
	}
	return minIdx
}

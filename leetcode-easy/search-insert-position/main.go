package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}

	fmt.Println(searchInsert(nums, 5))
}

func searchInsert(nums []int, target int) int {
	// right has to be len(nums) to cope with the missing target
	left, right := 0, len(nums)
	for left < right {
		fmt.Println(left, right)
		mid := left + (right-left) / 2
		if nums[mid] >= target {
			right = mid 
		} else {
			left = mid + 1
		}
	}
	return left
}


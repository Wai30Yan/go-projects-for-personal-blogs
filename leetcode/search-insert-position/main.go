package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}

	fmt.Println(searchInsert(nums, 5))
}

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		}
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] < target && nums[i+1] > target {
			return i
		}
	}
	return 0
}
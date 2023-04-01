package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	fmt.Println(search(arr, 6))
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	
	for l <= r {
		mid := l + (r-l) / 2
		if nums[mid] > target {
			r = mid - 1
		}
		if nums[mid] < target {
			l = mid + 1
		}
		if nums[mid] == target {
			return mid
		}
	}

	return -1
}
package main

import "fmt"

func binarySearch(arr []int, num int) bool {
	fmt.Println("arr", arr)
	if len(arr) == 2 {
		return arr[0] == num || arr[1] == num
	}

	mid := (len(arr) - 1) / 2
	fmt.Println("mid", arr[mid], "index", mid)
	if arr[mid] == num {
		return true
	}

	if arr[mid] < num {
		arr = arr[mid:]
	} 
	if arr[mid] > num {
		arr = arr[:mid]
	}
	return binarySearch(arr, num)
}

func binarySearch2(arr []int, num int) bool {
	if arr[0] == num {
		return true
	}

	left, right := 0, len(arr) - 1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == num {
			return true
		} else if arr[mid] < num {
			left = mid+1
		} else {
			right = mid - 1
		}
	}
	return false
}
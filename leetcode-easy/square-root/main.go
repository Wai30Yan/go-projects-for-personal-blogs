package main

import "fmt"

func main() {
	fmt.Println(mySqrt(16))
	fmt.Println(mySqrt(25))
	fmt.Println(mySqrt(27))
}

// binary search
func mySqrt(x int) int {
	low, high := 0, x

	for low <= high {
		mid := (low + high) / 2
		if mid*mid == x {
			return mid
		}
		if mid*mid < x {
			low = mid + 1
		}
		if mid*mid > x {
			high = mid - 1
		}
	}
	return high
}
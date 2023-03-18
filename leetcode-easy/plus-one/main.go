package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1,2,3}))
	fmt.Println(plusOne([]int{4,3,2,1}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{8,9,9,9}))
	fmt.Println(plusOne([]int{4,3,9,9}))

}

func plusOne(digits []int) []int {
	n := len(digits)-1

	if n == 0 && digits[n] < 9 {
		digits[n]++
		return digits
	}

	for digits[n] == 9 && n > 0 {
		digits[n] = 0
		n--
	}

	if n == 0 && digits[n] == 9 {
		digits[n] = 0
		digits = append([]int{1}, digits...)
	} else {
		digits[n]++
	}
	return digits
}

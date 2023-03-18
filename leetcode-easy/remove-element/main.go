package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 3, 3, 3, 3, 4, 4, 5}
	//  output: [1, 2, 3, 4, 5, _, _, _, _, _]
	fmt.Println(removeElement(arr, 3))
}

func removeElement(nums []int, val int) int {
    idx := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[idx] = nums[i]
			idx++
        } 

		fmt.Println(nums)
    }
	return idx
}
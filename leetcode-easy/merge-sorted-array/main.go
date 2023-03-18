package main

import (
	"fmt"
)

func main() {
	a1 := []int{0,0,0,0,0}
	a2 := []int{1,2,3,4,5}
	merge(a1,a2, 0, 5)
	fmt.Println(a1)
}

func merge(nums1, nums2 []int, m, n int) {
	i, j, k := m-1, n-1, m+n-1

	// iterate from right to insert at the end
	for ; j >= 0; k-- {
		if i >= 0 && nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}

// for i <= n-1 {
// 	if nums1[i] >= nums2[j] {
// 		nums1[i], nums2[j] = nums2[j], nums1[i]
// 	} 
// 	if nums1[i] == 0 {
// 		nums1[i], nums2[j] = nums2[j], nums1[i]
// 		j++
// 	}
// 	i++

// }
// j = 0
// for k := n; k < len(nums1); k++ {
// 	if nums1[k] == 0 {
// 		nums1[k] = nums2[j]
// 		j++
// 	}
// }
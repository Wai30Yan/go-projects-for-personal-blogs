package main

import "math"

func minimizeArrayValue(nums []int) int {
	answer, prefixSum := 0, 0

	for i := range nums {
		prefixSum += nums[i]		
		answer = int(math.Max(float64(answer), float64((prefixSum+i)/(i+1))))
	}
	
	return answer
}
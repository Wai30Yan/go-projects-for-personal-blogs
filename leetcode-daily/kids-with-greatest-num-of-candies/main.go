package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := []bool{}
	max := max(candies)
	
	for _, v := range candies {
		if v + extraCandies >= max {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}

func max(arr []int) int {

	for _, v := range arr {
		if arr[0] < v {
			arr[0] = v
		}
	}
	return arr[0]
}
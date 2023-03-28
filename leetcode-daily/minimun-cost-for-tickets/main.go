package main

func mincostTickets(days []int, costs []int) int {
	var dp [396]int
	for _, d := range days {
		dp[d] = 1
	}
	for d := 365; d > 0; d-- {
		if dp[d] == 0 {
			dp[d] = dp[d+1]
			continue
		}
		dp[d] = min(
			dp[d+1]+costs[0],
			dp[d+7]+costs[1],
			dp[d+30]+costs[2])
	}
	return dp[1]
}

func min(nums ...int) int {
	for _, n := range nums {
		if n < nums[0] {
			nums[0] = n
		}
	}
	return nums[0]
}
package main

import "fmt"

func main() {
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(6))
	fmt.Println(climbStairs(45))
}

/**
	You are climbing a staircase. It takes n steps to reach the top.

	Each time you can either climb 1 or 2 steps. In how many distinct ways can you
	climb to the top?

	Input: n = 2
	Output: 2
	Explanation: There are two ways to climb to the top.
	1. 1 step + 1 step
	2. 2 steps
**/

var m = make(map[int]int)

func climbStairs(n int) int {
	if n <= 2 {
		m[n] = n
		return n
	}

	if _, ok := m[n]; !ok {
		ans := climbStairs(n-1) + climbStairs(n-2)
		m[n] = ans
	}

	return m[n]
}
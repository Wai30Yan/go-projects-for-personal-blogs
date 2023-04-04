package main

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	i, j, ans := 0, len(people) - 1, 0

	for i <= j {
		ans++
		if people[i] + people[j] <= limit {
			i++
		} 
		j--
		
	}

	return ans
}
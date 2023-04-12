package main

import "sort"

func findContentChildrenSort(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	content, i, j := 0, 0, 0

	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			content++
			i++
			j++
		} else {
			j++
		}
	}
	return content
}

func findContentChildrenIterate(g []int, s []int) int {
	content := 0
	m := make(map[int]int)
	set := map[int]bool{}

	for i := range s {
		m[s[i]] = s[i]
	}

	for _, v := range g {
		if _, ok := m[v]; ok && !set[v] {
			set[v] = true
			content++
		}
	}

	return content
}
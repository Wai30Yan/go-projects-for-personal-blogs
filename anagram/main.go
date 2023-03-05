package main

import "fmt"

func main() {
	s1 := "players"
	s2 := "parsley"

	fmt.Println(anagram(s1, s2))
}

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	sum1, sum2 := 0, 0
	for k, v := range s1 {
		sum1 += int(v)
		sum2 += int(s2[k])
	}

	return sum1 == sum2
}
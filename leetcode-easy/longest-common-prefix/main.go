package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"flower", "flower", "flower", "flower"}
	strs1 := []string{"dog", "racecar", "car"}

	fmt.Println(longestCommonPrefix(strs))
	fmt.Println(longestCommonPrefix(strs1))
}

func longestCommonPrefix(strs []string) string {
	first := strs[0]
	result := ""
	for _, l := range first {
		result += string(l)
		for d, each := range strs[1:] {
			if strings.HasPrefix(each, result) {
				d++
			} else {
				return result[:len(result)-1]
			}
		}

	}
	return result
}
package main

import (
	"fmt"
	"strings"
)

func main() {
	haystack := "leetcode"
	needle := "leeto"

	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		s := haystack[i:]
		if strings.HasPrefix(s, needle) {
			return i
		}
	}
	return -1
}
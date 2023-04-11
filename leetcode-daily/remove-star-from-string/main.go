package main

import "fmt"

func main() {
	s := removeStars("abc**d")
	fmt.Println(s)
}

func removeStars(s string) string {
	ans := ""
	for i := range s {
		r := rune(s[i])
		if r == '*' {
			ans = ans[:len(ans)-1]
		} else {
			ans += string(r)
		}
	}
	return ans
}

func twoPointers(s string) string {
	ch := make([]byte, len(s))
	j := 0

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '*' {
			j--
		} else {
			ch[j] = c
			j++
		}
	}

	answer := string(ch[:j])

	return answer
}
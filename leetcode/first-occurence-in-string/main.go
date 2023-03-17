package main

import (
	"fmt"
	"strings"
)

func main() {
	haystack := "leetcode"
	needle := "leeto"

	fmt.Println(strStr(haystack, needle))

	m := makeHashMap()
	fmt.Println(m)
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

/**
 * Multiply the hash value of the previous m-substring by radix 26.
 * Subtract the value of the first character of the previous m-substring, 
   since we are moving the window by one character, and that character is now 
   out of the window.
 * Add the value of the last character of the current m-substring, since we are moving the window by one character
**/
func RabinKarpHash(haystack, needle string) int {
	if len(haystack) < len(needle){
		return -1
	}
	return -1
}

func hashValue(str string, RADIX int, MOD int, m int) int {
    ans := 0
    factor := 1
    for i := m - 1; i >= 0; i-- {
        ans += ((int(str[i]) - int('a')) * factor) % MOD
        factor = (factor * RADIX) % MOD
    }
    return ans % MOD
}

func makeHashMap() map[string]int {
	m := make(map[string]int)

	c := 'a'
	
	for i := 0; i < 26; i++ {
		m[string(c)] = i
		c = c + 1
	}

	return m
}
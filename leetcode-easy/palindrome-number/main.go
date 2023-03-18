package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
}

func isPalindrome(x int) bool {
	s := fmt.Sprint(x)
	reverse := ""

	for i := len(s)-1; i >= 0; i-- {
		reverse += string(s[i])
	}

	fmt.Println(s)

	return s == reverse
}
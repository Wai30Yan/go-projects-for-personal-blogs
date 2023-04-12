package main

import "strings"

func licenseKeyFormatting(s string, k int) string {
	count, n := 0, len(s)
	ans := ""

	for i := n-1; i >= 0; i-- {
		r := string(s[i])
		if r != "-" {
			ans += strings.ToUpper(r)
			count++
			if count == k {
				ans += "-"
				count = 0
			}
		}
	}

	if len(ans) > 0 && ans[len(ans)-1] == '-' {
		ans = ans[:len(ans)-1]
	}


	return reverseString(ans)
}

func reverseString(s string) string {
	// Convert the string to a slice of bytes
	bytes := []byte(s)
	// Get the length of the slice
	length := len(bytes)
	// Swap elements from both ends of the slice
	for i := 0; i < length/2; i++ {
		bytes[i], bytes[length-1-i] = bytes[length-1-i], bytes[i]
	}
	// Convert the slice back to a string
	return string(bytes)
}
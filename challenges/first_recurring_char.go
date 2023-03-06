package main

// O(n^2)
func firstRecurringCharNaive(str string) string {

	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str); j++ {
			if i != j && str[i] == str[j] {
				return string(str[i])
			}
		}
	}
	return ""
}

var m map[string]bool

// linear O(n)
func firstRecurringChar(str string) string {
	m = make(map[string]bool)
	for _, v := range str {
		if !m[string(v)] {
			m[string(v)] = true
		} else {
			return string(v)
		}
	}
	return ""
}
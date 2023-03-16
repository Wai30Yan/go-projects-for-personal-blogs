package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(isValid("(){}[]"))
	// fmt.Println(isValid("(){}["))
	// fmt.Println(isValid("){}["))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	m := make(map[rune]rune, 0)
	m['{'] = '}'
	m['['] = ']'
	m['('] = ')'

	r := []rune(s)
	fmt.Println(r)
	if _, ok := m[r[0]]; !ok {
		return false
	}
	queue := list.New()
	count := 0
	for i := 0; i < len(s); i++ {
		if _, ok := m[r[i]]; ok {
			queue.PushBack(m[r[i]])
		} else if queue.Len() != 0 {
			back := queue.Back()
			fmt.Println(back.Value, string(r[i]))
			if back.Value != r[i] {
				return false
			}
			queue.Remove(back)
			count++
		}
	}
	return count == len(s)/2
}

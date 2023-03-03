package main

import "fmt"

type Stack []string

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() {
	idx := len(*s) - 1
	element := (*s)[idx]
	fmt.Println(element)
	*s = (*s)[:idx]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
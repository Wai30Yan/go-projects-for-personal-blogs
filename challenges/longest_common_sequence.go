package main

import "fmt"

func longestCommonSequence(s1, s2 string) int {
	row := len(s2) 
	col := len(s1) 
	m := make([][]int, row)

	fmt.Println(len(m))
	for i := 0; i < row; i++ {
		arr := make([]int, col)
		m[i] = append(m[i], arr...)
		for j := 0; j < col; j++ {
			if s2[i] == s1[j] {
				m[i][j] = 1
			} else {
				m[i][j] = 0
			}
		}
	}

	sum := 0
	for _, v := range m {
		for _, value := range v {
			if  value == 1 {
				sum += 1
			}
		}
		fmt.Println(v)
	}
	return sum
}
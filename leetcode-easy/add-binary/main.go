package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println(addBinary("10", "01"))
	// fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("100", "110010"))
}

func addBinary(a, b string) string {
	s := ""
	carr := 0
	m, n := len(a)-1, len(b)-1
	for m >= 0 && n >= 0 {
		i, _ := strconv.Atoi(string(a[m]))
		j, _ := strconv.Atoi(string(b[n]))
		str, carry := binary(i, j, carr)
		carr = carry
		s = str + s
		m--
		n--
	}
	
	for m >= 0 {
		i, _ := strconv.Atoi(string(a[m]))
		if carr == 1 {
			strconv.Itoa(carr)
			str, carry := binary(i, 0, carr)
			s = str + s
			carr = carry
		} else {
			s = string(a[m]) + s
		}
		m--
	}
	for n >= 0 {
		j, _ := strconv.Atoi(string(b[n]))
		if carr == 1 {
			strconv.Itoa(carr)
			str, carry := binary(0, j, carr)
			s = str + s
			carr = carry
		} else {
			s = string(b[n]) + s
		}
		n--
	}
	if carr == 1 {
		s = strconv.Itoa(carr) + s
	}
	return s
}

func binary(i, j, carr int) (string, int) {
	s := ""
	bi := i + j + carr
	carr = 0
	if bi <= 1 {
		s = fmt.Sprint(bi) + s
	}
	if bi == 2 {
		s = "0" + s
		carr = 1
	}
	if bi == 3 {
		s = "1" + s
		carr = 1
	}

	return s, carr
}

// func remainingString(str string, carr int) (string, int) {
// 	s := ""
// 	i, _ := strconv.Atoi(str)
// 	if carr == 1 {
// 		strconv.Itoa(carr)
// 		str, carry := binary(i, 0, carr)
// 		s = str + s
// 		carr = carry
// 	} else {
// 		s = str + s
// 	}

// 	return s, carr
// }
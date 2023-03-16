package main

import "fmt"

func main() {
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	var v, lv, cv int
    m := map[uint8]int{
        'I': 1,
		'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    for i := len(s)-1; i >= 0; i-- {
		cv = m[s[i]]
		if cv < lv {
			v -= cv
		} else {
			v += cv
		}
		lv = cv
    }
    return v
}
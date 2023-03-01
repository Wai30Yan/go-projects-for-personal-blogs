package main

import (
	"fmt"
	"sort"
)

// has to assign to an empty value with make()
// or it will panic assignment to entry in nil map
var m map[uint64]uint64 = make(map[uint64]uint64)

func main() {
	fmt.Println("Enter a number: ")
	var n uint64
	fmt.Scanln(&n)
	var i uint64
	for i = 0; i <= n; i++ {
		fmt.Println("fib of", i, ":", fibMemo(i))
	}

	// making an ordered map by creating and sorting a slice
	keys := make([]uint64, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for k := range keys {
		fmt.Println(k,":",m[uint64(k)])
	}

	// fmt.Println(m)
}

func fibMemo(n uint64) uint64 {
	
	if n < 2 {
		return n
	}
	if _, ok := m[n]; !ok {
		ans := fibMemo(n-1) + fibMemo(n-2)
		m[n] = ans
	} 

	return m[n]
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-1)
}


package main

// has to assign to an empty value with make()
// or it will panic assignment to entry in nil map
var fibM map[uint64]uint64

// Naive Approach - Recursive
func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-1)
}

// Dynamic Programming with Memoization
// linear O(n)
func fibMemo(n uint64) uint64 {
	fibM = make(map[uint64]uint64)
	if n < 2 {
		return n
	}
	if _, ok := fibM[n]; !ok {
		ans := fibMemo(n-1) + fibMemo(n-2)
		fibM[n] = ans
	} 

	return fibM[n]
}


// func main() {
// 	fmt.Println("Enter a number: ")
// 	var n uint64
// 	fmt.Scanln(&n)
// 	var i uint64
// 	for i = 0; i <= n; i++ {
// 		fmt.Println("fib of", i, ":", fibMemo(i))
// 	}

// 	// making an ordered map by creating and sorting a slice
// 	keys := make([]uint64, 0, len(m))

// 	for k := range m {
// 		keys = append(keys, k)
// 	}

// 	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

// 	for k := range keys {
// 		fmt.Println(k,":",m[uint64(k)])
// 	}

// 	// fmt.Println(m)
// }

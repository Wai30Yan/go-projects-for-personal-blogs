package main

func isHappy(n int) bool {
	seen := make(map[int]bool)

	for n != 1 && !seen[n] {
		seen[n] = true
		n = getNext(n)
	}

	return n == 1
}

func getNext(n int) int {
	totalSum := 0
	for n > 0 {
		d := n % 10
		n = n / 10
		totalSum += d*d
	}
	return totalSum
}
package main

func longestPalindromeSubseq(s string) int {
	n := len(s)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
	}

	return lps(s, 0, n-1, memo)    
}

func lps(s string, i, j int, memo [][]int) int {
	if memo[i][j] != 0 {
		return memo[i][j]
	}
	if i > j {
		return 0
	}
	if i == j {
		return 1
	}

	if s[i] == s[j] {
		memo[i][j] = lps(s, i+1, j-1, memo) + 2
	} else {
		memo[i][j] = max(lps(s, i+1, j, memo), lps(s, i, j-1, memo))
	}
	return memo[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
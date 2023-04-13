package main

func validateStackSequences(pushed []int, popped []int) bool {
    N := len(pushed)
	stack := []int{}

	j := 0

	for _, v := range pushed {
		stack = append(stack, v)

		for len(stack) > 0 && j < N && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}

	return j == N
}
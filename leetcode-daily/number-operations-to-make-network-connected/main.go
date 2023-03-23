package main

func makeConnected(n int, connections [][]int) int {
	set := n - 1
	if len(connections) < n-1 {
		return -1
	}

	u := make([]int, n, n)

	find := func(e int) int {
		for u[e] > 0 {
			e = u[e]
		}
		return e
	}

	union := func(e1, e2 int) {
		pe1 := find(e1)
		pe2 := find(e2)
		if pe1 != pe2 {
			if u[pe1] < u[pe2] {
				u[pe2] = pe1
			} else {
				u[pe1] = pe2
			}
			set--
		}
	}

	for i := 0; i < len(connections); i++ {
		union(connections[i][0], connections[i][1])
	}
	return set
}
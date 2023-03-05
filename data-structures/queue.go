package main

type Queue []int

func (q *Queue) Enqueue(v int) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() {
	// element := (*q)[0]
	// fmt.Println(element)
	*q = (*q)[1:]
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Find(n int) bool {
	for _, v := range *q {
		if v == n {
			return true
		}
	}
	return false
}

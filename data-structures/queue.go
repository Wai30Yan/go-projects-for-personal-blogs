package main

import "fmt"

type Queue []string

func (q *Queue) Enqueue(v string) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() {
	element := (*q)[0]
	fmt.Println(element)
	*q = (*q)[1:]
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}




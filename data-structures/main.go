package main

import "fmt"

func main() {
	l := &LinkedList{}

	for i := 1; i < 6; i++ {
		node := &Node{value: i}
		l.InsertFromTail(node)
	}

	l.Display()
	fmt.Println("reversing...")
	l.ReverseList()

	l.Display()

}
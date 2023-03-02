package main

import "fmt"

func main() {
	l := &LinkedList{head: nil}
	for i := 1; i < 6; i++ {
		node := &Node{value: i}
		l.InsertFromTail(node)
	}
	l.Display()
	fmt.Println(l.head)
	fmt.Println(l.Search(l.head, 4))
}
package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) InsertFromHead(node *Node) {
	if l.head != nil {
		node.next = l.head
		l.head = node
	} else {
		l.head = node
	}
}

func (l *LinkedList) InsertFromTail(node *Node) {
	if l.head != nil {
		c := l.head
		for c.next != nil {
			c = c.next
		}
		c.next = node	
	} else {
		l.head = node
	}
}

func (l *LinkedList) Search(head *Node, value int) bool {
	if head == nil {
		return false
	}
	
	if head.value == value {
		fmt.Println("found", head)
		return true
	} else {
		return l.Search(head.next, value)
	}
}

func (l *LinkedList) DeleteNode(v int) {
	node := l.head

	for node != nil {
		if node.next.value == v {
			deleteNode := node.next
			node.next = deleteNode.next
			fmt.Println("Deleting", deleteNode)
			l.Display()
			return
		}
		node = node.next
	}
}

func (l *LinkedList) Display() {
	node := l.head
	if node == nil {
		return
	}

	for node != nil {
		fmt.Println("display", node.value, "\t", node.next)
		node = node.next
	}
}

// func main() {
// 	l := &LinkedList{}
// 	for i := 1; i < 6; i++ {
// 		node := &Node{value: i}
// 		l.InsertFromTail(node)
// 	}
// 	for i := 7; i < 15; i++ {
// 		node := &Node{value: i}
// 		l.InsertFromHead(node)
// 	}
// 	l.Display()
// 	fmt.Println(l.Search(l.head, 12))
// 	l.DeleteNode(12)
// }
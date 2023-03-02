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

func (l *LinkedList) Display() {
	if l.head == nil {
		return
	}

	for l.head != nil {
		fmt.Println("display", l.head.value, l.head.next)
		l.head = l.head.next
	}
}
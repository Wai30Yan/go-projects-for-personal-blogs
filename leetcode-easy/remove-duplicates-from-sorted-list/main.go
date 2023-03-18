package main

func main() {

}


type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	node := head
	prev := &ListNode{}
	listSet := make(map[int]bool)
	for node != nil {
		if !listSet[node.Val] {
			listSet[node.Val] = true
			prev = node
			node = node.Next
		} else {
			node = node.Next
			prev.Next = node
		}
	}
	return head
}
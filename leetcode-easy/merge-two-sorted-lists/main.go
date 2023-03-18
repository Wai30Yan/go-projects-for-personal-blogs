package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	l1, l2 := list1, list2
    res := &ListNode{}
    cur := res

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
            cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
            l2 = l2.Next
		}
        cur = cur.Next
	}
    
    if l1 == nil {
        cur.Next = l2
    } else {
        cur.Next = l1
    }

	return res.Next
}
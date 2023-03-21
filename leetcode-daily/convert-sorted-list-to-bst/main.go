package main

func main() {

}


type ListNode struct {
    Val int
    Next *ListNode
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
    listNode := head
    s := make([]int, 0)
    for listNode != nil {
        s = append(s, listNode.Val)
        listNode = listNode.Next
    }

	root := buildBinaryTree(s)
    return root
}

func buildBinaryTree(s []int) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	
	mid := len(s) / 2
	root := &TreeNode{Val: s[mid]}

	if mid != 0 {
		root.Left = buildBinaryTree(s[:mid])
	}
	if mid != len(s)-1 {
		root.Right = buildBinaryTree(s[mid+1:])
	}

	return root
}
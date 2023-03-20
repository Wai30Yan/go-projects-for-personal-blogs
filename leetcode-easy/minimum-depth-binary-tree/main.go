package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	return minimum(minDepth(root.Left), minDepth(root.Right)) + 1
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}
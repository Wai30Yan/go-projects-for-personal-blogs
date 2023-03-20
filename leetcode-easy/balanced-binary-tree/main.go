package main

import "math"

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)

	if math.Abs(float64(leftHeight)- float64(rightHeight)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(height(node.Left), height(node.Right))
}
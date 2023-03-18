package main

func main() {

}

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var s = make([]int, 0)

	if root == nil {
		return s
	}

	s = append(s, inorderTraversal(root.Left)...)
	s = append(s, root.Val)
	s = append(s, inorderTraversal(root.Right)...)

	return s
}
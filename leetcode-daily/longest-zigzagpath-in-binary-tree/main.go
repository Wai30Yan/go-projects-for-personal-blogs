package main

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	return max(dfs(root.Left, true, 0), dfs(root.Right, false, 0))
}

func dfs(node *TreeNode, isLeft bool, count int) int {
	if node == nil {
		return count
	}

	if isLeft {
		return max(dfs(node.Left, true, 0), dfs(node.Right, false, count+1))
	}

	return max(dfs(node.Left, true, count+1), dfs(node.Right, false, 0))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
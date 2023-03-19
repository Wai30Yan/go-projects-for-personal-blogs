package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	i := len(nums) / 2
	root := &TreeNode{Val: nums[i]}
	if i != 0 {
		root.Left = sortedArrayToBST(nums[:i])
	}
	if i != len(nums)-1 {
		root.Right = sortedArrayToBST(nums[i+1:])
	}
	return root
}
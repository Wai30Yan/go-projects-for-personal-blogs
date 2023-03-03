package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	value               int
	parent, left, right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func (t *Tree) BuildTree(root, node *TreeNode) {
	if t.root != nil {
		if node.value >= root.value {
			// go right
			if root.right != nil {
				t.BuildTree(root.right, node)
			} else {
				root.right = node
				node.parent = root
			}
		} else {
			// go left
			if root.left != nil {
				t.BuildTree(root.left, node)
			} else {
				root.left = node
				node.parent = root
			}
		}
	} else {
		t.root = node
	}
}

func (t *Tree) PreOrder(root *TreeNode) {
	fmt.Println(root)
	if root.left != nil {
		t.PreOrder(root.left)
	}
	if root.right != nil {
		t.PreOrder(root.right)
	}
}

func (t *Tree) InOrder(root *TreeNode) {
	if root == nil {
		return
	}
	t.InOrder(root.left)
	fmt.Println(root)
	t.InOrder(root.right)
}

func (t *Tree) PostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	t.PostOrder(root.left)
	t.PostOrder(root.right)
	fmt.Println(root)
}

func (t *Tree) Search(root *TreeNode, v int) (bool, *TreeNode) {
	if root == nil {
		return false, nil
	}

	if root.value == v {
		return true, root
	}

	if root.value < v {
		return t.Search(root.right, v)
	} else {
		return t.Search(root.left, v)
	}
}

func (t *Tree) DeleteNode(root, node *TreeNode) {
	if node == nil {
		return 
	}

	if root.value == node.value {
		if root.right == nil && root.left == nil {
			root = nil
			return
		}
		p := root.parent
		if root == p.right {
			fmt.Println(root, "is right child of", p)
			return
		}
		if root == p.left {
			fmt.Println(root, "is left child of", p)
			return
		}
	}

	if root.value < node.value {
		t.DeleteNode(root.right, node)
	} else {
		t.DeleteNode(root.left, node)
	}
}

func printTreeHoriz(node *TreeNode, level int) {
    if node == nil {
        return
    }
    printTreeHoriz(node.right, level+1)
    fmt.Printf("%s%d\n", strings.Repeat("\t", level), node.value)
    printTreeHoriz(node.left, level+1)
}

// func main() {
// 	t := Tree{}
// 	arr := []int{2, 4, 18, 3, 3, 1, 16, 4, 16, 1}
// 	fmt.Println(arr)
// 	for i := 0; i < len(arr); i++ {
// 		n := &TreeNode{value: arr[i]}
// 		t.BuildTree(t.root, n)
// 	}

// 	printTreeHoriz(t.root, 0)
// 	found, node := t.Search(t.root, 1)
// 	if found {
// 		t.DeleteNode(t.root, node)
// 		printTreeHoriz(t.root, 0)
// 	}
// }
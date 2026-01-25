package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	minValue := math.MinInt
	var inorder func(*TreeNode) bool
	inorder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !inorder(node.Left) {
			return false
		}
		if node.Val <= minValue {
			return false
		}
		minValue = node.Val
		return inorder(node.Right)
	}
	return inorder(root)
}

func main() {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(isValidBST(root))
}

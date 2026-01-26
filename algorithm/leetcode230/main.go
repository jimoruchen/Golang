package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var result []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return result[k-1]
}

func kthSmallest1(root *TreeNode, k int) (ans int) {
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		k--
		if k == 0 {
			ans = node.Val
		}
		inorder(node.Right)
	}
	inorder(root)
	return
}

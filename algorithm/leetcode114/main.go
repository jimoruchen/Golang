package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var tmp []*TreeNode
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		tmp = append(tmp, node)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	current := root
	for i := 1; i < len(tmp); i++ {
		current.Left = nil
		current.Right = tmp[i]
		current = tmp[i]
	}
}

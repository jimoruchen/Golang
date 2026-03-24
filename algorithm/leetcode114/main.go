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

func flatten1(root *TreeNode) {
	if root == nil {
		return
	}
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
	cur := root
	for i := 1; i < len(tmp); i++ {
		cur.Right = tmp[i]
		cur.Left = nil
		cur = cur.Right
	}
}

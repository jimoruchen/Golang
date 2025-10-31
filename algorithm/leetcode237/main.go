package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func findNode(head *ListNode, val int) *ListNode {
	for head != nil {
		if head.Val == val {
			return head
		} else {
			head = head.Next
		}
	}
	return nil
}

func CreateLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i], Next: nil}
		cur = cur.Next
	}
	return head
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	var nums = []int{4, 5, 1, 9}
	head := CreateLinkedList(nums)
	nodeToDelete := findNode(head, 1)
	if nodeToDelete != nil && nodeToDelete.Next != nil {
		deleteNode(nodeToDelete)
	}
	PrintList(head)
}

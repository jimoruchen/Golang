package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedList(nums []int) *ListNode {
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head
}

func PrintLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	cur := head.Next
	tmp := head.Next
	for cur != nil && cur.Next != nil {
		pre.Next = cur.Next
		pre = pre.Next
		cur.Next = pre.Next
		cur = cur.Next
	}
	pre.Next = tmp
	return head
}

func main() {
	var nums = []int{1, 2, 3, 4, 5}
	head := CreateLinkedList(nums)
	PrintLinkedList(head)
	PrintLinkedList(oddEvenList(head))
}

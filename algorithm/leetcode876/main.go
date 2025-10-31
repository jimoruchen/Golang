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

func middleNode(head *ListNode) *ListNode {
	dummy := &ListNode{-1, head}
	pre := dummy
	cur := pre
	length := 0
	for pre.Next != nil {
		length++
		pre = pre.Next
	}
	mid := length/2 + 1
	for mid != 0 {
		cur = cur.Next
		mid--
	}
	return cur
}

func middleNode1(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6}
	head := CreateLinkedList(nums)
	PrintLinkedList(head)
	middle := middleNode(head)
	PrintLinkedList(middle)
}

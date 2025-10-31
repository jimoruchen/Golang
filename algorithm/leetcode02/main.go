package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		count := carry
		if l1 != nil {
			count += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			count += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: count % 10}
		carry = count / 10
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummy.Next
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
	cur := head
	for cur != nil {
		fmt.Printf("%d->", cur.Val)
		cur = cur.Next
	}
	fmt.Println("nil")
}

func main() {
	var nums1 = []int{2, 4, 3}
	var nums2 = []int{5, 6, 4}
	l1 := CreateLinkedList(nums1)
	l2 := CreateLinkedList(nums2)
	PrintLinkedList(l1)
	PrintLinkedList(l2)
	head := addTwoNumbers(l1, l2)
	PrintLinkedList(head)
}

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	for cur := dummy; cur.Next != nil; {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

func removeElements1(head *ListNode, val int) *ListNode {
	dummy := &ListNode{-1, head}
	pre := dummy
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = pre.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}

func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

func CreateLinkedList(nums []int) (head *ListNode) {
	if len(nums) == 0 {
		return nil
	}
	head = &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return
}

func PrintList(head *ListNode) {
	if head == nil {
		return
	}
	cur := head
	for cur != nil {
		fmt.Printf("%d->", cur.Val)
		cur = cur.Next
	}
	fmt.Printf("nil\n")
}

func main() {
	var nums = []int{1, 2, 6, 3, 4, 5, 6}
	var target int
	fmt.Scan(&target)
	head := CreateLinkedList(nums)
	PrintList(head)
	cur := removeElements(head, target)
	for cur != nil {
		fmt.Printf("%d->", cur.Val)
		cur = cur.Next
	}
}

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		tmp := cur.Next
		cur.Next = tmp.Next
		tmp.Next = pre.Next
		pre.Next = tmp
	}
	return dummy.Next
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

func PrintLinkedList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d->", cur.Val)
		cur = cur.Next
	}
	fmt.Println("nil")
}

func main() {
	var nums = []int{1, 2, 3, 4, 5}
	var left, right int
	fmt.Scan(&left, &right)
	list := CreateLinkedList(nums)
	PrintLinkedList(list)
	list = reverseBetween(list, left, right)
	PrintLinkedList(list)
}

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	len := 0
	dummy := &ListNode{Val: -1, Next: head}
	cur := dummy
	tmp := dummy
	for cur.Next != nil {
		len++
		cur = cur.Next
	}
	for i := 0; i < len-n; i++ {
		tmp = tmp.Next
	}
	tmp.Next = tmp.Next.Next
	return dummy.Next
}

func CreatLinkList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head
}

func PrintList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d->", cur.Val)
		cur = cur.Next
	}
	fmt.Printf("nil\n")
}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7}
	head := CreatLinkList(nums)
	PrintList(head)
	removeNthFromEnd(head, 2)
	PrintList(head)
}

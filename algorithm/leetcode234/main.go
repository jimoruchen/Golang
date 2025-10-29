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
	cur := head
	for cur != nil {
		fmt.Printf("%d->", cur.Val)
		cur = cur.Next
	}
	fmt.Println("nil")
}

func copy(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{}
	cur := dummy
	for cur1 := head; cur1 != nil; cur1 = cur1.Next {
		newNode := &ListNode{Val: cur1.Val}
		cur.Next = newNode
		cur = newNode
	}
	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	dummy := &ListNode{-1, head}
	pre := dummy
	cur := head
	for cur != nil && cur.Next != nil {
		tmp := cur.Next
		cur.Next = tmp.Next
		tmp.Next = pre.Next
		pre.Next = tmp
	}
	return dummy.Next
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	copy := copy(head)
	reverse := reverse(copy)
	p1 := head
	p2 := reverse
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func isPalindrome1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var nums []int
	for cur := head; cur != nil; cur = cur.Next {
		nums = append(nums, cur.Val)
	}
	length := len(nums)
	for i := 0; i < length/2; i++ {
		if nums[i] != nums[length-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var nums = []int{1, 2, 2, 1}
	head := CreateLinkedList(nums)
	PrintLinkedList(head)
	fmt.Println(isPalindrome1(head))
}

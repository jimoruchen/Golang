package main

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var list []*ListNode
	cur := head
	for cur != nil {
		list = append(list, cur)
		cur = cur.Next
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Val < list[j].Val
	})
	for i := 0; i < len(list)-1; i++ {
		list[i].Next = list[i+1]
	}
	list[len(list)-1].Next = nil
	return list[0]
}

# 算法

## 1、两数之和
### **题目**
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
你可以按任意顺序返回答案。

* 示例1：
>输入：nums = [2,7,11,15], target = 9
>输出：[0,1]
>解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

### 代码
```go
func twoSum(nums []int, target int) []int {
    var mapAns = map[int]int{}
    len := len(nums)
    for i := 0; i < len; i++ {
        value, ok := mapAns[target - nums[i]]
        if (ok) {
            return []int{i, value}
        } else {
            mapAns[nums[i]] = i
        }
    }
    return []int{}
}
```
```go
func twoSum(nums []int, target int) []int {
    var mapAns = map[int]int{}
    len := len(nums)
    for i := 0; i < len; i++ {
        if value, ok := mapAns[target - nums[i]]; ok {
            return []int{i, value}
        } else {
            mapAns[nums[i]] = i
        }
    }
    return []int{}
}
```

<hr>

## 19、删除链表的倒数第 N 个结点
### **题目**
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

* 示例1：
>输入：head = [1,2,3,4,5], n = 2
>输出：[1,2,3,5]

### 代码
```go
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
```

<hr>

## 24、两两交换链表中的节点
### **题目**
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

* 示例1：
>输入：head = [1,2,3,4]
>输出：[2,1,4,3]

### 代码
```go
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{-1, head}
	prev := dummy
	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := prev.Next.Next
		prev.Next = second
		first.Next = second.Next
		second.Next = first
		prev = first
	}
	return dummy.Next
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
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
	fmt.Printf("nil\n")
}

func main() {
	var nums = []int{1, 2, 3, 4}
	head := CreateLinkedList(nums)
	PrintList(head)
	head = swapPairs(head)
	PrintList(head)
}
```

<hr>

## 26、删除有序数组中的重复项
### **题目**
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，
返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

* 示例1：
>输入：nums = [1,1,2]
>输出：2, nums = [1,2,_]

### 代码
```go
package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[slow] != nums[i] {
			slow++
			nums[slow] = nums[i]
		}
	}
	return slow + 1
}

func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			slow++
			nums[slow] = nums[i]
		}
	}
	return slow + 1
}
```

<hr>

## 27、移除元素
### **题目**
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。

* 示例1：
>输入：nums = [3,2,2,3], val = 3
>输出：2, nums = [2,2,_,_]

### 代码
```go
package main

func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	return j
}

func removeElement1(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
```

<hr>

## 160、相交链表
### **题目**
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。

* 示例1：
>输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
>输出：Intersected at '8'

### 代码
```go
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	curA, curB := headA, headB
	for curA != curB {
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headB
		}
		if curB != nil {
			curB = curB.Next
		} else {
			curB = headA
		}
	}
	return curA
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	var maps = make(map[*ListNode]struct{})
	curA, curB := headA, headB
	for curA != nil {
		maps[curA] = struct{}{}
		curA = curA.Next
	}
	for curB != nil {
		if _, ok := maps[curB]; ok {
			return curB
		} else {
			curB = curB.Next
		}
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	var maps = make(map[*ListNode]bool)
	curA, curB := headA, headB
	for curA != nil {
		maps[curA] = true
		curA = curA.Next
	}
	for curB != nil {
		if maps[curB] {
			return curB
		} else {
			curB = curB.Next
		}
	}
	return nil
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
	// 先创建公共的相交部分
	common := &ListNode{Val: 8}
	common.Next = &ListNode{Val: 4}
	common.Next.Next = &ListNode{Val: 5}

	// 构造 headA: 4 -> 1 -> [8 -> 4 -> 5]
	headA := &ListNode{Val: 4}
	headA.Next = &ListNode{Val: 1}
	headA.Next.Next = common

	// 构造 headB: 5 -> 6 -> 1 -> [8 -> 4 -> 5]
	headB := &ListNode{Val: 5}
	headB.Next = &ListNode{Val: 6}
	headB.Next.Next = &ListNode{Val: 1}
	headB.Next.Next.Next = common

	// 打印验证
	PrintList(headA) // 4->1->8->4->5->nil
	PrintList(headB) // 5->6->1->8->4->5->nil

	// 查找交点
	headC := getIntersectionNode(headA, headB)
	if headC != nil {
		fmt.Printf("相交于节点，值为: %d\n", headC.Val) // 应该输出 8
	} else {
		fmt.Println("无交点")
	}
}
```

<hr>

## 203、移除链表元素
### **题目**
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

* 示例1：
>输入：nums = [1, 2, 6, 3, 4, 5, 6], val = 6
>输出：[1, 2, 3, 4, 5]

### 代码
```go
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
```

<hr>

## 283、移动零
### **题目**
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

* 示例1：
>输入：nums = [0,1,0,3,12]
>输出：[1,3,12,0,0]

### 代码
```go
package main

func moveZeroes(nums []int) {
	var len = len(nums)
	var tmp = 0
	for i := 0; i < len; i++ {
		if nums[i] != 0 {
			nums[tmp] = nums[i]
			tmp++
		}
	}
	for i := tmp; i < len; i++ {
		nums[i] = 0
	}
}

func moveZeroes1(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
```
切片本身是引用类型，可以直接交换。

<hr>


## 485、最大连续 1 的个数
### **题目**
给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。

* 示例1：
>输入：nums = [1,1,0,1,1,1]
>输出：3

### 代码
```go
package main

func findMaxConsecutiveOnes(nums []int) int {
	var ans = 0
	var tmp = 0
	var length = len(nums)
	for i := 0; i < length; i++ {
		if nums[i] == 1 {
			tmp++
			if tmp > ans {
				ans = tmp
			}
		} else {
			tmp = 0
		}
	}
	return ans
}
```


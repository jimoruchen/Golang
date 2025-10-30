# 算法

## 1、两数之和
### 题目
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

## 3、无重复字符的最长子串
### 题目
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。

* 示例1：
>输入：s = "abcabcbb"
>输出：3
>解释：因为无重复字符的最长子串是 "abc"，所以其长度为 3。注意 "bca" 和 "cab" 也是正确答案。

### 代码
```go
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maps := make(map[byte]struct{})
	left := 0
	count := 0
	for right := 0; right < len(s); right++ {
		for {
			if _, ok := maps[s[right]]; ok {
				delete(maps, s[left])
				left++
			} else {
				break
			}
		}
		maps[s[right]] = struct{}{}
		count = max(count, right-left+1)
	}
	return count
}

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}
```

<hr>

## 19、删除链表的倒数第 N 个结点
### 题目
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

## 21、合并两个有序链表
### 题目
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

* 示例1：
>输入：l1 = [1,2,4], l2 = [1,3,4]
>输出：[1,1,2,3,4,4]

### 代码
```go
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	pre := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}
	if list1 != nil {
		pre.Next = list1
	}
	if list2 != nil {
		pre.Next = list2
	}
	return dummy.Next
}

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list2.Next, list1)
		return list2
	}
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
	var nums1 = []int{1, 2, 4}
	var nums2 = []int{1, 3, 4}
	list1 := CreatLinkList(nums1)
	list2 := CreatLinkList(nums2)
	head := mergeTwoLists(list1, list2)
	PrintList(head)
}
```

<hr>

## 24、两两交换链表中的节点
### 题目
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

* 示例1：
>输入：head = [1,2,3,4]
>输出：[2,1,4,3]

<img src="https://s2.loli.net/2025/10/10/73USJjm1BebrIQh.png" >

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

func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head.Next
	head.Next = swapPairs(tmp.Next)
	tmp.Next = head
	return tmp
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
### 题目
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

## 49、字母异位词分组
### 题目
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

* 示例1：
>输入：strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
>输出：[["bat"],["nat","tan"],["ate","eat","tea"]]

### 代码
```go
package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	maps := make(map[string][]string)
	for _, str := range strs {
		chars := []byte(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		sortStr := string(chars)
		maps[sortStr] = append(maps[sortStr], str)
	}
	var ans [][]string
	for _, v := range maps {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ans := groupAnagrams(strs)
	fmt.Println(ans)
}
```

<hr>

## 92、反转链表 II
### 题目
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

* 示例1：
>输入：head = [1,2,3,4,5], left = 2, right = 4
>输出：[1,4,3,2,5]

<a href="https://sm.ms/image/zmte8AKy9p27GvS" target="_blank"><img src="https://s2.loli.net/2025/08/26/zmte8AKy9p27GvS.png" alt="image.png"></a>

### 代码
```go
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
```

<hr>

## 138、随机链表的复制
### 题目
给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。

* 示例1：
>输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
>输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

### 代码
```go
package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	maps := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		maps[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		newNode := maps[cur]
		newNode.Next = maps[cur.Next]
		newNode.Random = maps[cur.Random]
		cur = cur.Next
	}
	return maps[head]
}

var maps map[*Node]*Node

func deepCopy(head *Node) *Node {
	if head == nil {
		return nil
	}
	if n, ok := maps[head]; ok {
		return n
	}
	newNode := &Node{Val: head.Val}
	maps[head] = newNode
	newNode.Next = deepCopy(head.Next)
	newNode.Random = deepCopy(head.Random)
	return newNode
}

func copyRandomList1(head *Node) *Node {
	maps = make(map[*Node]*Node)
	return deepCopy(head)
}
```

<hr>

## 141、环形链表
### 题目
给你一个链表的头节点 head ，判断链表中是否有环。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。  
如果链表中存在环 ，则返回 true 。 否则，返回 false 。

* 示例1：
>输入：head = [3,2,0,-4], pos = 1
>输出：true

### 代码
```go
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var maps = make(map[*ListNode]bool)
	cur := head
	for cur != nil {
		if maps[cur] {
			return true
		} else {
			maps[cur] = true
			cur = cur.Next
		}
	}
	return false
}

func hasCycle1(head *ListNode) bool {
	var maps = make(map[*ListNode]struct{})
	cur := head
	for cur != nil {
		if _, ok := maps[cur]; ok {
			return true
		} else {
			maps[cur] = struct{}{}
			cur = cur.Next
		}
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
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
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = head
	fmt.Println(hasCycle(head))
}
```

<hr>

## 142、环形链表 II
### 题目
给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

* 示例1：
>输入：head = [3,2,0,-4], pos = 1
>输出：返回索引为 1 的链表节点

### 代码
```go
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	maps := make(map[*ListNode]struct{})
	cur := head
	for cur != nil {
		if _, ok := maps[cur]; ok {
			return cur
		} else {
			maps[cur] = struct{}{}
			cur = cur.Next
		}
	}
	return nil
}

func detectCycle1(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func detectCycle2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
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
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = head
	fmt.Println(detectCycle(head).Val)
}
```

<hr>

## 160、相交链表
### 题目
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
### 题目
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

## 206、反转链表
### 题目
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

* 示例1：
>输入：head = [1,2,3,4,5]
>输出：[5,4,3,2,1]

<img src="https://s2.loli.net/2025/10/10/zkbaH3mh4OKyXYx.png" >

### 代码
```go
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
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
	list := CreateLinkedList(nums)
	PrintLinkedList(list)
	list = reverseList(list)
	PrintLinkedList(list)
}
```

<hr>

## 234、回文链表
### 题目
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

* 示例1：
>输入：head = [1,2,2,1]
>输出：true

### 代码
```go
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
```

<hr>

## 237、删除链表的节点
### 题目
有一个单链表的 head，我们想删除它其中的一个节点 node。

* 示例1：
>输入：head = [4,5,1,9], node = 5
>输出：[4,1,9]

### 代码
```go
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
```

<hr>

## 283、移动零
### 题目
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

## 328、奇偶链表
### 题目
给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别分组，保持它们原有的相对顺序，然后把偶数索引节点分组连接到奇数索引节点分组之后，返回重新排序的链表。

* 示例1：
>输入：head = [1,2,3,4,5]
>输出：[1,3,5,2,4]

### 代码
```go
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
```

<hr>

## 485、最大连续 1 的个数
### 题目
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

<hr>

## 876、链表的中间结点
### 题目
给定一个头结点为 head 的非空单链表，返回链表的中间结点。  
如果有两个中间结点，则返回第二个中间结点。  

* 示例1：
>输入：head = [1,2,3,4,5]
>输出：[3,4,5]

* 示例2：
>输入：head = [1,2,3,4,5,6]
>输出：[4,5,6]

### 代码
```go
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
```
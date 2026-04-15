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
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mapAns := map[int]int{}
	for i, num := range nums {
		if value, ok := mapAns[target-num]; ok {
			return []int{value, i}
		} else {
			mapAns[num] = i
		}
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	var Mymap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := Mymap[target-nums[i]]; ok {
			return []int{val, i}
		} else {
			Mymap[nums[i]] = i
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	var maps = map[int]int{}
	for i, num := range nums {
		if _, ok := maps[target-num]; ok {
			return []int{i, maps[target-num]}
		}
		maps[num] = i
	}
	return nil
}

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 9
	ans := twoSum2(nums, target)
	fmt.Println(ans)
}
```

<hr>

## 2、两数相加
### 题目
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

* 示例1：
>输入：l1 = [2,4,3], l2 = [5,6,4]
>输出：[7,0,8]
>解释：342 + 465 = 807.

### 代码
```go
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

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	pre := dummy
	carry := 0
	for l1 != nil && l2 != nil {
		tmp := l1.Val + l2.Val + carry
		pre.Next = &ListNode{Val: tmp % 10}
		carry = tmp / 10
		l1 = l1.Next
		l2 = l2.Next
		pre = pre.Next
	}
	for l1 != nil {
		tmp := carry + l1.Val
		pre.Next = &ListNode{Val: tmp % 10}
		carry = tmp / 10
		l1 = l1.Next
		pre = pre.Next
	}
	for l2 != nil {
		tmp := carry + l2.Val
		pre.Next = &ListNode{Val: tmp % 10}
		carry = tmp / 10
		l2 = l2.Next
		pre = pre.Next
	}
	if carry != 0 {
		pre.Next = &ListNode{Val: carry}
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
```

<hr>

## 3、无重复字符的最长子串
### 题目
给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。

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

func lengthOfLongestSubstring1(s string) int {
	maps := make(map[byte]bool)
	left := 0
	res := 0
	for right := 0; right < len(s); right++ {
		for maps[s[right]] {
			maps[s[left]] = false
			left++
		}
		maps[s[right]] = true
		res = max(res, right - left + 1)
	}
	return res
}

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}
```

<hr>

## 5、最长回文子串
### 题目
给你一个字符串 s，找到 s 中最长的 回文 子串。

* 示例1：
>输入：s = "babad"
>输出："bab"

### 代码
```go
package main

func longestPalindrome(s string) string {
	var ans string
	length := 0
	expand := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > length {
				length = right - left + 1
				ans = s[left : right+1]
			}
			left--
			right++
		}
	}
	for i := 0; i < len(s); i++ {
		expand(i, i)
		expand(i, i+1)
	}
	return ans
}
```

<hr>

## 11、盛最多水的容器
### 题目
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。  
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。  
返回容器可以储存的最大水量。  

* 示例1：
>输入：[1,8,6,2,5,4,8,3,7]
>输出：49

### 代码
```go
package main

import "fmt"

func maxArea(height []int) int {
	length := len(height)
	ans := 0
	left := 0
	right := length - 1
	for left != right {
		tmp := 0
		if height[left] < height[right] {
			tmp = height[left] * (right - left)
			ans = max(ans, tmp)
			left++
		} else {
			tmp = height[right] * (right - left)
			ans = max(ans, tmp)
			right--
		}
	}
	return ans
}

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
}
```

<hr>

## 15、三数之和
### 题目
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

* 示例1：
>输入：nums = [-1,0,1,2,-1,-4]
>输出：[[-1,-1,2],[-1,0,1]]

### 代码
```go
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left+1] == nums[left] {
					left++
				}
				for left < right && nums[right-1] == nums[right] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ans
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ans := threeSum(nums)
	fmt.Println(ans)
}
```

<hr>

## 17、电话号码的字母组合
### 题目
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

* 示例1：
>输入：digits = "23"
>输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

### 代码
```go
package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	mapping := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var path []byte
	var result []string
	var backtracking func(int)
	backtracking = func(index int) {
		if len(digits) == index {
			result = append(result, string(path))
			return
		}
		for _, digit := range mapping[digits[index]-'0'] {
			path = append(path, byte(digit))
			backtracking(index + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return result
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

## 20、有效的括号
### 题目
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

* 示例1：
>输入：s = "()[]{}"
>输出：true

### 代码
```go
package main

import (
	"container/list"
	"fmt"
)

func isValid(s string) bool {
	stack := list.New()
	for _, str := range s {
		if str == '(' || str == '[' || str == '{' {
			stack.PushBack(str)
		} else if str == ')' {
			if stack.Len() == 0 {
				return false
			}
			if stack.Back().Value != '(' {
				return false
			}
			stack.Remove(stack.Back())
		} else if str == ']' {
			if stack.Len() == 0 {
				return false
			}
			if stack.Back().Value != '[' {
				return false
			}
			stack.Remove(stack.Back())
		} else {
			if stack.Len() == 0 {
				return false
			}
			if stack.Back().Value != '{' {
				return false
			}
			stack.Remove(stack.Back())
		}
	}
	if stack.Len() != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("()"))
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

## 22、括号生成
### 题目
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

* 示例1：
>输入：n = 3
>输出：["((()))","(()())","(())()","()(())","()()()"]

### 代码
```go
package main

func generateParenthesis(n int) []string {
	var path []byte
	var result []string
	var backtracking func(int, int)
	backtracking = func(left, right int) {
		if len(path) == 2*n {
			result = append(result, string(path))
			return
		}
		if left < n {
			path = append(path, '(')
			backtracking(left+1, right)
			path = path[:len(path)-1]
		}
		if right < left {
			path = append(path, ')')
			backtracking(left, right+1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0, 0)
	return result
}
```

<hr>E

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

## 31、下一个排列
### 题目
整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。
例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。
更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。
如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

* 示例1：
>输入：nums = [1,2,3]
>输出：[1,3,2]

### 代码
```go
package main

func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums, i+1, n-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
```

<hr>

## 33、搜索旋转排序数组
### 题目
整数数组 nums 按升序排列，数组中的值 互不相同 。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 向左旋转，
使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
例如， [0,1,2,4,5,6,7] 下标 3 上向左旋转后可能变为 [4,5,6,7,0,1,2] 。

* 示例1：
>输入：nums = [4,5,6,7,0,1,2], target = 0
>输出：4

### 代码
```go
package main

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	l, r := 0, n - 1
	for l <= r {
		mid := l + (r - l) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[0] {
			if target >= nums[0] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[n - 1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
```

<hr>

## 34、在排序数组中查找元素的第一个和最后一个位置
### 题目
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。

* 示例1：
>输入：nums = [5,7,7,8,8,10], target = 8
>输出：[3,4]

### 代码
```go
package main

import (
	"fmt"
	"sort"
)

func searchRange(nums []int, target int) []int {
	left := findFirst(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}
	right := findLast(nums, target)
	return []int{left, right}
}

// 找第一个等于 target 的索引
func findFirst(nums []int, target int) int {
	i, j := 0, len(nums)-1
	res := -1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] == target {
			res = mid   // 记录可能的左边界
			j = mid - 1 // 继续向左找
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return res
}

// 找最后一个等于 target 的索引
func findLast(nums []int, target int) int {
	i, j := 0, len(nums)-1
	res := -1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] == target {
			res = mid   // 记录可能的右边界
			i = mid + 1 // 继续向右找
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return res
}

func searchRange1(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange1(nums, target))
}
```

<hr>

## 35、搜索插入位置
### 题目
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

* 示例1：
>输入：nums = [1,3,5,6], target = 5
>输出：2

### 代码
```go
package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right { //最后一轮没命中刚好left+1就是要插入的位置
		mid := (right-left)/2 + left
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
}
```

<hr>

## 39、组合总和
### 题目
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，
找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

* 示例1：
>输入：candidates = [2,3,6,7], target = 7
>输出：[[2,2,3],[7]]

### 代码
```go
package main

func combinationSum(candidates []int, target int) [][]int {
	var path []int
	var result [][]int
	var backtracking func(int, int)
	backtracking = func(startIndex, currentSum int) {
		if currentSum > target {
			return
		}
		if currentSum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := startIndex; i < len(candidates); i++ {
			path = append(path, candidates[i])
			backtracking(i, currentSum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	backtracking(0, 0)
	return result
}
```

<hr>

## 45、跳跃游戏 Ⅱ
### 题目
给定一个长度为 n 的 0 索引整数数组 nums。初始位置在下标 0。
每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在索引 i 处，你可以跳转到任意 (i + j) 处：
0 <= j <= nums[i] 且
i + j < n
返回到达 n - 1 的最小跳跃次数。测试用例保证可以到达 n - 1。

* 示例1：
>输入：nums = [2,3,1,1,4]
>输出：2

### 代码
```go
package main

func jump(nums []int) int {
	maxLength := 0
	right := 0
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		maxLength = max(maxLength, i+nums[i])
		if i == right {
			right = maxLength
			count++
		}
	}
	return count
}
```

<hr>

## 46、全排列
### 题目
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

* 示例1：
>输入：nums = [1,2,3]
>输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

### 代码
```go
package main

import "fmt"

func permute(nums []int) [][]int {
	var path []int
	var result [][]int
	used := make([]bool, len(nums))
	var backtracking func()
	backtracking = func() {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtracking()
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backtracking()
	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(permute(nums))
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

## 53、最大子数组和
### 题目 
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。  
子数组是数组中的一个连续部分。  

* 示例1：
>输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
>输出：6
>解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

### 代码
```go
package main

import "fmt"

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > maxSum {
			maxSum = nums[i]
		}
	}
	return maxSum
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
```

<hr>

## 54、螺旋矩阵
### 题目
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

* 示例1：
>输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
>输出：[1,2,3,6,9,8,7,4,5]

### 代码
```go
package main

func spiralOrder(matrix [][]int) []int {
	var ans []int
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return nil
	}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	row, col, index := 0, 0, 0
	for i := 0; i < m*n; i++ {
		ans = append(ans, matrix[row][col])
		visited[row][col] = true
		nextRow, nextCol := row+dir[index][0], col+dir[index][1]
		if nextRow < 0 || nextRow >= m || nextCol < 0 || nextCol >= n || visited[nextRow][nextCol] {
			index = (index + 1) % 4
		}
		row += dir[index][0]
		col += dir[index][1]
	}
	return ans
}

func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	var ans []int
	m, n := len(matrix), len(matrix[0])
	top, bottom, left, right := 0, m-1, 0, n-1
	for len(ans) < m*n {
		// 1. 从左到右 (Top Row)
		for col := left; col <= right && len(ans) < m*n; col++ {
			ans = append(ans, matrix[top][col])
		}
		top++
		// 2. 从上到下 (Right Col)
		for row := top; row <= bottom && len(ans) < m*n; row++ {
			ans = append(ans, matrix[row][right])
		}
		right--
		// 3. 从右到左 (Bottom Row)
		for col := right; col >= left && len(ans) < m*n; col-- {
			ans = append(ans, matrix[bottom][col])
		}
		bottom--
		// 4. 从下到上 (Left Col)
		for row := bottom; row >= top && len(ans) < m*n; row-- {
			ans = append(ans, matrix[row][left])
		}
		left++
	}
	return ans
}
```

<hr>

## 55、跳跃游戏
### 题目
给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。

* 示例1：
>输入：nums = [2,3,1,1,4]
>输出：true

### 代码
```go
package main

func canJump(nums []int) bool {
	maxLength := 0
	for i := 0; i < len(nums); i++ {
		if maxLength >= len(nums)-1 {
			return true
		}
		if i > maxLength {
			return false
		}
		maxLength = max(maxLength, i+nums[i])
	}
	return true
}
```

<hr>

## 56、合并区间
### 题目
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。  
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

* 示例1：
>输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
>输出：[[1,6],[8,10],[15,18]]
>解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6]

### 代码
```go
package main

import (
	"fmt"
	"slices"
)

func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	// sort.Slice(intervals, func(i, j int) bool {
	//     return intervals[i][0] < intervals[j][0]
	// })
	slices.SortFunc(intervals, func(p, q []int) int {
		return p[0] - q[0]
	})
	for _, nums := range intervals {
		length := len(ans)
		if length > 0 &&  nums[0] <= ans[length - 1][1] {
			ans[length -1][1] = max(nums[1], ans[length -1][1])
		} else {
			ans = append(ans, nums)
		}
	}
	return ans
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals = merge(intervals)
	for _, interval := range intervals {
		fmt.Println(interval)
	}
}
```

<hr>

## 54、螺旋矩阵 II
### 题目
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

* 示例1：
>输入：n = 3
>输出：[[1,2,3],[8,9,4],[7,6,5]]

### 代码
```go
package main

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	left, top := 0, 0
	right, bottom := n-1, n-1
	tmp := 1
	for left <= right && top <= bottom {
		for j := left; j <= right; j++ {
			matrix[top][j] = tmp
			tmp++
		}
		top++
		for i := top; i <= bottom; i++ {
			matrix[i][right] = tmp
			tmp++
		}
		right--
		for j := right; j >= left; j-- {
			matrix[bottom][j] = tmp
			tmp++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			matrix[i][left] = tmp
			tmp++
		}
		left++
	}
	return matrix
}
```

<hr>

## 62、不同路径
### 题目
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径？

* 示例1：
>输入：m = 3, n = 7
>输出：28

### 代码
```go
package main

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 1; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 1; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
```

<hr>

## 64、最小路径和
### 题目
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。

* 示例1：
>输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
>输出：7

### 代码
```go
package main

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}
```

<hr>

## 69、x 的平方根
### 题目
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。

* 示例1：
>输入：x = 4
>输出：2

### 代码
```go
package main

// func mySqrt(x int) int {
//     if x == 1 {
//         return 1
//     }
//     var ans int
//     for i := 0; i <= x / 2; i++ {
//         if i * i <= x && (i + 1) * (i + 1) > x {
//             ans = i
//             break
//         }
//     }
//     return ans
// }

// m^2 <= x (m+1)^2 > x
// 0 1 2 3 4  x^1/2

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left, right := 0, x
	for left < right {
		mid := (right-left)/2 + left
		if x < mid*mid { //找到第一个满足 mid * mid > x 的整数 mid，最后mid再减一。
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left - 1
}
```

<hr>

## 70、爬楼梯
### 题目
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

* 示例1：
>输入：n = 2
>输出：2

### 代码
```go
package main

// func climbStairs(n int) int {
//     if n <= 2 {
//         return n
//     } 
//     dp := make([]int, n + 1)
//     dp[1], dp[2] = 1, 2
//     for i := 3; i <= n; i++ {
//         dp[i] = dp[i - 1] + dp[i - 2]
//     }
//     return dp[n]
// }

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	prev, cur := 1, 2
	for i := 3; i <= n; i++ {
		prev, cur = cur, prev + cur
	}
	return cur
}
```

<hr>

## 74、搜索二维矩阵
### 题目
给你一个满足下述两条属性的 m x n 整数矩阵：
每行中的整数从左到右按非严格递增顺序排列。
每行的第一个整数大于前一行的最后一个整数。
给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。

* 示例1：
>输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
>输出：true

### 代码
```go
package main

import (
	"fmt"
	"sort"
)

func searchMatrix(matrix [][]int, target int) bool {
	var nums []int
	for _, ma := range matrix {
		nums = append(nums, ma...)
	}
	left, right := 0, len(nums)
	for left < right {
		mid := (right-left)/2 + left
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left < len(nums) && nums[left] == target {
		return true
	}
	return false
}

func searchMatrix1(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n
	for left < right {
		mid := (right-left)/2 + left
		row := mid / n
		col := mid % n
		value := matrix[row][col]
		if target <= value {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left < m*n && matrix[left/n][left%n] == target {
		return true
	}
	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool {
		return matrix[i][0] > target
	}) - 1
	if row < 0 {
		return false
	}
	// col := sort.Search(len(matrix[0]), func(i int) bool {
	//     return matrix[row][i] >= target
	// })
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && target == matrix[row][col]
}

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	fmt.Println(searchMatrix(matrix, 3))
}

```

<hr>

## 75、颜色分类
### 题目
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地 对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

* 示例1：
>输入：nums = [2,0,2,1,1,0]
>输出：[0,0,1,1,2,2]

### 代码
```go
package main

import "fmt"

func sortColors(nums []int) {
	maps := make(map[int]int)
	for _, num := range nums {
		maps[num]++
	}
	for i := 0; i < maps[0]; i++ {
		nums[i] = 0
	}
	for i := maps[0]; i < maps[0]+maps[1]; i++ {
		nums[i] = 1
	}
	for i := maps[0] + maps[1]; i < len(nums); i++ {
		nums[i] = 2
	}
}

func sortColors1(nums []int) {
	p0, p1 := 0, 0
	for i, num := range nums {
		nums[i] = 2
		if num <= 1 {
			nums[p1] = 1
			p1++
		}
		if num == 0 {
			nums[p0] = 0
			p0++
		}
	}
}

func main() {
	nums := []int{1, 1, 2, 0, 2, 0}
	sortColors(nums)
	fmt.Println(nums)
}
```

<hr>

## 77、组合
### 题目
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。

* 示例1：
>输入：n = 4, k = 2
>输出：[[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]

### 代码
```go
package main

import "fmt"

func combine(n int, k int) [][]int {
	var path []int
	var result [][]int
	backtracking(n, k, 1, path, &result)
	return result
}

func backtracking(n, k, startIndex int, path []int, result *[][]int) {
	if len(path) == k {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}
	for i := startIndex; i <= n; i++ {
		path = append(path, i)
		backtracking(n, k, i+1, path, result)
		path = path[:len(path)-1]
	}
}

func combine1(n int, k int) [][]int {
	var result [][]int
	var path []int
	var backtracking func(int)
	backtracking = func(startIndex int) {
		if len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := startIndex; i <= n; i++ {
			path = append(path, i)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(1)
	return result
}

func main() {
	n := 4
	k := 2
	fmt.Println(combine(n, k))
}
```

<hr>

## 78、子集
### 题目
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

* 示例1：
>输入：nums = [1,2,3]
>输出：[[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]

### 代码
```go
package main

import "fmt"

func subsets(nums []int) [][]int {
	var path []int
	var result [][]int
	var backtracking func(int)
	backtracking = func(startIndex int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return result
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
```

<hr>

## 79、删除有序数组中的重复项
### 题目
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

* 示例1：
>输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "ABCCED"
>输出：true

### 代码
```go
package main

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	if m == 0 || n == 0 || len(word) == 0 {
		return false
	}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var backtracking func(index, i, j int) bool
	backtracking = func(index, i, j int) bool {
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || board[i][j] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		visited[i][j] = true
		found := backtracking(index+1, i-1, j) || backtracking(index+1, i, j-1) || backtracking(index+1, i+1, j) || backtracking(index+1, i, j+1)
		if !found {
			visited[i][j] = false
		}
		return found
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtracking(0, i, j) {
				return true
			}
		}
	}
	return false
}
```

<hr>

## 88、合并两个有序数组
### 题目
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

* 示例1：
>输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
>输出：[1,2,2,3,5,6]

### 代码
```go
package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, right := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[right] = nums1[i]
			right--
			i--
		} else {
			nums1[right] = nums2[j]
			right--
			j--
		}
	}
	for j >= 0 {
		nums1[right] = nums2[j]
		right--
		j--
	}
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

## 94、二叉树的中序遍历
### 题目
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

* 示例1：
>输入：root = [1,null,2,3]
>输出：[1,3,2]

### 代码
```go
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return result
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	result := inorderTraversal(root)
	fmt.Println(result)
}
```

<hr>

## 98、验证二叉搜索树
### 题目
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

* 示例1：
>输入：root = [2,1,3]
>输出：true

### 代码
```go
package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	minValue := math.MinInt
	var inorder func(*TreeNode) bool
	inorder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !inorder(node.Left) {
			return false
		}
		if node.Val <= minValue {
			return false
		}
		minValue = node.Val
		return inorder(node.Right)
	}
	return inorder(root)
}

func main() {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(isValidBST(root))
}
```

<hr>

## 101、对称二叉树
### 题目
给你一个二叉树的根节点 root ， 检查它是否轴对称。

* 示例1：
>输入：root = [1,2,2,3,4,4,3]
>输出：true

### 代码
```go
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isSameTree(root.Left, root.Right)
}

func isSameTree(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Right) && isSameTree(p.Right, q.Left)
}

func isSymmetric1(root *TreeNode) bool {
	u, v := root, root
	var q []*TreeNode
	q = append(q, u)
	q = append(q, v)
	for len(q) > 0 {
		u, v = q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		q = append(q, u.Left)
		q = append(q, v.Right)
		q = append(q, u.Right)
		q = append(q, v.Left)
	}
	return true
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 3}
	fmt.Println(isSymmetric(root))
}
```

<hr>

## 102、二叉树的层序遍历
### 题目
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

* 示例1：
>输入：root = [1,2,2,3,4,4,3]
>输出：[[1] [2 2] [3 4 4 3]]

### 代码
```go
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		tmp := make([]int, size)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			tmp[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		result = append(result, tmp)
	}
	return result
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 3}
	fmt.Println(levelOrder(root))
}
```

<hr>

## 104、二叉树的最大深度
### 题目
给定一个二叉树 root ，返回其最大深度。

* 示例1：
>输入：root = [1,null,2,3]
>输出：3

### 代码
```go
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		depth++
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return depth
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	result := maxDepth(root)
	fmt.Println(result)
}
```

<hr>

## 108、将有序数组转换为二叉搜索树
### 题目
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 平衡 二叉搜索树。

* 示例1：
>输入：nums = [-10,-3,0,5,9]
>输出：[0,-3,9,-10,null,5]

### 代码
```go
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return buildBST(nums, 0, len(nums)-1)
}

func buildBST(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = buildBST(nums, left, mid-1)
	root.Right = buildBST(nums, mid+1, right)
	return root
}

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		ans = append(ans, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return ans
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(nums)
	ans := inorderTraversal(root)
	fmt.Println(ans)
}
```

<hr>

## 114、二叉树展开为链表
### 题目
给你二叉树的根结点 root ，请你将它展开为一个单链表：
展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

* 示例1：
>输入：root = [1,2,5,3,4,null,6]
>输出：[1,null,2,null,3,null,4,null,5,null,6]

### 代码
```go
func flatten(root *TreeNode)  {
    var tmp []*TreeNode 
    var preorder func(*TreeNode)
    preorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        tmp = append(tmp, node)
        preorder(node.Left)
        preorder(node.Right)
    }
    preorder(root)
    current := root
    for i := 1; i < len(tmp); i++ {
        current.Left = nil
        current.Right = tmp[i]
        current = tmp[i]
    }
}
```

<hr>

## 118、杨辉三角
### 题目
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。

* 示例1：
>输入：numRows = 5
>输出：[[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

### 代码
```go
package main

// func generate(numRows int) [][]int {
//     if numRows == 1 {
//         return [][]int{{1}}
//     }
//     if numRows == 2 {
//         return [][]int{{1}, {1, 1}}
//     }
//     var ans [][]int
//     dp := make([][]int, numRows)
//     for i := range dp {
//         dp[i] = make([]int, numRows)
//     }
//     dp[0][0] = 1
//     dp[1][0] = 1
//     dp[1][1] = 1
//     for i := 2; i < numRows; i++ {
//         dp[i][0] = 1
//         dp[i][i] = 1
//         for j := 1; j <= i - 1; j++ {
//             dp[i][j] = dp[i - 1][j - 1] + dp[i - 1][j]
//         }
//     }
//     for i := 0; i < numRows; i++ {
//         ans = append(ans, dp[i][:i + 1])
//     }
//     return ans
// }

func generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	for i := range dp {
		dp[i] = make([]int, i+1)
		dp[i][0] = 1
		dp[i][i] = 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}
	return dp
}
```

<hr>

## 121、买卖股票的最佳时机
### 题目
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

* 示例1：
>输入：[7,1,5,3,6,4]
>输出：5

### 代码
```go
package main

func maxProfit(prices []int) int {
	profit := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			profit = max(profit, prices[i]-minPrice)
		}
	}
	return profit
}
```

<hr>

## 128、最长连续序列
### 题目
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

* 示例1：
>输入：nums = [100,4,200,1,3,2]
>输出：4
>解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

### 代码
```go
package main

import "fmt"

func longestConsecutive(nums []int) int {
	ans := 0
	maps := make(map[int]bool)
	for _, num := range nums {
		maps[num] = true
	}
	for k, _ := range maps {
		if maps[k-1] {
			continue
		}
		y := k + 1
		for maps[y] {
			y++
		}
		ans = max(ans, y-k)
	}
	return ans
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
```

<hr>

## 131、分割回文串
### 题目
给你一个字符串 s，请你将 s 分割成一些 子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

* 示例1：
>输入：s = "aab"
>输出：[["a","a","b"],["aa","b"]]

### 代码
```go
package main

func partition(s string) [][]string {
	var path []string
	var res [][]string
	var backtracking func(startIndex int)
	backtracking = func(startIndex int) {
		if startIndex == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		for i := startIndex; i < len(s); i++ {
			str := s[startIndex : i+1]
			if isPalindrome(str) {
				path = append(path, str)
				backtracking(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	backtracking(0)
	return res
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
```

<hr>

## 136、只出现一次的数字
### 题目
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

* 示例1：
>输入：nums = [2,2,1]
>输出：1

### 代码
```go
package main

import "fmt"

func singleNumber(nums []int) int {
	ans := 0
	maps := make(map[int]int)
	for _, num := range nums {
		maps[num]++
	}
	for k, v := range maps {
		if v == 1 {
			ans = k
		}
	}
	return ans
}

func singleNumber1(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}

func main() {
	nums := []int{1, 2, 2, 3, 3}
	fmt.Println(singleNumber1(nums))
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

## 139、单词拆分
### 题目
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

* 示例1：
>输入：s = "leetcode", wordDict = ["leet", "code"]
>输出：true

### 代码
```go
package main

func wordBreak(s string, wordDict []string) bool {
	maps := make(map[string]bool)
	for _, str := range wordDict {
		maps[str] = true
	}
	dp := make([]bool, len(s)+1)
	// dp[i] 表示 s[0:i] 能否被拆分
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && maps[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
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

## 144、二叉树的前序遍历
### 题目
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

* 示例1：
>输入：root = [1,null,2,3]
>输出：[1,2,3]

### 代码
```go
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var result []int
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return result
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	result := preorderTraversal(root)
	fmt.Println(result)
}
```

<hr>

## 144、二叉树的后序遍历
### 题目
给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。

* 示例1：
>输入：root = [1,null,2,3]
>输出：[3,2,1]

### 代码
```go
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		result = append(result, node.Val)
	}
	postorder(root)
	return result
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	result := postorderTraversal(root)
	fmt.Println(result)
}
```

<hr>

## 146、LRU 缓存
### 题目
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

* 示例1：
>输入：l1 = [2,4,3], l2 = [5,6,4]
>输出：[7,0,8]
>解释：342 + 465 = 807.

### 代码



## 151、反转字符串中的单词
### 题目
给你一个字符串 s ，请你反转字符串中 单词 的顺序。
单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

* 示例1：
>输入：["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"] 
>      [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
>输出：[null, null, null, 1, null, -1, null, -1, 3, 4]

### 代码
```go
package main

type Node struct {
	key, val   int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	cache := make(map[int]*Node)
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    cache,
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		node.val = value
		this.moveToHead(node)
	} else {
		newNode := &Node{
			key: key,
			val: value,
		}
		this.cache[key] = newNode
		this.addToHead(newNode)
		if len(this.cache) > this.capacity {
			node := this.deleteTail()
			delete(this.cache, node.key)
		}
	}
}

func (this *LRUCache) moveToHead(node *Node) {
	this.deleteNode(node)
	this.addToHead(node)
}

func (this *LRUCache) deleteNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) deleteTail() *Node {
	res := this.tail.prev
	this.deleteNode(res)
	return res
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
```

<hr>

## 152、乘积最大数组
### 题目
给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续 子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
测试用例的答案是一个 32-位 整数。
请注意，一个只包含一个元素的数组的乘积是这个元素的值。

* 示例1：
>输入：nums = [2,3,-2,4]
>输出：6

### 代码
```go
package main

func maxProduct(nums []int) int {
	ans := nums[0]
	maxdp := make([]int, len(nums))
	mindp := make([]int, len(nums))
	maxdp[0] = nums[0]
	mindp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		curmax := maxdp
		curmin := mindp
		maxdp[i] = max(curmax[i-1]*nums[i], max(curmin[i-1]*nums[i], nums[i]))
		mindp[i] = min(curmax[i-1]*nums[i], min(curmin[i-1]*nums[i], nums[i]))
		ans = max(ans, maxdp[i])
	}
	return ans
}
```

<hr>

## 153、寻找旋转排序数组中的最小值
### 题目
已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。
给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。

* 示例1：
>输入：nums = [3,4,5,1,2]
>输出：1

### 代码
```go
package main

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right { //当 left == right 时已找到唯一候选，无需继续；避免死循环
		mid := left + (right-left)/2
		if nums[mid] > nums[len(nums)-1] {
			left = mid + 1
		} else {
			right = mid //mid 可能是最小值，必须保留
		}
	}
	return nums[left]
}
```

<hr>

## 154、寻找旋转排序数组中的最小值 II
### 题目
已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,4,4,5,6,7] 在变化后可能得到：
若旋转 4 次，则可以得到 [4,5,6,7,0,1,4]
若旋转 7 次，则可以得到 [0,1,4,4,5,6,7]
注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
给你一个可能存在 重复 元素值的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
你必须尽可能减少整个过程的操作步骤。

* 示例1：
>输入：nums = [1,3,5]
>输出：1

### 代码
```go
package main

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] == nums[right] {
			right--
		} else {
			right = mid
		}
	}
	return nums[left]
}
```

<hr>

## 155、最小栈
### 题目
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。  
实现 MinStack 类:  
MinStack() 初始化堆栈对象。  
void push(int val) 将元素val推入堆栈。  
void pop() 删除堆栈顶部的元素。  
int top() 获取堆栈顶部的元素。  
int getMin() 获取堆栈中的最小元素。  

* 示例1：
>输入：["MinStack","push","push","push","getMin","pop","top","getMin"] [[],[-2],[0],[-3],[],[],[],[]]
>输出：[null,null,null,null,-3,null,0,-2]

### 代码
```go
package main

import "fmt"

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	tmp := val
	if len(this.minStack) > 0 {
		tmp = min(this.minStack[len(this.minStack)-1], val)
	}
	this.minStack = append(this.minStack, tmp)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func main() {
	MinStack := Constructor()
	MinStack.Push(1)
	MinStack.Push(2)
	MinStack.Push(3)
	fmt.Println(MinStack.Top(), MinStack.GetMin())
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

## 169、多数元素
### 题目
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

* 示例1：
>输入：nums = [3,2,3]
>输出：3

### 代码
```go
package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	ans := 0
	maps := make(map[int]int)
	for _, num := range nums {
		maps[num]++
	}
	for k, v := range maps {
		if v > len(nums)/2 {
			ans = k
		}
	}
	return ans
}

func majorityElement1(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(majorityElement(nums))
}
```

<hr>

## 189、轮转数组
### 题目
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。  

* 示例1：
>输入：nums = [1,2,3,4,5,6,7], k = 3
>输出：[5,6,7,1,2,3,4]

### 代码
```go
package main

import "fmt"

func rotate(nums []int, k int) {
	var tmp []int
	k = k % len(nums)
	for i := len(nums) - k; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
	}
	fmt.Println(tmp)
	for i := 0; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
	}
	copy(nums, tmp)
}

func rotate1(nums []int, k int) {
	tmp := make([]int, len(nums))
	k = k % len(nums)
	for i, num := range nums {
		tmp[(k+i)%len(nums)] = num
	}
	copy(nums, tmp)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}
```

<hr>

## 198、打家劫舍
### 题目
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

* 示例1：
>输入：[1,2,3,1]
>输出：4

### 代码
```go
package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}
```

<hr>

## 199、二叉树的右视图
### 题目
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

* 示例1：
>输入：root = [1,2,3,null,5,null,4]
>输出：[1,3,4]

### 代码
```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if i == length-1 {
				ans = append(ans, node.Val)
			}
		}
	}
	return ans
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

## 209、长度最小的子数组
### 题目
给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

* 示例1：
>输入：target = 7, nums = [2,3,1,2,4,3]
>输出：2

### 代码
```go
package main

func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	n := len(nums) + 1
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			n = min(n, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if n == len(nums)+1 {
		n = 0
	}
	return n
}
```

<hr>

## 215、数组中的第K个最大元素
### 题目
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

* 示例1：
>输入：[3,2,1,5,6,4], k = 2
>输出：5

### 代码
```go
package main

import "math/rand"

func findKthLargest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, len(nums)-k)
}

//func findKthLargest(nums []int, k int) int {
//	quickSort(nums, 0, len(nums) - 1)
//	return nums[len(nums) - k]
//}
//
//func quickSort(nums []int, left, right int) {
//	if left >= right {
//		return
//	}
//	pivot := partition(nums, left, right)
//	quickSort(nums, left, pivot - 1)
//	quickSort(nums, pivot + 1, right)
//}

func quickSort(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}
	pivot := partition(nums, left, right)
	if pivot == k {
		return nums[pivot]
	} else if pivot > k {
		return quickSort(nums, left, pivot-1, k)
	} else {
		return quickSort(nums, pivot+1, right, k)
	}
}

func partition(nums []int, left, right int) int {
	i, j := left, right
	index := left + rand.Intn(right-left+1)
	nums[index], nums[i] = nums[i], nums[index]
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}
```

<hr>

## 226、翻转二叉树
### 题目
给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

* 示例1：
>输入：root = [4,2,7,1,3,6,9]
>输出：[4,7,2,9,6,3,1]

### 代码
```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return root
}
```

<hr>

## 230、二叉树搜索树中第 K 小的元素
### 题目
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（k 从 1 开始计数）。

* 示例1：
>输入：root = [3,1,4,null,2], k = 1
>输出：1

### 代码
```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var result []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return result[k-1]
}

func kthSmallest1(root *TreeNode, k int) (ans int) {
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		k--
		if k == 0 {
			ans = node.Val
		}
		inorder(node.Right)
	}
	inorder(root)
	return
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

## 238、除自身以外数组的乘积
### 题目
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

* 示例1：
>输入：nums = [1,2,3,4]
>输出：[24,12,8,6]

### 代码
```go
package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	ans := make([]int, len(nums))
	tmp := 1
	for i := 0; i < len(nums); i++ {
		prefix[i] = tmp
		tmp = tmp * nums[i]
	}
	tmp = 1
	for i := len(nums) - 1; i >= 0; i-- {
		suffix[i] = tmp
		tmp = tmp * nums[i]
	}
	for i := 0; i < len(nums); i++ {
		ans[i] = prefix[i] * suffix[i]
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4}
	ans := productExceptSelf(nums)
	fmt.Println(ans)
}
```

<hr>

## 240、搜索二维矩阵 II
### 题目
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

* 示例1：
>输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
>输出：true

### 代码
```go
package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := 0
	col := len(matrix[0]) - 1
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 15}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(searchMatrix(matrix, 5))
}
```

<hr>

## 279、完全平方数
### 题目
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

* 示例1：
>输入：n = 12
>输出：3
>解释：12 = 4 + 4 + 4

### 代码
```go
package main

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = n + 1
	}
	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
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

## 287、寻找重复数
### 题目
给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。  

假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。  

* 示例1：
>输入：nums = [1,3,4,2,2]
>输出：2

### 代码
```go
package main

import "fmt"

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}
```

<hr>

## 300、最长递增子序列
### 题目
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

* 示例1：
>输入：nums = [10,9,2,5,3,7,101,18]
>输出：4

### 代码
```go
package main

func lengthOfLIS(nums []int) int {
	result := 1
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(result, dp[i])
	}
	return result
}
```

<hr>

## 301、删除无效的括号
### 题目
给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。
返回所有可能的结果。答案可以按 任意顺序 返回。

* 示例1：
>输入：s = "()())()"
>输出：["(())()","()()()"]

### 代码
```go
package main

func removeInvalidParentheses(s string) []string {
	left, right := 0, 0 //要去除的 ( 和 )
	for _, char := range s {
		if char == '(' {
			left++
		} else if char == ')' {
			if left > 0 {
				left--
			} else {
				right++
			}
		}
	}
	var path []byte
	var result []string
	visited := make(map[string]bool)
	var backtracking func(index, leftCount, rightCount, leftRem, rightRem int)
	backtracking = func(index, leftCount, rightCount, leftRem, rightRem int) {
		if index == len(s) {
			if leftRem == 0 && rightRem == 0 {
				str := string(path)
				if !visited[str] {
					visited[str] = true
					result = append(result, str)
				}
			}
			return
		}
		char := s[index]
		if char == '(' && leftRem > 0 {
			backtracking(index+1, leftCount, rightCount, leftRem-1, rightRem)
		}
		if char == ')' && rightRem > 0 {
			backtracking(index+1, leftCount, rightCount, leftRem, rightRem-1)
		}
		path = append(path, char)
		if char == '(' {
			backtracking(index+1, leftCount+1, rightCount, leftRem, rightRem)
		} else if char == ')' {
			if rightCount < leftCount {
				backtracking(index+1, leftCount, rightCount+1, leftRem, rightRem)
			}
		} else {
			backtracking(index+1, leftCount, rightCount, leftRem, rightRem)
		}
		path = path[:len(path)-1]
	}
	backtracking(0, 0, 0, left, right)
	return result
}
```

<hr>

## 322、零钱兑换
### 题目
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。

* 示例1：
>输入：coins = [1, 2, 5], amount = 11
>输出：3

### 代码
```go
package main

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) // dp[i],金额为i所需的硬币数
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for _, num := range coins {
		for j := num; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-num]+1)
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
```

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

## 344、反转字符串
### 题目
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

* 示例1：
>输入：s = ["h","e","l","l","o"]
>输出：["o","l","l","e","h"]

### 代码
```go
func reverseString(s []byte)  {
    left, right := 0, len(s)-1
    for left < right {
        s[left], s[right] = s[right], s[left]
        left++
        right--
    }
}
```

<hr>

## 347、前 K 个高频元素
### 题目
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

* 示例1：
>输入：nums = [1,1,1,2,2,3], k = 2
>输出：[1,2]

### 代码
```go
// func topKFrequent(nums []int, k int) []int {
//     maps := make(map[int]int)
//     for _, v := range nums {
//         maps[v]++
//     }
//     var keys []int
//     for k := range maps {
//         keys = append(keys, k)
//     }
//     sort.Slice(keys, func(i, j int) bool {
//         return maps[keys[i]] > maps[keys[j]]
//     })
//     return keys[:k]
// }

func topKFrequent(nums []int, k int) []int {
    var ans []int
    maps := make(map[int]int)
    for _, v := range nums {
        maps[v]++
    }
    type kv struct {
        k int
        v int
    }
    var KV []kv
    for k, v := range maps {
        KV = append(KV, kv{k, v})
    }
    sort.Slice(KV, func(i, j int) bool {
        return KV[i].v > KV[j].v
    })
    for i := 0; i < k; i++ {
        ans = append(ans, KV[i].k)
    }
    return ans
}
```

<hr>

## 367、有效的完全平方数
### 题目
给你一个正整数 num 。如果 num 是一个完全平方数，则返回 true ，否则返回 false 。
完全平方数 是一个可以写成某个整数的平方的整数。换句话说，它可以写成某个整数和自身的乘积。
不能使用任何内置的库函数，如  sqrt 。

* 示例1：
>输入：num = 16
>输出：true

### 代码
```go
package main

func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left < right {
		mid := (right-left)/2 + left
		if num <= mid*mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left*left == num {
		return true
	}
	return false
}
```

<hr>

## 377、组合总和 Ⅳ
### 题目
给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素排列的个数。
题目数据保证答案符合 32 位整数范围。

* 示例1：
>输入：nums = [1,2,3], target = 4
>输出：7

### 代码
```go
package main

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] = dp[i] + dp[i-num]
			}
		}
	}
	return dp[target]
}
```

<hr>

## 394、字符串解码
### 题目
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

* 示例1：
>输入：s = "3[a]2[bc]"
>输出："aaabcbc"
* 示例2：
>输入：s = "3[a2[c]]"
>输出："accaccacc"

### 代码
```go
package main

import (
	"container/list"
	"fmt"
	"strings"
)

func decodeString(s string) string {
	stack1 := list.New()
	stack2 := list.New()
	tmp := 0
	res := ""
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			tmp = tmp*10 + int(ch-'0')
		} else if ch == '[' {
			stack1.PushBack(tmp)
			stack2.PushBack(res)
			tmp = 0
			res = ""
		} else if ch == ']' {
			count := stack1.Remove(stack1.Back()).(int)
			ans := stack2.Remove(stack2.Back()).(string)
			str := strings.Repeat(res, count)
			res = ans + str
		} else {
			res += string(ch)
		}
	}
	return res
}

func decodeString1(s string) string {
	stack1 := list.New()
	stack2 := list.New()
	tmp := 0
	current := ""
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			tmp = tmp*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			stack1.PushBack(tmp)
			stack2.PushBack(current)
			tmp = 0
			current = ""
		} else if s[i] == ']' {
			// 弹出重复次数
			countElem := stack1.Back()
			count := countElem.Value.(int)
			stack1.Remove(countElem)

			// 弹出之前的字符串
			strElem := stack2.Back()
			ans := strElem.Value.(string)
			stack2.Remove(strElem)

			// 重复 current count 次
			tmpstr := strings.Repeat(current, count) // 推荐用 strings.Repeat
			current = ans + tmpstr
		} else {
			current += string(s[i])
		}
	}
	return current
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func decodeString2(s string) string {
	stack1 := list.New()
	stack2 := list.New()
	tmp := 0
	var builder strings.Builder
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			tmp = tmp*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			stack1.PushBack(tmp)
			stack2.PushBack(builder.String())
			tmp = 0
			builder.Reset()
		} else if s[i] == ']' {
			count := stack1.Remove(stack1.Back()).(int)
			ans := stack2.Remove(stack2.Back()).(string)

			current := builder.String()
			builder.Reset()
			builder.WriteString(ans)
			for j := 0; j < count; j++ {
				builder.WriteString(current)
			}
		} else {
			builder.WriteByte(s[i])
		}
	}
	return builder.String()
}

func decodeString3(s string) string {
	stack1 := list.New()
	stack2 := list.New()
	tmp := 0
	var builder strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			tmp = tmp*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			stack1.PushBack(tmp)
			stack2.PushBack(builder.String())
			tmp = 0
			builder.Reset()
		} else if s[i] == ']' {
			count := stack1.Remove(stack1.Back()).(int)
			ans := stack2.Remove(stack2.Back()).(string)

			current := builder.String()
			builder.Reset()
			builder.WriteString(ans)
			for j := 0; j < count; j++ {
				builder.WriteString(current)
			}
		} else {
			builder.WriteByte(s[i])
		}
	}
	return builder.String()
}

func main() {
	s := "3[a2[c]]"
	fmt.Println(decodeString(s))
}
```

<hr>

## 416、分割等和子集
### 题目
给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

* 示例1：
>输入：nums = [1,5,11,5]
>输出：true

### 代码
```go
package main

// func canPartition(nums []int) bool {
//     sum := 0
//     for _, num := range nums {
//         sum += num
//     }
//     if sum % 2 != 0 {
//         return false
//     }
//     sum = sum / 2
//     dp := make([][]int, len(nums))
//     for i := range dp {
//         dp[i] = make([]int, sum+1)
//     }
//     for j := 0; j <= sum; j++ {
//         if j >= nums[0] {
//             dp[0][j] = nums[0]
//         }
//     }
//     for i := 1; i < len(nums); i++ {
//         for j := 0; j <= sum; j++ {
//             if j < nums[i] {
//                 dp[i][j] = dp[i-1][j]
//             } else {
//                 dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]] + nums[i])
//             }
//         }
//     }
//     return dp[len(nums)-1][sum] == sum
// }

// func canPartition(nums []int) bool {
//     sum := 0
//     for _, num := range nums {
//         sum += num
//     }
//     if sum % 2 != 0 {
//         return false
//     }
//     sum = sum / 2
//     dp := make([]int, sum+1)
//     for i := 0; i < len(nums); i++ {
//         for j := sum; j >= nums[i]; j-- {
//             dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
//         }
//     }
//     return dp[sum] == sum
// }

// func canPartition(nums []int) bool {
//     sum := 0
//     for _, num := range nums {
//         sum += num
//     }
//     if sum % 2 != 0 {
//         return false
//     }
//     sum = sum / 2
//     dp := make([][]bool, len(nums))
//     for i := range dp {
//         dp[i] = make([]bool, sum+1)
//         dp[i][0] = true 
//     }
//     if nums[0] <= sum {
//         dp[0][nums[0]] = true
//     }
//     for i := 1; i < len(nums); i++ {
//         for j := 0; j <= sum; j++ {
//             if j < nums[i] {
//                 dp[i][j] = dp[i-1][j]
//             } else {
//                 dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
//             }
//         }
//     }
//     return dp[len(nums)-1][sum]
// }

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum % 2 != 0 {
		return false
	}
	sum = sum / 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := sum; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[sum]
}

```

<hr>

## 455、分发饼干
### 题目
假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，
都有一个尺寸 s[j] 。如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。
你的目标是满足尽可能多的孩子，并输出这个最大数值。

* 示例1：
>输入：g = [1,2,3], s = [1,1]
>输出：1

### 代码
```go
package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	for _, cookie := range s {
		if count < len(g) && cookie >= g[count] {
			count++
		}
	}
	return count
}
```

<hr>

## 459、重复的子字符串
### 题目
给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。

* 示例1：
>输入：s = "abab"
>输出：true

### 代码
```go
package main

import "strings"

func repeatedSubstringPattern(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i != 0 {
			continue
		}
		flag := true
		sub := s[:i]
		for j := i; j < len(s); j += i {
			if sub != s[j:j+i] {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	return false
}

func repeatedSubstringPattern1(s string) bool {
	ss := s + s
	// 去掉首尾字符
	ss = ss[1 : len(ss)-1]
	// 判断 s 是否在 ss 中
	return strings.Contains(ss, s)
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

## 494、目标和
### 题目
给你一个非负整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

* 示例1：
>输入：nums = [1,1,1,1,1], target = 3
>输出：5

### 代码
```go
package main

func findTargetSumWays(nums []int, target int) int {
	// P N  P+N=S P-N=T P=(S+T)/2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if (sum+target)%2 != 0 || sum < abs(target) {
		return 0
	}
	ans := (sum + target) / 2
	dp := make([]int, ans+1)
	dp[0] = 1
	for _, num := range nums {
		for j := ans; j >= num; j-- {
			dp[j] = dp[j] + dp[j-num]
		}
	}
	return dp[ans]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

<hr>

## 509、斐波拉契数
### 题目
斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
F(0) = 0，F(1) = 1
F(n) = F(n - 1) + F(n - 2)，其中 n > 1
给定 n ，请计算 F(n) 。

* 示例1：
>输入：n = 2
>输出：1

### 代码
```go
package main

func fib(n int) int {
	// maps := make(map[int]int)
	// return fibDp1(n, maps)
	return fibDp3(n)
}

func fibDp(n int, maps map[int]int) int {
	if n <= 1 {
		return n
	}
	if v, ok := maps[n]; ok {
		return v
	}
	maps[n] = fibDp(n-1, maps) + fibDp(n-2, maps)
	return maps[n]
}

func fibDp2(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fibDp3(n int) int {
	if n <= 1 {
		return n
	}
	prev, cur := 0, 1
	for i := 2; i <= n; i++ {
		prev, cur = cur, prev+cur
	}
	return cur
}
```

<hr>

## 518、零钱兑换 II
### 题目
给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
假设每一种面额的硬币有无限个。
题目数据保证结果符合 32 位带符号整数。

* 示例1：
>输入：amount = 5, coins = [1, 2, 5]
>输出：4

### 代码
```go
package main

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] = dp[j] + dp[j-coin]
		}
	}
	return dp[amount]
}
```

<hr>

## 541、反转字符串 II
### 题目
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。
如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

* 示例1：
>输入：s = "abcdefg", k = 2
>输出："bacdfeg"

### 代码
```go
package main

func reverseStr(s string, k int) string {
	n := len(s)
	chars := []byte(s)
	right := 0
	for right < n {
		if n-right <= k {
			reverse(chars, right, n-1)
			break
		} else if n-right > k && n-right <= 2*k {
			reverse(chars, right, right+k-1)
			break
		} else {
			reverse(chars, right, right+k-1)
			right += 2 * k
		}
	}
	return string(chars)
}

func reverse(chars []byte, left, right int) {
	for left < right {
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}
}
```

<hr>

## 543、二叉树的直径
### 题目
给你一棵二叉树的根节点，返回该树的 直径 。
二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
两节点之间路径的 长度 由它们之间边数表示。

* 示例1：
>输入：root = [1,2,3,4,5]
>输出：3

### 代码
```go
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := depth(node.Left)
		r := depth(node.Right)
		ans = max(ans, l+r)
		return max(l, r) + 1
	}
	depth(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 3}
	fmt.Println(diameterOfBinaryTree(root))
}
```

<hr>

## 560、和为 K 的子数组
### 题目
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
子数组是数组中元素的连续非空序列。

* 示例1：
>输入：nums = [1,1,1], k = 2
>输出：2

### 代码
```go
package main

import "fmt"

func subarraySum(nums []int, k int) int {
	count, sum := 0, 0
	prefixMap := map[int]int{0: 1}
	for _, num := range nums {
		sum += num
		if tmp, ok := prefixMap[sum-k]; ok {
			count += tmp
		}
		prefixMap[sum]++
	}
	return count
}

func main() {
	var nums = []int{1, 2, 3}
	k := 3
	fmt.Println(subarraySum(nums, k))
}
```

<hr>

## 674、最长连续递增序列
### 题目
给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。

* 示例1：
>输入：nums = [1,3,5,4,7]
>输出：3

### 代码
```go
package main

// func findLengthOfLCIS(nums []int) int {
//     ans := 1
//     tmp := 1
//     for i := 1; i < len(nums); i++ {
//         for j := i; j > 0; j-- {
//             if nums[j] > nums[j-1] {
//                 tmp++
//             } else {
//                 break
//             }
//         }
//         ans = max(ans, tmp)
//         tmp = 1
//     }
//     return ans
// }

func findLengthOfLCIS(nums []int) int {
	ans := 1
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		ans = max(ans, dp[i])
	}
	return ans
}
```

<hr>

## 718、最长重复子数组
### 题目
给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。

* 示例1：
>输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
>输出：3

### 代码
```go
package main

func findLength(nums1 []int, nums2 []int) int {
	ans := 0
	m := len(nums1)
	n := len(nums2)
	dp := make([][]int, m+1) //dp[i][j]表示以nums1[i-1],nums2[j-1]结尾的最长重复子数组
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		dp[i][0] = 0
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = 0
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans
}
```

<hr>

## 739、每日温度
### 题目
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

* 示例1：
>输入：temperatures = [73,74,75,71,69,72,76,73]
>输出：[1,1,4,2,1,1,0,0]

### 代码
```go
package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	ans := dailyTemperatures(temperatures)
	fmt.Println(ans)
}
```

<hr>

## 763、跳跃游戏 Ⅱ
### 题目
给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。
例如，字符串 "ababcc" 能够被分为 ["abab", "cc"]，但类似 ["aba", "bcc"] 或 ["ab", "ab", "cc"] 的划分是非法的。
注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。
返回一个表示每个字符串片段的长度的列表。

* 示例1：
>输入：s = "ababcbacadefegdehijhklij"
>输出：[9,7,8]

### 代码
```go
package main

func partitionLabels(s string) []int {
	var ans []int
	last := [26]int{}
	for i, c := range s {
		last[c-'a'] = i
	}
	start, end := 0, 0
	for i, c := range s {
		if last[c-'a'] > end {
			end = last[c-'a']
		}
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}
	return ans
}
```

<hr>

## 844、比较含退格的字符串
### 题目
给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。
注意：如果对空文本输入退格字符，文本继续为空。

* 示例1：
>输入：s = "ab#c", t = "ad#c"
>输出：true

### 代码
```go
package main

func backspaceCompare1(s string, t string) bool {
	var ch1 []byte
	for i := range s {
		if s[i] != '#' {
			ch1 = append(ch1, s[i])
		} else if len(ch1) > 0 {
			ch1 = ch1[:len(ch1)-1]
		}
	}
	var ch2 []byte
	for i := range t {
		if t[i] != '#' {
			ch2 = append(ch2, t[i])
		} else if len(ch2) > 0 {
			ch2 = ch2[:len(ch2)-1]
		}
	}
	return string(ch1) == string(ch2)
}

func backspaceCompare2(s string, t string) bool {
	return helper(s) == helper(t)
}
func helper(s string) string {
	c := []byte(s)
	left := 0
	for right := 0; right < len(c); right++ {
		if c[right] != '#' {
			c[left] = c[right]
			left++
		} else {
			if left > 0 {
				left--
			}
		}
	}
	return string(c[:left])
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

<hr>

## 912、排序数组
### 题目
给你一个整数数组 nums，请你将该数组升序排列。
你必须在 不使用任何内置函数 的情况下解决问题，时间复杂度为 O(nlog(n))，并且空间复杂度尽可能小。

* 示例1：
>输入：nums = [5,2,3,1]
>输出：[1,2,3,5]

### 代码
```go
package main

import (
	"math/rand"
	"slices"
)

func sortArray(nums []int) []int {
	if slices.IsSorted(nums) {
		return nums
	}
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(nums, left, right)
	quickSort(nums, pivot+1, right)
	quickSort(nums, left, pivot-1)
}

func partition(nums []int, left, right int) int {
	index := rand.Intn(right-left+1) + left
	nums[left], nums[index] = nums[index], nums[left]
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[left], nums[i] = nums[i], nums[left]
	return i
}

func quickSort1(nums []int) {
	if len(nums) <= 1 {
		return
	}
	index := rand.Intn(len(nums))
	nums[0], nums[index] = nums[index], nums[0]
	pivot := nums[0]
	left, right := 1, len(nums) - 1
	for left <= right {
		if nums[left] <= pivot {
			left++
		} else if nums[right] >= pivot {
			right--
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	nums[0], nums[right] = nums[right], nums[0]
	quickSort1(nums[:right])
	quickSort1(nums[right+1:])
}
```

<hr>

## 977、有序数组的平方
### 题目
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

* 示例1：
>输入：nums = [-4,-1,0,3,10]
>输出：[0,1,9,16,100]

### 代码
```go
package main

func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	left, right := 0, len(nums)-1
	for pos := len(nums) - 1; pos >= 0; pos-- {
		if nums[left] > nums[right] {
			ans[pos] = nums[left]
			left++
		} else {
			ans[pos] = nums[right]
			right--
		}
	}
	return ans
}
```

<hr>

## 1049、最后一块石头的重量 II
### 题目
有一堆石头，用整数数组 stones 表示。其中 stones[i] 表示第 i 块石头的重量。
每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
如果 x == y，那么两块石头都会被完全粉碎；
如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。

* 示例1：
>输入：stones = [2,7,4,1,8,1]
>输出：1

### 代码
```go
package main

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, num := range stones {
		sum += num
	}
	target := sum / 2
	dp := make([]int, target+1)
	for _, num := range stones {
		for j := target; j >= num; j-- {
			dp[j] = max(dp[j], dp[j-num]+num)
		}
	}
	return sum - 2*dp[target]
}
```

<hr>

## 1143、最长公共子序列
### 题目
给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

* 示例1：
>输入：text1 = "abcde", text2 = "ace"
>输出：3

### 代码
```go
package main

func longestCommonSubsequence(text1 string, text2 string) int {
	ans := 0
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
			ans = max(ans, dp[i][j])
		}
	}
	return ans
}

//      a b c d e
//    a 1 1 1 1 1
//    c 1 1 2 2 2
//    e 1 1 2 2 3
```

<hr>

## LCR018、验证回文串
### 题目
给定一个字符串 s ，验证 s 是否是 回文串 ，只考虑字母和数字字符，可以忽略字母的大小写。
本题中，将空字符串定义为有效的 回文串 。

* 示例1：
>输入：s = "A man, a plan, a canal: Panama"
>输出：true

### 代码
````go
package main

import "strings"

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	s = strings.ToLower(s)
	for left < right {
		for left < right && !isVaild(s[left]) {
			left++
		}
		for left < right && !isVaild(s[right]) {
			right--
		}
		if left < right && s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isVaild(b byte) bool {
	return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}
````

<hr>

## LCR020、回文子串
### 题目
给定一个字符串 s ，请计算这个字符串中有多少个回文子字符串。
具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

* 示例1：
>输入：s = "abc"
>输出：3

### 代码
```go
package main

// func countSubstrings(s string) int {
//     ans := 0
//     for i := 0; i < len(s); i++ {
//         for j := i + 1; j <= len(s); j++ {
//             if isPalindrome(s[i:j]) {
//                 ans++
//             }
//         }
//     }
//     return ans
// }
// func isPalindrome(s string) bool {
//     left := 0
//     right := len(s) - 1
//     for left < right {
//         if s[left] != s[right] {
//             return false
//         }
//         left++
//         right--
//     }
//     return true
// }

// func countSubstrings(s string) int {
//     ans := 0
//     for i := 0; i < len(s); i++ {
//         for j := i; j < len(s); j++ {
//             if isPalindrome(s, i, j) {
//                 ans++
//             }
//         }
//     }
//     return ans
// }
// func isPalindrome(s string, left, right int) bool {
//     for left < right {
//         if s[left] != s[right] {
//             return false
//         }
//         left++
//         right--
//     }
//     return true
// }

func countSubstrings(s string) int {
	ans := 0
	expand := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
			ans++
		}
	}
	for i := 0; i < len(s); i++ {
		expand(i, i)
		expand(i, i+1)
	}
	return ans
}
```

<hr>
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
	var tmp []int
	for _, row := range matrix {
		tmp = append(tmp, row...)
	}
	left := 0
	right := len(tmp) - 1
	for left <= right {
		mid := (right-left)/2 + left
		if tmp[mid] == target {
			return true
		} else if tmp[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func searchMatrix1(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := (left + right) / 2
		row, col := mid/n, mid%n
		val := matrix[row][col]
		if val == target {
			return true
		} else if val > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
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
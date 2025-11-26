# Golang

## 交换两个数
```go
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
```
```go
nums[i], nums[j] = nums[j], nums[i]
```
如果nums是切片可以直接交换。

<hr>

## sort.Slice
```go
nums := []int{5, 2, 6, 3, 1}
sort.Slice(nums, func(i, j int) bool {
    return nums[i] < nums[j] // 升序
})
// 结果: [1, 2, 3, 4, 5]
```
```go
sort.Slice(nums, func(i, j int) bool {
    return nums[i] > nums[j] // 降序
})
// 结果: [5, 4, 3, 2, 1]
```
```go
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Alice", 25},
    {"Bob",   30},
    {"John",  25},
}
sort.Slice(people, func(i, j int) bool {
return people[i].Age < people[j].Age
})
// 结果: John(25), Alice(25), Bob(30)
```
```go
sort.Slice(people, func(i, j int) bool {
    if people[i].Age != people[j].Age {
        return people[i].Age < people[j].Age // 年龄升序
    }
    return people[i].Name < people[j].Name // 年龄相同时按名字升序
})
```

<hr>

## slices.SortFunc
```go
package main

import (
    "fmt"
    "slices"
)

func main() {
    words := []string{"banana", "apple", "cherry"}
    slices.SortFunc(words, func(a, b string) int {
        return len(a) - len(b) // 按长度升序
    })
    fmt.Println(words) // 输出: [apple banana cherry]
}
```
```go
package main

import (
    "fmt"
    "sort"
    "slices"
)

func main() {
    // 原始数据：表示区间 [start, end]
    intervals := [][]int{{3, 6}, {1, 4}, {2, 8}, {5, 7}}

    // 方法1: 使用 sort.Slice (Go 1.8+)
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    fmt.Println("sort.Slice:", intervals)

    // 重置数据（因为上面已修改原切片）
    intervals = [][]int{{3, 6}, {1, 4}, {2, 8}, {5, 7}}

    // 方法2: 使用 slices.SortFunc (Go 1.21+)
    slices.SortFunc(intervals, func(a, b []int) int {
        return a[0] - b[0]
    })
    fmt.Println("slices.SortFunc:", intervals)
}
```

<hr>

## 判断map为空
直接判断`len(maps) == 0`

<hr>

## 去除map所有元素
直接重新构建map，`maps = make(map[int]int)`

<hr>

## errors 包 (错误处理)
`errors.New(text)`	创建一个简单的错误。
`fmt.Errorf(...)`	创建一个格式化的错误（更常用）。

<hr>

## 栈
栈是一种后进先出（LIFO）的数据结构，Golang 标准库没有单独提供一个栈的类型，但可以使用切片或者双链表来模拟栈的功能，因为切片和双链表在尾部添加和删除元素都比较高效。
```go
package main

import (
    "fmt"
)

func main() {
    // 初始化一个空的整型栈 s
    var s []int
    // 向栈顶（切片末尾）添加元素
    s = append(s, 10)
    s = append(s, 20)
    s = append(s, 30)
    // 检查栈是否为空，输出：false
    fmt.Println(len(s) == 0)
    // 获取栈的大小，输出：3
    fmt.Println(len(s))
    // 获取栈顶元素，输出：30
    fmt.Println(s[len(s)-1])
    // 删除栈顶元素
    s = s[:len(s)-1]
    // 输出新的栈顶元素：20
    fmt.Println(s[len(s)-1])
}
```
```go
package main

import (
    "container/list"
    "fmt"
)

func main() {
    stack := list.New()
    stack.PushBack(1)
    stack.PushBack(2)
    stack.PushBack(3)

    for stack.Len() > 0 {
        e := stack.Back()
        fmt.Println(e.Value) // 3, 2, 1
        stack.Remove(e)
		//e := stack.Remove(stack.Back()).(int)  //Remove()会返回Value，返回的是值，不是节点！
    }
}
```

<hr>

## 队列
```go
package main

import (
    "container/list"
    "fmt"
)

func main() {
    // 初始化一个空的整型队列 q
    q := list.New()
    // 在队尾添加元素
    q.PushBack(10)
    q.PushBack(20)
    q.PushBack(30)
    // 检查队列是否为空，输出：false
    fmt.Println(q.Len() == 0)
    // 获取队列的大小，输出：3
    fmt.Println(q.Len())
    // 获取队列的队头元素
    // 输出：10
    front := q.Front().Value.(int)
    fmt.Println(front)
    // 删除队头元素
    q.Remove(q.Front())
    // 输出新的队头元素：20
    newFront := q.Front().Value.(int)
    fmt.Println(newFront)
}
```

<hr>

## 二维数组转一维
leetcode74

### 显示
```go
func searchMatrix(matrix [][]int, target int) bool {
	var tmp []int
	for _, row := range matrix {
		tmp = append(tmp, row...)
	}
	left := 0
	right := len(tmp) - 1
	for left <= right {
		mid := (right - left) / 2 + left
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
```

### 隐式
```go
func searchMatrix(matrix [][]int, target int) bool {
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
```

## sort.Search(n, f)
在区间 [0, n) 中查找第一个满足 f(i) == true 的最小索引 i。
leetcode74
```go
func searchMatrix(matrix [][]int, target int) bool {
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
```

## sort.SearchInts(slice, target)
在已排序的 []int 中找第一个 ≥ target 的位置，没找到返回切片长度。
等价于
```go
sort.Search(len(slice), func(i int) bool { return slice[i] >= target })
```

## 拼接字符、字符串
```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	chars := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	var ch strings.Builder
	for i := 0; i < len(chars); i++ {
		ch.WriteByte(chars[i])
	}
	fmt.Println(ch.String())
}
```
```go
var b strings.Builder
for i := 0; i < 5; i++ {
    b.WriteString("Go ")
}
fmt.Println(b.String()) // Go Go Go Go Go 
```

## 字符串重复k次
```go
strings.Repeat(current, count)
```

## 字符转为数字
```go
tmp = int(s[i] - '0')
```

## string和byte拼接
```go
str += string(ch)
```
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
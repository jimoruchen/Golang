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


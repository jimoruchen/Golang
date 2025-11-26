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
			tmp := strings.Repeat(res, count)
			res = ans + tmp
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

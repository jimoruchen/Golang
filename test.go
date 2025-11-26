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

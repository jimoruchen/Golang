package main

import "fmt"

func main() {
	var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	num := removeDuplicates(nums)
	fmt.Println(num)
}

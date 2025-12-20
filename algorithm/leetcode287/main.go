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

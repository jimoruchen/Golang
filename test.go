package main

import "fmt"

func main() {

	for i := 1; i <= 9; i++ {
		for j := i + 1; j <= 9; j++ {
			fmt.Println(i, j)
			if j == 5 {
				break
			}
		}
		fmt.Println("111")
	}
}

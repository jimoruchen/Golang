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

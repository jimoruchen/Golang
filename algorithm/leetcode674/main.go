package main

// func findLengthOfLCIS(nums []int) int {
//     ans := 1
//     tmp := 1
//     for i := 1; i < len(nums); i++ {
//         for j := i; j > 0; j-- {
//             if nums[j] > nums[j-1] {
//                 tmp++
//             } else {
//                 break
//             }
//         }
//         ans = max(ans, tmp)
//         tmp = 1
//     }
//     return ans
// }

func findLengthOfLCIS(nums []int) int {
	ans := 1
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		ans = max(ans, dp[i])
	}
	return ans
}

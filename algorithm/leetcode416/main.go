package main

// func canPartition(nums []int) bool {
//     // 1 5 12 3
//     sum := 0
//     for _, num := range nums {
//         sum += num
//     }
//     if sum % 2 != 0 {
//         return false
//     }
//     sum = sum / 2
//     dp := make([][]int, len(nums))
//     for i := range dp {
//         dp[i] = make([]int, sum+1)
//     }
//     for j := 0; j <= sum; j++ {
//         if j >= nums[0] {
//             dp[0][j] = nums[0]
//         }
//     }
//     for i := 1; i < len(nums); i++ {
//         for j := 0; j <= sum; j++ {
//             if j < nums[i] {
//                 dp[i][j] = dp[i-1][j]
//             } else {
//                 dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]] + nums[i])
//             }
//         }
//     }
//     return dp[len(nums)-1][sum] == sum
// }

// func canPartition(nums []int) bool {
//     // 1 5 12 3
//     sum := 0
//     for _, num := range nums {
//         sum += num
//     }
//     if sum % 2 != 0 {
//         return false
//     }
//     sum = sum / 2
//     dp := make([]int, sum+1)
//     for i := 1; i < len(nums); i++ {
//         for j := sum; j >= nums[i]; j-- {
//             dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
//         }
//     }
//     return dp[sum] == sum
// }

func canPartition(nums []int) bool {
	// 1 5 12 3
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	dp := make([]int, sum+1)
	for i := 1; i < len(nums); i++ {
		for j := sum; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[sum] == sum
}

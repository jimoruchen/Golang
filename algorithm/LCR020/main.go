package main

// func countSubstrings(s string) int {
//     ans := 0
//     for i := 0; i < len(s); i++ {
//         for j := i + 1; j <= len(s); j++ {
//             if isPalindrome(s[i:j]) {
//                 ans++
//             }
//         }
//     }
//     return ans
// }
// func isPalindrome(s string) bool {
//     left := 0
//     right := len(s) - 1
//     for left < right {
//         if s[left] != s[right] {
//             return false
//         }
//         left++
//         right--
//     }
//     return true
// }

// func countSubstrings(s string) int {
//     ans := 0
//     for i := 0; i < len(s); i++ {
//         for j := i; j < len(s); j++ {
//             if isPalindrome(s, i, j) {
//                 ans++
//             }
//         }
//     }
//     return ans
// }
// func isPalindrome(s string, left, right int) bool {
//     for left < right {
//         if s[left] != s[right] {
//             return false
//         }
//         left++
//         right--
//     }
//     return true
// }

func countSubstrings(s string) int {
	ans := 0
	expand := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
			ans++
		}
	}
	for i := 0; i < len(s); i++ {
		expand(i, i)
		expand(i, i+1)
	}
	return ans
}

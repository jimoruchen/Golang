package main

func reverseStr(s string, k int) string {
	n := len(s)
	chars := []byte(s)
	right := 0
	for right < n {
		if n-right <= k {
			reverse(chars, right, n-1)
			break
		} else if n-right > k && n-right <= 2*k {
			reverse(chars, right, right+k-1)
			break
		} else {
			reverse(chars, right, right+k-1)
			right += 2 * k
		}
	}
	return string(chars)
}

func reverse(chars []byte, left, right int) {
	for left < right {
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}
}

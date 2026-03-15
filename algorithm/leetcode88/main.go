package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, right := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[right] = nums1[i]
			right--
			i--
		} else {
			nums1[right] = nums2[j]
			right--
			j--
		}
	}
	for j >= 0 {
		nums1[right] = nums2[j]
		right--
		j--
	}
}

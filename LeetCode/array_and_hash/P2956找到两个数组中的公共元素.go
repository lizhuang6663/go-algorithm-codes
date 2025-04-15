package main

// 签到：哈希表

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	s1, s2 := map[int]bool{}, map[int]bool{}
	for _, x := range nums1 {
		s1[x] = true
	}
	for _, x := range nums2 {
		s2[x] = true
	}
	res := make([]int, 2)
	for _, x := range nums1 {
		if s2[x] {
			res[0]++
		}
	}
	for _, x := range nums2 {
		if s1[x] {
			res[1]++
		}
	}
	return res
}

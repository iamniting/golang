package main

func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	res := []int{}

	for _, n := range nums1 {
		m[n] = true
	}

	for _, n := range nums2 {
		if m[n] {
			res = append(res, n)
			m[n] = false
		}
	}
	return res
}

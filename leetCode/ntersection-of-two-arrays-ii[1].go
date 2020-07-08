package main

func intersect(nums1 []int, nums2 []int) []int {

	m := map[int]int{}
	res := []int{}

	for _, n := range nums1 {
		m[n]++
	}

	for _, n := range nums2 {
		if _, ok := m[n]; ok && m[n] > 0 {
			res = append(res, n)
			m[n]--
		}
	}

	return res
}

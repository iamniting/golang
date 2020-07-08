package main

func intersect(nums1 []int, nums2 []int) []int {

	Append := func(array *[]int, item int, n int) {
		for i := 0; i < n; i++ {
			*array = append(*array, item)
		}
	}

	m1, m2 := map[int]int{}, map[int]int{}

	for _, n := range nums1 {
		m1[n]++
	}

	for _, n := range nums2 {
		m2[n]++
	}

	res := []int{}

	for key := range m1 {
		if _, ok := m2[key]; ok && m1[key] <= m2[key] {
			Append(&res, key, m1[key])
		} else if _, ok := m2[key]; ok && m1[key] > m2[key] {
			Append(&res, key, m2[key])
		}
	}

	return res
}

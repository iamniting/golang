package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	res := make([]int, len(nums1))

	for i, n1 := range nums1 {
		for j, n2 := range nums2 {

			if n1 != n2 {
				continue
			}

			res[i] = -1

			for _, n := range nums2[j:] {
				if n > n2 {
					res[i] = n
					break
				}
			}
			break
		}
	}

	return res
}

package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i := 0
	j := 0
	tmp := []int{}

	// merge sort
	for {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				tmp = append(tmp, nums1[i])
				i++
			} else {
				tmp = append(tmp, nums2[j])
				j++
			}
		} else if i < len(nums1) || j < len(nums2) {
			if i < len(nums1) {
				tmp = append(tmp, nums1[i])
				i++
			} else {
				tmp = append(tmp, nums2[j])
				j++
			}
		} else {
			break
		}
	}

	return float64(tmp[len(tmp)>>1]+tmp[(len(tmp)-1)>>1]) / 2.0
}

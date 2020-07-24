package main

func containsNearbyDuplicate(nums []int, k int) bool {

	m := map[int]int{}

	for i, n := range nums {

		// if n exist in map and diff b/w index is <= k return true
		if index, ok := m[n]; ok && i-index <= k {
			return true
		}
		// else put the new index of n to make the diff shorter
		m[n] = i
	}

	return false
}

package main

func uniqueOccurrences(arr []int) bool {

	// get the occurrence count
	m := make(map[int]int)

	for _, n := range arr {
		m[n]++
	}

	// get the values in the set
	set := make(map[int]bool)

	for _, v := range m {
		set[v] = true
	}

	return len(set) == len(m)
}

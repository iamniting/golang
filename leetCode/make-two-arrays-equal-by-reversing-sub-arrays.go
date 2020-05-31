package main

func canBeEqual(target []int, arr []int) bool {

	m := map[int]bool{}

	for _, n := range target {
		m[n] = true
	}

	for _, n := range arr {
		if !m[n] {
			return false
		}
	}

	return true
}

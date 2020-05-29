package main

func repeatedNTimes(A []int) int {
	m := make(map[int]bool)

	for _, n := range A {
		if m[n] {
			return n
		}
		m[n] = true
	}
	return 0
}

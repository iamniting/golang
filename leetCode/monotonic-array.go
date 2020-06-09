package main

func isMonotonic(A []int) bool {
	inc, dec := false, false
	prev := A[0]

	for _, n := range A {
		inc = inc || (prev < n)
		dec = dec || (prev > n)
		prev = n

		if inc && dec {
			return false
		}
	}
	return !(inc && dec)
}

package main

func validMountainArray(A []int) bool {

	var incr, decr bool

	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] && !decr {
			incr = true
		} else if A[i] < A[i-1] && incr {
			decr = true
		} else {
			return false
		}
	}
	return incr && decr
}

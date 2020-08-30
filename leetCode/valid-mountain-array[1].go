package main

func validMountainArray(A []int) bool {

	i, n := 1, len(A)

	// increasing
	for ; i < n && A[i] > A[i-1]; i++ {
	}

	// peak can not be first or last
	if i == 1 || i == n {
		return false
	}

	// decreasing
	for ; i < n && A[i] < A[i-1]; i++ {
	}

	return i == n
}

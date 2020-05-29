package main

func sortArrayByParityII(A []int) []int {

	i, j := 0, 1

	for i < len(A) {
		if A[i]&1 == 0 {
			i += 2
		} else {
			A[i], A[j] = A[j], A[i]
			j += 2
		}
	}

	return A
}

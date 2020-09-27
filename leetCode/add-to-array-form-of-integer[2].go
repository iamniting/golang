package main

func addToArrayForm(A []int, K int) []int {

	for i := len(A) - 1; i > -1; i-- {
		sum := A[i] + K
		A[i] = sum % 10
		K = sum / 10
	}

	for ; K > 0; K /= 10 {
		A = append([]int{K % 10}, A...)
	}

	return A
}

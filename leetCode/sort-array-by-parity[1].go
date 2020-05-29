package main

func sortArrayByParity(A []int) []int {
	low, high := 0, len(A)-1

	for low <= high {
		if A[low]&1 == 0 {
			low++
			continue
		}

		if A[high]&1 == 1 {
			high--
			continue
		}

		A[low], A[high] = A[high], A[low]
	}
	return A
}

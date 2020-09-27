package main

func addToArrayForm(A []int, K int) []int {

	var carry int

	for i := len(A) - 1; i >= 0 || K > 0 || carry == 1; {
		sum := carry

		if K > 0 {
			sum += K % 10
			K /= 10
		}

		if i >= 0 {
			sum += A[i]
			A[i] = sum % 10
			i--
		} else {
			A = append([]int{sum % 10}, A...)
		}

		carry = sum / 10
	}

	return A
}

package main

func addToArrayForm(A []int, K int) []int {

	var carry int
	var res []int

	for i := len(A) - 1; i >= 0 || K > 0 || carry == 1; i, K = i-1, K/10 {
		sum := carry + (K % 10)
		carry = 0

		if i >= 0 {
			sum += A[i]
		}

		if sum > 9 {
			carry = 1
			sum = sum % 10
		}

		res = append([]int{sum}, res...)
	}

	return res
}

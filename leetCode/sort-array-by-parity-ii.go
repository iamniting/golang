// https://leetcode.com/problems/sort-array-by-parity-ii
// Just sol to the problem, It does not include the I/O part

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

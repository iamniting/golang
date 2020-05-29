// https://www.geeksforgeeks.org/sort-array-converting-elements-squares
// https://www.youtube.com/watch?v=4eWKHLSRHPY

package main

import "math"

func sortedSquares(A []int) []int {
	// left most index
	i := 0
	// right most index
	j := len(A) - 1
	// result array right most index
	length := j
	// result array
	res := make([]int, j+1)

	// iterate both the pointer untill they meet
	for i <= j {
		// if left index is greater than right, then add A[i] in res
		if math.Abs(float64(A[i])) > math.Abs(float64(A[j])) {
			res[length] = A[i] * A[i]
			i++
			// if right index is greater then left, then add A[j] in res
		} else {
			res[length] = A[j] * A[j]
			j--
		}
		length--
	}

	return res
}

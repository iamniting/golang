package main

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	res := make([]int, len(queries))
	sum := 0

	// calculate sum of even integers
	for _, n := range A {
		if n&1 == 0 {
			sum += n
		}
	}

	for i, q := range queries {
		val, index := q[0], q[1]

		old := A[index]
		A[index] += val

		if old&1 == 0 && val&1 == 0 {
			sum += val

		} else if old&1 == 1 && val&1 == 1 {
			sum += old + val

		} else if old&1 == 0 && val&1 == 1 {
			sum -= old
		}

		res[i] = sum
	}
	return res
}

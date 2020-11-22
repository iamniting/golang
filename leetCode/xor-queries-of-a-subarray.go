package main

func xorQueries(arr []int, queries [][]int) []int {

	res := make([]int, len(queries))

	for i, q := range queries {

		for j := q[0]; j <= q[1]; j++ {
			res[i] ^= arr[j]
		}
	}

	return res
}

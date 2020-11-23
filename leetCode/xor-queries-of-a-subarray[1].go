package main

func xorQueries(arr []int, queries [][]int) []int {

	/*`
	If the question is to get the subarray sum of an array, we use prefix sum.
	Here the question is to get the subarray XOR, so we use prefix xor.
	`*/

	prefix := make([]int, len(arr))
	res := make([]int, len(queries))

	prefix[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		prefix[i] = prefix[i-1] ^ arr[i]
	}

	for i, q := range queries {
		l, r := q[0], q[1]
		if l == 0 {
			res[i] = prefix[r]
			continue
		}
		res[i] = prefix[l-1] ^ prefix[r]
	}

	return res
}

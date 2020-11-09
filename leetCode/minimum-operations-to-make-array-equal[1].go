package main

func minOperations(n int) int {

	var res int

	// n is target which needs to be achieved
	for i := 0; i < n/2; i++ {
		res += n - ((2 * i) + 1)
	}

	return res
}

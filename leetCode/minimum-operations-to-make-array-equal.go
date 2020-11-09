package main

func minOperations(n int) int {

	var sum, target, res int
	arr := make([]int, n)

	// Build the array arr using the given formula, define target = sum(arr) / n
	for i := range arr {
		arr[i] = (2 * i) + 1
		sum += arr[i]
	}

	target = sum / n

	// number of operations needed to convert arr so that all elements equal target
	for i := 0; i < n/2; i++ {
		res += target - arr[i]
	}

	return res
}

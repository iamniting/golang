package main

import "fmt"

func maxSubarraySum(arr []int) int {
	maxEndingHere := arr[0]
	maxSoFar := arr[0]
	for i := 1; i < len(arr); i++ {
		// maxEndingHere stores the maximum sum contiguous subarray ending at the current index.
		maxEndingHere = maxEndingHere + arr[i]
		// maxSoFar stores the maximum sum of contiguous subarray found so far.
		maxSoFar = max(maxSoFar, maxEndingHere)

		// if maxEndingHere is negative
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
	}
	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Find the subarray with the maximum sum in an array of integers
	arr := []int{-2, 1, -3, 4, -1, 6, -1, 2, 1, -5, 4}
	maxSum := maxSubarraySum(arr)
	fmt.Println("Max subarray sum:", maxSum)
}

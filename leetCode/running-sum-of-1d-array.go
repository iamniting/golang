package main

func runningSum(nums []int) []int {
	res := make([]int, len(nums))
	sum := 0

	for i, n := range nums {
		sum += n
		res[i] = sum
	}
	return res
}

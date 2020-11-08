package main

func arrangeCoins(n int) int {

	left, right := 0, n

	for left <= right {

		mid := (left + right) >> 1

		// sum of n natural ints
		sum := (mid * (mid + 1)) / 2

		if sum == n {
			return mid
		} else if sum > n {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left - 1 // or right
}

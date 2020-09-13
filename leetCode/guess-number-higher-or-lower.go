package main

func guessNumber(n int) int {

	left, right := 1, n

	for left <= right {

		mid := (left + right) / 2
		val := guess(mid)

		if val == 0 {
			return mid
		} else if val == 1 {
			left = mid + 1
		} else if val == -1 {
			right = mid - 1
		}
	}

	return 0
}

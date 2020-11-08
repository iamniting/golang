package main

func arrangeCoins(n int) int {

	for i := 2; i <= n; i++ {
		sum := (i * (i + 1)) / 2

		if sum == n {
			return i
		} else if sum > n {
			return i - 1
		}
	}

	return n
}

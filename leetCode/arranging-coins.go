package main

func arrangeCoins(n int) int {

	sum, i := 0, 1

	for sum <= n {
		sum += i
		i++
	}

	return i - 2
}

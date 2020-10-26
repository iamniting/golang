package main

func dailyTemperatures(T []int) []int {

	res := make([]int, len(T))

	for i, n := range T {
		for j := i + 1; j < len(T); j++ {
			if T[j] > n {
				res[i] = j - i
				break
			}
		}
	}

	return res
}

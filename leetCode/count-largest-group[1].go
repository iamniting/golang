package main

func countLargestGroup(n int) int {
	m := map[int]int{}
	maxCount, res := 0, 0

	for i := 1; i <= n; i++ {
		sum := 0

		for n := i; n > 0; n /= 10 {
			sum += n % 10
		}

		m[sum]++

		if m[sum] > maxCount {
			maxCount = m[sum]
			res = 0
		}
		if m[sum] == maxCount {
			res++
		}
	}

	return res
}

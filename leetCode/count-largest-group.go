package main

func countLargestGroup(n int) int {
	m := map[int][]int{}
	maxLen, res := 0, 0

	for i := 1; i <= n; i++ {
		sum := 0

		for n := i; n > 0; n /= 10 {
			sum += n % 10
		}

		m[sum] = append(m[sum], i)

		if len(m[sum]) > maxLen {
			maxLen = len(m[sum])
			res = 0
		}
		if len(m[sum]) == maxLen {
			res++
		}
	}

	return res
}

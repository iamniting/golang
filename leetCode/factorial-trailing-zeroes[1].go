package main

func trailingZeroes(n int) int {

	var res int
	dp := make([]int, n+1)

	for i := 5; i <= n; i += 5 {

		for j := i; j%5 == 0; j /= 5 {
			if dp[j] != 0 {
				dp[i] += dp[j]
				break
			}
			dp[i]++
		}

		res += dp[i]
	}

	return res
}

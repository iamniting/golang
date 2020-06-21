package main

func maxProfit(prices []int) int {
	var res int

	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

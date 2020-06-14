package main

func finalPrices(prices []int) []int {
	res := make([]int, len(prices))

	for i, p := range prices {
		res[i] = p

		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= p {
				res[i] = p - prices[j]
				break
			}
		}
	}
	return res
}

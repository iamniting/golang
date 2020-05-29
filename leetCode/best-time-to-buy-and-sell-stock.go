package main

import "math"

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	max, min := 0, prices[0]

	for _, p := range prices {
		if p < min {
			min = p
		} else {
			cur := p - min
			max = int(math.Max(float64(max), float64(cur)))
		}
	}

	return max
}

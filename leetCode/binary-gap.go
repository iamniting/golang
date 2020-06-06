package main

import "math"

func binaryGap(N int) int {
	max, cnt := 0, math.MinInt32
	for ; N > 0; N >>= 1 {
		if N&1 == 1 {
			max = int(math.Max(float64(max), float64(cnt)))
			cnt = 0
		}
		cnt++
	}
	return max
}

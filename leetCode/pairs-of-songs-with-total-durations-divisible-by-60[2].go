package main

func numPairsDivisibleBy60(time []int) int {

	var res int
	var cnt [60]int

	for _, t := range time {

		remainder := t % 60
		// % 60 again to stop overflow
		// eg. remainder is 0 then 60 - 0 will be 60 which will cause overflow
		res += cnt[(60-remainder)%60]
		cnt[remainder]++
	}

	return res
}

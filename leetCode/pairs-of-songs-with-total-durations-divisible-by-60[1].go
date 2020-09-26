package main

func numPairsDivisibleBy60(time []int) int {

	var res int
	var cnt [60]int

	for _, t := range time {

		remainder := t % 60
		var required int

		if remainder > 0 {
			required = 60 - remainder
		}

		res += cnt[required]
		cnt[remainder]++
	}

	return res
}

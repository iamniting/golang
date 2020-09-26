package main

func numPairsDivisibleBy60(time []int) int {

	var res int

	for i := range time {
		for j := i + 1; j < len(time); j++ {
			if (time[i]+time[j])%60 == 0 {
				res++
			}
		}
	}

	return res
}

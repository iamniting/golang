package main

func maximumWealth(accounts [][]int) int {

	var res int

	for _, row := range accounts {
		var sum int

		for _, n := range row {
			sum += n
		}

		if sum > res {
			res = sum
		}
	}

	return res
}

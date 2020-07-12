package main

func numIdenticalPairs(nums []int) int {
	var res int
	m := map[int]int{}

	for _, n := range nums {
		m[n]++
	}

	for _, v := range m {
		if v > 1 {
			res += (v * (v - 1)) / 2
		}
	}

	return res
}

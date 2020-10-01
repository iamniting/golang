package main

func findPairs(nums []int, k int) int {

	var res int
	m := map[int]int{}

	for _, n := range nums {
		m[n]++
	}

	for n, count := range m {
		if k > 0 && m[k+n] != 0 {
			res++
		} else if k == 0 && count > 1 {
			res++
		}
	}

	return res
}

package main

func largeGroupPositions(S string) [][]int {

	var res [][]int
	var prev rune
	var prevIdx int

	for i, c := range S + "1" {
		if c != prev {
			if i-prevIdx > 2 {
				res = append(res, []int{prevIdx, i - 1})
			}
			prev = c
			prevIdx = i
		}
	}
	return res
}

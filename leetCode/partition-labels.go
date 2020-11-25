package main

import "math"

func partitionLabels(S string) []int {

	var res []int
	var last [26]int

	for i, c := range S {
		last[c-'a'] = i
	}

	var start, end int
	for i, c := range S {
		end = int(math.Max(float64(end), float64(last[c-'a'])))

		if end == i {
			res = append(res, end-start+1)
			start = i + 1
		}
	}

	return res
}

package main

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	type strong struct {
		row, sol int
	}
	tmp := make([]strong, len(mat))
	res := make([]int, k)

	for i, row := range mat {
		sol := 0
		for _, col := range row {
			sol += col
		}
		tmp[i] = strong{i, sol}
	}

	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].sol == tmp[j].sol {
			return tmp[i].row < tmp[j].row
		}
		return tmp[i].sol < tmp[j].sol
	})

	i := 0
	for k > 0 {
		res[i] = tmp[i].row
		i++
		k--
	}
	return res
}

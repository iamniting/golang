package main

import "sort"

func getKth(lo int, hi int, k int) int {

	type numPower struct {
		num, power int
	}

	var slice []numPower

	for i := lo; i <= hi; i++ {
		slice = append(slice, numPower{i, getPower(i)})
	}

	sort.Slice(slice, func(i, j int) bool {
		if slice[i].power == slice[j].power {
			return slice[i].num < slice[j].num
		}
		return slice[i].power < slice[j].power
	})

	return slice[k-1].num
}

func getPower(num int) int {
	var steps int

	for num != 1 {
		if num&1 == 0 {
			num = num / 2
		} else {
			num = 3*num + 1
		}
		steps++
	}

	return steps
}

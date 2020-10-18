package main

import "sort"

func getKth(lo int, hi int, k int) int {

	type numPower struct {
		num, power int
	}

	var slice []numPower
	m := map[int]int{}

	for i := lo; i <= hi; i++ {
		slice = append(slice, numPower{i, getPower(i, m)})
	}

	sort.Slice(slice, func(i, j int) bool {
		if slice[i].power == slice[j].power {
			return slice[i].num < slice[j].num
		}
		return slice[i].power < slice[j].power
	})

	return slice[k-1].num
}

func getPower(num int, m map[int]int) int {
	var steps int
	n := num

	for num != 1 && m[num] == 0 {
		if num&1 == 0 {
			num = num / 2
		} else {
			num = 3*num + 1
		}
		steps++
	}

	if num != 1 {
		m[n] = steps + m[num]
		return m[n]
	}

	return steps
}

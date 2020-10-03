package main

import "math"

func powerfulIntegers(x int, y int, bound int) []int {

	var res []int
	m := map[int]bool{}

	for i := 0; i < 25 && int(math.Pow(float64(x), float64(i))) <= bound; i++ {
		for j := 0; j < 25 && int(math.Pow(float64(y), float64(j))) <= bound; j++ {

			val := int(math.Pow(float64(x), float64(i))) +
				int(math.Pow(float64(y), float64(j)))
			if val <= bound && !m[val] {
				m[val] = true
				res = append(res, val)
			}
		}
	}

	return res
}

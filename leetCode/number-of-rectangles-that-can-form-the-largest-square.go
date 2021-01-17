package main

import "math"

func countGoodRectangles(rectangles [][]int) int {

	var square, max, res int

	for _, r := range rectangles {
		square = int(math.Min(float64(r[0]), float64(r[1])))

		if square > max {
			max = square
			res = 1
		} else if square == max {
			res++
		}
	}

	return res
}

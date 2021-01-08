package main

import "math"

func leastBricks(wall [][]int) int {

	var max int
	mymap := map[int]int{}

	for _, row := range wall {
		var sum int

		for _, c := range row[:len(row)-1] {
			sum += c
			mymap[sum]++
			max = int(math.Max(float64(max), float64(mymap[sum])))
		}
	}

	return len(wall) - max
}

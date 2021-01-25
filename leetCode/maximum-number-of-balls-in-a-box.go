package main

import "math"

func countBalls(lowLimit int, highLimit int) int {

	var res int
	mymap := map[int]int{}

	for i := lowLimit; i <= highLimit; i++ {

		var sum int

		for n := i; n > 0; n /= 10 {
			sum += n % 10
		}

		mymap[sum]++

		res = int(math.Max(float64(res), float64(mymap[sum])))
	}

	return res
}

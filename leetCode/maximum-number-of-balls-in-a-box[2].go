package main

import "math"

func countBalls(lowLimit int, highLimit int) int {

	/*
		With the increase of ball number, the corresponding box id will
		generally increase till encountering those ball numbers with trailing 0s,
		in which case the box id will decrease by 9 for each trailing 0. e.g.,

		ball # --- box #
		9 --------> 9
		10 ------> 1

		99 -----> 18
		100 -----> 1
	*/

	res, sum := 1, 0
	var freq [46]int

	// compute box id for lowLimit.
	for n := lowLimit; n > 0; n /= 10 {
		sum += n % 10
	}

	freq[sum]++

	// compute all other box ids.
	for i := lowLimit + 1; i <= highLimit; i++ {

		// for ball numbers with trailing 0's, decrease 9 for each 0.
		for n := i; n%10 == 0; n /= 10 {
			sum -= 9
		}

		sum++
		freq[sum]++

		res = int(math.Max(float64(res), float64(freq[sum])))
	}

	return res
}

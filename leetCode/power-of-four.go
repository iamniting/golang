package main

import "math"

func isPowerOfFour(num int) bool {

	for i := 1; i <= math.MaxInt32; i <<= 2 {
		if i == num {
			return true
		}
	}

	return false
}

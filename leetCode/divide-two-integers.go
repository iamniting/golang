package main

import "math"

func divide(dividend int, divisor int) int {

	sign1, sign2 := false, false
	if dividend < 0 {
		sign1 = true
	}
	if divisor < 0 {
		sign2 = true
	}

	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))

	res := 0

	for dividend >= divisor {
		dividend -= divisor
		res++
	}

	if !sign1 && sign2 {
		res = -res
	} else if sign1 && !sign2 {
		res = -res
	}

	if res < math.MinInt32 {
		return math.MinInt32
	}
	if res > math.MaxInt32 {
		return math.MaxInt32
	}

	return res
}

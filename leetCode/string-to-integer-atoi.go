package main

import "math"

func myAtoi(str string) int {
	res := 0
	sign := 1
	start := false

	for _, s := range str {

		if s >= '0' && s <= '9' {
			res = (res * 10) + int(s-'0')
			start = true
			if res > math.MaxInt32 {
				break
			}
		} else if !start && s == '-' {
			start = true
			sign = -1
		} else if !start && s == '+' {
			start = true
			sign = +1
		} else if !start && s == ' ' {
			continue
		} else {
			break
		}
	}

	res = res * sign

	if res > math.MaxInt32 {
		return math.MaxInt32
	}

	if res < math.MinInt32 {
		return math.MinInt32
	}

	return res
}

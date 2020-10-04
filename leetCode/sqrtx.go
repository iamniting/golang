package main

func mySqrt(x int) int {

	for i := 1; i <= x; i++ {
		if val := i * i; val == x {
			return i
		} else if val > x {
			return i - 1
		}
	}

	return 0
}

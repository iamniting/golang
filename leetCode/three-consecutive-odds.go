package main

func threeConsecutiveOdds(arr []int) bool {

	var odds int

	for _, n := range arr {
		if n&1 == 1 {
			odds++
		} else {
			odds = 0
		}

		if odds == 3 {
			return true
		}
	}
	return false
}

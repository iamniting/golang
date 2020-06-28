package main

func lemonadeChange(bills []int) bool {
	var five, ten int

	for _, b := range bills {

		if b == 5 {
			five++
		} else if b == 10 && five > 0 {
			five, ten = five-1, ten+1
		} else if b == 10 && five < 1 {
			return false
		} else if b == 20 && five > 0 && ten > 0 {
			five, ten = five-1, ten-1
		} else if b == 20 && five > 2 {
			five -= 3
		} else {
			return false
		}
	}

	return true
}

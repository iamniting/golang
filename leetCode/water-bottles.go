package main

func numWaterBottles(numBottles int, numExchange int) int {

	var res, emptyBottles int

	for numBottles > 0 || emptyBottles >= numExchange {

		res += numBottles
		emptyBottles += numBottles
		numBottles = emptyBottles / numExchange
		emptyBottles %= numExchange
	}
	return res
}

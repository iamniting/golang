package main

func totalMoney(n int) int {

	weeks := n / 7
	days := n % 7

	/*
		eg. n = 34
		weeks = 4
		days = 5
		[1, 2, 3, 4, 5, 6, 7] = 28
		[2, 3, 4, 5, 6, 7, 8] = 35
		[3, 4, 5, 6, 7, 8, 9] = 42
		[4, 5, 6, 7, 8, 9, 10] = 49
		[5, 6, 7, 8, 9, 10] = 45

		calculate sum of weeks
		28 + 35 + 42 + 49

		(4 * 28) + (0 + 7 + 14 + 21)
		(4 * 28) + (7 * (0 + 1 + 2 + 3))
	*/
	weeksMoney := (28 * weeks) + (7 * (weeks * (weeks - 1)) / 2)

	/*
		calculate sum of days
		5 + 6 + 7 + 8 + 9 + 10

		(5 * 6) + (0 + 1 + 2 + 3 + 4 + 5)
	*/
	daysMoney := ((weeks + 1) * days) + ((days * (days - 1)) / 2)

	return weeksMoney + daysMoney
}

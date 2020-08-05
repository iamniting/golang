package main

import "strconv"

func dayOfYear(date string) int {

	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])

	res := 0

	for _, n := range days[:month-1] {
		res += n
	}

	if (month > 2) && (year%4 == 0 && year%100 != 0 || year%400 == 0) {
		res++
	}

	return res + day
}

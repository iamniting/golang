package main

import "strconv"

func daysBetweenDates(date1 string, date2 string) int {

	isLeapYear := func(year int) bool {
		return year%4 == 0 && year%100 != 0 || year%400 == 0
	}

	if date1 > date2 {
		date1, date2 = date2, date1
	}

	days := []int{31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334, 365}

	// parse date1
	year1, _ := strconv.Atoi(date1[:4])
	month1, _ := strconv.Atoi(date1[5:7])
	day1, _ := strconv.Atoi(date1[8:])

	// parse date2
	year2, _ := strconv.Atoi(date2[:4])
	month2, _ := strconv.Atoi(date2[5:7])
	day2, _ := strconv.Atoi(date2[8:])

	var res int

	// calculate days b/w years including starting year excluding last year
	for year := year1; year < year2; year++ {
		res += 365
		if isLeapYear(year) {
			res++
		}
	}

	// remove extra added days of starting year
	if month1 > 1 {
		res -= days[month1-2]
	}
	res -= day1

	if isLeapYear(year1) && month1 > 2 {
		res--
	}

	// add days of last year
	if month2 > 1 {
		res += days[month2-2]
	}
	res += day2

	if isLeapYear(year2) && month2 > 2 {
		res++
	}

	return res
}

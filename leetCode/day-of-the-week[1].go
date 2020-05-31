package main

func dayOfTheWeek(day int, month int, year int) string {
	// starting day is friday on 1 Jan 1971
	week := []string{
		"Thursday", "Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday",
	}
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	isLeapYear := func(year int) bool {
		return year%4 == 0 && year%100 != 0 || year%400 == 0
	}

	days := 0

	// days in years except current year
	for i := 1971; i < year; i++ {
		days += 365
		if isLeapYear(i) {
			days++
		}
	}

	// days in months in current year except current month
	for i := 1; i < month; i++ {
		days += months[i-1]
	}

	// if current year is leap year
	if month > 2 && isLeapYear(year) {
		days++
	}

	// days in current month
	days += day

	days %= 7

	return week[days]
}

package main

import (
	"strconv"
	"time"
)

func dayOfTheWeek(day int, month int, year int) string {
	layout := "2/1/2006"
	date := strconv.Itoa(day) + "/" + strconv.Itoa(month) + "/" + strconv.Itoa(year)
	t, _ := time.Parse(layout, date)

	return t.Weekday().String()
}

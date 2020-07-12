package main

import "strings"

func reformatDate(date string) string {

	months := map[string]string{
		"Jan": "01", "Feb": "02", "Mar": "03", "Apr": "04",
		"May": "05", "Jun": "06", "Jul": "07", "Aug": "08",
		"Sep": "09", "Oct": "10", "Nov": "11", "Dec": "12"}

	d := strings.Fields(date)

	res := d[2] + "-" + months[d[1]] + "-"

	if len(d[0]) == 3 {
		res += "0" + d[0][0:1]
	} else {
		res += d[0][0:2]
	}

	return res
}

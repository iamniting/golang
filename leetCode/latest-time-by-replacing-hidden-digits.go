package main

func maximumTime(time string) string {

	var res string
	max := "23:59"

	for i := range time {

		if time[i] != '?' {
			res += time[i : i+1]
		} else if i == 0 && time[1] != '?' && time[1] > '3' {
			res += "1"
		} else if i == 1 && time[0] != '?' && time[0] < '2' {
			res += "9"
		} else {
			res += max[i : i+1]
		}
	}

	return res
}

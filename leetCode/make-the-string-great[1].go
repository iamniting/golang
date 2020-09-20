package main

import "math"

func makeGood(s string) string {

	j := -1
	str := []byte(s)

	for i := 0; i < len(str); i++ {
		if j > -1 && int(math.Abs(float64(str[j])-float64(str[i]))) == 32 {
			j--
		} else {
			j++
			str[j] = str[i]
		}
	}

	return string(str[:j+1])
}

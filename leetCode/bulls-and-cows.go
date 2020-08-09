package main

import (
	"math"
	"strconv"
)

func getHint(secret string, guess string) string {

	var bull, cow int
	arr1, arr2 := [10]int{}, [10]int{}

	for i := range secret {
		arr1[secret[i]-'0']++
		arr2[guess[i]-'0']++

		if secret[i] == guess[i] {
			bull++
		}
	}

	for i := range arr1 {
		cow += int(math.Min(float64(arr1[i]), float64(arr2[i])))
	}

	return strconv.Itoa(bull) + "A" + strconv.Itoa(cow-bull) + "B"
}

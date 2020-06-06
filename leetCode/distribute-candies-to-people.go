package main

func distributeCandies(candies int, numPeople int) []int {
	res := make([]int, numPeople)

	c, index := 1, 0

	for candies > 0 {
		if c <= candies {
			res[index] += c
		} else {
			res[index] += candies
		}

		candies -= c
		c++
		index++

		if index == numPeople {
			index = 0
		}
	}

	return res
}

package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {

	athletes := make([][]int, len(nums))

	for i, num := range nums {
		athletes[i] = []int{i, num}
	}

	sort.Slice(athletes, func(i, j int) bool {
		return athletes[i][1] > athletes[j][1]
	})

	res := make([]string, len(athletes))
	for i, athlete := range athletes {
		switch i {
		case 0:
			res[athlete[0]] = "Gold Medal"
		case 1:
			res[athlete[0]] = "Silver Medal"
		case 2:
			res[athlete[0]] = "Bronze Medal"
		default:
			res[athlete[0]] = strconv.Itoa(i + 1)
		}
	}

	return res
}

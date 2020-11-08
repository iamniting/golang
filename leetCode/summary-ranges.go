package main

import (
	"strconv"
	"strings"
)

func summaryRanges(nums []int) []string {

	var res []string

	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i-1]+1 != nums[i] {
			res = append(res, strconv.Itoa(nums[i]))
		} else {
			res[len(res)-1] = strings.Split(res[len(res)-1], "->")[0] + "->" + strconv.Itoa(nums[i])
		}
	}

	return res
}

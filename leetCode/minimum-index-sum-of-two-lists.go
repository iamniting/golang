package main

import "math"

func findRestaurant(list1 []string, list2 []string) []string {

	var res []string
	set := map[string]int{}

	for i, r := range list1 {
		set[r] = i
	}

	var sum = math.MaxInt32
	for i, r := range list2 {
		if ind, ok := set[r]; ok && i+ind < sum {
			sum = i + ind
			res = []string{r}
		} else if ok && i+ind == sum {
			sum = i + ind
			res = append(res, r)
		}
	}

	return res
}

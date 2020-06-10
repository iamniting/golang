package main

import "sort"

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})

	var sum int
	for i, cost := range costs {
		if i < len(costs)/2 {
			sum += cost[0]
		} else {
			sum += cost[1]
		}
	}
	return sum
}

package main

func groupThePeople(groupSizes []int) [][]int {

	var res [][]int
	m := map[int][]int{}

	for id, gSize := range groupSizes {
		m[gSize] = append(m[gSize], id)
		if len(m[gSize]) == gSize {
			res = append(res, m[gSize])
			m[gSize] = []int{}
		}
	}

	return res
}

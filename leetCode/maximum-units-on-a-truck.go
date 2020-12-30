package main

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {

	sort.Slice(boxTypes, func(i, j int) bool { return boxTypes[i][1] > boxTypes[j][1] })

	var res int

	for _, box := range boxTypes {
		if box[0] < truckSize {
			res += box[0] * box[1]
		} else {
			res += truckSize * box[1]
			break
		}
		truckSize -= box[0]
	}

	return res
}

package main

func maximumUnits(boxTypes [][]int, truckSize int) int {

	var boxes [1001]int
	var res int

	for _, boxType := range boxTypes {
		boxes[boxType[1]] += boxType[0]
	}

	for i := 1000; i > 0 && truckSize > 0; i-- {
		if boxes[i] == 0 {
			continue
		}
		if boxes[i] < truckSize {
			truckSize -= boxes[i]
			res += boxes[i] * i
		} else {
			res += truckSize * i
			break
		}
	}

	return res
}

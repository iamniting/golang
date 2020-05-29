package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	res := 0

	for i := range startTime {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			res++
		}
	}
	return res
}

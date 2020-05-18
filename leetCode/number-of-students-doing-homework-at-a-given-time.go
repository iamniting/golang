// https://leetcode.com/problems/number-of-students-doing-homework-at-a-given-time
// Just sol to the problem, It does not include the I/O part

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	res := 0

	for i := range startTime {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			res++
		}
	}
	return res
}

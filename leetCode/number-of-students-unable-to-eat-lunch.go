package main

func countStudents(students []int, sandwiches []int) int {
	var counters [2]int

	for i := 0; i < len(students); i++ {
		counters[students[i]]++
	}

	// go over the sandwiches queue
	for i := 0; i < len(sandwiches); i++ {
		if counters[sandwiches[i]] == 0 {
			break // stop if there is no student of the kind left
		}
		counters[sandwiches[i]]-- // decrement number of students of the kind
	}

	// return number of students left
	return counters[0] + counters[1]
}

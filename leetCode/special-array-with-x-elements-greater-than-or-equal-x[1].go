package main

func specialArray(nums []int) int {

	var count int
	freq := make([]int, 1001)

	for _, n := range nums {
		freq[n]++
	}

	for x := 1000; x > 0; x-- {
		count += freq[x]
		// count == x mean count of elements greater or equal to the x
		if count == x {
			return x
		}
	}

	return -1
}

package main

func sumOfUnique(nums []int) int {

	freq := map[int]int{}
	var sum int

	for _, n := range nums {
		freq[n]++

		if freq[n] == 1 {
			sum += n
		} else if freq[n] == 2 {
			sum -= n
		}
	}

	return sum
}

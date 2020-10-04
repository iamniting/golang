package main

func specialArray(nums []int) int {

	for x := 1; x <= len(nums); x++ {
		var count int

		for _, n := range nums {
			if n >= x {
				count++
			}
		}
		if count == x {
			return x
		}
	}
	return -1
}

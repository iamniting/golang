package main

func findKthPositive(arr []int, k int) int {

	var missing []int
	var track [3000]bool

	for _, n := range arr {
		track[n] = true
	}

	for i := 1; i < 3000; i++ {
		if !track[i] {
			missing = append(missing, i)
			if len(missing) == k {
				return missing[k-1]
			}
		}
	}

	return -1
}

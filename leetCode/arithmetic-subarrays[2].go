package main

import "math"

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {

	res := make([]bool, len(l))

	for i := range l {
		start, end := l[i], r[i]
		min, max := math.MaxInt32, math.MinInt32

		// find min and max from sub array
		for j := start; j <= end; j++ {
			min = int(math.Min(float64(min), float64(nums[j])))
			max = int(math.Max(float64(max), float64(nums[j])))
		}

		// if min == max means all elements are equal
		if min == max {
			res[i] = true
			continue
		}
		// if can not divide
		if (max-min)%(end-start) != 0 {
			res[i] = false
			continue
		}

		slice := make([]bool, end-start+1)
		interval := (max - min) / (end - start)
		j := start

		// iterate sub array
		for ; j <= end; j++ {
			// if element can not be divided by interval after subtracting min
			// means element can not be part of arithmetic series || element already present
			if (nums[j]-min)%interval != 0 || slice[(nums[j]-min)/interval] {
				break
			}
			slice[(nums[j]-min)/interval] = true
		}

		// if j < end which means broken in b/w then false
		res[i] = j > end
	}

	return res
}

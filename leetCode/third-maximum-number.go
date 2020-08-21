package main

func thirdMax(nums []int) int {

	min := ^(1 << 32) + 1
	first, second, third := min, min, min

	for _, n := range nums {

		if n > first {
			first, second, third = n, first, second
		} else if n < first && n > second {
			second, third = n, second
		} else if n < first && n < second && n > third {
			third = n
		}
	}

	if third == min {
		return first
	}

	return third
}

package main

func countBits(num int) []int {

	res := make([]int, num+1)

	for i := 1; i <= num; i++ {
		if i&1 == 1 {
			// if i is odd it will be prev val + 1
			res[i] = res[i-1] + 1
		} else {
			// if i is even it will be i/2 eg 3 = 0011 and 6 = 0110
			// as we does multiplication with 2 without adding more bits eg 3 << 1 = 6
			res[i] = res[i>>1]
		}
	}

	return res
}

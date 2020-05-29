package main

func selfDividingNumbers(left int, right int) []int {

	IsSelfDiv := func(n int) bool {
		tmp := n
		for tmp > 0 {
			r := tmp % 10
			tmp = tmp / 10
			if r == 0 {
				return false
			}
			if n%r != 0 {
				return false
			}
		}
		return true
	}

	res := []int{}
	for i := left; i <= right; i++ {
		if IsSelfDiv(i) {
			res = append(res, i)
		}
	}
	return res
}

package main

func numJewelsInStones(J string, S string) int {

	res := 0

	for _, j := range J {
		for _, s := range S {
			if j == s {
				res++
			}
		}
	}

	return res
}

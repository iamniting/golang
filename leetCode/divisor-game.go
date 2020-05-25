package main

func divisorGame(N int) bool {
	flag := true
	alice := 0
	bob := 0

	for N > 1 {
		x := 1
		if N%x != 0 {
			return alice > bob
		}
		if flag {
			alice++
		} else {
			bob++
		}
		flag = !flag
		N = N - x
	}
	return alice > bob
}
